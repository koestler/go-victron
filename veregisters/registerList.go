package veregisters

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
