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

func NewSolarChargerBatteryVoltageEnum(v int) (Enum, error) {
	s := SolarChargerBatteryVoltage(v)
	if _, ok := solarChargerBatteryVoltageMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func SolarChargerBatteryVoltageMap() map[int]string {
	ret := make(map[int]string, len(solarChargerBatteryVoltageMap))
	for k, v := range solarChargerBatteryVoltageMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerBatteryVoltage) Idx() int {
	return int(s)
}

func (s SolarChargerBatteryVoltage) Value() string {
	if v, ok := solarChargerBatteryVoltageMap[s]; ok {
		return v
	}
	return ""
}
