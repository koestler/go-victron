package veconst

type InverterFrequency uint8
type InverterFrequencyFactoryType struct{}

const (
	InverterFrequency60Hz InverterFrequency = 0
	InverterFrequency50Hz InverterFrequency = 1
)

var inverterFrequencyMap = map[InverterFrequency]string{
	InverterFrequency60Hz: "60 Hz",
	InverterFrequency50Hz: "50 Hz",
}
var InverterFrequencyFactory InverterFrequencyFactoryType

func (f InverterFrequencyFactoryType) New(v uint8) (InverterFrequency, error) {
	s := InverterFrequency(v)
	if _, ok := inverterFrequencyMap[s]; !ok {
		return InverterFrequency60Hz, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f InverterFrequencyFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
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
