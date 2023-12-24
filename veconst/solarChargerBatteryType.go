package veconst

type SolarChargerBatteryType uint8

const (
	SolarChargerBatteryTypeGelLongLife           SolarChargerBatteryType = 1
	SolarChargerBatteryTypeGelDeepDischargeMode1 SolarChargerBatteryType = 2
	SolarChargerBatteryTypeGelDeepDischargeMode2 SolarChargerBatteryType = 3
	SolarChargerBatteryTypeAGMDeepDischargeMode3 SolarChargerBatteryType = 4
	SolarChargerBatteryTypeTubularMode1          SolarChargerBatteryType = 5
	SolarChargerBatteryTypeTubularMode2          SolarChargerBatteryType = 6
	SolarChargerBatteryTypeTubularMode3          SolarChargerBatteryType = 7
	SolarChargerBatteryTypeLiFEPO4               SolarChargerBatteryType = 8
	SolarChargerBatteryTypeUserDefined           SolarChargerBatteryType = 255
)

var solarChargerBatteryTypeMap = map[SolarChargerBatteryType]string{
	SolarChargerBatteryTypeGelLongLife:           "Gel Victron Long Life (14.1V)",
	SolarChargerBatteryTypeGelDeepDischargeMode1: "Gel Victron Deep discharge (14.3V)",
	SolarChargerBatteryTypeGelDeepDischargeMode2: "Gel Victron Deep discharge (14.4V)",
	SolarChargerBatteryTypeAGMDeepDischargeMode3: "AGM Victron Deep discharge (14.7V)",
	SolarChargerBatteryTypeTubularMode1:          "Tubular plate cyclic mode 1 (14.9V)",
	SolarChargerBatteryTypeTubularMode2:          "Tubular plate cyclic mode 2 (15.1V)",
	SolarChargerBatteryTypeTubularMode3:          "Tubular plate cyclic mode 3 (15.3V)",
	SolarChargerBatteryTypeLiFEPO4:               "LiFEPO4 (14.2V)",
	SolarChargerBatteryTypeUserDefined:           "User defined",
}

// GetSolarChargerBatteryTypeStringMap returns a map of SolarChargerBatteryType values to their string representation.
func GetSolarChargerBatteryTypeStringMap() map[SolarChargerBatteryType]string {
	ret := make(map[SolarChargerBatteryType]string, len(solarChargerBatteryTypeMap))
	for k, v := range solarChargerBatteryTypeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerBatteryType exists.
func (s SolarChargerBatteryType) Exists() bool {
	_, ok := solarChargerBatteryTypeMap[s]
	return ok
}

// String returns the string representation of a SolarChargerBatteryType.
func (s SolarChargerBatteryType) String() string {
	if v, ok := solarChargerBatteryTypeMap[s]; ok {
		return v
	}
	return ""
}
