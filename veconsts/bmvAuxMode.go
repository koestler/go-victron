package veconsts

type AuxMode uint8

const (
	AuxModeVoltage         AuxMode = 0
	AuxModeMidpointVoltage AuxMode = 1
	AuxModeTemperature     AuxMode = 2
	AuxModeDisabled        AuxMode = 3
)

func AuxModeMap() map[AuxMode]string {
	return map[AuxMode]string{
		AuxModeVoltage:         "Aux voltage",
		AuxModeMidpointVoltage: "Mid-point voltage",
		AuxModeTemperature:     "Temperature",
		AuxModeDisabled:        "Disabled",
	}
}

func (s AuxMode) String() string {
	m := AuxModeMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
