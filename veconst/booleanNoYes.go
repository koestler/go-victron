package veconst

type BooleanNoYes uint8

const (
	BooleanNo  BooleanNoYes = 0
	BooleanYes BooleanNoYes = 1
)

var booleanNoYesMap = map[BooleanNoYes]string{
	BooleanNo:  "No",
	BooleanYes: "Yes",
}

type BooleanNoYesFactoryType struct{}

var BooleanNoYesFactory BooleanNoYesFactoryType

func (f BooleanNoYesFactoryType) NewEnum(v int) (Enum, error) {
	s := BooleanNoYes(v)
	if _, ok := booleanNoYesMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f BooleanNoYesFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(booleanNoYesMap))
	for k, v := range booleanNoYesMap {
		ret[int(k)] = v
	}
	return ret
}

func (s BooleanNoYes) Idx() int {
	return int(s)
}

func (s BooleanNoYes) String() string {
	if v, ok := booleanNoYesMap[s]; ok {
		return v
	}
	return ""
}

func (s BooleanNoYes) Exists() bool {
	_, ok := booleanNoYesMap[s]
	return ok
}
