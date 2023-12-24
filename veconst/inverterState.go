package veconst

type InverterState uint8

const (
	InverterStateOff       InverterState = 0
	InverterStateLowPower  InverterState = 1
	InverterStateFault     InverterState = 2
	InverterStateInverting InverterState = 9
)

var inverterStateMap = map[InverterState]string{
	InverterStateOff:       "Off",
	InverterStateLowPower:  "Low Power",
	InverterStateFault:     "Fault",
	InverterStateInverting: "Inverting",
}

// InverterStateStringMap returns a map of InverterState values to their string representation.
func InverterStateStringMap() map[InverterState]string {
	ret := make(map[InverterState]string, len(inverterStateMap))
	for k, v := range inverterStateMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the InverterState exists.
func (s InverterState) Exists() bool {
	_, ok := inverterStateMap[s]
	return ok
}

// String returns the string representation of a InverterState.
func (s InverterState) String() string {
	if v, ok := inverterStateMap[s]; ok {
		return v
	}
	return ""
}
