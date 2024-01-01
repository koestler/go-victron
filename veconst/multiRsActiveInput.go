package veconst

type MultiRsActiveInput uint8

const (
	MultiRsActiveInputAC1          MultiRsActiveInput = 0
	MultiRsActiveInputAC2          MultiRsActiveInput = 1
	MultiRsActiveInputNotConnected MultiRsActiveInput = 2
	MultiRsActiveInputUnknown      MultiRsActiveInput = 3
)

var multiRsActiveInputMap = map[MultiRsActiveInput]string{
	MultiRsActiveInputAC1:          "AC in 1",
	MultiRsActiveInputAC2:          "AC in 2",
	MultiRsActiveInputNotConnected: "Not Connected",
	MultiRsActiveInputUnknown:      "Unknown",
}

type MultiRsActiveInputFactoryType struct{}

var MultiRsActiveInputFactory MultiRsActiveInputFactoryType

func (f MultiRsActiveInputFactoryType) NewEnum(v int) (Enum, error) {
	s := MultiRsActiveInput(v)
	if _, ok := multiRsActiveInputMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f MultiRsActiveInputFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(multiRsActiveInputMap))
	for k, v := range multiRsActiveInputMap {
		ret[int(k)] = v
	}
	return ret
}

func (s MultiRsActiveInput) Idx() int {
	return int(s)
}

func (s MultiRsActiveInput) String() string {
	if v, ok := multiRsActiveInputMap[s]; ok {
		return v
	}
	return ""
}

func (s MultiRsActiveInput) Exists() bool {
	_, ok := multiRsActiveInputMap[s]
	return ok
}
