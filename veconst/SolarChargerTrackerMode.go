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

type SolarChargerTrackerModeFactoryType struct{}

var SolarChargerTrackerModeFactory SolarChargerTrackerModeFactoryType

func (f SolarChargerTrackerModeFactoryType) NewEnum(v int) (Enum, error) {
	s := SolarChargerTrackerMode(v)
	if _, ok := solarChargerTrackerModeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
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

func (s SolarChargerTrackerMode) Exists() bool {
	_, ok := solarChargerTrackerModeMap[s]
	return ok
}
