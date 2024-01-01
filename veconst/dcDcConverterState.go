package veconst

type DcDcConverterState uint8
type DcDcConverterStateFactoryType struct{}

const (
	DcDcConverterStateNotCharging        DcDcConverterState = 0
	DcDcConverterStateFault              DcDcConverterState = 2
	DcDcConverterStateBulkCharging       DcDcConverterState = 3
	DcDcConverterStateAbsorptionCharging DcDcConverterState = 4
	DcDcConverterStateFloatCharging      DcDcConverterState = 5
	DcDcConverterStateManualEqualise     DcDcConverterState = 7
	DcDcConverterStateWakeUp             DcDcConverterState = 245
	DcDcConverterStateAutoEqualise       DcDcConverterState = 247
	DcDcConverterStateExternalControl    DcDcConverterState = 252
	DcDcConverterStateUnavailable        DcDcConverterState = 255
)

var dcDcConverterStateMap = map[DcDcConverterState]string{
	DcDcConverterStateNotCharging:        "Not charging",
	DcDcConverterStateFault:              "Fault",
	DcDcConverterStateBulkCharging:       "Bulk Charging",
	DcDcConverterStateAbsorptionCharging: "Absorption Charging",
	DcDcConverterStateFloatCharging:      "Float Charging",
	DcDcConverterStateManualEqualise:     "Manual Equalise",
	DcDcConverterStateWakeUp:             "Wake-Up",
	DcDcConverterStateAutoEqualise:       "Auto Equalise",
	DcDcConverterStateExternalControl:    "External Control",
	DcDcConverterStateUnavailable:        "Unavailable",
}
var DcDcConverterStateFactory DcDcConverterStateFactoryType

func (f DcDcConverterStateFactoryType) New(v int) (DcDcConverterState, error) {
	s := DcDcConverterState(v)
	if _, ok := dcDcConverterStateMap[s]; !ok {
		return DcDcConverterStateUnavailable, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f DcDcConverterStateFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
}

func (f DcDcConverterStateFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(dcDcConverterStateMap))
	for k, v := range dcDcConverterStateMap {
		ret[int(k)] = v
	}
	return ret
}

func (s DcDcConverterState) Idx() int {
	return int(s)
}

func (s DcDcConverterState) String() string {
	if v, ok := dcDcConverterStateMap[s]; ok {
		return v
	}
	return ""
}

func (s DcDcConverterState) Exists() bool {
	_, ok := dcDcConverterStateMap[s]
	return ok
}
