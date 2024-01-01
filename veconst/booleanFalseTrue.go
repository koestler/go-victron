package veconst

type BooleanFalseTrue uint8

const (
	BooleanFalse BooleanFalseTrue = 0
	BooleanTrue  BooleanFalseTrue = 1
)

var booleanFalseTrueMap = map[BooleanFalseTrue]string{
	BooleanFalse: "False",
	BooleanTrue:  "True",
}

type BooleanFalseTrueFactoryType struct{}

var BooleanFalseTrueFactory BooleanFalseTrueFactoryType

func (f BooleanFalseTrueFactoryType) NewEnum(v int) (Enum, error) {
	s := BooleanFalseTrue(v)
	if _, ok := booleanFalseTrueMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanFalseTrueFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(booleanFalseTrueMap))
	for k, v := range booleanFalseTrueMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanFalseTrue) Idx() int {
	return int(s)
}

func (s BooleanFalseTrue) String() string {
	if v, ok := booleanFalseTrueMap[s]; ok {
		return v
	}
	return ""
}

func (s BooleanFalseTrue) Exists() bool {
	_, ok := booleanFalseTrueMap[s]
	return ok
}
