package veregisters

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestAppendBmv(t *testing.T) {
	rl := RegisterList{}
	AppendBmv(&rl)

	expected := []string{
		"Number: category=Essential, name=Power, desription=Power, sort=0, address=0xED8E, static=false, writable=false, signed=true, factor=1, offset=0.000000, unit=W",
		"Number: category=Essential, name=CurrentHighRes, desription=Current, sort=1, address=0xED8C, static=false, writable=false, signed=true, factor=1000, offset=0.000000, unit=A",
		"Number: category=Essential, name=MainVoltage, desription=Main voltage, sort=2, address=0xED8D, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=V",
		"Number: category=Essential, name=SOC, desription=State of charge, sort=3, address=0xFFF, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=%",
		"Number: category=Essential, name=BatteryTemperature, desription=Battery Temperature, sort=4, address=0xEDEC, static=false, writable=false, signed=false, factor=100, offset=-273.150000, unit=°C",
		"Number: category=Monitor, name=AuxVoltage, desription=Aux (starter) voltage, sort=100, address=0xED7D, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Monitor, name=Consumed, desription=Consumed, sort=101, address=0xEEFF, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=Ah",
		"Number: category=Monitor, name=TTG, desription=Time to go, sort=102, address=0xFFE, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=min",
		"Number: category=Monitor, name=MidPointVoltage, desription=Mid-point voltage, sort=104, address=0x382, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Monitor, name=MidPointVoltageDeviation, desription=Mid-point voltage deviation, sort=105, address=0x383, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=%",
		"Enum: category=Monitor, name=SynchronizationState, desription=Synchronization state, sort=106, address=0xEEB6, static=false, writable=false, bit=-1, enum=map[0:false 1:true]",
		"Number: category=Product, name=ProductId, desription=Product id, sort=200, address=0x100, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Product, name=ProductRevision, desription=Product revision, sort=201, address=0x101, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Text: category=Product, name=SerialNumber, desription=Serial number, sort=202, address=0x10A, static=true, writable=false",
		"Text: category=Product, name=ModelName, desription=Model name, sort=203, address=0x10B, static=true, writable=false",
		"Text: category=Product, name=Description, desription=Description, sort=204, address=0x10C, static=true, writable=false",
		"Number: category=Product, name=Uptime, desription=Device uptime, sort=205, address=0x120, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=s",
		"Number: category=Historic, name=DepthOfTheDeepestDischarge, desription=Depth of the deepest discharge, sort=300, address=0x300, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=Ah",
		"Number: category=Historic, name=DepthOfTheLastDischarge, desription=Depth of the last discharge, sort=301, address=0x301, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=Ah",
		"Number: category=Historic, name=DepthOfTheAverageDischarge, desription=Depth of the average discharge, sort=302, address=0x302, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=Ah",
		"Number: category=Historic, name=NumberOfCycles, desription=Number of cycles, sort=303, address=0x303, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Historic, name=NumberOfFullDischarges, desription=Number of full discharges, sort=304, address=0x304, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Historic, name=CumulativeAmpHours, desription=Cumulative amp hours, sort=305, address=0x305, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=Ah",
		"Number: category=Historic, name=MainVoltageMinimum, desription=Minimum voltage, sort=306, address=0x306, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Historic, name=MainVoltageMaximum, desription=Maximum voltage, sort=307, address=0x307, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Historic, name=TimeSinceFullCharge, desription=Time since full charge, sort=308, address=0x308, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=s",
		"Number: category=Historic, name=NumberOfAutomaticSynchronizations, desription=Number of automatic synchronizations, sort=309, address=0x309, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Historic, name=NumberOfLowMainVoltageAlarms, desription=Number of low voltage alarms, sort=310, address=0x30A, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Historic, name=NumberOfHighMainVoltageAlarms, desription=Number of high voltage alarms, sort=311, address=0x30B, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Historic, name=AuxVoltageMinimum, desription=Minimum starter voltage, sort=312, address=0x30E, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=V",
		"Number: category=Historic, name=AuxVoltageMaximum, desription=Maximum starter voltage, sort=313, address=0x30F, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=V",
		"Number: category=Historic, name=AmountOfDischargedEnergy, desription=Amount of discharged energy, sort=314, address=0x310, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Historic, name=AmountOfChargedEnergy, desription=Amount of charged energy, sort=315, address=0x311, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
	}

	got := rl.testStrings()
	if diff := cmp.Diff(got, expected); len(diff) > 0 {
		t.Errorf("register list does not match :\nexpected:\n%s\n\ngot:\n%s\n\ndiff:\n%s",
			strings.Join(expected, "\n"),
			strings.Join(got, "\n"),
			diff,
		)
	}
}
