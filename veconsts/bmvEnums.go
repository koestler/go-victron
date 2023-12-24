package veconsts

type BmvAuxMode uint8

const (
	BmvAuxModeStarterVoltage  BmvAuxMode = 0
	BmvAuxModeMidpointVoltage BmvAuxMode = 1
	BmvAuxModeTemperature     BmvAuxMode = 2
	BmvAuxModeDisabled        BmvAuxMode = 3
)

func GetBmvAuxModeMap() map[BmvAuxMode]string {
	return map[BmvAuxMode]string{
		BmvAuxModeStarterVoltage:  "Starter voltage",
		BmvAuxModeMidpointVoltage: "Mid-point voltage",
		BmvAuxModeTemperature:     "Temperature",
		BmvAuxModeDisabled:        "Disabled",
	}
}

func (s BmvAuxMode) String() string {
	m := GetBmvAuxModeMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
