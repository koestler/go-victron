package vedirectapi

import (
	"cmp"
	"fmt"
	"github.com/koestler/go-victron/veconst"
	"github.com/koestler/go-victron/veregister"
	"slices"
	"sort"
	"strings"
)

// RegisterValue is a common interface for register values, including number, text, enum, and field-list values.
type RegisterValue interface {
	veregister.Register
	GenericValue() interface{}
	String() string
}

// NumberRegisterValue is a container for Number Registers and their values.
type NumberRegisterValue struct {
	veregister.NumberRegisterStruct
	value float64
}

// TextRegisterValue is a container for Text Registers and their values.
type TextRegisterValue struct {
	veregister.TextRegisterStruct
	value string
}

// EnumRegisterValue is a container for Enum Registers and their values.
type EnumRegisterValue struct {
	veregister.EnumRegisterStruct
	value veconst.Enum
}

// FieldListValue is a container for FieldList Registers and their values.
type FieldListValue struct {
	veregister.FieldListRegisterStruct
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
	set := make([]veconst.Field, 0, len(fields))
	for f, isSet := range fields {
		if isSet {
			set = append(set, f)
		}
	}
	sort.Slice(set, func(i, j int) bool { return set[i].Idx() < set[j].Idx() })

	strs := make([]string, len(set))
	for i, f := range set {
		strs[i] = f.String()
	}
	return strings.Join(strs, ", ")
}

func (v FieldListValue) GenericValue() interface{} {
	return v.value
}

func (v FieldListValue) String() string {
	return fmt.Sprintf("%s=%s", v.Name(), v.CommaString())
}

// RegisterValues contains the values of number, text, enum, and field-list registers.
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

	slices.SortFunc(list, func(a, b RegisterValue) int {
		return cmp.Or(
			cmp.Compare(a.Sort(), b.Sort()),
			strings.Compare(a.Name(), b.Name()),
		)
	})

	return list
}
