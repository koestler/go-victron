package veconst

type InverterMode uint8
type InverterModeFactoryType struct{}

const (
	InverterModeInverterOn InverterMode = 2
	InverterModeDeviceOn   InverterMode = 3
	InverterModeDeviceOff  InverterMode = 4
	InverterModeEcoMode    InverterMode = 5
	InverterModeHibernate  InverterMode = 0xFD
)

var inverterModeMap = map[InverterMode]string{
	InverterModeInverterOn: "Inverter On",
	InverterModeDeviceOn:   "Device On",
	InverterModeDeviceOff:  "Device Off",
	InverterModeEcoMode:    "Eco mode",
	InverterModeHibernate:  "Hibernate",
}
var InverterModeFactory InverterModeFactoryType

func (f InverterModeFactoryType) New(v int) (InverterMode, error) {
	s := InverterMode(v)
	if _, ok := inverterModeMap[s]; !ok {
		return InverterModeHibernate, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f InverterModeFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
}

func (f InverterModeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(inverterModeMap))
	for k, v := range inverterModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s InverterMode) Idx() int {
	return int(s)
}

func (s InverterMode) String() string {
	if v, ok := inverterModeMap[s]; ok {
		return v
	}
	return ""
}

func (s InverterMode) Exists() bool {
	_, ok := inverterModeMap[s]
	return ok
}
