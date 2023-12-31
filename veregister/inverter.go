package veregister

import "github.com/koestler/go-victron/veconst"

// AppendInverter appends all registers of phoenix inverters to the given RegisterList.
// The list is based on:
// https://www.victronenergy.com/upload/documents/VE.Direct-HEX-Protocol-Phoenix-Inverter.pdf
func AppendInverter(rl *RegisterList) {
	AppendInverterProduct(rl)
	AppendInverterGeneric(rl)
	AppendInverterHistory(rl)
	AppendInverterOperation(rl)
	AppendInverterAcOutControl(rl)
	AppendInverterBatteryControl(rl)
	AppendInverterDynamicCutoff(rl)
	// appendInverterRelayControl(rl)
}

// AppendInverterProduct appends all registers of the Product category to the given RegisterList.
func AppendInverterProduct(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"ProductId",
			"Product id",
			100,
			0x0100,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Product",
			"ProductRevision",
			"Hardware Revision",
			101,
			0x0101,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Product",
			"AppVer",
			"Software Revision",
			102,
			0x0102,
			true,
			false,
			false,
			1,
			0,
			"",
		),
	)
	rl.AppendTextRegisterStruct(
		newTextRegisterStruct(
			"Product",
			"SerialNumber",
			"Serial number",
			103,
			0x010A,
			true,
			false,
		),
		newTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			104,
			0x010B,
			true,
			false,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"ACOutRatedPower",
			"AC Out Rated Power",
			105,
			0x2203,
			true,
			false,
			true,
			1,
			0,
			"VA",
		),
		// skipping capabilities
		newNumberRegisterStruct(
			"Product",
			"ACOutNomVoltage",
			"AC Out Nominal Voltage",
			106,
			0x2202,
			true,
			false,
			false,
			1,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Product",
			"BatVoltage",
			"Battery Voltage",
			107,
			0xEDEF,
			true,
			false,
			false,
			1,
			0,
			"V",
		),
	)
}

// AppendInverterGeneric appends all registers of the Generic category to the given RegisterList.
func AppendInverterGeneric(rl *RegisterList) {
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Essential",
			"DeviceState",
			"Device state",
			10,
			0x0201,
			false,
			false,
			veconst.InverterStateFactory,
		),
		newEnumRegisterStruct(
			"Operation",
			"DeviceMode",
			"Device mode",
			300,
			0x0200,
			false,
			false,
			veconst.InverterModeFactory,
		),
	)
	rl.AppendFieldListRegisterStruct(
		newFieldListRegisterStruct(
			"Operation",
			"OffReason",
			"Device off reasons",
			200,
			0x0207,
			false,
			false,
			veconst.InverterOffReasonsFactory,
		),
		newFieldListRegisterStruct(
			"Operation",
			"WarningReason",
			"Warning reasons",
			210,
			0x031C,
			false,
			false,
			veconst.InverterWarningReasonFactory,
		),
	)
}

// AppendInverterHistory appends all registers of the History category to the given RegisterList.
func AppendInverterHistory(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"History",
			"HistoryTime",
			"Time",
			310,
			0x1040,
			false,
			false,
			false,
			1,
			0,
			"s",
		),
		newNumberRegisterStruct(
			"History",
			"HistoryEnergy",
			"Energy",
			311,
			0x1041,
			false,
			false,
			false,
			100,
			0,
			"kVAh",
		),
	)
}

// AppendInverterOperation appends all registers of the Operation category to the given RegisterList.
func AppendInverterOperation(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Essential",
			"AcOutCurrent",
			"AC Output Current",
			2,
			0x2201,
			false,
			false,
			true,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Essential",
			"AcOutVoltage",
			"AC Output Voltage",
			1,
			0x2200,
			false,
			false,
			true,
			100,
			0,
			"V",
		),
		/*
			todo: use capabilities register to determine if this is needed (howto test?)
			newNumberRegisterStruct(
				"Essential",
				"AcOutApparentPower",
				"AC Output Apparent Power",
				0,
				0x2205,
				false,
				true,
				1,
				0,
				"VA",
			),
		*/
		newNumberRegisterStruct(
			"Operation",
			"InvLoopGetIinv",
			"Inverter Loop get I inv",
			301,
			0xEB4E,
			false,
			false,
			true,
			1000,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Essential",
			"DcChannel1Voltage",
			"Input Battery Voltage",
			3,
			0xED8D,
			false,
			false,
			true,
			100,
			0,
			"V",
		),
	)
}

// AppendInverterAcOutControl appends all registers of the AC-out settings category to the given RegisterList.
func AppendInverterAcOutControl(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpoint",
			"Voltage Setpoint",
			400,
			0x0230,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpointMin",
			"Voltage Setpoint Minimum",
			401,
			0x0231,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"AC-out settings",
			"AcOutVoltageSetpointMax",
			"Voltage Setpoint Maximum",
			402,
			0x0232,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		/*
			// todo: activate via capabilities?
			newNumberRegisterStruct(
				"AC-out settings",
				"AcLoadSensePowerThreshold",
				"Load Sense Power Threshold",
				403,
				0x2206,
				true,
				false,
				1,
				0,
				"VA",
			),
			newNumberRegisterStruct(
				"AC-out settings",
				"AcLoadSensePowerClear",
				"Load Sense Power Clear",
				404,
				0x2207,
				true,
				false,
				1,
				0,
				"VA",
			),
		*/
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"AC-out settings",
			"InvWaveSet50HzNot60Hz",
			"Frequency",
			405,
			0xEB03,
			true,
			false,
			veconst.InverterFrequencyFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoModeInvMin",
			"Inverter Eco Mode Inv Min",
			406,
			0xEB04,
			true,
			false,
			true,
			1000,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoModeRetryTime",
			"Inverter Eco Mode Retry Time",
			407,
			0xEB06,
			true,
			false,
			false,
			4,
			0,
			"s",
		),
		newNumberRegisterStruct(
			"AC-out settings",
			"InvOperEcoLoadDetectPeriods",
			"Inverter Eco Load Detect Periods",
			408,
			0xEB10,
			true,
			false,
			false,
			1,
			0,
			"",
		),
	)
}

// AppendInverterBatteryControl appends all registers of the Battery settings category to the given RegisterList.
func AppendInverterBatteryControl(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Battery settings",
			"ShutdownLowVoltageSet",
			"Shutdown Low Voltage",
			500,
			0x2210,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"AlarmLowVoltageSet",
			"Alarm Low Voltage Set",
			501,
			0x0320,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"AlarmLowVoltageClear",
			"Alarm Low Voltage Clear",
			502,
			0x0321,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"VoltageRangeMin",
			"Voltage Range Min",
			503,
			0x2211,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"VoltageRangeMax",
			"Voltage Range Max",
			504,
			0x2212,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
	)
}

// AppendInverterDynamicCutoff appends all registers of the Dynamic Cutoff category to the given RegisterList.
func AppendInverterDynamicCutoff(rl *RegisterList) {
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Dynamic Cutoff",
			"InvProtUbatDynCutoffEnable",
			"Dynamic Cutoff Enable",
			600,
			0xEBBA,
			true,
			false,
			veconst.BooleanDisabledEnabledFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor",
			"Factor",
			601,
			0xEBB7,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor2000",
			"Factor 2000",
			602,
			0xEBB5,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor250",
			"Factor 250",
			603,
			0xEBB3,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffFactor5",
			"Factor 5",
			604,
			0xEBB2,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Battery settings",
			"InvProtUbatDynCutoffVoltage",
			"Voltage",
			605,
			0xEBB1,
			true,
			false,
			false,
			1000,
			0,
			"V",
		),
	)
}
