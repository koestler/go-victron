package vebleapi

import (
	"fmt"
	"github.com/koestler/go-victron/veconst"
	"sort"
	"strings"
)

// RegisterStruct is the base struct for all register types.
type RegisterStruct struct {
	category    string // the category of the register (e.g. "Essential")
	name        string // the technical name of the register (e.g. "PanelVoltage")
	description string // the description of the register (e.g. "Panel Voltage")
	sort        int    // the sort order of the registers
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

// NumberRegisterValue is a container for Number Registers and their values.
type NumberRegisterValue struct {
	RegisterStruct
	value float64
}

// TextRegisterValue is a container for Text Registers and their values.
type TextRegisterValue struct {
	RegisterStruct
	value string
}

// EnumRegisterValue is a container for Enum Registers and their values.
type EnumRegisterValue struct {
	RegisterStruct
	value veconst.Enum
}

// FieldListValue is a container for FieldList Registers and their values.
type FieldListValue struct {
	RegisterStruct
	value veconst.FieldList
}

func (v NumberRegisterValue) Value() float64 {
	return v.value
}

func (v NumberRegisterValue) GenericValue() interface{} {
	return v.value
}

func (v NumberRegisterValue) String() string {
	return fmt.Sprintf("%s=%f%s", v.Name(), v.Value(), v.Unit())
}

func (v TextRegisterValue) Value() string {
	return v.value
}

func (v TextRegisterValue) GenericValue() interface{} {
	return v.value
}

func (v TextRegisterValue) String() string {
	return fmt.Sprintf("%s=%s", v.Name(), v.Value())
}

func (v EnumRegisterValue) Value() veconst.Enum {
	return v.value
}

func (v EnumRegisterValue) GenericValue() interface{} {
	return v.value
}

func (v EnumRegisterValue) String() string {
	return fmt.Sprintf("%s=%d:%s", v.Name(), v.value.Idx(), v.value.String())
}

func (v FieldListValue) Value() veconst.FieldList {
	return v.value
}

func (v FieldListValue) CommaString() string {
	fields := v.value.Fields()
	strs := make([]string, 0, len(fields))
	for f, set := range fields {
		if !set {
			continue
		}
		strs = append(strs, f.String())
	}
	return strings.Join(strs, ", ")
}

func (v FieldListValue) GenericValue() interface{} {
	return v.value
}

func (v FieldListValue) String() string {
	return fmt.Sprintf("%s=%s", v.Name(), v.CommaString())
}

// RegisterValues is a container for Number, Text and Enum Registers and their values.
type RegisterValues struct {
	NumberValues    map[string]NumberRegisterValue
	TextValues      map[string]TextRegisterValue
	EnumValues      map[string]EnumRegisterValue
	FieldListValues map[string]FieldListValue
}

// GetList returns a sorted list of all RegisterValues.
func (rv RegisterValues) GetList() []RegisterValue {
	list := make([]RegisterValue, 0, len(rv.NumberValues)+len(rv.TextValues)+len(rv.EnumValues)+len(rv.FieldListValues))

	for _, v := range rv.NumberValues {
		list = append(list, v)
	}
	for _, v := range rv.TextValues {
		list = append(list, v)
	}
	for _, v := range rv.EnumValues {
		list = append(list, v)
	}
	for _, v := range rv.FieldListValues {
		list = append(list, v)
	}

	sort.SliceStable(list, func(i, j int) bool { return list[i].Sort() < list[j].Sort() })

	return list
}
