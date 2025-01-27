package veregister

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestAppendSolar(t *testing.T) {
	rl := RegisterList{}
	AppendSolar(&rl)

	expected := []string{
		"Number: category=Essential, name=PanelPower, description=Panel power, sort=0, address=0xEDBC, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=W",
		"Number: category=Essential, name=ChargerCurrent, description=Charger current, sort=1, address=0xEDD7, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Essential, name=ChargerVoltage, description=Charger voltage, sort=2, address=0xEDD5, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Essential, name=YieldToday, description=Yield today, sort=3, address=0xEDD3, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Essential, name=YieldYesterday, description=Yield yesterday, sort=4, address=0xEDD1, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Essential, name=BatteryTemperature, description=Battery temperature, sort=5, address=0xEDEC, static=false, writable=false, signed=false, factor=100, offset=-273.150000, unit=°C",
		"Enum: category=Essential, name=State, description=Device state, sort=6, address=0x201, static=false, writable=false, enum=map[0:Not charging 2:Fault 3:Bulk Charging 4:Absorption Charging 5:Float Charging 7:Manual Equalise 245:Wake-Up 247:Auto Equalise 252:External Control 255:Unavailable]",
		"Number: category=Panel, name=PanelVoltage, description=Panel voltage, sort=100, address=0xEDBB, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Panel, name=PanelCurrent, description=Panel current, sort=101, address=0xEDBD, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Panel, name=PanelMaximumVoltage, description=Panel maximum voltage, sort=102, address=0xEDB8, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Enum: category=Panel, name=TrackerMode, description=Tracker mode, sort=103, address=0xEDB3, static=false, writable=false, enum=map[0:Off 1:Voltage/current limited 2:MPP tracker]",
		"Number: category=Charger, name=ChargerMaximumCurrent, description=Charger maximum current, sort=201, address=0xEDDF, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Charger, name=SystemYield, description=System yield, sort=202, address=0xEDDD, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Charger, name=SystemYieldResettable, description=System yield (resettable), sort=203, address=0xEDDC, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kWh",
		"Number: category=Charger, name=ChargerInternalTemperature, description=Charger internal temperature, sort=204, address=0xEDDB, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=°C",
		"Number: category=Charger, name=MaximumPowerToday, description=Maximum power today, sort=205, address=0xEDD2, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=W",
		"Number: category=Charger, name=MaximumPowerYesterday, description=Maximum power yesterday, sort=207, address=0xEDD0, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=W",
		"Number: category=Product, name=ProductId, description=Product id, sort=300, address=0x100, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Product, name=GroupId, description=Group id, sort=301, address=0x104, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Text: category=Product, name=SerialNumber, description=Serial number, sort=302, address=0x10A, static=true, writable=false",
		"Text: category=Product, name=ModelName, description=Model name, sort=303, address=0x10B, static=true, writable=false",
		"Enum: category=Generic, name=DeviceMode, description=Device mode, sort=400, address=0x200, static=false, writable=false, enum=map[0:Charger Off 1:Charger On 4:Charger Off]",
		"FieldList: category=Generic, name=OffReason, description=Device off reasons, sort=401, address=0x207, static=false, writable=false, fieldList=map[0:No input power 2:Soft power button or SW controller 3:HW remote input connector 4:Internal reason (see alarm reason) 5:PayGo, out of credit, need token]",
		"Enum: category=Generic, name=ChargerErrorCode, description=Charger error, sort=401, address=0xEDDA, static=false, writable=false, enum=map[0:No error 2:Battery voltage too high 17:Charger internal temperature too high 18:Charger excessive output current 19:Charger current polarity reversed 20:Charger bulk time expired (when 10 hour bulk time protection active) 21:Charger current sensor issue (bias not within expected limits during off state) 26:Charger terminals overheated 28:Converter issue (dual converter models, one of the converters is not working) 33:Input voltage too high 34:Input excessive current 38:Input shutdown (due to excessive battery voltage) 39:Input shutdown (current flowing while the converter is switched off) 66:Incompatible device in the network (for synchronized charging) 67:BMS connection lost 68:Network misconfigured (e.g. combining ESS with ve.smart networking) 116:Calibration data lost 117:Incompatible firmware (i.e. not for this model) 119:Settings data invalid / corrupted (use restore to defaults and set to recover) 255:Unknown error]",
		"Number: category=Load, name=LoadCurrent, description=Load Current, sort=500, address=0xEDAD, static=false, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Load, name=LoadOffsetVoltage, description=Load Offset Voltage, sort=501, address=0xEDAC, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Load, name=LoadOutputVoltage, description=Load Output Voltage, sort=502, address=0xEDA9, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Enum: category=Load, name=LoadOutputState, description=Load Output State, sort=503, address=0xEDA8, static=false, writable=false, enum=map[0:Off 1:On]",
		"Number: category=Load, name=LoadSwitchHighLevel, description=Load Switch High Level, sort=504, address=0xED9D, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=v",
		"Number: category=Load, name=LoadSwitchLowLevel, description=Load Switch Low Level, sort=505, address=0xED9C, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=v",
		"Enum: category=Load, name=LoadOffReason, description=Load Off Reason, sort=506, address=0xED91, static=false, writable=false, enum=map[0:Battery low 1:Short circuit 2:Timer program 3:Remote input 4:Pay-as-you-go out of credit 7:Device starting up]",
		"Enum: category=Settings, name=AdaptiveMode, description=Adaptive mode, sort=600, address=0xEDFE, static=true, writable=false, enum=map[0:Off 1:On]",
		"Number: category=Settings, name=AutomaticEqualisationMode, description=Automatic equalisation mode, sort=601, address=0xEDFD, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Settings, name=BatteryAbsorptionTimeLimit, description=Battery absorption time limit, sort=602, address=0xEDFB, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=h",
		"Number: category=Settings, name=BatteryAbsorptionVoltage, description=Battery absorption voltage, sort=603, address=0xEDF7, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryFloatVoltage, description=Battery float voltage, sort=604, address=0xEDF6, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryEqualisationVoltage, description=Battery equalisation voltage, sort=605, address=0xEDF4, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryTempCompensation, description=Battery temperature compensation, sort=606, address=0xEDF2, static=true, writable=false, signed=true, factor=100, offset=0.000000, unit=mV/K",
		"Enum: category=Settings, name=BatteryType, description=Battery type, sort=607, address=0xEDF1, static=false, writable=false, enum=map[1:Gel Victron Long Life (14.1V) 2:Gel Victron Deep discharge (14.3V) 3:Gel Victron Deep discharge (14.4V) 4:AGM Victron Deep discharge (14.7V) 5:Tubular plate cyclic mode 1 (14.9V) 6:Tubular plate cyclic mode 2 (15.1V) 7:Tubular plate cyclic mode 3 (15.3V) 8:LiFEPO4 (14.2V) 255:User defined]",
		"Number: category=Settings, name=BatteryMaximumCurrent, description=Battery maximum current, sort=608, address=0xEDF0, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Settings, name=BatteryVoltage, description=Battery voltage, sort=609, address=0xEDEF, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=V",
		"Enum: category=Settings, name=BatteryVoltageSetting, description=Battery voltage setting, sort=611, address=0xEDEA, static=false, writable=false, enum=map[0:Auto detection at startup 12:12V battery 24:24V battery 36:36V battery 48:48V battery]",
		"Enum: category=Settings, name=BmsPresent, description=BMS present, sort=612, address=0xEDE8, static=false, writable=false, enum=map[0:No 1:Yes]",
		"Number: category=Settings, name=TailCurrent, description=Tail current, sort=613, address=0xEDE7, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Number: category=Settings, name=LowTempCurrent, description=Low temperature charge current, sort=614, address=0xEDE6, static=true, writable=false, signed=false, factor=10, offset=0.000000, unit=A",
		"Enum: category=Settings, name=AutoEqualiseStop, description=Auto equalise stop on voltage, sort=615, address=0xEDE5, static=false, writable=false, enum=map[0:No 1:Yes]",
		"Number: category=Settings, name=EqualisationCurrentLevel, description=Equalisation current level, sort=616, address=0xEDE4, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=%",
		"Number: category=Settings, name=EqualisationDuration, description=Equalisation duration, sort=617, address=0xEDE3, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=h",
		"Number: category=Settings, name=ReBulkVoltageOffset, description=Re-bulk voltage offset, sort=618, address=0xED2E, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Settings, name=BatteryLowTemperatureLevel, description=Battery low temperature level, sort=619, address=0xEDE0, static=true, writable=false, signed=true, factor=100, offset=0.000000, unit=°C",
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
