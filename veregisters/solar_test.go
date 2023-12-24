package veregisters

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestAppendSolar(t *testing.T) {
	rl := RegisterList{}
	AppendSolar(&rl)

	expected := []string{
		"Number: category=Essential, name=PanelPower, desription=Panel power, sort=0, address=0xEDBC, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=W",
		"Number: category=Essential, name=ChargerCurrent, desription=Charger current, sort=1, address=0xEDD7, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Essential, name=ChargerVoltage, desription=Charger voltage, sort=2, address=0xEDD5, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Essential, name=YieldToday, desription=Yield today, sort=3, address=0xEDD3, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Essential, name=YieldYesterday, desription=Yield yesterday, sort=4, address=0xEDD1, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Essential, name=BatteryTemperature, desription=Battery temperature, sort=5, address=0xEDEC, static=false, writable=false, signed=false, factor=100, offset=-273.150000, unit=°C",
		"Enum: category=Essential, name=State, desription=Device state, sort=6, address=0x201, static=false, writable=false, bit=-1, enum=map[0:Not charging 2:Fault 3:Bulk Charging 4:Absorption Charging 5:Float Charging 7:Manual Equalise 245:Wake-Up 247:Auto Equalise 252:External Control 255:Unavailable]",
		"Number: category=Panel, name=PanelVoltage, desription=Panel voltage, sort=100, address=0xEDBB, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Panel, name=PanelCurrent, desription=Panel current, sort=101, address=0xEDBD, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Panel, name=PanelMaximumVoltage, desription=Panel maximum voltage, sort=102, address=0xEDB8, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Enum: category=Panel, name=TrackerMode, desription=Tracker mode, sort=103, address=0xEDB3, static=false, writable=false, bit=-1, enum=map[0:off 1:voltage/current limited 2:MPP tracker]",
		"Number: category=Charger, name=ChargerMaximumCurrent, desription=Charger maximum current, sort=201, address=0xEDDF, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Charger, name=SystemYield, desription=System yield, sort=202, address=0xEDDD, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Charger, name=SystemYieldResettable, desription=System yield (resettable), sort=203, address=0xEDDC, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Charger, name=ChargerInternalTemperature, desription=Charger internal temperature, sort=204, address=0xEDDB, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=°C",
		"Number: category=Charger, name=MaximumPowerToday, desription=Maximum power today, sort=205, address=0xEDD2, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=W",
		"Number: category=Charger, name=MaximumPowerYesterday, desription=Maximum power yesterday, sort=207, address=0xEDD0, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=W",
		"Number: category=Product, name=ProductId, desription=Product id, sort=300, address=0x100, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Product, name=GroupId, desription=Group id, sort=301, address=0x104, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Text: category=Product, name=SerialNumber, desription=Serial number, sort=302, address=0x10A, static=true, writable=false",
		"Text: category=Product, name=ModelName, desription=Model name, sort=303, address=0x10B, static=true, writable=false",
		"Enum: category=Generic, name=DeviceMode, desription=Device mode, sort=400, address=0x200, static=false, writable=false, bit=-1, enum=map[0:Charger off 1:Charger on 4:Charger off]",
		"Enum: category=Generic, name=ChargerErrorCode, desription=Charger error, sort=401, address=0xEDDA, static=false, writable=false, bit=-1, enum=map[0:No error 2:Battery voltage too high 17:Charger internal temperature too high 18:Charger excessive output current 19:Charger current polarity reversed 20:Charger bulk time expired (when 10 hour bulk time protection active) 21:Charger current sensor issue (bias not within expected limits during off state) 26:Charger terminals overheated 28:Converter issue (dual converter models, one of the converters is not working) 33:Input voltage too high 34:Input excessive current 38:Input shutdown (due to excessive battery voltage) 39:Input shutdown (current flowing while the converter is switched off) 66:Incompatible device in the network (for synchronized charging) 67:BMS connection lost 68:Network misconfigured (e.g. combining ESS with ve.smart networking) 116:Calibration data lost 117:Incompatible firmware (i.e. not for this model) 119:Settings data invalid / corrupted (use restore to defaults and set to recover) 255:Unknown error]",
		"Enum: category=Settings, name=AdaptiveMode, desription=Adaptive mode, sort=500, address=0xEDFE, static=true, writable=false, bit=-1, enum=map[0:off 1:on]",
		"Number: category=Settings, name=AutomaticEqualisationMode, desription=Automatic equalisation mode, sort=501, address=0xEDFD, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Settings, name=BatteryAbsorptionTimeLimit, desription=Battery absorption time limit, sort=502, address=0xEDFB, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=h",
		"Number: category=Settings, name=BatteryAbsorptionVoltage, desription=Battery absorption voltage, sort=503, address=0xEDF7, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryFloatVoltage, desription=Battery float voltage, sort=504, address=0xEDF6, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryEqualisationVoltage, desription=Battery equalisation voltage, sort=505, address=0xEDF4, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryTempCompensation, desription=Battery temperature compensation, sort=506, address=0xEDF2, static=true, writable=false, signed=true, factor=100, offset=0.000000, unit=mV/K",
		"Enum: category=Settings, name=BatteryType, desription=Battery type, sort=507, address=0xEDF1, static=false, writable=false, bit=-1, enum=map[1:Gel Victron Long Life (14.1V) 2:Gel Victron Deep discharge (14.3V) 3:Gel Victron Deep discharge (14.4V) 4:AGM Victron Deep discharge (14.7V) 5:Tubular plate cyclic mode 1 (14.9V) 6:Tubular plate cyclic mode 2 (15.1V) 7:Tubular plate cyclic mode 3 (15.3V) 8:LiFEPO4 (14.2V) 255:User defined]",
		"Number: category=Settings, name=BatteryMaximumCurrent, desription=Battery maximum current, sort=508, address=0xEDF0, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Settings, name=BatteryVoltage, desription=Battery voltage, sort=509, address=0xEDEF, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=V",
		"Enum: category=Settings, name=BatteryVoltageSetting, desription=Battery voltage setting, sort=511, address=0xEDEA, static=false, writable=false, bit=-1, enum=map[0:Auto detection at startup 12:12V battery 24:24V battery 36:36V battery 48:48V battery]",
		"Enum: category=Settings, name=BmsPresent, desription=BMS present, sort=512, address=0xEDE8, static=false, writable=false, bit=-1, enum=map[0:no 1:yes]",
		"Number: category=Settings, name=TailCurrent, desription=Tail current, sort=513, address=0xEDE7, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Settings, name=LowTempCurrent, desription=Low temperature charge current, sort=514, address=0xEDE6, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Enum: category=Settings, name=AutoEqualiseStop, desription=Auto equalise stop on voltage, sort=515, address=0xEDE5, static=false, writable=false, bit=-1, enum=map[0:no 1:yes]",
		"Number: category=Settings, name=EqualisationCurrentLevel, desription=Equalisation current level, sort=516, address=0xEDE4, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=%",
		"Number: category=Settings, name=EqualisationDuration, desription=Equalisation duration, sort=517, address=0xEDE3, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=h",
		"Number: category=Settings, name=ReBulkVoltageOffset, desription=Re-bulk voltage offset, sort=518, address=0xED2E, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryLowTemperatureLevel, desription=Battery low temperature level, sort=519, address=0xEDE0, static=true, writable=false, signed=true, factor=100, offset=0.000000, unit=°C",
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
