package veconst

type InverterState uint8
type InverterStateFactoryType struct{}

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
var InverterStateFactory InverterStateFactoryType

func (f InverterStateFactoryType) New(v uint8) (InverterState, error) {
	s := InverterState(v)
	if _, ok := inverterStateMap[s]; !ok {
		return InverterStateOff, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f InverterStateFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f InverterStateFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(inverterStateMap))
	for k, v := range inverterStateMap {
		ret[int(k)] = v
	}
	return ret
}

func (s InverterState) Idx() int {
	return int(s)
}

func (s InverterState) String() string {
	if v, ok := inverterStateMap[s]; ok {
		return v
	}
	return ""
}
