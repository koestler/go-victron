package ble

import (
	"log"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

type BleStruct struct {
	cfg Config
}

func New(cfg Config) (*BleStruct, error) {
	if cfg.LogDebug() {
		log.Printf("ble[%s]: create", cfg.Name())
	}

	if err := adapter.Enable(); err != nil {
		log.Printf("ble[%s]: error during enable: %s", cfg.Name(), err)
	}

	// Start scanning.
	log.Printf("ble[%s]: start scanning", cfg.Name())
	{
		err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
			log.Printf("ble[%s]: found device: %s %d %s", cfg.Name(), device.Address.String(), device.RSSI, device.LocalName())
		})
		if err != nil {
			log.Printf("ble[%s]: error during scanning: %s", cfg.Name(), err)
		}
	}

	return &BleStruct{
		cfg: cfg,
	}, nil
}

func (md *BleStruct) Name() string {
	return md.cfg.Name()
}

func (md *BleStruct) Shutdown() {

}
