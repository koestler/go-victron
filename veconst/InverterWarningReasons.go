package veconst

type InverterWarningReason uint16
type InverterWarningReasons uint32
type InverterWarningReasonsFactoryType struct{}

const (
	DeviceWarningReasonLowBatVoltage  InverterWarningReason = 0
	DeviceWarningReasonHighBatVoltage InverterWarningReason = 1
	DeviceWarningReasonLowTemp        InverterWarningReason = 5
	DeviceWarningReasonHighTemp       InverterWarningReason = 6
	DeviceWarningReasonOverload       InverterWarningReason = 8
	DeviceWarningReasonPoorDC         InverterWarningReason = 9
	DeviceWarningReasonLowAcVoltage   InverterWarningReason = 10
	DeviceWarningReasonHighAcVoltage  InverterWarningReason = 11
)

var inverterWarningReasonMap = map[InverterWarningReason]string{
	DeviceWarningReasonLowBatVoltage:  "Low battery voltage",
	DeviceWarningReasonHighBatVoltage: "High battery voltage",
	DeviceWarningReasonLowTemp:        "Low temperature",
	DeviceWarningReasonHighTemp:       "High temperature",
	DeviceWarningReasonOverload:       "Overload",
	DeviceWarningReasonPoorDC:         "Poor DC connection",
	DeviceWarningReasonLowAcVoltage:   "Low AC-output voltage",
	DeviceWarningReasonHighAcVoltage:  "High AC-output voltage",
}
var InverterWarningReasonFactory InverterWarningReasonsFactoryType

func (f InverterWarningReasonsFactoryType) New(v uint32) (InverterWarningReasons, error) {
	return InverterWarningReasons(v), nil
}

func (f InverterWarningReasonsFactoryType) NewFieldList(v uint) (FieldList, error) {
	return f.New(uint32(v))
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
