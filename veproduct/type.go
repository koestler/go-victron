package veproduct

type Type uint8

const (
	TypeUnknown                 Type = iota
	TypeBMV                     Type = iota
	TypeBMVSmart                Type = iota
	TypeBlueSolarMPPT           Type = iota
	TypeSmartSolarMPPT          Type = iota
	TypeBlueSolarMPPTVECan      Type = iota
	TypeSmartSolarMPPTVECan     Type = iota
	TypePhoenixInverter         Type = iota
	TypePhoenixInverterSmart    Type = iota
	TypePhoenixSmartIP43Charger Type = iota
	TypeSmartShunt              Type = iota
)

func (t Type) String() string {
	switch t {
	case TypeBMV:
		return "BMV"
	case TypeBMVSmart:
		return "BMV Smart"
	case TypeBlueSolarMPPT:
		return "BlueSolar MPPT"
	case TypeSmartSolarMPPT:
		return "SmartSolar MPPT"
	case TypeBlueSolarMPPTVECan:
		return "BlueSolar MPPT VE.Can"
	case TypeSmartSolarMPPTVECan:
		return "SmartSolar MPPT VE.Can"
	case TypePhoenixInverter:
		return "Phoenix Inverter"
	case TypePhoenixInverterSmart:
		return "Phoenix Inverter Smart"
	case TypePhoenixSmartIP43Charger:
		return "Phoenix Smart IP43 Charger"
	case TypeSmartShunt:
		return "SmartShunt"
	default:
		return ""
	}
}
