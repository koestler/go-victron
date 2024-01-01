package veconst

type BooleanInactiveActive uint8
type BooleanInactiveActiveFactoryType struct{}

const (
	BooleanInactive BooleanInactiveActive = 0
	BooleanActive   BooleanInactiveActive = 1
)

var booleanInactiveActiveMap = map[BooleanInactiveActive]string{
	BooleanInactive: "Inactive",
	BooleanActive:   "Active",
}
var BooleanInactiveActiveFactory BooleanInactiveActiveFactoryType

func (f BooleanInactiveActiveFactoryType) New(v uint8) (BooleanInactiveActive, error) {
	s := BooleanInactiveActive(v)
	if _, ok := booleanInactiveActiveMap[s]; !ok {
		return BooleanInactive, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanInactiveActiveFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f BooleanInactiveActiveFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(booleanInactiveActiveMap))
	for k, v := range booleanInactiveActiveMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanInactiveActive) Idx() int {
	return int(s)
}

func (s BooleanInactiveActive) String() string {
	if v, ok := booleanInactiveActiveMap[s]; ok {
		return v
	}
	return ""
}
