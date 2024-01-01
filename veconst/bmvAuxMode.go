package veconst

type BmvAuxMode uint8

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

type BmvAuxModeFactoryType struct{}

var BmvAuxModeFactory BmvAuxModeFactoryType

func (f BmvAuxModeFactoryType) NewEnum(v int) (Enum, error) {
	s := BmvAuxMode(v)
	if _, ok := bmvAuxModeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
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

func (s BmvAuxMode) Exists() bool {
	_, ok := bmvAuxModeMap[s]
	return ok
}
