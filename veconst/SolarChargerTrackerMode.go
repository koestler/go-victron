package veconst

type SolarChargerTrackerMode uint8
type SolarChargerTrackerModeFactoryType struct{}

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
var SolarChargerTrackerModeFactory SolarChargerTrackerModeFactoryType

func (f SolarChargerTrackerModeFactoryType) New(v uint8) (SolarChargerTrackerMode, error) {
	s := SolarChargerTrackerMode(v)
	if _, ok := solarChargerTrackerModeMap[s]; !ok {
		return SolarChargerTrackerModeOff, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f SolarChargerTrackerModeFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f SolarChargerTrackerModeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarChargerTrackerModeMap))
	for k, v := range solarChargerTrackerModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerTrackerMode) Idx() int {
	return int(s)
}

func (s SolarChargerTrackerMode) String() string {
	if v, ok := solarChargerTrackerModeMap[s]; ok {
		return v
	}
	return ""
}
