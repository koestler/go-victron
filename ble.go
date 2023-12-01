package ble

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/koestler/go-iotdevice/victronDefinitions"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"log"
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
		defer api.Exit()

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

				if cfg.LogDebug() {
					log.Printf("ble[%s]: device recovered: path=%s, name=%s, addr=%x, rssi=%d",
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
							log.Printf("ble[%s]: device %s panicked: %s", ble.Name(), deviceConfig.Name(), r)
						}
					}()
					if err = ble.connectDevice(dev, deviceConfig); err != nil {
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
		// unwatch properties calls recover
		if r := recover(); r != nil {
			panic(r)
		}

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
				rawBytes = md.([]uint8)
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
	if ble.cfg.LogDebug() {
		log.Printf("ble[%s]->%s: handle rawBytes=%x",
			ble.cfg.Name(), deviceConfig.Name(), rawBytes,
		)
	}

	if len(rawBytes) < 4 {
		log.Printf("ble[%s]->%s: len(rawBytes) is to low",
			ble.cfg.Name(), deviceConfig.Name(),
		)
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
	productId := uint16(rawBytes[2]) | uint16(rawBytes[3])<<8
	product := victronDefinitions.VeProduct(productId)
	recordType := rawBytes[4]
	nonce := rawBytes[5:7] // used ad iv for encryption; is only 16 bits

	firstByteOfEncryptionKey := rawBytes[7]
	encryptedBytes := rawBytes[8:]

	log.Printf("ble[%s]->%s: prefix=%x productId=%x productString=%s recordType=%x, nonce=%x, firstByteOfEncryptionKey=%x",
		ble.cfg.Name(), deviceConfig.Name(), prefix, productId, product.String(), recordType, nonce, firstByteOfEncryptionKey,
	)

	// decrypt rawBytes using aes-ctr algorithm
	// encryption key of config is fixed to 32 hex chars, so 16 bytes, so 128-bit AES is used here
	block, err := aes.NewCipher(deviceConfig.EncryptionKey())
	if err != nil {
		log.Printf("ble[%s]->%s: cannot create aes cipher: %s",
			ble.cfg.Name(), deviceConfig.Name(), err,
		)
		return
	}

	decryptedBytes := make([]byte, len(encryptedBytes))

	// iv needs to be 16 bytes for 128-bit AES, use nonce and pad with 0
	iv := make([]byte, 14, 16)
	iv = append(iv, nonce...)

	ctrStream := cipher.NewCTR(block, iv)
	ctrStream.XORKeyStream(decryptedBytes, encryptedBytes)

	log.Printf("ble[%s]->%s: handle decryptedBytes=%#v",
		ble.cfg.Name(), deviceConfig.Name(), decryptedBytes,
	)

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
// output []byte{0x, 0xd4, 0x9d, 0xd2, 0x92, 0x62, 0x02}
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
