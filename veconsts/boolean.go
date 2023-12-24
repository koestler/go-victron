package veconsts

type Boolean uint8

const (
	BooleanFalse Boolean = 0
	BooleanTrue  Boolean = 1
)

var booleanMap = map[Boolean]string{
	BooleanFalse: "False",
	BooleanTrue:  "True",
}

// GetBooleanStringMap returns a map of Boolean values to their string representation.
func GetBooleanStringMap() map[Boolean]string {
	ret := make(map[Boolean]string, len(booleanMap))
	for k, v := range booleanMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the Boolean exists.
func (s Boolean) Exists() bool {
	_, ok := booleanMap[s]
	return ok
}

// String returns the string representation of a Boolean.
func (s Boolean) String() string {
	if v, ok := booleanMap[s]; ok {
		return v
	}
	return ""
}
