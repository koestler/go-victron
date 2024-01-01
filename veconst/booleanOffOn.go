package veconst

type BooleanOffOn uint8
type BooleanOffOnFactoryType struct{}

const (
	BooleanOff BooleanOffOn = 0
	BooleanOn  BooleanOffOn = 1
)

var booleanOffOnMap = map[BooleanOffOn]string{
	BooleanOff: "Off",
	BooleanOn:  "On",
}
var BooleanOffOnFactory BooleanOffOnFactoryType

func (f BooleanOffOnFactoryType) New(v uint8) (BooleanOffOn, error) {
	s := BooleanOffOn(v)
	if _, ok := booleanOffOnMap[s]; !ok {
		return BooleanOff, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanOffOnFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f BooleanOffOnFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(booleanOffOnMap))
	for k, v := range booleanOffOnMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanOffOn) Idx() int {
	return int(s)
}

func (s BooleanOffOn) String() string {
	if v, ok := booleanOffOnMap[s]; ok {
		return v
	}
	return ""
}
