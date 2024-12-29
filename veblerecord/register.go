package veblerecord

import (
	"fmt"
	"github.com/koestler/go-victron/veconst"
	"sort"
	"strings"
)

type Registers interface {
	NumberRegisters() []NumberRegister
	EnumRegisters() []EnumRegister
	FieldListRegisters() []FieldListRegister
}

type Register struct {
	name        string
	description string
}

func (r Register) Name() string {
	return r.name
}

func (r Register) Description() string {
	return r.description
}

type NumberRegister struct {
	Register
	value float64
	unit  string
}

func (r NumberRegister) GenericValue() interface{} {
	return r.value
}

func (r NumberRegister) Value() float64 {
	return r.value
}

func (r NumberRegister) Unit() string {
	return r.unit
}

func (r NumberRegister) String() string {
	return fmt.Sprintf("%s=%f %s", r.name, r.value, r.unit)
}

type EnumRegister struct {
	Register
	value veconst.Enum
}

func (r EnumRegister) GenericValue() interface{} {
	return r.value
}

func (r EnumRegister) Value() veconst.Enum {
	return r.value
}

func (r EnumRegister) Unit() string {
	return ""
}

func (r EnumRegister) String() string {
	return fmt.Sprintf("%s=%s", r.name, r.value)
}

type FieldListRegister struct {
	Register
	value veconst.FieldList
}

func (r FieldListRegister) GenericValue() interface{} {
	return r.value
}

func (r FieldListRegister) Value() veconst.FieldList {
	return r.value
}

func (r FieldListRegister) Unit() string {
	return ""
}

func (r FieldListRegister) String() string {
	trueFields := make([]string, 0)
	for f, v := range r.value.Fields() {
		if !v {
			continue
		}
		trueFields = append(trueFields, f.String())
	}
	sort.Strings(trueFields)

	return fmt.Sprintf("%s=%s", r.name, strings.Join(trueFields, ","))
}
