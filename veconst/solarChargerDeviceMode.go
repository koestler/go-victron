package veconst

type SolarChargerDeviceMode uint8

const (
	SolarChargerDeviceModeOff  SolarChargerDeviceMode = 0
	SolarChargerDeviceModeOff4 SolarChargerDeviceMode = 4
	SolarChargerDeviceModeOn   SolarChargerDeviceMode = 1
)

var solarChargerDeviceModeMap = map[SolarChargerDeviceMode]string{
	SolarChargerDeviceModeOff:  "Charger Off",
	SolarChargerDeviceModeOff4: "Charger Off",
	SolarChargerDeviceModeOn:   "Charger On",
}

// SolarChargerDeviceModeStringMap returns a map of SolarChargerDeviceMode to their string representation.
func SolarChargerDeviceModeStringMap() map[SolarChargerDeviceMode]string {
	ret := make(map[SolarChargerDeviceMode]string, len(solarChargerDeviceModeMap))
	for k, v := range solarChargerDeviceModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerDeviceMode exists.
func (s SolarChargerDeviceMode) Exists() bool {
	_, ok := solarChargerDeviceModeMap[s]
	return ok
}

// String returns the string representation of a SolarChargerDeviceMode.
func (s SolarChargerDeviceMode) String() string {
	if v, ok := solarChargerDeviceModeMap[s]; ok {
		return v
	}
	return ""
}
