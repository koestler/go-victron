package veconst

type VeBusAlarm uint8
type VeBusAlarmFactoryType struct{}

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
var VeBusAlarmFactory VeBusAlarmFactoryType

func (f VeBusAlarmFactoryType) New(v uint8) (VeBusAlarm, error) {
	s := VeBusAlarm(v)
	if _, ok := veBusAlarmMap[s]; !ok {
		return VeBusAlarmUndefined, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f VeBusAlarmFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f VeBusAlarmFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(veBusAlarmMap))
	for k, v := range veBusAlarmMap {
		ret[int(k)] = v
	}
	return ret
}

func (s VeBusAlarm) Idx() int {
	return int(s)
}

func (s VeBusAlarm) String() string {
	if v, ok := veBusAlarmMap[s]; ok {
		return v
	}
	return ""
}
