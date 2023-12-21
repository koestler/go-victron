package vedirectapi

import (
	"fmt"
	"github.com/koestler/go-victron/veregisters"
	"sort"
)

// RegisterValue is a common interface for Number, Text and Enum Registers and their values.
type RegisterValue interface {
	veregisters.Register
	GenericValue() interface{}
	String() string
}

// NumberRegisterValue is a container for Number Registers and their values.
type NumberRegisterValue struct {
	veregisters.NumberRegisterStruct
	value float64
}

// TextRegisterValue is a container for Text Registers and their values.
type TextRegisterValue struct {
	veregisters.TextRegisterStruct
	value string
}

// EnumRegisterValue is a container for Enum Registers and their values.
type EnumRegisterValue struct {
	veregisters.EnumRegisterStruct
	enumIdx   int
	enumValue string
}

func (v NumberRegisterValue) Value() float64 {
	return v.value
}

func (v NumberRegisterValue) GenericValue() interface{} {
	return v.value
}

func (v TextRegisterValue) Value() string {
	return v.value
}

func (v TextRegisterValue) GenericValue() interface{} {
	return v.value

}

func (v EnumRegisterValue) Idx() int {
	return v.enumIdx
}

func (v EnumRegisterValue) Value() string {
	return v.enumValue
}

func (v EnumRegisterValue) GenericValue() interface{} {
	return v.enumValue
}

func (v NumberRegisterValue) String() string {
	return fmt.Sprintf("%s=%f%s", v.Name(), v.Value(), v.Unit())
}

func (v TextRegisterValue) String() string {
	return fmt.Sprintf("%s=%s", v.Name(), v.Value())
}

func (v EnumRegisterValue) String() string {
	return fmt.Sprintf("%s=%d:%s", v.Name(), v.Idx(), v.Value())
}

// RegisterValues is a container for Number, Text and Enum Registers and their values.
type RegisterValues struct {
	NumberValues map[string]NumberRegisterValue
	TextValues   map[string]TextRegisterValue
	EnumValues   map[string]EnumRegisterValue
}

// GetList returns a sorted list of all RegisterValues.
func (rv RegisterValues) GetList() []RegisterValue {
	list := make([]RegisterValue, 0, len(rv.NumberValues)+len(rv.TextValues)+len(rv.EnumValues))

	for _, v := range rv.NumberValues {
		list = append(list, v)
	}
	for _, v := range rv.TextValues {
		list = append(list, v)
	}
	for _, v := range rv.EnumValues {
		list = append(list, v)
	}

	sort.SliceStable(list, func(i, j int) bool { return list[i].Sort() < list[j].Sort() })

	return list
}
