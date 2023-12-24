package veconst

type SolarChargerState uint8

const (
	SolarChargerStateNotCharging        SolarChargerState = 0
	SolarChargerStateFault              SolarChargerState = 2
	SolarChargerStateBulkCharging       SolarChargerState = 3
	SolarChargerStateAbsorptionCharging SolarChargerState = 4
	SolarChargerStateFloatCharging      SolarChargerState = 5
	SolarChargerStateManualEqualise     SolarChargerState = 7
	SolarChargerStateWakeUp             SolarChargerState = 245
	SolarChargerStateAutoEqualise       SolarChargerState = 247
	SolarChargerStateExternalControl    SolarChargerState = 252
	SolarChargerStateUnavailable        SolarChargerState = 255
)

var solarChargerStateMap = map[SolarChargerState]string{
	SolarChargerStateNotCharging:        "Not charging",
	SolarChargerStateFault:              "Fault",
	SolarChargerStateBulkCharging:       "Bulk Charging",
	SolarChargerStateAbsorptionCharging: "Absorption Charging",
	SolarChargerStateFloatCharging:      "Float Charging",
	SolarChargerStateManualEqualise:     "Manual Equalise",
	SolarChargerStateWakeUp:             "Wake-Up",
	SolarChargerStateAutoEqualise:       "Auto Equalise",
	SolarChargerStateExternalControl:    "External Control",
	SolarChargerStateUnavailable:        "Unavailable",
}

// SolarChargerStateStringMap returns a map of SolarChargerState values to their string representation.
func SolarChargerStateStringMap() map[SolarChargerState]string {
	ret := make(map[SolarChargerState]string, len(solarChargerStateMap))
	for k, v := range solarChargerStateMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerState exists.
func (s SolarChargerState) Exists() bool {
	_, ok := solarChargerStateMap[s]
	return ok
}

// String returns the string representation of a SolarChargerState.
func (s SolarChargerState) String() string {
	if v, ok := solarChargerStateMap[s]; ok {
		return v
	}
	return ""
}
