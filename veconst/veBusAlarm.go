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

func NewVeBusAlarmEnum(v int) (Enum, error) {
	s := VeBusAlarm(v)
	if _, ok := veBusAlarmMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func VeBusAlarmMap() map[int]string {
	ret := make(map[int]string, len(veBusAlarmMap))
	for k, v := range veBusAlarmMap {
		ret[int(k)] = v
	}
	return ret
}

func (s VeBusAlarm) Idx() int {
	return int(s)
}

func (s VeBusAlarm) Value() string {
	if v, ok := veBusAlarmMap[s]; ok {
		return v
	}
	return ""
}
