package veconst

type InverterWarningReason uint8
type InverterWarningReasons uint16
type InverterWarningReasonsFactoryType struct{}

const (
	InverterWarningReasonLowBatVoltage  InverterWarningReason = 0
	InverterWarningReasonHighBatVoltage InverterWarningReason = 1
	InverterWarningReasonLowTemp        InverterWarningReason = 5
	InverterWarningReasonHighTemp       InverterWarningReason = 6
	InverterWarningReasonOverload       InverterWarningReason = 8
	InverterWarningReasonPoorDC         InverterWarningReason = 9
	InverterWarningReasonLowAcVoltage   InverterWarningReason = 10
	InverterWarningReasonHighAcVoltage  InverterWarningReason = 11
)

var inverterWarningReasonMap = map[InverterWarningReason]string{
	InverterWarningReasonLowBatVoltage:  "Low battery voltage",
	InverterWarningReasonHighBatVoltage: "High battery voltage",
	InverterWarningReasonLowTemp:        "Low temperature",
	InverterWarningReasonHighTemp:       "High temperature",
	InverterWarningReasonOverload:       "Overload",
	InverterWarningReasonPoorDC:         "Poor DC connection",
	InverterWarningReasonLowAcVoltage:   "Low AC-output voltage",
	InverterWarningReasonHighAcVoltage:  "High AC-output voltage",
}
var InverterWarningReasonFactory InverterWarningReasonsFactoryType

func (f InverterWarningReasonsFactoryType) New(v uint16) (InverterWarningReasons, error) {
	return InverterWarningReasons(v), nil
}

func (f InverterWarningReasonsFactoryType) NewFieldList(v uint) (FieldList, error) {
	return f.New(uint16(v))
}

func (f InverterWarningReasonsFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(inverterWarningReasonMap))
	for k, v := range inverterWarningReasonMap {
		ret[int(k)] = v
	}
	return ret
}

func (l InverterWarningReasons) Decode() map[InverterWarningReason]bool {
	ret := make(map[InverterWarningReason]bool, len(inverterWarningReasonMap))
	for i := range inverterWarningReasonMap {
		ret[i] = l&(1<<i) != 0
	}
	return ret
}

func (l InverterWarningReasons) Fields() map[Field]bool {
	m := l.Decode()
	ret := make(map[Field]bool, len(m))
	for i, v := range m {
		ret[i] = v
	}
	return ret
}

func (s InverterWarningReason) Idx() int {
	return int(s)
}

func (s InverterWarningReason) String() string {
	if v, ok := inverterWarningReasonMap[s]; ok {
		return v
	}
	return ""
}
