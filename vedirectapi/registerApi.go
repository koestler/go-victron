// Package vedirectapi implements an easy to use api to connect to VE.Direct devices and read their values.
// It utilizes the low level vedirect, veregister and veproduct packages and intatiates a connection using the serial package.
package vedirectapi

import (
	"context"
	"fmt"
	"github.com/koestler/go-victron/veconst"
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
	ioPort    vedirect.IOPort
	Vd        *vedirect.Vedirect
	Product   veproduct.Product
	Registers veregister.RegisterList
}

var ErrCtxDone = fmt.Errorf("context done")

// NewSerialRegisterApi open the given serial device, creates a new RegisterApi instance,
// and tries to connect to the device over a serial connection.
func NewSerialRegisterApi(serialDevice string, vdConfig vedirect.Config) (*RegisterApi, error) {
	serialConfig := serial.Config{
		Name:        serialDevice,
		Baud:        19200,
		ReadTimeout: time.Millisecond * 200,
	}

	if ioHandle, err := serial.OpenPort(&serialConfig); err != nil {
		return nil, fmt.Errorf("cannot open port %s: %w", serialDevice, err)
	} else {
		return NewRegisterApi(ioHandle, vdConfig)
	}
}

// NewRegisterApi creates a new RegisterApi instance and tries to connect to the device:
// - sends a ping
// - gets the device id and uses it to get the product type
// - gets the register list for the product type
// Make sure to defer RegisterApi.Close() after creating a new RegisterApi instance.
func NewRegisterApi(ioPort vedirect.IOPort, vdConfig vedirect.Config) (*RegisterApi, error) {
	sa := RegisterApi{
		ioPort: ioPort,
	}

	if vd, err := vedirect.NewVedirect(sa.ioPort, vdConfig); err != nil {
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
	if rl, err := veregister.GetRegisterListByProduct(sa.Product); err != nil {
		return nil, fmt.Errorf("cannot get register list: %w", err)
	} else {
		sa.Registers = rl
	}

	return &sa, nil
}

// Close closes the underlying serial port.
func (sa *RegisterApi) Close() error {
	return sa.ioPort.Close()
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
func (sa *RegisterApi) ReadEnumRegister(r veregister.EnumRegisterStruct) (veconst.Enum, error) {
	intValue, err := sa.Vd.GetUint(r.Address())

	if err != nil {
		return nil, fmt.Errorf("fetching enum register '%s' failed: %w", r.Name(), err)
	}

	if e, err := r.Factory().NewEnum(int(intValue)); err != nil {
		return nil, fmt.Errorf("decoding enum register '%s' failed: %w", r.Name(), err)
	} else {
		return e, nil
	}
}

// ReadFieldListRegister fetches a single field list register and decodes the field list to a map of field name to field value.
func (sa *RegisterApi) ReadFieldListRegister(r veregister.FieldListRegisterStruct) (veconst.FieldList, error) {
	intValue, err := sa.Vd.GetUint(r.Address())

	if err != nil {
		return nil, fmt.Errorf("fetching field list register '%s' failed: %w", r.Name(), err)
	}

	if fl, err := r.Factory().NewFieldList(uint(intValue)); err != nil {
		return nil, fmt.Errorf("decoding field list register '%s' failed: %w", r.Name(), err)
	} else {
		return fl, nil
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
		NumberValues:    make(map[string]NumberRegisterValue, len(rl.NumberRegisters)),
		TextValues:      make(map[string]TextRegisterValue, len(rl.TextRegisters)),
		EnumValues:      make(map[string]EnumRegisterValue, len(rl.EnumRegisters)),
		FieldListValues: make(map[string]FieldListValue, len(rl.FieldListRegisters)),
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
			FieldList: func(v FieldListValue) {
				rv.FieldListValues[v.Name()] = v
			},
		})

	return
}

// ValueHandler is a container for handlers for number, text and enum registers values.
// This is made such that new register types can be added without breaking the api.
type ValueHandler struct {
	Number    func(v NumberRegisterValue)
	Text      func(v TextRegisterValue)
	Enum      func(v EnumRegisterValue)
	FieldList func(v FieldListValue)
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
			v, err := sa.ReadEnumRegister(r)
			if err != nil {
				return err
			}
			handlers.Enum(EnumRegisterValue{
				EnumRegisterStruct: r,
				value:              v,
			})
		}
	}

	if handlers.FieldList != nil {
		for _, r := range rl.FieldListRegisters {
			select {
			case <-ctx.Done():
				return ErrCtxDone
			default:
			}
			v, err := sa.ReadFieldListRegister(r)
			if err != nil {
				return err
			}
			handlers.FieldList(FieldListValue{
				FieldListRegisterStruct: r,
				value:                   v,
			})
		}
	}

	return nil
}
