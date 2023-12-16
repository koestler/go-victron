package veregisters

import (
	"golang.org/x/exp/constraints"
)

type RegisterStruct struct {
	Category    string
	Name        string
	Description string
	Sort        int

	Address uint16
	Static  bool
}

type NumberRegisterStruct struct {
	RegisterStruct
	Unit string

	Signed bool
	Factor int
	Offset float64
}

type TextRegisterStruct struct {
	RegisterStruct
}

type EnumRegisterStruct struct {
	RegisterStruct
	Bit  int // when positive, only the given bit is used as 0/1
	Enum map[int]string
}

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
