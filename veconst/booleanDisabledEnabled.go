package veconst

type BooleanDisabledEnabled byte

const (
	BooleanDisabled BooleanDisabledEnabled = 0
	BooleanEnabled  BooleanDisabledEnabled = 1
)

var booleanDisabledEnabledMap = map[BooleanDisabledEnabled]string{
	BooleanDisabled: "Disabled",
	BooleanEnabled:  "Enabled",
}

// GetBooleanDisabledEnabledStringMap returns a map of BooleanDisabledEnabled values to their string representation.
func GetBooleanDisabledEnabledStringMap() map[BooleanDisabledEnabled]string {
	ret := make(map[BooleanDisabledEnabled]string, len(booleanDisabledEnabledMap))
	for k, v := range booleanDisabledEnabledMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the BooleanDisabledEnabled exists.
func (s BooleanDisabledEnabled) Exists() bool {
	_, ok := booleanDisabledEnabledMap[s]
	return ok
}

// String returns the string representation of a BooleanDisabledEnabled.
func (s BooleanDisabledEnabled) String() string {
	if v, ok := booleanDisabledEnabledMap[s]; ok {
		return v
	}
	return ""
}
