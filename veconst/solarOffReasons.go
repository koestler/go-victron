package veconst

type SolarOffReason uint8
type SolarOffReasons uint32
type SolarOffReasonsFactoryType struct{}

const (
	SolarOffReasonNoInputPower        SolarOffReason = 0
	SolarOffReasonPhysicalPowerSwitch SolarOffReason = 1
	SolarOffReasonSoftPowerSwitch     SolarOffReason = 2
	SolarOffReasonRemoteInput         SolarOffReason = 3
	SolarOffReasonInternal            SolarOffReason = 4
	SolarOffReasonPayGo               SolarOffReason = 5
	SolarOffReasonBmsShutdown         SolarOffReason = 6
	SolarOffReasonLowTemp             SolarOffReason = 9
)

var solarOffReasonMap = map[SolarOffReason]string{
	SolarOffReasonNoInputPower:        "No input power",
	SolarOffReasonPhysicalPowerSwitch: "Physical power switch",
	SolarOffReasonSoftPowerSwitch:     "Soft power switch",
	SolarOffReasonRemoteInput:         "Remote input",
	SolarOffReasonInternal:            "Internal reason",
	SolarOffReasonPayGo:               "Pay-as-you-go out of credit",
	SolarOffReasonBmsShutdown:         "BMS shutdown",
	SolarOffReasonLowTemp:             "Battery temperature too low",
}
var SolarOffReasonsFactory SolarOffReasonsFactoryType

func (f SolarOffReasonsFactoryType) New(v uint32) (SolarOffReasons, error) {
	return SolarOffReasons(v), nil
}

func (f SolarOffReasonsFactoryType) NewFieldList(v uint) (FieldList, error) {
	return f.New(uint32(v))
}

func (f SolarOffReasonsFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(solarOffReasonMap))
	for k, v := range solarOffReasonMap {
		ret[int(k)] = v
	}
	return ret
}

func (l SolarOffReasons) Decode() map[SolarOffReason]bool {
	ret := make(map[SolarOffReason]bool, len(solarOffReasonMap))
	for i := range solarOffReasonMap {
		ret[i] = l&(1<<i) != 0
	}
	return ret
}

func (l SolarOffReasons) Fields() map[Field]bool {
	m := l.Decode()
	ret := make(map[Field]bool, len(m))
	for i, v := range m {
		ret[i] = v
	}
	return ret
}

func (s SolarOffReason) Idx() int {
	return int(s)
}

func (s SolarOffReason) String() string {
	if v, ok := solarOffReasonMap[s]; ok {
		return v
	}
	return ""
}
