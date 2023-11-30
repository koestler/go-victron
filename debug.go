package ble

import (
	"fmt"
	"log"
)

func (md *BleStruct) debugPrintf(format string, v ...interface{}) {
	// check if debug output is enabled
	if !md.cfg.LogDebug() {
		return
	}

	s := fmt.Sprintf(format, v...)
	log.Printf("ble[%s]: %s", md.cfg.Name(), s)
}
