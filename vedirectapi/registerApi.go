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

// RegisterApi is the main struct for the VE.Direct serial protocol api. It allows to communicate with the device
// by reading the device id, using it to get the list of available registers and than allow to read those registers.
type RegisterApi struct {
	IoHandle  *serial.Port
	Vd        *vedirect.Vedirect
	Product   veproduct.Product
	Registers veregisters.RegisterList
}

// NewRegistertApi creates a new RegisterApi instance and tries to connect to the device:
// - opens the serial port
// - sends a ping
// - gets the device id and uses it to get the product type
// - gets the register list for the product type
// Make sure to defer RegisterApi.Close() after creating a new RegisterApi instance.
func NewRegistertApi(serialDevice string, vdConfig vedirect.Config) (*RegisterApi, error) {
	sa := RegisterApi{}

	serialConfig := serial.Config{
		Name:        serialDevice,
		Baud:        19200,
		ReadTimeout: time.Millisecond * 200,
	}

	if ioHandle, err := serial.OpenPort(&serialConfig); err != nil {
		return nil, fmt.Errorf("cannot open port %s: %w", serialDevice, err)
	} else {
		sa.IoHandle = ioHandle
	}

	if vd, err := vedirect.NewVedirect(sa.IoHandle, vdConfig); err != nil {
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
func (sa *RegisterApi) Close() error {
	return sa.IoHandle.Close()
}

// ReadNumberRegister fetches a single number register and converts it to a float64.
func (sa *RegisterApi) ReadNumberRegister(r veregisters.NumberRegisterStruct) (value float64, err error) {
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

// ReadTextRegister fetches a single text register.
func (sa *RegisterApi) ReadTextRegister(r veregisters.TextRegisterStruct) (value string, err error) {
	value, err = sa.Vd.GetString(r.Address)
	if err != nil {
		return "", fmt.Errorf("fetching text register failed: %w", err)
	}
	return
}

// ReadEnumRegister fetches a single enum register and decodes the enum to enum index and enum value.
func (sa *RegisterApi) ReadEnumRegister(r veregisters.EnumRegisterStruct) (enumIdx int, enumValue string, err error) {
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

// RegisterValues is a container for Number, Text and Enum Registers and their values.
type RegisterValues struct {
	NumberValues map[string]NumberRegisterValue
	TextValues   map[string]TextRegisterValue
	EnumValues   map[string]EnumRegisterValue
}

// ReadRegisterList fetches all registers from the given list and returns them as a RegisterValues struct.
// When an error occurs, fetching is aborted and the error is returned.
func (sa *RegisterApi) ReadRegisterList(rl veregisters.RegisterList) (RegisterValues, error) {
	rv := RegisterValues{
		NumberValues: make(map[string]NumberRegisterValue),
		TextValues:   make(map[string]TextRegisterValue),
		EnumValues:   make(map[string]EnumRegisterValue),
	}

	for _, r := range rl.NumberRegisters {
		v, err := sa.ReadNumberRegister(r)
		if err != nil {
			return rv, err
		}
		rv.NumberValues[r.Name] = NumberRegisterValue{
			Register: r,
			Value:    v,
		}
	}

	for _, r := range rl.TextRegisters {
		v, err := sa.ReadTextRegister(r)
		if err != nil {
			return rv, err
		}
		rv.TextValues[r.Name] = TextRegisterValue{
			Register: r,
			Value:    v,
		}
	}

	for _, r := range rl.EnumRegisters {
		idx, v, err := sa.ReadEnumRegister(r)
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

// ReadAllRegisters fetches all available registers and returns them as a RegisterValues struct.
func (sa *RegisterApi) ReadAllRegisters() (RegisterValues, error) {
	return sa.ReadRegisterList(sa.Registers)
}

// StreamRegisterList fetches all registers from the given list and calls the given handlers for each register.
// When an error occurs, fetching is aborted and the error is returned.
// This is useful since fetching all registers of a device can take up to a second. This way, you can start processing
// the values as soon as they are available.
func (sa *RegisterApi) StreamRegisterList(
	rl veregisters.RegisterList,
	handlers struct {
		number func(register veregisters.NumberRegisterStruct, value float64)
		text   func(register veregisters.TextRegisterStruct, value string)
		enum   func(register veregisters.EnumRegisterStruct, enumIdx int, enumValue string)
	},
) error {
	if handlers.number != nil {
		for _, r := range rl.NumberRegisters {
			v, err := sa.ReadNumberRegister(r)
			if err != nil {
				return err
			}
			handlers.number(r, v)
		}
	}

	if handlers.text != nil {
		for _, r := range rl.TextRegisters {
			v, err := sa.ReadTextRegister(r)
			if err != nil {
				return err
			}
			handlers.text(r, v)
		}
	}

	if handlers.enum != nil {
		for _, r := range rl.EnumRegisters {
			v, ev, err := sa.ReadEnumRegister(r)
			if err != nil {
				return err
			}
			handlers.enum(r, v, ev)
		}
	}

	return nil
}
