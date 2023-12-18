// Package vedirectapi implements an easy to use api to connect to VE.Direct devices and read their values.
// It utilizes  the low level vedirect, veregisters and veproduct packages and intatiates a connection using the serial package.
package vedirectapi

import (
	"fmt"
	"github.com/koestler/go-victron/vedirect"
	"github.com/tarm/serial"
	"log"
	"time"
)

type Config struct {
	PortName    string
	IoLogger    *log.Logger
	DebugLogger *log.Logger
}

func OpenSerialPort(portName string, debugLogger *log.Logger) (*vedirect.Vedirect, error) {
	if debugLogger != nil {
		debugLogger.Printf("vedirect: Open portName=%v", portName)
	}

	options := serial.Config{
		Name:        portName,
		Baud:        19200,
		ReadTimeout: time.Millisecond * 200,
	}

	ioHandle, err := serial.OpenPort(&options)
	if err != nil {
		return nil, fmt.Errorf("cannot open port: %v", portName)
	}

}
