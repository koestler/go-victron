package veconst

type InverterFrequency uint8

const (
	InverterFrequency60Hz InverterFrequency = 0
	InverterFrequency50Hz InverterFrequency = 1
)

var inverterFrequencyMap = map[InverterFrequency]string{
	InverterFrequency60Hz: "60 Hz",
	InverterFrequency50Hz: "50 Hz",
}

type InverterFrequencyFactoryType struct{}

var InverterFrequencyFactory InverterFrequencyFactoryType

func (f InverterFrequencyFactoryType) NewEnum(v int) (Enum, error) {
	s := InverterFrequency(v)
	if _, ok := inverterFrequencyMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f InverterFrequencyFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(inverterFrequencyMap))
	for k, v := range inverterFrequencyMap {
		ret[int(k)] = v
	}
	return ret
}

func (s InverterFrequency) Idx() int {
	return int(s)
}

func (s InverterFrequency) String() string {
	if v, ok := inverterFrequencyMap[s]; ok {
		return v
	}
	return ""
}

func (s InverterFrequency) Exists() bool {
	_, ok := inverterFrequencyMap[s]
	return ok
}
