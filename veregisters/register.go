// Package veregisters contains the register definitions for the VE.Direct protocol.
package veregisters

import (
	"golang.org/x/exp/constraints"
)

// RegisterStruct is the base struct for all register types.
type RegisterStruct struct {
	category    string // the category of the register (e.g. "Essential")
	name        string // the technical name of the register (e.g. "PanelVoltage")
	description string // the description of the register (e.g. "Panel Voltage")
	sort        int    // the sort order of the registers
	address     uint16 // the memory address used to fetch register from the device
	static      bool   // when true, the register is static and not updated by the device
	writable    bool   // when true, the register can be written to
}

// NumberRegisterStruct is for registers that store numeric (integer or float) values.
type NumberRegisterStruct struct {
	RegisterStruct
	signed bool    // when true, the number is signed
	factor int     // the factor to multiply the number with (e.g. 10 when the number is in 0.1V resolution)
	offset float64 // the offset to add to the number
	unit   string  // the unit of the number (e.g. "V" for volt)
}

// TextRegisterStruct is for registers that store text values.
type TextRegisterStruct struct {
	RegisterStruct
}

// EnumRegisterStruct is for registers that store enum values.
type EnumRegisterStruct struct {
	RegisterStruct
	bit int // when positive, the specified bit is used as a boolean 0/1
	// todo: remove this bit stuff; add a new multi enum register type instead
	enum map[int]string // a map uf enum index to enum string value
}

// NewNumberRegisterStruct is a shortcut to create a RegisterStruct and embed it into a NumberRegisterStruct.
func NewNumberRegisterStruct(
	category, name, description string,
	sort int,
	address uint16,
	static bool,
	writable bool,
	signed bool,
	factor int,
	offset float64,
	unit string,
) NumberRegisterStruct {
	return NumberRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
			writable,
		},
		signed,
		factor,
		offset,
		unit,
	}
}

// NewTextRegisterStruct is a shortcut to create a RegisterStruct and embed it into a TextRegisterStruct.
func NewTextRegisterStruct(
	category, name, description string,
	sort int,
	address uint16,
	static bool,
	writable bool,
) TextRegisterStruct {
	return TextRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
			writable,
		},
	}
}

// NewEnumRegisterStruct is a shortcut to create a RegisterStruct and embed it into a EnumRegisterStruct.
// Also, different key types for the enum map are supported as long as they are integers.
func NewEnumRegisterStruct[K constraints.Integer, M map[K]string](
	category, name, description string,
	sort int,
	address uint16,
	bit int,
	static bool,
	writable bool,
	enum M,
) EnumRegisterStruct {
	return EnumRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
			writable,
		},
		bit,
		func() map[int]string {
			intEnum := make(map[int]string)
			for k, v := range enum {
				intEnum[int(k)] = v
			}
			return intEnum
		}(),
	}
}

type Register interface {
	Category() string
	Name() string
	Description() string
	Sort() int
	Address() uint16
	Static() bool
	Writable() bool
}

type NumberRegister interface {
	Register
	Unit() string
	Signed() bool
	Factor() int
	Offset() float64
}

type TextRegister interface {
	Register
}

type EnumRegister interface {
	Register
	Bit() int
	Enum() map[int]string
}

// Category return the category of the register (e.g. "Essential").
func (r RegisterStruct) Category() string {
	return r.category
}

// Name return the technical name of the register (e.g. "PanelVoltage").
func (r RegisterStruct) Name() string {
	return r.name
}

// Description return the description of the register (e.g. "Panel Voltage").
func (r RegisterStruct) Description() string {
	return r.description
}

// Sort return the sort order of the registers.
func (r RegisterStruct) Sort() int {
	return r.sort
}

// Address return the memory address used to fetch register from the device.
func (r RegisterStruct) Address() uint16 {
	return r.address
}

// Static return true when the register is static and not updated by the device.
func (r RegisterStruct) Static() bool {
	return r.static
}

// Writable return true when the register can be written to.
func (r RegisterStruct) Writable() bool {
	return r.writable
}

// Signed return true when the number is signed.
func (r NumberRegisterStruct) Signed() bool {
	return r.signed
}

// Factor return the factor to multiply the number with (e.g. 10 when the number is in 0.1V resolution).
func (r NumberRegisterStruct) Factor() int {
	return r.factor
}

// Offset return the offset to add to the number.
func (r NumberRegisterStruct) Offset() float64 {
	return r.offset
}

// Unit return the unit of the number (e.g. "V" for volt).
func (r NumberRegisterStruct) Unit() string {
	return r.unit
}

// Bit return the bit to use as a boolean 0/1.
func (r EnumRegisterStruct) Bit() int {
	return r.bit
}

// Enum return the enum map of enum index to enum string value.
func (r EnumRegisterStruct) Enum() map[int]string {
	return r.enum
}
