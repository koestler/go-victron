package ble

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
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

				dev, err := device.NewDevice1(ev.Path)
				if err != nil {
					log.Printf("ble[%s]: error with path %s: %s", ble.Name(), ev.Path, err)
					continue
				}

				if cfg.LogDebug() {
					log.Printf("ble[%s] device recovered: path=%s, name=%s, addr=%x, rssi=%d",
						ble.Name(), ev.Path,
						dev.Properties.Name, dev.Properties.Address, dev.Properties.RSSI,
					)
				}

				deviceConfig := ble.getDeviceConfig(dev.Properties.Address)
				if deviceConfig == nil {
					continue
				}

				go func() {
					err = ble.connectDevice(dev, deviceConfig)
					if err != nil {
						log.Printf("ble[%s]: device %s failed: %s", ble.Name(), ev.Path, err)
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
	// 02 - 03 : model id
	// 04 - 04 : record type
	// 05 - 06 : Nonce/Data counter in LSB order
	// 07 - 07 : first byte of encryption key

	prefix := rawBytes[0:1]
	modelId := rawBytes[2:3]
	recordType := rawBytes[4]
	nonce := rawBytes[5:6]
	firstByteOfEncryptionKey := rawBytes[7]

	log.Printf("ble[%s]->%s: recordType=%x, nonce=%x, firstByteOfEncryptionKey=%x",
		ble.cfg.Name(), deviceConfig.Name(), recordType, nonce, firstByteOfEncryptionKey,
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
