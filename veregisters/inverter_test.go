package veregisters

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestAppendInverter(t *testing.T) {
	rl := RegisterList{}
	AppendInverter(&rl)

	expected := []string{
		"Number: category=Essential, name=AcOutVoltage, description=AC Output Voltage, sort=1, address=0x2200, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=V",
		"Number: category=Essential, name=AcOutCurrent, description=AC Output Current, sort=2, address=0x2201, static=false, writable=false, signed=true, factor=10, offset=0.000000, unit=A",
		"Number: category=Essential, name=DcChannel1Voltage, description=Input Battery Voltage, sort=3, address=0xED8D, static=false, writable=false, signed=true, factor=100, offset=0.000000, unit=V",
		"Enum: category=Essential, name=DeviceState, description=Device state, sort=10, address=0x201, static=false, writable=false, bit=-1, enum=map[0:Off 1:Low Power 2:Fault 9:Inverting]",
		"Number: category=Product, name=ProductId, description=Product id, sort=100, address=0x100, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Product, name=ProductRevision, description=Hardware Revision, sort=101, address=0x101, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Product, name=AppVer, description=Software Revision, sort=102, address=0x102, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Text: category=Product, name=SerialNumber, description=Serial number, sort=103, address=0x10A, static=true, writable=false",
		"Text: category=Product, name=ModelName, description=Model name, sort=104, address=0x10B, static=true, writable=false",
		"Number: category=Product, name=ACOutRatedPower, description=AC Out Rated Power, sort=105, address=0x2203, static=true, writable=false, signed=true, factor=1, offset=0.000000, unit=VA",
		"Number: category=Product, name=BatVoltage, description=Battery Voltage, sort=106, address=0xEDEF, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=V",
		"Number: category=Product, name=ACOutNomVoltage, description=AC Out Nominal Voltage, sort=106, address=0x2202, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=V",
		"Enum: category=Off Reasons, name=DeviceOffReasonNoInputPower, description=No input power, sort=200, address=0x207, static=false, writable=false, bit=0, enum=map[0:inactive 1:active]",
		"Enum: category=Off Reasons, name=DeviceOffReasonPowerButton, description=Soft power button or SW controller, sort=201, address=0x207, static=false, writable=false, bit=2, enum=map[0:inactive 1:active]",
		"Enum: category=Off Reasons, name=DeviceOffReasonRemoteInput, description=HW remote input connector, sort=202, address=0x207, static=false, writable=false, bit=3, enum=map[0:inactive 1:active]",
		"Enum: category=Off Reasons, name=DeviceOffReasonInternal, description=Internal reason (see alarm reason), sort=203, address=0x207, static=false, writable=false, bit=4, enum=map[0:inactive 1:active]",
		"Enum: category=Off Reasons, name=DeviceOffReasonPayGo, description=PayGo, out of credit, need token, sort=204, address=0x207, static=false, writable=false, bit=5, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonLowBatVoltage, description=Low battery voltage, sort=210, address=0x31C, static=false, writable=false, bit=0, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonHighBatVoltage, description=High battery voltage, sort=211, address=0x31C, static=false, writable=false, bit=1, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonLowTemp, description=Low temperature, sort=212, address=0x31C, static=false, writable=false, bit=5, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonHighTemp, description=High temperature, sort=213, address=0x31C, static=false, writable=false, bit=6, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonOverload, description=Overload, sort=214, address=0x31C, static=false, writable=false, bit=8, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonPoorDC, description=Poor DC connection, sort=215, address=0x31C, static=false, writable=false, bit=9, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonLowAcVoltage, description=Low AC-output voltage, sort=216, address=0x31C, static=false, writable=false, bit=10, enum=map[0:inactive 1:active]",
		"Enum: category=Warning Reasons, name=DeviceWarningReasonHighAcVoltage, description=High AC-output voltage, sort=217, address=0x31C, static=false, writable=false, bit=11, enum=map[0:inactive 1:active]",
		"Enum: category=Operation, name=DeviceMode, description=Device mode, sort=300, address=0x200, static=false, writable=false, bit=-1, enum=map[2:Inverter On 3:Device On 4:Device Off 5:Eco mode 253:Hibernate]",
		"Number: category=Operation, name=InvLoopGetIinv, description=Inverter Loop get I inv, sort=301, address=0xEB4E, static=false, writable=false, signed=true, factor=1000, offset=0.000000, unit=A",
		"Number: category=History, name=HistoryTime, description=Time, sort=310, address=0x1040, static=false, writable=false, signed=false, factor=1, offset=0.000000, unit=s",
		"Number: category=History, name=HistoryEnergy, description=Energy, sort=311, address=0x1041, static=false, writable=false, signed=false, factor=100, offset=0.000000, unit=kVAh",
		"Number: category=AC-out settings, name=AcOutVoltageSetpoint, description=Voltage Setpoint, sort=400, address=0x230, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=AC-out settings, name=AcOutVoltageSetpointMin, description=Voltage Setpoint Minimum, sort=401, address=0x231, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=AC-out settings, name=AcOutVoltageSetpointMax, description=Voltage Setpoint Maximum, sort=402, address=0x232, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Enum: category=AC-out settings, name=InvWaveSet50HzNot60Hz, description=Frequency, sort=405, address=0xEB03, static=true, writable=false, bit=-1, enum=map[0:60 Hz 1:50 Hz]",
		"Number: category=AC-out settings, name=InvOperEcoModeInvMin, description=Inverter Eco Mode Inv Min, sort=406, address=0xEB04, static=true, writable=false, signed=true, factor=1000, offset=0.000000, unit=A",
		"Number: category=AC-out settings, name=InvOperEcoModeRetryTime, description=Inverter Eco Mode Retry Time, sort=407, address=0xEB06, static=true, writable=false, signed=false, factor=4, offset=0.000000, unit=s",
		"Number: category=AC-out settings, name=InvOperEcoLoadDetectPeriods, description=Inverter Eco Load Detect Periods, sort=408, address=0xEB10, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Battery settings, name=ShutdownLowVoltageSet, description=Shutdown Low Voltage, sort=500, address=0x2210, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Battery settings, name=AlarmLowVoltageSet, description=Alarm Low Voltage Set, sort=501, address=0x320, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Battery settings, name=AlarmLowVoltageClear, description=Alarm Low Voltage Clear, sort=502, address=0x321, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Battery settings, name=VoltageRangeMin, description=Voltage Range Min, sort=503, address=0x2211, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Number: category=Battery settings, name=VoltageRangeMax, description=Voltage Range Max, sort=504, address=0x2212, static=true, writable=false, signed=false, factor=100, offset=0.000000, unit=V",
		"Enum: category=Dynamic Cutoff, name=InvProtUbatDynCutoffEnable, description=Dynamic Cutoff Enable, sort=600, address=0xEBBA, static=true, writable=false, bit=-1, enum=map[0:Disabled 1:Enabled]",
		"Number: category=Battery settings, name=InvProtUbatDynCutoffFactor, description=Factor, sort=601, address=0xEBB7, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Battery settings, name=InvProtUbatDynCutoffFactor2000, description=Factor 2000, sort=602, address=0xEBB5, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Battery settings, name=InvProtUbatDynCutoffFactor250, description=Factor 250, sort=603, address=0xEBB3, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Battery settings, name=InvProtUbatDynCutoffFactor5, description=Factor 5, sort=604, address=0xEBB2, static=true, writable=false, signed=false, factor=1, offset=0.000000, unit=",
		"Number: category=Battery settings, name=InvProtUbatDynCutoffVoltage, description=Voltage, sort=605, address=0xEBB1, static=true, writable=false, signed=false, factor=1000, offset=0.000000, unit=V",
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
