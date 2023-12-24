package veconsts

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

// GetMultiRsActiveInputStringMap returns a map of MultiRsActiveInput values to their string representation.
func GetMultiRsActiveInputStringMap() map[MultiRsActiveInput]string {
	ret := make(map[MultiRsActiveInput]string, len(multiRsActiveInputMap))
	for k, v := range multiRsActiveInputMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the MultiRsActiveInput exists.
func (s MultiRsActiveInput) Exists() bool {
	_, ok := multiRsActiveInputMap[s]
	return ok
}

// String returns the string representation of a MultiRsActiveInput.
func (s MultiRsActiveInput) String() string {
	if v, ok := multiRsActiveInputMap[s]; ok {
		return v
	}
	return ""
}
