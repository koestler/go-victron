package veconst

type BmvAuxMode uint8
type BmvAuxModeFactoryType struct{}

const (
	BmvAuxModeStarterVoltage  BmvAuxMode = 0
	BmvAuxModeMidpointVoltage BmvAuxMode = 1
	BmvAuxModeTemperature     BmvAuxMode = 2
	BmvAuxModeDisabled        BmvAuxMode = 3
)

var bmvAuxModeMap = map[BmvAuxMode]string{
	BmvAuxModeStarterVoltage:  "Starter voltage",
	BmvAuxModeMidpointVoltage: "Mid-point voltage",
	BmvAuxModeTemperature:     "Temperature",
	BmvAuxModeDisabled:        "Disabled",
}
var BmvAuxModeFactory BmvAuxModeFactoryType

func (f BmvAuxModeFactoryType) New(v uint8) (BmvAuxMode, error) {
	s := BmvAuxMode(v)
	if _, ok := bmvAuxModeMap[s]; !ok {
		return BmvAuxModeDisabled, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BmvAuxModeFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f BmvAuxModeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(bmvAuxModeMap))
	for k, v := range bmvAuxModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BmvAuxMode) Idx() int {
	return int(s)
}

func (s BmvAuxMode) String() string {
	if v, ok := bmvAuxModeMap[s]; ok {
		return v
	}
	return ""
}
