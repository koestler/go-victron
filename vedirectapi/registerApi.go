// Package vedirectapi implements an easy to use api to connect to VE.Direct devices and read their values.
// It utilizes the low level vedirect, veregister and veproduct packages and intatiates a connection using the serial package.
package vedirectapi

import (
	"context"
	"fmt"
	"github.com/koestler/go-victron/vedirect"
	"github.com/koestler/go-victron/veproduct"
	"github.com/koestler/go-victron/veregister"
	"github.com/tarm/serial"
	"strings"
	"time"
)

// RegisterApi is the main struct for the VE.Direct serial protocol api. It allows to communicate with the device
// by reading the device id, using it to get the list of available registers and than allow to read those registers.
type RegisterApi struct {
	IoHandle  *serial.Port
	Vd        *vedirect.Vedirect
	Product   veproduct.Product
	Registers veregister.RegisterList
}

var ErrCtxDone = fmt.Errorf("context done")

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
	if rl, err := veregister.GetRegisterListByProductType(sa.Product.Type()); err != nil {
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
func (sa *RegisterApi) ReadNumberRegister(r veregister.NumberRegisterStruct) (value float64, err error) {
	if r.Signed() {
		var intValue int64
		intValue, err = sa.Vd.GetInt(r.Address())
		value = float64(intValue)
	} else {
		var intValue uint64
		intValue, err = sa.Vd.GetUint(r.Address())
		value = float64(intValue)
	}

	if err != nil {
		return 0.0, fmt.Errorf("fetching number register '%s' failed: %w", r.Name(), err)
	}

	value = value/float64(r.Factor()) + r.Offset()
	return
}

// ReadTextRegister fetches a single text register.
func (sa *RegisterApi) ReadTextRegister(r veregister.TextRegisterStruct) (value string, err error) {
	value, err = sa.Vd.GetString(r.Address())

	if err != nil {
		return "", fmt.Errorf("fetching text register '%s' failed: %w", r.Name(), err)
	}

	value = strings.TrimSpace(value) // various devices have trailing spaces in their text registers
	return
}

// ReadEnumRegister fetches a single enum register and decodes the enum to enum index and enum value.
func (sa *RegisterApi) ReadEnumRegister(r veregister.EnumRegisterStruct) (enumIdx int, enumValue string, err error) {
	var intValue uint64
	intValue, err = sa.Vd.GetUint(r.Address())

	if err != nil {
		return 0, "", fmt.Errorf("fetching enum register '%s' failed: %w", r.Name(), err)
	}

	if bit := r.Bit(); bit >= 0 {
		intValue = (intValue >> bit) & 1
	}
	enumIdx = int(intValue)

	if e, err := r.Factory()(enumIdx); err != nil {
		return e.Idx(), e.Value(), nil
	} else {
		return 0, "", fmt.Errorf("decoding enum register '%s' failed: %w", r.Name(), err)
	}
}

// ReadAllRegisters fetches all available registers and returns them as a RegisterValues struct.
func (sa *RegisterApi) ReadAllRegisters(ctx context.Context) (RegisterValues, error) {
	return sa.ReadRegisterList(ctx, sa.Registers)
}

// ReadRegisterList fetches all registers from the given list and returns them as a RegisterValues struct.
// When an error occurs, fetching is aborted and the error is returned.
func (sa *RegisterApi) ReadRegisterList(
	ctx context.Context,
	rl veregister.RegisterList,
) (rv RegisterValues, err error) {
	rv = RegisterValues{
		NumberValues: make(map[string]NumberRegisterValue, len(rl.NumberRegisters)),
		TextValues:   make(map[string]TextRegisterValue, len(rl.TextRegisters)),
		EnumValues:   make(map[string]EnumRegisterValue, len(rl.EnumRegisters)),
	}

	err = sa.StreamRegisterList(
		ctx, rl, ValueHandler{
			Number: func(v NumberRegisterValue) {
				rv.NumberValues[v.Name()] = v
			},
			Text: func(v TextRegisterValue) {
				rv.TextValues[v.Name()] = v
			},
			Enum: func(v EnumRegisterValue) {
				rv.EnumValues[v.Name()] = v
			},
		})

	return
}

// ValueHandler is a container for handlers for number, text and enum registers values.
// This is made such that new register types can be added without breaking the api.
type ValueHandler struct {
	Number func(v NumberRegisterValue)
	Text   func(v TextRegisterValue)
	Enum   func(v EnumRegisterValue)
}

// StreamRegisterList fetches all registers from the given list and calls the given handlers for each register.
// When an error occurs, fetching is aborted and the error is returned.
// This is useful since fetching all registers of a device can take up to a second. This way, you can start processing
// the values as soon as they are available.
func (sa *RegisterApi) StreamRegisterList(
	ctx context.Context,
	rl veregister.RegisterList,
	handlers ValueHandler,
) error {
	if handlers.Number != nil {
		for _, r := range rl.NumberRegisters {
			select {
			case <-ctx.Done():
				return ErrCtxDone
			default:
			}
			v, err := sa.ReadNumberRegister(r)
			if err != nil {
				return err
			}
			handlers.Number(NumberRegisterValue{
				NumberRegisterStruct: r,
				value:                v,
			})
		}
	}

	if handlers.Text != nil {
		for _, r := range rl.TextRegisters {
			select {
			case <-ctx.Done():
				return ErrCtxDone
			default:
			}
			v, err := sa.ReadTextRegister(r)
			if err != nil {
				return err
			}
			handlers.Text(TextRegisterValue{
				TextRegisterStruct: r,
				value:              v,
			})
		}
	}

	if handlers.Enum != nil {
		for _, r := range rl.EnumRegisters {
			select {
			case <-ctx.Done():
				return ErrCtxDone
			default:
			}
			idx, v, err := sa.ReadEnumRegister(r)
			if err != nil {
				return err
			}
			handlers.Enum(EnumRegisterValue{
				EnumRegisterStruct: r,
				enumIdx:            idx,
				enumValue:          v,
			})
		}
	}

	return nil
}
