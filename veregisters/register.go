package veregisters

import (
	"github.com/koestler/go-victron/dataflow"
	"golang.org/x/exp/constraints"
)

type VictronRegister struct {
	dataflow.RegisterStruct
	address uint16
	bit     int // only used for enums, when positive, only the given bit is used as 0/1
	static  bool

	// only relevant for number registers
	signed bool
	factor int
	offset float64
}

func MergeRegisters(maps ...[]VictronRegister) (output []VictronRegister) {
	if len(maps) == 0 {
		return
	}
	output = maps[0]
	for i := 1; i < len(maps); i++ {
		output = append(output, maps[i]...)
	}

	return
}

func FilterRegisters(input []VictronRegister, registerFilter dataflow.RegisterFilterConf) (output []VictronRegister) {
	output = make([]VictronRegister, 0, len(input))
	f := dataflow.RegisterFilter(registerFilter)
	for _, r := range input {
		if f(r) {
			output = append(output, r)
		}
	}
	return
}

func FilterRegistersByName(input []VictronRegister, names ...string) (output []VictronRegister) {
	output = make([]VictronRegister, 0, len(input))
	for _, r := range input {
		if registerNameExcluded(names, r) {
			continue
		}
		output = append(output, r)
	}
	return
}

func registerNameExcluded(exclude []string, r dataflow.Register) bool {
	for _, e := range exclude {
		if e == r.Name() {
			return true
		}
	}
	return false
}

func NewTextRegisterStruct(
	category, name, description string,
	address uint16,
	static bool,
	sort int,
) VictronRegister {
	return VictronRegister{
		dataflow.NewRegisterStruct(
			category, name, description,
			dataflow.TextRegister,
			nil,
			"",
			sort,
			false,
		),
		address,
		-1,
		static,
		false, // unused
		1,     // unused
		0,     // unused
	}
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
) VictronRegister {
	return VictronRegister{
		dataflow.NewRegisterStruct(
			category, name, description,
			dataflow.NumberRegister,
			nil,
			unit,
			sort,
			false,
		),
		address,
		-1,
		static,
		signed,
		factor,
		offset,
	}
}

func NewEnumRegisterStruct[K constraints.Integer, M map[K]string](
	category, name, description string,
	address uint16,
	bit int,
	static bool,
	enum M,
	sort int,
) VictronRegister {
	intEnum := make(map[int]string)
	for k, v := range enum {
		intEnum[int(k)] = v
	}

	return VictronRegister{
		dataflow.NewRegisterStruct(
			category, name, description,
			dataflow.EnumRegister,
			intEnum,
			"",
			sort,
			false,
		),
		address,
		bit,
		static,
		false, // unused
		1,     // unused
		0,     // unused
	}
}