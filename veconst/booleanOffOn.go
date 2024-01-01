package veconst

type BooleanOffOn uint8

const (
	BooleanOff BooleanOffOn = 0
	BooleanOn  BooleanOffOn = 1
)

var booleanOffOnMap = map[BooleanOffOn]string{
	BooleanOff: "Off",
	BooleanOn:  "On",
}

func NewBooleanOffOnEnum(v int) (Enum, error) {
	s := BooleanOffOn(v)
	if _, ok := booleanOffOnMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func BooleanOffOnMap() map[int]string {
	ret := make(map[int]string, len(booleanOffOnMap))
	for k, v := range booleanOffOnMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanOffOn) Idx() int {
	return int(s)
}

func (s BooleanOffOn) Value() string {
	if v, ok := booleanOffOnMap[s]; ok {
		return v
	}
	return ""
}
