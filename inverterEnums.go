package victronDefinitions

type InverterState uint8

const (
	InverterStateOff       InverterState = 0
	InverterStateLowPower  InverterState = 1
	InverterStateFault     InverterState = 2
	InverterStateInverting InverterState = 9
)

func GetInverterStateMap() map[InverterState]string {
	return map[InverterState]string{
		InverterStateOff:       "Off",
		InverterStateLowPower:  "Low Power",
		InverterStateFault:     "Fault",
		InverterStateInverting: "Inverting",
	}
}

func (s InverterState) String() string {
	m := GetInverterStateMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}

type InverterMode uint8

const (
	InverterModeInverterOn InverterMode = 2
	InverterModeDeviceOn   InverterMode = 3
	InverterModeDeviceOff  InverterMode = 4
	InverterModeEcoMode    InverterMode = 5
	InverterModeHibernate  InverterMode = 0xFD
)

func GetInverterModeMap() map[InverterMode]string {
	return map[InverterMode]string{
		InverterModeInverterOn: "Inverter On",
		InverterModeDeviceOn:   "Device On",
		InverterModeDeviceOff:  "Device Off",
		InverterModeEcoMode:    "Eco mode",
		InverterModeHibernate:  "Hibernate",
	}
}

func (s InverterMode) String() string {
	m := GetInverterModeMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
