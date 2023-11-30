package ble

import (
	"context"
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
	MacAddress() string
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

	for pu := range propUpdates {
		if pu.Name == "ManufacturerData" {
			ble.handleNewManufacturerData(deviceConfig, dev.Properties.ManufacturerData)
		}
	}

	return nil
}

func (ble *BleStruct) handleNewManufacturerData(deviceConfig DeviceConfig, data map[uint16]interface{}) {
	if ble.cfg.LogDebug() {
		log.Printf("ble[%s]->%s: handle data=%#v",
			ble.cfg.Name(), deviceConfig.Name(), data,
		)
	}

	var rawBytes []uint8

	if md, ok := data[VictronManufacturerId]; !ok {
		log.Printf("ble[%s]->%s: invalid manufacturer data record",
			ble.cfg.Name(), deviceConfig.Name(),
		)
		return
	} else {
		rawBytes = md.([]uint8)
	}

	log.Printf("ble[%s]->%s: handle rawBytes=%x",
		ble.cfg.Name(), deviceConfig.Name(), rawBytes,
	)

	if len(rawBytes) < 4 {
		log.Printf("ble[%s]->%s: len(rawBytes) is to low",
			ble.cfg.Name(), deviceConfig.Name(),
		)
		return
	}

	// map rawBytes:
	// 00 - 03 : unknown
	// 04 - 04 : record type
	// 05 - 06 : Nonce/Data counter in LSB order
	// 07 - 07 : first byte of encryption key

	recordType := rawBytes[4]
	nonce := rawBytes[5:6]
	firstByteOfEncryptionKey := rawBytes[7]

	log.Printf("ble[%s]->%s: recordType=%x, nonce=%x, firstByteOfEncryptionKey=%x",
		ble.cfg.Name(), deviceConfig.Name(), recordType, nonce, firstByteOfEncryptionKey,
	)

}

func (ble *BleStruct) getDeviceConfig(bluezAddr string) DeviceConfig {
	for _, d := range ble.cfg.Devices() {
		if d.MacAddress() == bluezAddrToOurAddr(bluezAddr) {
			return d
		}
	}
	return nil
}

// input: D4:9D:D2:92:62:02
// output d49dd2926202
func bluezAddrToOurAddr(i string) string {
	return strings.ToLower(strings.ReplaceAll(i, ":", ""))
}

func (ble *BleStruct) Name() string {
	return ble.cfg.Name()
}

func (ble *BleStruct) Shutdown() {
	ble.cancel()
}
