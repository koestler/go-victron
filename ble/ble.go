package ble

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/koestler/go-victron/bleparser"
	"github.com/koestler/go-victron/veproduct"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"log"
	"runtime/debug"
	"strings"
)

const VictronManufacturerId = 0x2e1

type Config interface {
	Name() string
	LogDebug() bool
	Devices() []DeviceConfig
}

type DeviceConfig interface {
	Name() string
	MacAddress() []byte
	EncryptionKey() []byte
}

type BleStruct struct {
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc
}

func New(cfg Config) (*BleStruct, error) {
	ctx, cancel := context.WithCancel(context.Background())

	ble := &BleStruct{
		cfg:    cfg,
		ctx:    ctx,
		cancel: cancel,
	}

	adapterID := "hci0"

	go func() {
		//clean up connection on exit
		defer api.Exit() //nolint:errcheck

		a, err := adapter.GetAdapter(adapterID)
		if err != nil {
			log.Printf("ble[%s]: error while getting adapter: %s", ble.Name(), err)
			return
		}

		err = a.FlushDevices()
		if err != nil {
			log.Printf("ble[%s]: error during flush: %s", ble.Name(), err)
			return
		}

		if cfg.LogDebug() {
			log.Printf("ble[%s]: start discovery", ble.Name())
		}

		discoveryFilter := adapter.NewDiscoveryFilter()
		discoveryFilter.DuplicateData = true
		discovery, cancel, err := api.Discover(a, &discoveryFilter)
		if err != nil {
			log.Printf("ble[%s]: cannot start discovery: %s", ble.Name(), err)
			return
		}
		defer cancel()

		go func() {
			for ev := range discovery {
				if ev.Type == adapter.DeviceRemoved {
					continue
				}

				devicePath := ev.Path
				dev, err := device.NewDevice1(devicePath)
				if err != nil {
					log.Printf("ble[%s]: error with path %s: %s", ble.Name(), devicePath, err)
					continue
				}

				// Filter out all non-Victron devices.
				if _, ok := dev.Properties.ManufacturerData[VictronManufacturerId]; !ok {
					continue
				}

				if cfg.LogDebug() {
					log.Printf("ble[%s]: device discovered: path=%s, name=%s, addr=%x, rssi=%d",
						ble.Name(), devicePath,
						dev.Properties.Name, dev.Properties.Address, dev.Properties.RSSI,
					)
				}

				deviceConfig := ble.getDeviceConfig(dev.Properties.Address)
				if deviceConfig == nil {
					continue
				}

				go func() {
					defer func() {
						if r := recover(); r != nil {
							log.Printf("ble[%s]: device %s panicked: %s %s", ble.Name(), deviceConfig.Name(), r, debug.Stack())

						}
					}()
					if err := ble.connectDevice(dev, deviceConfig); err != nil {
						log.Printf("ble[%s]: device %s failed: %s", ble.Name(), deviceConfig.Name(), err)
					}
				}()
			}
		}()

		// wait for cancel
		<-ble.ctx.Done()
	}()

	return ble, nil
}

func (ble *BleStruct) connectDevice(dev *device.Device1, deviceConfig DeviceConfig) error {
	propUpdates, err := dev.WatchProperties()
	if err != nil {
		return fmt.Errorf("cannot watch props: %s", err)
	}

	defer func() {
		if err := dev.UnwatchProperties(propUpdates); err != nil {
			log.Printf("error during unwatch: %s", err)
		} else {
			log.Printf("unwatch ok")
		}
	}()

	var lastRawBytes []uint8

	for pu := range propUpdates {
		if pu.Name == "ManufacturerData" {
			data := dev.Properties.ManufacturerData

			var rawBytes []uint8
			if md, ok := data[VictronManufacturerId]; !ok {
				log.Printf("ble[%s]->%s: invalid manufacturer data record",
					ble.cfg.Name(), deviceConfig.Name(),
				)
				continue
			} else {
				var isBytes bool
				if rawBytes, isBytes = md.([]uint8); !isBytes {
					log.Printf("ble[%s]->%s: manufacturer data has unexpected type %T",
						ble.cfg.Name(), deviceConfig.Name(), md,
					)
					continue
				}
			}

			if bytes.Equal(rawBytes, lastRawBytes) {
				continue
			}

			ble.handleNewManufacturerData(deviceConfig, rawBytes)

			lastRawBytes = rawBytes
		}
	}

	return nil
}

func (ble *BleStruct) handleNewManufacturerData(deviceConfig DeviceConfig, rawBytes []uint8) {
	debug := ble.cfg.LogDebug()
	prefixFmt := "ble[%s]->%s: "

	if debug {
		log.Printf(prefixFmt+"handle len=%d, rawBytes=%x",
			ble.cfg.Name(), deviceConfig.Name(), len(rawBytes), rawBytes,
		)
	}

	if len(rawBytes) < 9 {
		log.Printf(prefixFmt+"len(rawBytes) is too low", ble.cfg.Name(), deviceConfig.Name())
		return
	}

	// map rawBytes:
	// 00 - 01 : prefix
	// 02 - 03 : product id
	// 04 - 04 : record type
	// 05 - 06 : Nonce/Data counter in LSB order
	// 07 - 07 : first byte of encryption key
	// 08 -    : encrypted data

	prefix := rawBytes[0:2]
	productId := binary.LittleEndian.Uint16(rawBytes[2:4])
	product := veproduct.Product(productId)
	recordType := rawBytes[4]
	nonce := rawBytes[5:7] // used as iv for encryption; is only 16 bits
	iv := binary.LittleEndian.Uint16(nonce)

	firstByteOfEncryptionKey := rawBytes[7]
	encryptedBytes := rawBytes[8:]

	if debug {
		log.Printf(prefixFmt+"prefix=%x productId=%x productString=%s recordType=%x, nonce=%d, firstByteOfEncryptionKey=%x",
			ble.cfg.Name(), deviceConfig.Name(), prefix, productId, product.String(), recordType, nonce, firstByteOfEncryptionKey,
		)
		log.Printf(prefixFmt+"encryptedBytes=%x, len=%d",
			ble.cfg.Name(), deviceConfig.Name(), encryptedBytes, len(encryptedBytes),
		)
	}

	// decrypt using aes-ctr algorithm
	// encryption key of config is fixed to 32 hex chars, so 16 bytes, so 128-bit AES is used here
	key := deviceConfig.EncryptionKey()
	if len(key) == 0 {
		log.Printf(prefixFmt+"no encryption key configured", ble.cfg.Name(), deviceConfig.Name())
		return
	}

	// the advertisement contains the first byte of the key to detect a wrong key
	if key[0] != firstByteOfEncryptionKey {
		log.Printf(prefixFmt+"encryption key mismatch: device uses a key starting with %02x but configured key starts with %02x",
			ble.cfg.Name(), deviceConfig.Name(), firstByteOfEncryptionKey, key[0],
		)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf(prefixFmt+"cannot create aes cipher: %s", ble.cfg.Name(), deviceConfig.Name(), err)
		return
	}

	// iv needs to be 16 bytes for 128-bit AES, use nonce and pad with 0
	ivBytes := make([]byte, 16)
	binary.LittleEndian.PutUint16(ivBytes, iv)

	decryptedBytes := make([]byte, len(encryptedBytes))
	cipher.NewCTR(block, ivBytes).XORKeyStream(decryptedBytes, encryptedBytes)

	if debug {
		log.Printf(prefixFmt+"iv=%d ivBytes=%x, decryptedBytes=%x, len=%d",
			ble.cfg.Name(), deviceConfig.Name(), iv, ivBytes, decryptedBytes, len(decryptedBytes),
		)
	}

	name, record, err := decodeRecord(recordType, decryptedBytes)
	if err != nil {
		log.Printf(prefixFmt+"cannot decode %s record: %s", ble.cfg.Name(), deviceConfig.Name(), name, err)
		return
	}
	if record == nil {
		if debug {
			log.Printf(prefixFmt+"unsupported record type %#x", ble.cfg.Name(), deviceConfig.Name(), recordType)
		}
		return
	}
	log.Printf(prefixFmt+"%s record=%#v", ble.cfg.Name(), deviceConfig.Name(), name, record)
}

// decodeRecord decodes the decrypted data according to the record type.
// It returns a nil record and no error for unsupported record types.
func decodeRecord(recordType byte, data []byte) (name string, record any, err error) {
	switch recordType {
	case 0x01:
		name = "solar charger"
		record, err = bleparser.DecodeSolarChargerRecord(data)
	case 0x02:
		name = "battery monitor"
		record, err = bleparser.DecodeBatteryMonitorRecord(data)
	case 0x03:
		name = "inverter"
		record, err = bleparser.DecodeInverterRecord(data)
	case 0x04:
		name = "dc-dc converter"
		record, err = bleparser.DecodeDcDcConverterRecord(data)
	case 0x05:
		name = "smart lithium"
		record, err = bleparser.DecodeSmartLithiumRecord(data)
	case 0x06:
		name = "inverter rs"
		record, err = bleparser.DecodeInverterRsRecord(data)
	case 0x07:
		name = "gx device"
		record, err = bleparser.DecodeGxDeviceRecord(data)
	case 0x08:
		name = "ac charger"
		record, err = bleparser.DecodeAcChargerRecord(data)
	case 0x09:
		name = "smart battery protect"
		record, err = bleparser.DecodeSmartBatteryProtectRecord(data)
	case 0x0A:
		name = "lynx smart bms"
		record, err = bleparser.DecodeLynxSmartBmsRecord(data)
	case 0x0B:
		name = "multi rs"
		record, err = bleparser.DecodeMultiRsRecord(data)
	case 0x0C:
		name = "ve.bus"
		record, err = bleparser.DecodeVeBusRecord(data)
	case 0x0D:
		name = "dc energy meter"
		record, err = bleparser.DecodeDcEnergyMeterRecord(data)
	}
	return
}

func (ble *BleStruct) getDeviceConfig(bluezAddr string) DeviceConfig {
	for _, d := range ble.cfg.Devices() {
		if bytes.Equal(d.MacAddress(), bluezAddrBytes(bluezAddr)) {
			return d
		}
	}
	return nil
}

// input: D4:9D:D2:92:62:02
// output []byte{0xd4, 0x9d, 0xd2, 0x92, 0x62, 0x02}
func bluezAddrBytes(i string) []byte {
	b, err := hex.DecodeString(strings.ReplaceAll(i, ":", ""))
	if err != nil {
		log.Printf("cannot decode %s, got: %s", i, err)
	}
	return b
}

func (ble *BleStruct) Name() string {
	return ble.cfg.Name()
}

func (ble *BleStruct) Shutdown() {
	ble.cancel()
}
