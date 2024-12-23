package veregister

import (
	"fmt"
	"github.com/koestler/go-victron/veproduct"
)

var ErrUnsupportedType = fmt.Errorf("unsuported product type")

func GetRegisterListByProduct(p veproduct.Product) (rl RegisterList, err error) {
	rl = NewRegisterList()

	switch p.Type() {
	case veproduct.TypeBMV:
		AppendBmv(&rl)
		rl.FilterByName(
			"AuxVoltage",
			"BatteryTemperature",
			"MidPointVoltage",
			"MidPointVoltageDeviation",
			"AuxVoltageMinimum",
			"AuxVoltageMaximum",
		)
	case veproduct.TypeBMVSmart, veproduct.TypeSmartShunt:
		AppendBmv(&rl)
		rl.FilterByName(
			"ProductRevision",
			"Description",
		)
	case veproduct.TypeBlueSolarMPPT, veproduct.TypeSmartSolarMPPT:
		AppendSolar(&rl)

		// The panel current is not available in the 10A/15A/20A chargers
		if c := p.MaxPanelCurrent(); c == 10 || c == 15 || c == 20 {
			rl.FilterByName("PanelCurrent")
		}
	case veproduct.TypePhoenixInverter, veproduct.TypePhoenixInverterSmart:
		AppendInverter(&rl)
	default:
		err = ErrUnsupportedType
	}

	return
}
