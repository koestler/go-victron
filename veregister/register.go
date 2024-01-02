// Package veregister contains the register definitions for the VE.Direct protocol.
package veregister

import (
	"github.com/koestler/go-victron/veconst"
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
	factory veconst.EnumFactory
}

// FieldListRegisterStruct is for registers that store a list of fields.
type FieldListRegisterStruct struct {
	RegisterStruct
	factory veconst.FieldListFactory
}

// newNumberRegisterStruct is a shortcut to create a RegisterStruct and embed it into a NumberRegisterStruct.
func newNumberRegisterStruct(
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

// newTextRegisterStruct is a shortcut to create a RegisterStruct and embed it into a TextRegisterStruct.
func newTextRegisterStruct(
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

// newEnumRegisterStruct is a shortcut to create a RegisterStruct and embed it into a EnumRegisterStruct.
func newEnumRegisterStruct(
	category, name, description string,
	sort int,
	address uint16,
	static bool,
	writable bool,
	factory veconst.EnumFactory,
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
		factory,
	}
}

// newFieldListRegisterStruct is a shortcut to create a RegisterStruct and embed it into a FieldListRegisterStruct.
func newFieldListRegisterStruct(
	category, name, description string,
	sort int,
	address uint16,
	static bool,
	writable bool,
	factory veconst.FieldListFactory,
) FieldListRegisterStruct {
	return FieldListRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
			writable,
		},
		factory,
	}
}

type Type uint

const (
	Undefined Type = iota
	Number
	Text
	Enum
	FieldList
)

type Register interface {
	Type() Type
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

type EnumRegister interface {
	Register
	Enum() map[int]string
}

// Type returns the type of the register.
func (r RegisterStruct) Type() Type {
	return Undefined
}

// Category returns the category of the register (e.g. "Essential").
func (r RegisterStruct) Category() string {
	return r.category
}

// Name returns the technical name of the register (e.g. "PanelVoltage").
func (r RegisterStruct) Name() string {
	return r.name
}

// Description returns the description of the register (e.g. "Panel Voltage").
func (r RegisterStruct) Description() string {
	return r.description
}

// Sort returns the sort order of the registers.
func (r RegisterStruct) Sort() int {
	return r.sort
}

// Address returns the memory address used to fetch register from the device.
func (r RegisterStruct) Address() uint16 {
	return r.address
}

// Static returns true when the register is static and not updated by the device.
func (r RegisterStruct) Static() bool {
	return r.static
}

// Writable returns true when the register can be written to.
func (r RegisterStruct) Writable() bool {
	return r.writable
}

// Type returns the type of the register.
func (r NumberRegisterStruct) Type() Type {
	return Number
}

// Signed returns true when the number is signed.
func (r NumberRegisterStruct) Signed() bool {
	return r.signed
}

// Factor returns the factor to multiply the number with (e.g. 10 when the number is in 0.1V resolution).
func (r NumberRegisterStruct) Factor() int {
	return r.factor
}

// Offset returns the offset to add to the number.
func (r NumberRegisterStruct) Offset() float64 {
	return r.offset
}

// Unit returns the unit of the number (e.g. "V" for volt).
func (r NumberRegisterStruct) Unit() string {
	return r.unit
}

// Type returns the type of the register.
func (r TextRegisterStruct) Type() Type {
	return Text
}

// Type returns the type of the register.
func (r EnumRegisterStruct) Type() Type {
	return Enum
}

// Factory returns the enum factory.
func (r EnumRegisterStruct) Factory() veconst.EnumFactory {
	return r.factory
}

// Type returns the type of the register.
func (r FieldListRegisterStruct) Type() Type {
	return FieldList
}

// Factory returns the field list factory.
func (r FieldListRegisterStruct) Factory() veconst.FieldListFactory {
	return r.factory
}
