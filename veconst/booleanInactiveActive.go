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

func NewBooleanInactiveActiveEnum(v int) (Enum, error) {
	s := BooleanInactiveActive(v)
	if _, ok := booleanInactiveActiveMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func BooleanInactiveActiveMap() map[int]string {
	ret := make(map[int]string, len(booleanInactiveActiveMap))
	for k, v := range booleanInactiveActiveMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanInactiveActive) Idx() int {
	return int(s)
}

func (s BooleanInactiveActive) Value() string {
	if v, ok := booleanInactiveActiveMap[s]; ok {
		return v
	}
	return ""
}
