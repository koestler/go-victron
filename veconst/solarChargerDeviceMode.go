package veconst

type SolarChargerDeviceMode uint8
type SolarChargerDeviceModeFactoryType struct{}

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
var SolarChargerDeviceModeFactory SolarChargerDeviceModeFactoryType

func (f SolarChargerDeviceModeFactoryType) New(v int) (SolarChargerDeviceMode, error) {
	s := SolarChargerDeviceMode(v)
	if _, ok := solarChargerDeviceModeMap[s]; !ok {
		return SolarChargerDeviceModeOff, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f SolarChargerDeviceModeFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
}

func (f SolarChargerDeviceModeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarChargerDeviceModeMap))
	for k, v := range solarChargerDeviceModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerDeviceMode) Idx() int {
	return int(s)
}

func (s SolarChargerDeviceMode) String() string {
	if v, ok := solarChargerDeviceModeMap[s]; ok {
		return v
	}
	return ""
}

func (s SolarChargerDeviceMode) Exists() bool {
	_, ok := solarChargerDeviceModeMap[s]
	return ok
}
