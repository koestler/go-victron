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

// GetInverterFrequencyStringMap returns a map of InverterFrequency values to their string representation.
func GetInverterFrequencyStringMap() map[InverterFrequency]string {
	ret := make(map[InverterFrequency]string, len(inverterFrequencyMap))
	for k, v := range inverterFrequencyMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the InverterFrequency exists.
func (s InverterFrequency) Exists() bool {
	_, ok := inverterFrequencyMap[s]
	return ok
}

// String returns the string representation of a InverterFrequency.
func (s InverterFrequency) String() string {
	if v, ok := inverterFrequencyMap[s]; ok {
		return v
	}
	return ""
}
