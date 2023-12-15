package ble

import (
	"fmt"
	"log"
)

func (ble *BleStruct) printf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Printf("bleparser[%s]: %s", ble.cfg.Name(), s)
}

func (ble *BleStruct) debugPrintf(format string, v ...interface{}) {
	// check if debug output is enabled
	if !ble.cfg.LogDebug() {
		return
	}

	ble.printf(format, v...)
}
