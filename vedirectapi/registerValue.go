package vedirectapi

import (
	"fmt"
	"github.com/koestler/go-victron/veregisters"
)

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

func (v NumberRegisterValue) String() string {
	return fmt.Sprintf("%s=%f%s", v.Register.Name, v.Value, v.Register.Unit)
}

func (v TextRegisterValue) String() string {
	return fmt.Sprintf("%s=%s", v.Register.Name, v.Value)
}

func (v EnumRegisterValue) String() string {
	return fmt.Sprintf("%s=%d:%s", v.Register.Name, v.EnumIdx, v.EnumValue)
}
