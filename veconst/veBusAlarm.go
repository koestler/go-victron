package veconst

type VeBusAlarm uint8

const (
	VeBusAlarmNoAlarm   VeBusAlarm = 0
	VeBusAlarmWarning   VeBusAlarm = 1
	VeBusAlarmAlarm     VeBusAlarm = 2
	VeBusAlarmUndefined VeBusAlarm = 3
)

var veBusAlarmMap = map[VeBusAlarm]string{
	VeBusAlarmNoAlarm:   "No Alarm",
	VeBusAlarmWarning:   "Warning",
	VeBusAlarmAlarm:     "Alarm",
	VeBusAlarmUndefined: "Undefined",
}

// VeBusAlarmStringMap returns a map of VeBusAlarm values to their string representation.
func VeBusAlarmStringMap() map[VeBusAlarm]string {
	ret := make(map[VeBusAlarm]string, len(veBusAlarmMap))
	for k, v := range veBusAlarmMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the VeBusAlarm exists.
func (s VeBusAlarm) Exists() bool {
	_, ok := veBusAlarmMap[s]
	return ok
}

// String returns the string representation of a VeBusAlarm.
func (s VeBusAlarm) String() string {
	if v, ok := veBusAlarmMap[s]; ok {
		return v
	}
	return ""
}
