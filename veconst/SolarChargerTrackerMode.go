package veconst

type SolarChargerTrackerMode uint8

const (
	SolarChargerTrackerModeOff                   SolarChargerTrackerMode = 0
	SolarChargerTrackerModeVoltageCurrentLimited SolarChargerTrackerMode = 1
	SolarChargerTrackerModeMPPTracker            SolarChargerTrackerMode = 2
)

var solarChargerTrackerModeMap = map[SolarChargerTrackerMode]string{
	SolarChargerTrackerModeOff:                   "Off",
	SolarChargerTrackerModeVoltageCurrentLimited: "Voltage/current limited",
	SolarChargerTrackerModeMPPTracker:            "MPP tracker",
}

// SolarChargerTrackerModeStringMap returns a map of SolarChargerTrackerMode values to their string representation.
func SolarChargerTrackerModeStringMap() map[SolarChargerTrackerMode]string {
	ret := make(map[SolarChargerTrackerMode]string, len(solarChargerTrackerModeMap))
	for k, v := range solarChargerTrackerModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerTrackerMode exists.
func (s SolarChargerTrackerMode) Exists() bool {
	_, ok := solarChargerTrackerModeMap[s]
	return ok
}

// String returns the string representation of a SolarChargerTrackerMode.
func (s SolarChargerTrackerMode) String() string {
	if v, ok := solarChargerTrackerModeMap[s]; ok {
		return v
	}
	return ""
}
