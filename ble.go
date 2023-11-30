package ble

import (
	"log"
)

type BleStruct struct {
	cfg Config
}

func New(cfg Config) (*BleStruct, error) {
	if cfg.LogDebug() {
		log.Printf("ble[%s]: create", cfg.Name())
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
