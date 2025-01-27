package veconst

type SolarChargerLoadOffReason uint8
type SolarChargerLoadOffReasonFactoryType struct{}

const (
	SolarChargerLoadOffReasonBatteryLow          SolarChargerLoadOffReason = 0
	SolarChargerLoadOffReasonBatteryShortCircuit SolarChargerLoadOffReason = 1
	SolarChargerLoadOffReasonBatteryTimerProgram SolarChargerLoadOffReason = 2
	SolarChargerLoadOffReasonBatteryRemoteInput  SolarChargerLoadOffReason = 3
	SolarChargerLoadOffReasonBatteryCredit       SolarChargerLoadOffReason = 4
	SolarChargerLoadOffReasonBatteryStartUp      SolarChargerLoadOffReason = 7
)

var solarChargerLoadOffReasonMap = map[SolarChargerLoadOffReason]string{
	SolarChargerLoadOffReasonBatteryLow:          "Battery low",
	SolarChargerLoadOffReasonBatteryShortCircuit: "Short circuit",
	SolarChargerLoadOffReasonBatteryTimerProgram: "Timer program",
	SolarChargerLoadOffReasonBatteryRemoteInput:  "Remote input",
	SolarChargerLoadOffReasonBatteryCredit:       "Pay-as-you-go out of credit",
	SolarChargerLoadOffReasonBatteryStartUp:      "Device starting up",
}
var SolarChargerLoadOffReasonFactory SolarChargerLoadOffReasonFactoryType

func (f SolarChargerLoadOffReasonFactoryType) New(v uint8) (SolarChargerLoadOffReason, error) {
	s := SolarChargerLoadOffReason(v)
	if _, ok := solarChargerLoadOffReasonMap[s]; !ok {
		return SolarChargerLoadOffReasonBatteryStartUp, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f SolarChargerLoadOffReasonFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f SolarChargerLoadOffReasonFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarChargerLoadOffReasonMap))
	for k, v := range solarChargerLoadOffReasonMap {
		ret[int(k)] = v
	}
	return ret
}

func (s SolarChargerLoadOffReason) Idx() int {
	return int(s)
}

func (s SolarChargerLoadOffReason) String() string {
	if v, ok := solarChargerLoadOffReasonMap[s]; ok {
		return v
	}
	return ""
}
