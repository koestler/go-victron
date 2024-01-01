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

type SolarChargerBatteryTypeFactoryType struct{}

var SolarChargerBatteryTypeFactory SolarChargerBatteryTypeFactoryType

func (f SolarChargerBatteryTypeFactoryType) NewEnum(v int) (Enum, error) {
	s := SolarChargerBatteryType(v)
	if _, ok := solarChargerBatteryTypeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f SolarChargerBatteryTypeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarChargerBatteryTypeMap))
	for k, v := range solarChargerBatteryTypeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerBatteryType) Idx() int {
	return int(s)
}

func (s SolarChargerBatteryType) String() string {
	if v, ok := solarChargerBatteryTypeMap[s]; ok {
		return v
	}
	return ""
}

func (s SolarChargerBatteryType) Exists() bool {
	_, ok := solarChargerBatteryTypeMap[s]
	return ok
}
