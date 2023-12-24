package veconst

type BooleanInactiveActive byte

const (
	BooleanInactive BooleanInactiveActive = 0
	BooleanActive   BooleanInactiveActive = 1
)

var booleanInactiveActiveMap = map[BooleanInactiveActive]string{
	BooleanInactive: "Inactive",
	BooleanActive:   "Active",
}

// GetBooleanInactiveActiveStringMap returns a map of BooleanInactiveActive values to their string representation.
func GetBooleanInactiveActiveStringMap() map[BooleanInactiveActive]string {
	ret := make(map[BooleanInactiveActive]string, len(booleanInactiveActiveMap))
	for k, v := range booleanInactiveActiveMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BooleanInactiveActive exists.
func (s BooleanInactiveActive) Exists() bool {
	_, ok := booleanInactiveActiveMap[s]
	return ok
}

// String returns the string representation of a BooleanInactiveActive.
func (s BooleanInactiveActive) String() string {
	if v, ok := booleanInactiveActiveMap[s]; ok {
		return v
	}
	return ""
}
