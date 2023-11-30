package ble

import (
	"fmt"
	"log"
)

func (ble *BleStruct) debugPrintf(format string, v ...interface{}) {
	// check if debug output is enabled
	if !ble.cfg.LogDebug() {
		return
	}

	s := fmt.Sprintf(format, v...)
	log.Printf("ble[%s]: %s", ble.cfg.Name(), s)
}
