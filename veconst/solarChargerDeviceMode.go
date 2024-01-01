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

func NewSolarChargerDeviceModeEnum(v int) (Enum, error) {
	s := SolarChargerDeviceMode(v)
	if _, ok := solarChargerDeviceModeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func SolarChargerDeviceModeMap() map[int]string {
	ret := make(map[int]string, len(solarChargerDeviceModeMap))
	for k, v := range solarChargerDeviceModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerDeviceMode) Idx() int {
	return int(s)
}

func (s SolarChargerDeviceMode) Value() string {
	if v, ok := solarChargerDeviceModeMap[s]; ok {
		return v
	}
	return ""
}
