package veconst

type BooleanInactiveActive uint8

const (
	BooleanInactive BooleanInactiveActive = 0
	BooleanActive   BooleanInactiveActive = 1
)

var booleanInactiveActiveMap = map[BooleanInactiveActive]string{
	BooleanInactive: "Inactive",
	BooleanActive:   "Active",
}

type BooleanInactiveActiveFactoryType struct{}

var BooleanInactiveActiveFactory BooleanInactiveActiveFactoryType

func (f BooleanInactiveActiveFactoryType) NewEnum(v int) (Enum, error) {
	s := BooleanInactiveActive(v)
	if _, ok := booleanInactiveActiveMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
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

func (s BooleanInactiveActive) Exists() bool {
	_, ok := booleanInactiveActiveMap[s]
	return ok
}
