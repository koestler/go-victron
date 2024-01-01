package veconst

type BooleanDisabledEnabled uint8

const (
	BooleanDisabled BooleanDisabledEnabled = 0
	BooleanEnabled  BooleanDisabledEnabled = 1
)

var booleanDisabledEnabledMap = map[BooleanDisabledEnabled]string{
	BooleanDisabled: "Disabled",
	BooleanEnabled:  "Enabled",
}

type BooleanDisabledEnabledFactoryType struct{}

var BooleanDisabledEnabledFactory BooleanDisabledEnabledFactoryType

func (f BooleanDisabledEnabledFactoryType) NewEnum(v int) (Enum, error) {
	s := BooleanDisabledEnabled(v)
	if _, ok := booleanDisabledEnabledMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanDisabledEnabledFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(booleanDisabledEnabledMap))
	for k, v := range booleanDisabledEnabledMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanDisabledEnabled) Idx() int {
	return int(s)
}

func (s BooleanDisabledEnabled) String() string {
	if v, ok := booleanDisabledEnabledMap[s]; ok {
		return v
	}
	return ""
}

func (s BooleanDisabledEnabled) Exists() bool {
	_, ok := booleanDisabledEnabledMap[s]
	return ok
}
