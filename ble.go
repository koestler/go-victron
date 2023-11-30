package ble

import (
	"log"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

type Config interface {
	Name() string
	LogDebug() bool
	Devices() []DeviceConfig
}

type DeviceConfig interface {
	MacAddress() string
	EncryptionKey() string
}

type BleStruct struct {
	cfg Config
}

func New(cfg Config) (*BleStruct, error) {
	if cfg.LogDebug() {
		log.Printf("ble[%s]: create", cfg.Name())
	}

	ble := &BleStruct{
		cfg: cfg,
	}

	go func() {
		if err := adapter.Enable(); err != nil {
			log.Printf("ble[%s]: error during enable: %s", cfg.Name(), err)
		}

		// Start scanning.
		log.Printf("ble[%s]: start scanning", cfg.Name())

		if err := adapter.Scan(func(a *bluetooth.Adapter, result bluetooth.ScanResult) {
			ble.bleAdvHandler(a, result)
		}); err != nil {
			log.Printf("ble[%s]: error during scanning: %s", cfg.Name(), err)
		}
	}()

	return ble, nil
}

func (ble *BleStruct) bleAdvHandler(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
	if device.LocalName() != "SmartSolar HQ19499RHC5" {
		return
	}

	log.Printf("ble[%s]: found device: %s %d %s", ble.Name(), device.Address.String(), device.RSSI, device.LocalName())
	log.Printf("ble[%s]: address: %x, isRandom=%v", ble.Name(), device.Address.String(), device.Address.IsRandom())
	// log.Printf("ble[%s]: bytes received: %x", ble.Name(), device.Bytes())
	// log.Printf("ble[%s]: ManufacturerData: %#v", ble.Name(), device.ManufacturerData())
}

func (ble *BleStruct) getDevice() {

}

func (ble *BleStruct) Name() string {
	return ble.cfg.Name()
}

func (ble *BleStruct) Shutdown() {

}
