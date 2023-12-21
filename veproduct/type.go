package veproduct

// Type defines the product line. It is the sort / purpose of the product.
type Type uint8

const (
	TypeUnknown Type = iota
	TypeBMV
	TypeBMVSmart
	TypeBlueSolarMPPT
	TypeSmartSolarMPPT
	TypeBlueSolarMPPTVECan
	TypeSmartSolarMPPTVECan
	TypePhoenixInverter
	TypePhoenixInverterSmart
	TypePhoenixSmartIP43Charger
	TypeSmartShunt
)

// String returns a human-readable string for the product type (e.g. "SmartSolar MPPT").
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
