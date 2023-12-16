package veregisters

func (r RegisterStruct) Category() string {
	return r.category
}

func (r RegisterStruct) Name() string {
	return r.name
}

func (r RegisterStruct) Description() string {
	return r.description
}

func (r RegisterStruct) Sort() int {
	return r.sort
}

func (r RegisterStruct) Address() uint16 {
	return r.address
}

func (r RegisterStruct) Static() bool {
	return r.static
}

func (r NumberRegisterStruct) Unit() string {
	return r.unit
}

func (r EnumRegisterStruct) Enum() map[int]string {
	return r.enum
}
