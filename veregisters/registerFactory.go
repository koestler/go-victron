package veregisters

import (
	"fmt"
	"github.com/koestler/go-victron/veproduct"
)

var ErrUnsupportedType = fmt.Errorf("unsuported product type")

func GetRegisterListByProductType(t veproduct.Type) (rl RegisterList, err error) {
	rl = NewRegisterList()

	switch t {
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
	case veproduct.TypePhoenixInverter, veproduct.TypePhoenixInverterSmart:
		AppendInverter(&rl)
	default:
		err = ErrUnsupportedType
	}

	return
}
