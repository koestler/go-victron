// Package vedirectapi implements an easy to use api to connect to VE.Direct devices and read their values.
// It utilizes  the low level vedirect, veregisters and veproduct packages and intatiates a connection using the serial package.
package vedirectapi

import (
	"fmt"
	"github.com/koestler/go-victron/vedirect"
	"github.com/koestler/go-victron/veproduct"
	"github.com/koestler/go-victron/veregisters"
	"github.com/tarm/serial"
	"time"
)

// Config is the configuration for the Api.
// PortName is the only mandatory field.
type Config struct {
	PortName    string          // mandatory: the serial device: e.g. /dev/ttyUSB0
	DebugLogger vedirect.Logger // optional: a logger for debug output; if nil, no debug output is written
	IoLogger    vedirect.Logger // optional: a logger for all io operations; if nil, no io output is written
}

type Api struct {
	IoHandle  *serial.Port
	Vd        *vedirect.Vedirect
	Product   veproduct.Product
	Registers veregisters.RegisterList
}

// NewApi creates a new Api instance and tries to connect to the device:
// - opens the serial port
// - sends a ping
// - gets the device id and uses it to get the product type
// - gets the register list for the product type
// Make sure to defer Api.Close() after creating a new Api instance.
func NewApi(config Config) (*Api, error) {
	sa := Api{}

	serialConfig := serial.Config{
		Name:        config.PortName,
		Baud:        19200,
		ReadTimeout: time.Millisecond * 200,
	}

	if ioHandle, err := serial.OpenPort(&serialConfig); err != nil {
		return nil, fmt.Errorf("cannot open port %s: %w", config.PortName, err)
	} else {
		sa.IoHandle = ioHandle
	}

	vdConfig := vedirect.Config{
		IOPort:      sa.IoHandle,
		IoLogger:    config.IoLogger,
		DebugLogger: config.DebugLogger,
	}

	if vd, err := vedirect.NewVedirect(&vdConfig); err != nil {
		return nil, fmt.Errorf("cannot create vedirect: %w", err)
	} else {
		sa.Vd = vd
	}

	// send ping
	if err := sa.Vd.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	// get deviceId
	if deviceId, err := sa.Vd.GetDeviceId(); err != nil {
		return nil, fmt.Errorf("cannot get DeviceId: %w", err)
	} else {
		// get Product
		sa.Product = veproduct.Product(deviceId)
		if !sa.Product.Exists() {
			return nil, fmt.Errorf("unknown Product: %w", err)
		}
	}

	// get register list
	if rl, err := veregisters.GetRegisterListByProductType(sa.Product.Type()); err != nil {
		return nil, fmt.Errorf("cannot get register list: %w", err)
	} else {
		sa.Registers = rl
	}

	return &sa, nil
}

// Close closes the underlying serial port.
func (sa *Api) Close() error {
	return sa.IoHandle.Close()
}

// FetchNumberRegister fetches a single number register and converts it to a float64.
func (sa *Api) FetchNumberRegister(r veregisters.NumberRegisterStruct) (value float64, err error) {
	if r.Signed {
		var intValue int64
		intValue, err = sa.Vd.GetInt(r.Address)
		value = float64(intValue)
	} else {
		var intValue uint64
		intValue, err = sa.Vd.GetUint(r.Address)
		value = float64(intValue)
	}

	if err != nil {
		return 0.0, fmt.Errorf("fetching number register failed: %w", err)
	}

	value = value/float64(r.Factor) + r.Offset

	return
}

// FetchTextRegister fetches a single text register.
func (sa *Api) FetchTextRegister(r veregisters.TextRegisterStruct) (value string, err error) {
	value, err = sa.Vd.GetString(r.Address)
	if err != nil {
		return "", fmt.Errorf("fetching text register failed: %w", err)
	}
	return
}

// FetchEnumRegister fetches a single enum register and decodes the enum to enum index and enum value.
func (sa *Api) FetchEnumRegister(r veregisters.EnumRegisterStruct) (enumIdx int, enumValue string, err error) {
	var intValue uint64

	intValue, err = sa.Vd.GetUint(r.Address)
	if err != nil {
		return 0, "", fmt.Errorf("fetching enum register failed: %w", err)
	}

	if bit := r.Bit; bit >= 0 {
		intValue = (intValue >> bit) & 1
	}
	enumIdx = int(intValue)

	// decode enum
	var ok bool
	enumValue, ok = r.Enum[enumIdx]

	if !ok {
		return 0, "", fmt.Errorf("unknown enum value: %d", intValue)
	}

	return
}

type NumberRegisterValue struct {
	Register veregisters.NumberRegisterStruct
	Value    float64
}

type TextRegisterValue struct {
	Register veregisters.TextRegisterStruct
	Value    string
}

type EnumRegisterValue struct {
	Register  veregisters.EnumRegisterStruct
	EnumIdx   int
	EnumValue string
}

type RegisterValues struct {
	NumberValues map[string]NumberRegisterValue
	TextValues   map[string]TextRegisterValue
	EnumValues   map[string]EnumRegisterValue
}

// FetchRegisterList fetches all registers from the register list and returns them as a RegisterValues struct.
// When an error occurs, fetching is aborted and the error is returned.
func (sa *Api) FetchRegisterList(rl veregisters.RegisterList) (RegisterValues, error) {
	rv := RegisterValues{
		NumberValues: make(map[string]NumberRegisterValue),
		TextValues:   make(map[string]TextRegisterValue),
		EnumValues:   make(map[string]EnumRegisterValue),
	}

	for _, r := range rl.NumberRegisters {
		v, err := sa.FetchNumberRegister(r)
		if err != nil {
			return rv, err
		}
		rv.NumberValues[r.Name] = NumberRegisterValue{
			Register: r,
			Value:    v,
		}
	}

	for _, r := range rl.TextRegisters {
		v, err := sa.FetchTextRegister(r)
		if err != nil {
			return rv, err
		}
		rv.TextValues[r.Name] = TextRegisterValue{
			Register: r,
			Value:    v,
		}
	}

	for _, r := range rl.EnumRegisters {
		idx, v, err := sa.FetchEnumRegister(r)
		if err != nil {
			return rv, err
		}
		rv.EnumValues[r.Name] = EnumRegisterValue{
			Register:  r,
			EnumIdx:   idx,
			EnumValue: v,
		}
	}

	return rv, nil
}

// StreamRegisterList fetches all registers from the register list and calls the given handlers for each register.
// When an error occurs, fetching is aborted and the error is returned.
// This is useful since fetching all registers of a device can take up to a second. This way, you can start processing
// the values as soon as they are available.
func (sa *Api) StreamRegisterList(
	rl veregisters.RegisterList,
	handleNumber func(register veregisters.NumberRegisterStruct, value float64),
	handleText func(register veregisters.TextRegisterStruct, value string),
	handleEnum func(register veregisters.EnumRegisterStruct, enumIdx int, enumValue string),
) error {
	for _, r := range rl.NumberRegisters {
		v, err := sa.FetchNumberRegister(r)
		if err != nil {
			return err
		}
		handleNumber(r, v)
	}

	for _, r := range rl.TextRegisters {
		v, err := sa.FetchTextRegister(r)
		if err != nil {
			return err
		}
		handleText(r, v)
	}

	for _, r := range rl.EnumRegisters {
		v, ev, err := sa.FetchEnumRegister(r)
		if err != nil {
			return err
		}
		handleEnum(r, v, ev)
	}

	return nil
}
