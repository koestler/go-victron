package victronDefinitions

type MultiRsActiveInput uint8

const (
	MultiRsActiveInputAC1          MultiRsActiveInput = 0
	MultiRsActiveInputAC2          MultiRsActiveInput = 1
	MultiRsActiveInputNotConnected MultiRsActiveInput = 2
	MultiRsActiveInputUnknown      MultiRsActiveInput = 3
)

func GetMultiRsActiveInputMap() map[MultiRsActiveInput]string {
	return map[MultiRsActiveInput]string{
		MultiRsActiveInputAC1:          "AC in 1",
		MultiRsActiveInputAC2:          "AC in 2",
		MultiRsActiveInputNotConnected: "Not Connected",
		MultiRsActiveInputUnknown:      "Unknown",
	}
}

func (s MultiRsActiveInput) String() string {
	m := GetMultiRsActiveInputMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
