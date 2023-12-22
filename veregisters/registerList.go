package veregisters

import "sort"

// RegisterList is a container holding {Number,Text,Enum}RegisterStructs.
type RegisterList struct {
	NumberRegisters []NumberRegisterStruct
	TextRegisters   []TextRegisterStruct
	EnumRegisters   []EnumRegisterStruct
}

func NewRegisterList() RegisterList {
	return RegisterList{
		NumberRegisters: make([]NumberRegisterStruct, 0),
		TextRegisters:   make([]TextRegisterStruct, 0),
		EnumRegisters:   make([]EnumRegisterStruct, 0),
	}
}

func (rl *RegisterList) AppendNumberRegisterStruct(r ...NumberRegisterStruct) {
	rl.NumberRegisters = append(rl.NumberRegisters, r...)
}

func (rl *RegisterList) AppendTextRegisterStruct(r ...TextRegisterStruct) {
	rl.TextRegisters = append(rl.TextRegisters, r...)
}

func (rl *RegisterList) AppendEnumRegisterStruct(r ...EnumRegisterStruct) {
	rl.EnumRegisters = append(rl.EnumRegisters, r...)
}

// FilterRegister removes all registers from the list for which the filter function return false.
// This is useful since the user does not need to care about the different register types.
func (rl *RegisterList) FilterRegister(f func(r Register) bool) {
	rl.NumberRegisters = filterRegisters(rl.NumberRegisters, f)
	rl.TextRegisters = filterRegisters(rl.TextRegisters, f)
	rl.EnumRegisters = filterRegisters(rl.EnumRegisters, f)
}

// GetRegisters returns all types of registers sorted by their sort value in a common interface slice.
func (rl *RegisterList) GetRegisters() []Register {
	oup := make([]Register, 0, len(rl.NumberRegisters)+len(rl.TextRegisters)+len(rl.EnumRegisters))

	for _, r := range rl.NumberRegisters {
		oup = append(oup, r)
	}
	for _, r := range rl.TextRegisters {
		oup = append(oup, r)
	}
	for _, r := range rl.EnumRegisters {
		oup = append(oup, r)
	}

	sort.SliceStable(oup, func(i, j int) bool { return oup[i].Sort() < oup[j].Sort() })

	return oup
}
