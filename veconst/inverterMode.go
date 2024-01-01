package veconst

type InverterMode uint8

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

func NewInverterModeEnum(v int) (Enum, error) {
	s := InverterMode(v)
	if _, ok := inverterModeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func InverterModeMap() map[int]string {
	ret := make(map[int]string, len(inverterModeMap))
	for k, v := range inverterModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s InverterMode) Idx() int {
	return int(s)
}

func (s InverterMode) Value() string {
	if v, ok := inverterModeMap[s]; ok {
		return v
	}
	return ""
}
