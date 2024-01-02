package veconst

type InverterOffReason uint8
type InverterOffReasons uint32
type InverterOffReasonsFactoryType struct{}

const (
	InverterOffReasonNoInputPower InverterOffReason = 0
	InverterOffReasonPowerButton  InverterOffReason = 2
	InverterOffReasonRemoteInput  InverterOffReason = 3
	InverterOffReasonInternal     InverterOffReason = 4
	InverterOffReasonPayGo        InverterOffReason = 5
)

var inverterOffReasonMap = map[InverterOffReason]string{
	InverterOffReasonNoInputPower: "No input power",
	InverterOffReasonPowerButton:  "Soft power button or SW controller",
	InverterOffReasonRemoteInput:  "HW remote input connector",
	InverterOffReasonInternal:     "Internal reason (see alarm reason)",
	InverterOffReasonPayGo:        "PayGo, out of credit, need token",
}
var InverterOffReasonsFactory InverterOffReasonsFactoryType

func (f InverterOffReasonsFactoryType) New(v uint32) (InverterOffReasons, error) {
	return InverterOffReasons(v), nil
}

func (f InverterOffReasonsFactoryType) NewFieldList(v uint) (FieldList, error) {
	return f.New(uint32(v))
}

func (f InverterOffReasonsFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(inverterOffReasonMap))
	for k, v := range inverterOffReasonMap {
		ret[int(k)] = v
	}
	return ret
}

func (l InverterOffReasons) Decode() map[InverterOffReason]bool {
	ret := make(map[InverterOffReason]bool, len(inverterOffReasonMap))
	for i := range inverterOffReasonMap {
		ret[i] = l&(1<<i) != 0
	}
	return ret
}

func (l InverterOffReasons) Fields() map[Field]bool {
	m := l.Decode()
	ret := make(map[Field]bool, len(m))
	for i, v := range m {
		ret[i] = v
	}
	return ret
}

func (s InverterOffReason) Idx() int {
	return int(s)
}

func (s InverterOffReason) String() string {
	if v, ok := inverterOffReasonMap[s]; ok {
		return v
	}
	return ""
}
