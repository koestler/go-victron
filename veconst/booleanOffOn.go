package veconst

type BooleanOffOn byte

const (
	BooleanOff BooleanOffOn = 0
	BooleanOn  BooleanOffOn = 1
)

var booleanOffOnMap = map[BooleanOffOn]string{
	BooleanOff: "Off",
	BooleanOn:  "On",
}

// GetBooleanOffOnStringMap returns a map of BooleanOffOn values to their string representation.
func GetBooleanOffOnStringMap() map[BooleanOffOn]string {
	ret := make(map[BooleanOffOn]string, len(booleanOffOnMap))
	for k, v := range booleanOffOnMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BooleanOffOn exists.
func (s BooleanOffOn) Exists() bool {
	_, ok := booleanOffOnMap[s]
	return ok
}

// String returns the string representation of a BooleanOffOn.
func (s BooleanOffOn) String() string {
	if v, ok := booleanOffOnMap[s]; ok {
		return v
	}
	return ""
}
