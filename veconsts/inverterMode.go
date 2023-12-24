package veconsts

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

// GetInverterModeMap returns a map of InverterMode values to their string representation.
func GetInverterModeMap() map[InverterMode]string {
	ret := make(map[InverterMode]string, len(inverterModeMap))
	for k, v := range inverterModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the InverterMode exists.
func (s InverterMode) Exists() bool {
	_, ok := inverterModeMap[s]
	return ok
}

// String returns the string representation of a InverterMode.
func (s InverterMode) String() string {
	if v, ok := inverterModeMap[s]; ok {
		return v
	}
	return ""
}
