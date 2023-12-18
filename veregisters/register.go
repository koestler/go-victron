// Package veregisters contains the register definitions for the VE.Direct protocol.
package veregisters

import (
	"golang.org/x/exp/constraints"
)

// RegisterStruct is the base struct for all register types.
type RegisterStruct struct {
	Category    string // the category of the register (e.g. "Essential")
	Name        string // the technical name of the register (e.g. "PanelVoltage")
	Description string // the description of the register (e.g. "Panel Voltage")
	Sort        int    // the sort order of the registers
	Address     uint16 // the memory address used to fetch register from the device
	Static      bool   // when true, the register is static and not updated by the device
}

// NumberRegisterStruct is for registers that store numeric (integer or float) values.
type NumberRegisterStruct struct {
	RegisterStruct
	Unit   string  // the unit of the number (e.g. "V" for volt)
	Signed bool    // when true, the number is signed
	Factor int     // the factor to multiply the number with (e.g. 10 when the number is in 0.1V resolution)
	Offset float64 // the offset to add to the number
}

// TextRegisterStruct is for registers that store text values.
type TextRegisterStruct struct {
	RegisterStruct
}

// EnumRegisterStruct is for registers that store enum values.
type EnumRegisterStruct struct {
	RegisterStruct
	Bit  int            // when positive, the specified bit is used as a boolean 0/1
	Enum map[int]string // a map uf enum index to enum string value
}

// NewNumberRegisterStruct is a shortcut to create a RegisterStruct and embed it into a NumberRegisterStruct.
func NewNumberRegisterStruct(
	category, name, description string,
	address uint16,
	static bool,
	signed bool,
	factor int,
	offset float64,
	unit string,
	sort int,
) NumberRegisterStruct {
	return NumberRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
		},
		unit,
		signed,
		factor,
		offset,
	}
}

// NewTextRegisterStruct is a shortcut to create a RegisterStruct and embed it into a TextRegisterStruct.
func NewTextRegisterStruct(
	category, name, description string,
	address uint16,
	static bool,
	sort int,
) TextRegisterStruct {
	return TextRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
		},
	}
}

// NewEnumRegisterStruct is a shortcut to create a RegisterStruct and embed it into a EnumRegisterStruct.
// Also, different key types for the enum map are supported as long as they are integers.
func NewEnumRegisterStruct[K constraints.Integer, M map[K]string](
	category, name, description string,
	address uint16,
	bit int,
	static bool,
	enum M,
	sort int,
) EnumRegisterStruct {
	return EnumRegisterStruct{
		RegisterStruct{
			category,
			name,
			description,
			sort,
			address,
			static,
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
