package veconsts

type VeBusAlarm uint8

const (
	VeBusAlarmNoAlarm   VeBusAlarm = 0
	VeBusAlarmWarning   VeBusAlarm = 1
	VeBusAlarmAlarm     VeBusAlarm = 2
	VeBusAlarmUndefined VeBusAlarm = 3
)

func GetVeBusAlarmMap() map[VeBusAlarm]string {
	return map[VeBusAlarm]string{
		VeBusAlarmNoAlarm:   "No Alarm",
		VeBusAlarmWarning:   "Warning",
		VeBusAlarmAlarm:     "Alarm",
		VeBusAlarmUndefined: "Undefined",
	}
}

func (s VeBusAlarm) String() string {
	m := GetVeBusAlarmMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
