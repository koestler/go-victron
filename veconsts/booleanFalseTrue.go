package veconsts

type BooleanFalseTrue byte

const (
	BooleanFalse BooleanFalseTrue = 0
	BooleanTrue  BooleanFalseTrue = 1
)

var booleanFalseTrueMap = map[BooleanFalseTrue]string{
	BooleanFalse: "False",
	BooleanTrue:  "True",
}

// GetBooleanFalseTrueStringMap returns a map of BooleanFalseTrue values to their string representation.
func GetBooleanFalseTrueStringMap() map[BooleanFalseTrue]string {
	ret := make(map[BooleanFalseTrue]string, len(booleanFalseTrueMap))
	for k, v := range booleanFalseTrueMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BooleanFalseTrue exists.
func (s BooleanFalseTrue) Exists() bool {
	_, ok := booleanFalseTrueMap[s]
	return ok
}

// String returns the string representation of a BooleanFalseTrue.
func (s BooleanFalseTrue) String() string {
	if v, ok := booleanFalseTrueMap[s]; ok {
		return v
	}
	return ""
}
