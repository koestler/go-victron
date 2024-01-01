package veconst

type DcDcConverterState uint8

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

func NewDcDcConverterStateEnum(v int) (Enum, error) {
	s := DcDcConverterState(v)
	if _, ok := dcDcConverterStateMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func DcDcConverterStateMap() map[int]string {
	ret := make(map[int]string, len(dcDcConverterStateMap))
	for k, v := range dcDcConverterStateMap {
		ret[int(k)] = v
	}
	return ret
}

func (s DcDcConverterState) Idx() int {
	return int(s)
}

func (s DcDcConverterState) Value() string {
	if v, ok := dcDcConverterStateMap[s]; ok {
		return v
	}
	return ""
}
