package veconst

type DcDcConverterState uint16

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

// GetDcDcConverterStateStringMap returns a map of DcDcConverterState values to their string representation.
func GetDcDcConverterStateStringMap() map[DcDcConverterState]string {
	ret := make(map[DcDcConverterState]string, len(dcDcConverterStateMap))
	for k, v := range dcDcConverterStateMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the DcDcConverterState exists.
func (s DcDcConverterState) Exists() bool {
	_, ok := dcDcConverterStateMap[s]
	return ok
}

// String returns the string representation of a DcDcConverterState.
func (s DcDcConverterState) String() string {
	if v, ok := dcDcConverterStateMap[s]; ok {
		return v
	}
	return ""
}
