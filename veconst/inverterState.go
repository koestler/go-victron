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

func NewInverterStateEnum(v int) (Enum, error) {
	s := InverterState(v)
	if _, ok := inverterStateMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func InverterStateMap() map[int]string {
	ret := make(map[int]string, len(inverterStateMap))
	for k, v := range inverterStateMap {
		ret[int(k)] = v
	}
	return ret
}

func (s InverterState) Idx() int {
	return int(s)
}

func (s InverterState) Value() string {
	if v, ok := inverterStateMap[s]; ok {
		return v
	}
	return ""
}
