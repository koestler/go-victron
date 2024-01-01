package veconst

type BooleanFalseTrue uint8
type BooleanFalseTrueFactoryType struct{}

const (
	BooleanFalse BooleanFalseTrue = 0
	BooleanTrue  BooleanFalseTrue = 1
)

var booleanFalseTrueMap = map[BooleanFalseTrue]string{
	BooleanFalse: "False",
	BooleanTrue:  "True",
}
var BooleanFalseTrueFactory BooleanFalseTrueFactoryType

func (f BooleanFalseTrueFactoryType) New(v int) (BooleanFalseTrue, error) {
	s := BooleanFalseTrue(v)
	if _, ok := booleanFalseTrueMap[s]; !ok {
		return BooleanFalse, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanFalseTrueFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(v)
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
