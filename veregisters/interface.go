package veregisters

type Register interface {
	GetCategory() string
	GetName() string
	GetDescription() string
	GetSort() int
	GetAddress() uint16
	GetStatic() bool
}

func (r RegisterStruct) GetCategory() string {
	return r.Category
}

func (r RegisterStruct) GetName() string {
	return r.Name
}

func (r RegisterStruct) GetDescription() string {
	return r.Description
}

func (r RegisterStruct) GetSort() int {
	return r.Sort
}

func (r RegisterStruct) GetAddress() uint16 {
	return r.Address
}

func (r RegisterStruct) GetStatic() bool {
	return r.Static
}
