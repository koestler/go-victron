package victronDevice

import (
	"github.com/koestler/go-victron/veproduct"
)

func RegisterFactoryByProduct(product veproduct.Product) []VictronRegister {
	switch product.Type() {
	case veproduct.TypeBMV:
		return RegisterListBmv700
	case veproduct.TypeBMVSmart:
		return RegisterListBmv712
	case veproduct.TypeBlueSolarMPPT:
		return RegisterListSolar
	case veproduct.TypeSmartSolarMPPT:
		return RegisterListSolar
	case veproduct.TypePhoenixInverter:
		return RegisterListInverter
	case veproduct.TypePhoenixInverterSmart:
		return RegisterListInverter
	case veproduct.TypeSmartShunt:
		return RegisterListBmv712
	default:
		return nil
	}
}
