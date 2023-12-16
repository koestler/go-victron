package veregisters

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

func (rl *RegisterList) appendNumberRegisterStruct(r ...NumberRegisterStruct) {
	rl.NumberRegisters = append(rl.NumberRegisters, r...)
}

func (rl *RegisterList) appendTextRegisterStruct(r ...TextRegisterStruct) {
	rl.TextRegisters = append(rl.TextRegisters, r...)
}

func (rl *RegisterList) appendEnumRegisterStruct(r ...EnumRegisterStruct) {
	rl.EnumRegisters = append(rl.EnumRegisters, r...)
}

func (rl *RegisterList) filterByName(exclude ...string) {
	rl.NumberRegisters = filterRegistersByName(rl.NumberRegisters, exclude...)
	rl.TextRegisters = filterRegistersByName(rl.TextRegisters, exclude...)
	rl.EnumRegisters = filterRegistersByName(rl.EnumRegisters, exclude...)
}
