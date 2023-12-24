package veconsts

type AuxMode uint8

const (
	AuxModeVoltage         AuxMode = 0
	AuxModeMidpointVoltage AuxMode = 1
	AuxModeTemperature     AuxMode = 2
	AuxModeDisabled        AuxMode = 3
)

var auxModeMap = map[AuxMode]string{
	AuxModeVoltage:         "Aux voltage",
	AuxModeMidpointVoltage: "Mid-point voltage",
	AuxModeTemperature:     "Temperature",
	AuxModeDisabled:        "Disabled",
}

// AuxModeStringMap returns a map of AuxMode values to their string representation.
func AuxModeStringMap() map[AuxMode]string {
	ret := make(map[AuxMode]string, len(auxModeMap))
	for k, v := range auxModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the AuxMode exists.
func (s AuxMode) Exists() bool {
	_, ok := auxModeMap[s]
	return ok
}

// String returns the string representation of a AuxMode.
func (s AuxMode) String() string {
	if v, ok := auxModeMap[s]; ok {
		return v
	}
	return ""
}
