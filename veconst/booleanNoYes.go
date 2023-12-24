package veconst

type BooleanNoYes byte

const (
	BooleanNo  BooleanNoYes = 0
	BooleanYes BooleanNoYes = 1
)

var booleanNoYesMap = map[BooleanNoYes]string{
	BooleanNo:  "No",
	BooleanYes: "Yes",
}

// GetBooleanNoYesStringMap returns a map of BooleanNoYes values to their string representation.
func GetBooleanNoYesStringMap() map[BooleanNoYes]string {
	ret := make(map[BooleanNoYes]string, len(booleanNoYesMap))
	for k, v := range booleanNoYesMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BooleanNoYes exists.
func (s BooleanNoYes) Exists() bool {
	_, ok := booleanNoYesMap[s]
	return ok
}

// String returns the string representation of a BooleanNoYes.
func (s BooleanNoYes) String() string {
	if v, ok := booleanNoYesMap[s]; ok {
		return v
	}
	return ""
}
