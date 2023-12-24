package veconsts

type BmvAuxMode uint8

const (
	BmvAuxModeStarterVoltage  BmvAuxMode = 0
	BmvAuxModeMidpointVoltage BmvAuxMode = 1
	BmvAuxModeTemperature     BmvAuxMode = 2
	BmvAuxModeDisabled        BmvAuxMode = 3
)

var bmvAuxModeMap = map[BmvAuxMode]string{
	BmvAuxModeStarterVoltage:  "Starter voltage",
	BmvAuxModeMidpointVoltage: "Mid-point voltage",
	BmvAuxModeTemperature:     "Temperature",
	BmvAuxModeDisabled:        "Disabled",
}

// GetBmvAuxModeStringMap returns a map of BmvAuxMode values to their string representation.
func GetBmvAuxModeStringMap() map[BmvAuxMode]string {
	ret := make(map[BmvAuxMode]string, len(bmvAuxModeMap))
	for k, v := range bmvAuxModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BmvAuxMode exists.
func (s BmvAuxMode) Exists() bool {
	_, ok := bmvAuxModeMap[s]
	return ok
}

// String returns the string representation of a BmvAuxMode.
func (s BmvAuxMode) String() string {
	if v, ok := bmvAuxModeMap[s]; ok {
		return v
	}
	return ""
}
