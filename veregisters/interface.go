package veregisters

type RegisterType int

// Register types
const (
	NumberRegister RegisterType = iota
	TextRegister
	EnumRegister
)

// Register is the interface implemented by all register types.
// For parameters which do not make any sense (e.g. Unit for TextRegister) will return the null type.
type Register interface {
	Category() string
	Name() string
	Description() string
	Sort() int
	Address() uint16
	Static() bool
	RegisterType() RegisterType
	Enum() map[int]string
	Unit() string
}

// Category return the category of the register (e.g. "Essential").
func (r RegisterStruct) Category() string {
	return r.category
}

// Name return the technical name of the register (e.g. "PanelVoltage").
func (r RegisterStruct) Name() string {
	return r.name
}

// Description return the description of the register (e.g. "Panel Voltage").
func (r RegisterStruct) Description() string {
	return r.description
}

// Sort return the sort order of the registers.
func (r RegisterStruct) Sort() int {
	return r.sort
}

// Address return the memory address used to fetch register from the device.
func (r RegisterStruct) Address() uint16 {
	return r.address
}

// Static return true when the register is static and not updated by the device.
func (r RegisterStruct) Static() bool {
	return r.static
}

// RegisterType return the data type of the register.
func (r NumberRegisterStruct) RegisterType() RegisterType {
	return NumberRegister
}

// Enum return always nil for NumberRegister and is only here so that all register types implement the Register interface.
func (r NumberRegisterStruct) Enum() map[int]string {
	return nil
}

// Unit return the unit of the number (e.g. "V" for volt).
func (r NumberRegisterStruct) Unit() string {
	return r.unit
}

// Signed return true when the number is signed.
func (r NumberRegisterStruct) Signed() bool {
	return r.signed
}

// Factor return the factor to multiply the number with (e.g. 10 when the number is in 0.1V resolution).
func (r NumberRegisterStruct) Factor() int {
	return r.factor
}

// Offset return the offset to add to the number.
func (r NumberRegisterStruct) Offset() float64 {
	return r.offset
}

// RegisterType return the data type of the register.
func (r TextRegisterStruct) RegisterType() RegisterType {
	return TextRegister
}

// Enum return always nil for TextRegister and is only here so that all register types implement the Register interface.
func (r TextRegisterStruct) Enum() map[int]string {
	return nil
}

// Unit return always "" for TextRegister and is only here so that all register types implement the Register interface.
func (r TextRegisterStruct) Unit() string {
	return ""
}

// RegisterType return the data type of the register.
func (r EnumRegisterStruct) RegisterType() RegisterType {
	return EnumRegister
}

// Enum return the enum map of enum index to enum string value.
func (r EnumRegisterStruct) Enum() map[int]string {
	return r.enum
}

// Unit return always "" for EnumRegister and is only here so that all register types implement the Register interface.
func (r EnumRegisterStruct) Unit() string {
	return ""
}

// Bit return the bit to use as a boolean 0/1.
func (r EnumRegisterStruct) Bit() int {
	return r.bit
}
