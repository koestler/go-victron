package veconst

type SolarChargerState uint8
type SolarChargerStateFactoryType struct{}

const (
	SolarChargerStateNotCharging        SolarChargerState = 0
	SolarChargerStateFault              SolarChargerState = 2
	SolarChargerStateBulkCharging       SolarChargerState = 3
	SolarChargerStateAbsorptionCharging SolarChargerState = 4
	SolarChargerStateFloatCharging      SolarChargerState = 5
	SolarChargerStateManualEqualise     SolarChargerState = 7
	SolarChargerStateWakeUp             SolarChargerState = 245
	SolarChargerStateAutoEqualise       SolarChargerState = 247
	SolarChargerStateExternalControl    SolarChargerState = 252
	SolarChargerStateUnavailable        SolarChargerState = 255
)

var solarChargerStateMap = map[SolarChargerState]string{
	SolarChargerStateNotCharging:        "Not charging",
	SolarChargerStateFault:              "Fault",
	SolarChargerStateBulkCharging:       "Bulk Charging",
	SolarChargerStateAbsorptionCharging: "Absorption Charging",
	SolarChargerStateFloatCharging:      "Float Charging",
	SolarChargerStateManualEqualise:     "Manual Equalise",
	SolarChargerStateWakeUp:             "Wake-Up",
	SolarChargerStateAutoEqualise:       "Auto Equalise",
	SolarChargerStateExternalControl:    "External Control",
	SolarChargerStateUnavailable:        "Unavailable",
}
var SolarChargerStateFactory SolarChargerStateFactoryType

func (f SolarChargerStateFactoryType) New(v int) (SolarChargerState, error) {
	s := SolarChargerState(v)
	if _, ok := solarChargerStateMap[s]; !ok {
		return SolarChargerStateUnavailable, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f SolarChargerStateFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
}

func (f SolarChargerStateFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarChargerStateMap))
	for k, v := range solarChargerStateMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerState) Idx() int {
	return int(s)
}

func (s SolarChargerState) String() string {
	if v, ok := solarChargerStateMap[s]; ok {
		return v
	}
	return ""
}

func (s SolarChargerState) Exists() bool {
	_, ok := solarChargerStateMap[s]
	return ok
}
