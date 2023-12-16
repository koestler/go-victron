package veregisters

import "github.com/koestler/go-victron/victronDefinitions"

// appendInverter appends all registers of phoenix inverters to the given RegisterList.
// The list is based on:
// https://www.victronenergy.com/upload/documents/VE.Direct-HEX-Protocol-Phoenix-Inverter.pdf
func appendInverter(rl *RegisterList) {
	appendInverterProduct(rl)
	appendInverterGeneric(rl)
	appendInverterOffReasons(rl)
	appendInverterWarningReasons(rl)
	appendInverterHistory(rl)
	appendInverterOperation(rl)
	appendInverterAcOutControl(rl)
	appendInverterBatteryControl(rl)
	appendInverterDynamicCutoff(rl)
	// appendInverterRelayControl(rl)
}

func appendInverterProduct(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Product",
			"ProductId",
			"Product id",
			0x0100,
			true,
			false,
			1,
			0,
			"",
			100,
		),
		NewNumberRegisterStruct(
			"Product",
			"ProductRevision",
			"Hardware Revision",
			0x0101,
			true,
			false,
			1,
			0,
			"",
			101,
		),
		NewNumberRegisterStruct(
			"Product",
			"AppVer",
			"Software Revision",
			0x0102,
			true,
			false,
			1,
			0,
			"",
			102,
		),
	)
	rl.appendTextRegisterStruct(
		NewTextRegisterStruct(
			"Product",
			"SerialNumber",
			"Serial number",
			0x010A,
			true,
			103,
		),
		NewTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			0x010B,
			true,
			104,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Product",
			"ACOutRatedPower",
			"AC Out Rated Power",
			0x2203,
			true,
			true,
			1,
			0,
			"VA",
			105,
		),
		// skipping capabilities
		NewNumberRegisterStruct(
			"Product",
			"ACOutNomVoltage",
			"AC Out Nominal Voltage",
			0x2202,
			true,
			false,
			1,
			0,
			"V",
			106,
		),
		NewNumberRegisterStruct(
			"Product",
			"BatVoltage",
			"Battery Voltage",
			0xEDEF,
			true,
			false,
			1,
			0,
			"V",
			106,
		),
	)
}

func appendInverterGeneric(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Essential",
			"DeviceState",
			"Device state",
			0x0201, -1,
			false,
			victronDefinitions.GetInverterStateMap(),
			10,
		),
		// todo: add device off reason, device warning reason and alarm reason (all bit masks)
		NewEnumRegisterStruct(
			"Operation",
			"DeviceMode",
			"Device mode",
			0x0200, -1,
			false,
			victronDefinitions.GetInverterModeMap(),
			300,
		),
	)
}

var inactiveActiveEnum = map[int]string{
	0: "inactive",
	1: "active",
}

func appendInverterOffReasons(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Off Reasons",
			"DeviceOffReasonNoInputPower",
			"No input power",
			0x0207, 0,
			false,
			inactiveActiveEnum,
			200,
		),
		NewEnumRegisterStruct(
			"Off Reasons",
			"DeviceOffReasonPowerButton",
			"Soft power button or SW controller",
			0x0207, 2,
			false,
			inactiveActiveEnum,
			201,
		),
		NewEnumRegisterStruct(
			"Off Reasons",
			"DeviceOffReasonRemoteInput",
			"HW remote input connector",
			0x0207, 3,
			false,
			inactiveActiveEnum,
			202,
		),
		NewEnumRegisterStruct(
			"Off Reasons",
			"DeviceOffReasonInternal",
			"Internal reason (see alarm reason)",
			0x0207, 4,
			false,
			inactiveActiveEnum,
			203,
		),
		NewEnumRegisterStruct(
			"Off Reasons",
			"DeviceOffReasonPayGo",
			"PayGo, out of credit, need token",
			0x0207, 5,
			false,
			inactiveActiveEnum,
			204,
		),
	)
}

func appendInverterWarningReasons(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonLowBatVoltage",
			"Low battery voltage",
			0x031C, 0,
			false,
			inactiveActiveEnum,
			210,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonHighBatVoltage",
			"High battery voltage",
			0x031C, 1,
			false,
			inactiveActiveEnum,
			211,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonLowTemp",
			"Low temperature",
			0x031C, 5,
			false,
			inactiveActiveEnum,
			212,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonHighTemp",
			"High temperature",
			0x031C, 6,
			false,
			inactiveActiveEnum,
			213,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonOverload",
			"Overload",
			0x031C, 8,
			false,
			inactiveActiveEnum,
			214,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonPoorDC",
			"Poor DC connection",
			0x031C, 9,
			false,
			inactiveActiveEnum,
			215,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonLowAcVoltage",
			"Low AC-output voltage",
			0x031C, 10,
			false,
			inactiveActiveEnum,
			216,
		),
		NewEnumRegisterStruct(
			"Warning Reasons",
			"DeviceWarningReasonHighAcVoltage",
			"High AC-output voltage",
			0x031C, 11,
			false,
			inactiveActiveEnum,
			217,
		),
	)
}

func appendInverterHistory(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"History",
			"HistoryTime",
			"Time",
			0x1040,
			false,
			false,
			1,
			0,
			"s",
			310,
		),
		NewNumberRegisterStruct(
			"History",
			"HistoryEnergy",
			"Energy",
			0x1041,
			false,
			false,
			100,
			0,
			"kVAh",
			311,
		),
	)
}

func appendInverterOperation(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Essential",
			"AcOutCurrent",
			"AC Output Current",
			0x2201,
			false,
			true,
			10,
			0,
			"A",
			2,
		),
		NewNumberRegisterStruct(
			"Essential",
			"AcOutVoltage",
			"AC Output Voltage",
			0x2200,
			false,
			true,
			100,
			0,
			"V",
			1,
		),
		/*
			todo: use capabilities register to determine if this is needed (howto test?)
			NewNumberRegisterStruct(
				"Essential",
				"AcOutApparentPower",
				"AC Output Apparent Power",
				0x2205,
				false,
				true,
				1,
				0,
				"VA",
				0,
			),
		*/
		NewNumberRegisterStruct(
			"Operation",
			"InvLoopGetIinv",
			"Inverter Loop get I inv",
			0xEB4E,
			false,
			true,
			1000,
			0,
			"A",
			301,
		),
		NewNumberRegisterStruct(
			"Essential",
			"DcChannel1Voltage",
			"Input Battery Voltage",
			0xED8D,
			false,
			true,
			100,
			0,
			"V",
			3,
		),
	)
}

func appendInverterAcOutControl(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpoint",
			"Voltage Setpoint",
			0x0230,
			true,
			false,
			100,
			0,
			"V",
			400,
		),
		NewNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpointMin",
			"Voltage Setpoint Minimum",
			0x0231,
			true,
			false,
			100,
			0,
			"V",
			401,
		),
		NewNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpointMax",
			"Voltage Setpoint Maximum",
			0x0232,
			true,
			false,
			100,
			0,
			"V",
			402,
		),
		/*
			// todo: activate via capabilities?
			NewNumberRegisterStruct(
				"AC-out settings",
				"AcLoadSensePowerThreshold",
				"Load Sense Power Threshold",
				0x2206,
				true,
				false,
				1,
				0,
				"VA",
				403,
			),
			NewNumberRegisterStruct(
				"AC-out settings",
				"AcLoadSensePowerClear",
				"Load Sense Power Clear",
				0x2207,
				true,
				false,
				1,
				0,
				"VA",
				404,
			),
		*/
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"AC-out settings",
			"InvWaveSet50HzNot60Hz",
			"Frequency",
			0xEB03, -1,
			true,
			map[int]string{
				0: "60 Hz",
				1: "50 Hz",
			},
			405,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoModeInvMin",
			"Inverter Eco Mode Inv Min",
			0xEB04,
			true,
			true,
			1000,
			0,
			"A",
			406,
		),
		NewNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoModeRetryTime",
			"Inverter Eco Mode Retry Time",
			0xEB06,
			true,
			false,
			4,
			0,
			"s",
			407,
		),
		NewNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoLoadDetectPeriods",
			"Inverter Eco Load Detect Periods",
			0xEB10,
			true,
			false,
			1,
			0,
			"",
			408,
		),
	)
}

func appendInverterBatteryControl(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Battery settings",
			"ShutdownLowVoltageSet",
			"Shutdown Low Voltage",
			0x2210,
			true,
			false,
			100,
			0,
			"V",
			500,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"AlarmLowVoltageSet",
			"Alarm Low Voltage Set",
			0x0320,
			true,
			false,
			100,
			0,
			"V",
			501,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"AlarmLowVoltageClear",
			"Alarm Low Voltage Clear",
			0x0321,
			true,
			false,
			100,
			0,
			"V",
			502,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"VoltageRangeMin",
			"Voltage Range Min",
			0x2211,
			true,
			false,
			100,
			0,
			"V",
			503,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"VoltageRangeMax",
			"Voltage Range Max",
			0x2212,
			true,
			false,
			100,
			0,
			"V",
			504,
		),
	)
}

func appendInverterDynamicCutoff(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Dynamic Cutoff",
			"InvProtUbatDynCutoffEnable",
			"Dynamic Cutoff Enable",
			0xEBBA, -1,
			true,
			map[int]string{
				0: "Disabled",
				1: "Enabled",
			},
			600,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor",
			"Factor",
			0xEBB7,
			true,
			false,
			1,
			0,
			"",
			601,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor2000",
			"Factor 2000",
			0xEBB5,
			true,
			false,
			1,
			0,
			"",
			602,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor250",
			"Factor 250",
			0xEBB3,
			true,
			false,
			1,
			0,
			"",
			603,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor5",
			"Factor 5",
			0xEBB2,
			true,
			false,
			1,
			0,
			"",
			604,
		),
		NewNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffVoltage",
			"Voltage",
			0xEBB1,
			true,
			false,
			1000,
			0,
			"V",
			605,
		),
	)
}

/*
func appendSolarPanelData(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		// todo: activate via capabilities?
		NewEnumRegisterStruct(
			"Relay Control",
			"RelayControl",
			"Relay Status",
			0x034E,
			false,
			map[int]string{
				0: "Off: NO=open, NC=closed",
				1: "On: NO=closed, NC=open",
			},
			700,
		),
		NewEnumRegisterStruct(
			"Relay Control",
			"RelayMode",
			"Relay Mode",
			0x034F,
			true,
			map[int]string{
				4: "Normal operation",
				0: "Warnings and alarms",
				5: "Battery low",
				6: "External fan",
				3: "Disabled relay",
				2: "Remote",
			},
			701,
		),
	)
}
*/
