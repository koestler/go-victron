package veconst

type SolarChargerBatteryVoltage uint8

const (
	SolarChargerBatteryVoltageAutoDetect SolarChargerBatteryVoltage = 0
	SolarChargerBatteryVoltage12V        SolarChargerBatteryVoltage = 12
	SolarChargerBatteryVoltage24V        SolarChargerBatteryVoltage = 24
	SolarChargerBatteryVoltage36V        SolarChargerBatteryVoltage = 36
	SolarChargerBatteryVoltage48V        SolarChargerBatteryVoltage = 48
)

var solarChargerBatteryVoltageMap = map[SolarChargerBatteryVoltage]string{
	SolarChargerBatteryVoltageAutoDetect: "Auto detection at startup",
	SolarChargerBatteryVoltage12V:        "12V battery",
	SolarChargerBatteryVoltage24V:        "24V battery",
	SolarChargerBatteryVoltage36V:        "36V battery",
	SolarChargerBatteryVoltage48V:        "48V battery",
}

// SolarChargerBatteryVoltageStringMap returns a map of SolarChargerBatteryVoltage values to their string representation.
func SolarChargerBatteryVoltageStringMap() map[SolarChargerBatteryVoltage]string {
	ret := make(map[SolarChargerBatteryVoltage]string, len(solarChargerBatteryVoltageMap))
	for k, v := range solarChargerBatteryVoltageMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerBatteryVoltage exists.
func (s SolarChargerBatteryVoltage) Exists() bool {
	_, ok := solarChargerBatteryVoltageMap[s]
	return ok
}

// String returns the string representation of a SolarChargerBatteryVoltage.
func (s SolarChargerBatteryVoltage) String() string {
	if v, ok := solarChargerBatteryVoltageMap[s]; ok {
		return v
	}
	return ""
}
