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
		appendBmv(&rl)
		rl.FilterByName(
			"AuxVoltage",
			"BatteryTemperature",
			"MidPointVoltage",
			"MidPointVoltageDeviation",
			"AuxVoltageMinimum",
			"AuxVoltageMaximum",
		)
	case veproduct.TypeBMVSmart, veproduct.TypeSmartShunt:
		appendBmv(&rl)
		rl.FilterByName(
			"ProductRevision",
			"Description",
		)
	case veproduct.TypeBlueSolarMPPT, veproduct.TypeSmartSolarMPPT:
		appendSolar(&rl)
	case veproduct.TypePhoenixInverter, veproduct.TypePhoenixInverterSmart:
		appendInverter(&rl)
	default:
		err = ErrUnsupportedType
	}

	return
}
