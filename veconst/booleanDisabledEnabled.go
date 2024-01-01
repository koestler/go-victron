package veconst

type BooleanDisabledEnabled uint8
type BooleanDisabledEnabledFactoryType struct{}

const (
	BooleanDisabled BooleanDisabledEnabled = 0
	BooleanEnabled  BooleanDisabledEnabled = 1
)

var booleanDisabledEnabledMap = map[BooleanDisabledEnabled]string{
	BooleanDisabled: "Disabled",
	BooleanEnabled:  "Enabled",
}
var BooleanDisabledEnabledFactory BooleanDisabledEnabledFactoryType

func (f BooleanDisabledEnabledFactoryType) New(v int) (BooleanDisabledEnabled, error) {
	s := BooleanDisabledEnabled(v)
	if _, ok := booleanDisabledEnabledMap[s]; !ok {
		return BooleanDisabled, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanDisabledEnabledFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
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
