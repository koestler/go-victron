package veregister

import "github.com/koestler/go-victron/veconst"

// AppendSolar appends all registers of solar chargers to the given RegisterList.
// The list is based on:
// https://www.victronenergy.com/upload/documents/BlueSolar-HEX-protocol.pdf
func AppendSolar(rl *RegisterList) {
	AppendSolarProduct(rl)
	AppendSolarGeneric(rl)
	AppendSolarSettings(rl)
	AppendSolarChargerData(rl)
	AppendSolarPanelData(rl)
	AppendSolarLoadData(rl)
}

// AppendSolarProduct appends all registers of the Product category to the given RegisterList.
func AppendSolarProduct(rl *RegisterList) {
	const catSort = 300
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"ProductId",
			"Product id",
			catSort,
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
			"GroupId",
			"Group id",
			catSort+1,
			0x0104,
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
			catSort+2,
			0x010A,
			true,
			false,
		),
		newTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			catSort+3,
			0x010B,
			true,
			false,
		),
	)
	// skipped capabilities
}

// AppendSolarGeneric appends all registers of the Generic category to the given RegisterList.
func AppendSolarGeneric(rl *RegisterList) {
	const catSort = 400
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Generic",
			"DeviceMode",
			"Device mode",
			catSort,
			0x0200,
			false,
			false,
			veconst.SolarChargerDeviceModeFactory,
		),
		newEnumRegisterStruct(
			"Essential",
			"State",
			"Device state",
			6,
			0x0201,
			false,
			false,
			veconst.SolarChargerStateFactory,
		),
		// skipped Remote control used
	)
	rl.AppendFieldListRegisterStruct(
		newFieldListRegisterStruct(
			"Generic",
			"OffReason",
			"Device off reasons",
			catSort+1,
			0x0207,
			false,
			false,
			veconst.InverterOffReasonsFactory,
		),
	)
}

// AppendSolarSettings appends all registers of the Settings category to the given RegisterList.
func AppendSolarSettings(rl *RegisterList) {
	const catSort = 600
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Settings",
			"AdaptiveMode",
			"Adaptive mode",
			catSort,
			0xEDFE,
			true,
			false,
			veconst.BooleanOffOnFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"AutomaticEqualisationMode",
			"Automatic equalisation mode",
			catSort+1,
			0xEDFD,
			true,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryAbsorptionTimeLimit",
			"Battery absorption time limit",
			catSort+2,
			0xEDFB,
			true,
			false,
			false,
			100,
			0,
			"h",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryAbsorptionVoltage",
			"Battery absorption voltage",
			catSort+3,
			0xEDF7,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryFloatVoltage",
			"Battery float voltage",
			catSort+4,
			0xEDF6,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryEqualisationVoltage",
			"Battery equalisation voltage",
			catSort+5,
			0xEDF4,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryTempCompensation",
			"Battery temperature compensation",
			catSort+6,
			0xEDF2,
			true,
			false,
			true,
			100,
			0,
			"mV/K",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Settings",
			"BatteryType",
			"Battery type",
			catSort+7,
			0xEDF1,
			false,
			false,
			veconst.SolarChargerBatteryTypeFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"BatteryMaximumCurrent",
			"Battery maximum current",
			catSort+8,
			0xEDF0,
			true,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryVoltage",
			"Battery voltage",
			catSort+9,
			0xEDEF,
			true,
			false,
			false,
			1,
			0,
			"V",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Settings",
			"BatteryVoltageSetting",
			"Battery voltage setting",
			catSort+11,
			0xEDEA,
			false,
			false,
			veconst.SolarChargerBatteryVoltageFactory,
		),
		newEnumRegisterStruct(
			"Settings",
			"BmsPresent",
			"BMS present",
			catSort+12,
			0xEDE8,
			false,
			false,
			veconst.BooleanNoYesFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"TailCurrent",
			"Tail current",
			catSort+13,
			0xEDE7,
			true,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Settings",
			"LowTempCurrent",
			"Low temperature charge current",
			catSort+14,
			0xEDE6,
			true,
			false,
			false,
			10,
			0,
			"A",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Settings",
			"AutoEqualiseStop",
			"Auto equalise stop on voltage",
			catSort+15,
			0xEDE5,
			false,
			false,
			veconst.BooleanNoYesFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"EqualisationCurrentLevel",
			"Equalisation current level",
			catSort+16,
			0xEDE4,
			true,
			false,
			false,
			1,
			0,
			"%",
		),
		newNumberRegisterStruct(
			"Settings",
			"EqualisationDuration",
			"Equalisation duration",
			catSort+17,
			0xEDE3,
			true,
			false,
			false,
			100,
			0,
			"h",
		),
		newNumberRegisterStruct(
			"Settings",
			"ReBulkVoltageOffset",
			"Re-bulk voltage offset",
			catSort+18,
			0xED2E,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Settings",
			"BatteryLowTemperatureLevel",
			"Battery low temperature level",
			catSort+19,
			0xEDE0,
			true,
			false,
			true,
			100,
			0,
			"°C",
		),
	)
}

// AppendSolarChargerData appends all registers of the Charger category to the given RegisterList.
func AppendSolarChargerData(rl *RegisterList) {
	const catSort = 200
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Essential",
			"BatteryTemperature",
			"Battery temperature",
			5,
			0xEDEC,
			false,
			false,
			false,
			100,
			-273.15, // unit outputs temp in K
			"°C",
		),
		newNumberRegisterStruct(
			"Charger",
			"ChargerMaximumCurrent",
			"Charger maximum current",
			catSort+1,
			0xEDDF,
			false,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Charger",
			"SystemYield",
			"System yield",
			catSort+2,
			0xEDDD,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
		newNumberRegisterStruct(
			"Charger",
			"SystemYieldResettable",
			"System yield (resettable)",
			catSort+3,
			0xEDDC,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
		newNumberRegisterStruct(
			"Charger",
			"ChargerInternalTemperature",
			"Charger internal temperature",
			catSort+4,
			0xEDDB,
			false,
			false,
			true,
			100,
			0,
			"°C",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Generic",
			"ChargerErrorCode",
			"Charger error",
			401,
			0xEDDA,
			false,
			false,
			veconst.SolarChargerErrorFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Essential",
			"ChargerCurrent",
			"Charger current",
			1,
			0xEDD7,
			false,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Essential",
			"ChargerVoltage",
			"Charger voltage",
			2,
			0xEDD5,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		// skipped Additional charger state info (bitmask)
		newNumberRegisterStruct(
			"Essential",
			"YieldToday",
			"Yield today",
			3,
			0xEDD3,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
		newNumberRegisterStruct(
			"Charger",
			"MaximumPowerToday",
			"Maximum power today",
			catSort+5,
			0xEDD2,
			false,
			false,
			false,
			1,
			0,
			"W",
		),
		newNumberRegisterStruct(
			"Essential",
			"YieldYesterday",
			"Yield yesterday",
			4,
			0xEDD1,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
		newNumberRegisterStruct(
			"Charger",
			"MaximumPowerYesterday",
			"Maximum power yesterday",
			catSort+7,
			0xEDD0,
			false,
			false,
			false,
			1,
			0,
			"W",
		),
	)
	// skipped voltage setting range
	// skipped history version
	// skipped streetlight version
	// skipped adjustable voltage minimum
	// skipped adjustable voltage maximum
}

// AppendSolarPanelData appends all registers of the Panel category to the given RegisterList.
func AppendSolarPanelData(rl *RegisterList) {
	const catSort = 100
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Essential",
			"PanelPower",
			"Panel power",
			0,
			0xEDBC,
			false,
			false,
			false,
			100,
			0,
			"W",
		),
		newNumberRegisterStruct(
			"Panel",
			"PanelVoltage",
			"Panel voltage",
			catSort,
			0xEDBB,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Panel",
			"PanelCurrent",
			"Panel current",
			catSort+1,
			0xEDBD,
			false,
			false,
			false,
			10,
			0,
			"A",
			// The panel current is not available in the 10A/15A/20A chargers
		),
		newNumberRegisterStruct(
			"Panel",
			"PanelMaximumVoltage",
			"Panel maximum voltage",
			catSort+2,
			0xEDB8,
			true,
			false,
			false,
			100,
			0,
			"V",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Panel",
			"TrackerMode",
			"Tracker mode",
			catSort+3,
			0xEDB3,
			false,
			false,
			veconst.SolarChargerTrackerModeFactory,
		),
	)
}

// AppendSolarLoadData appends all registers of the Load category to the given RegisterList.
func AppendSolarLoadData(rl *RegisterList) {
	const catSort = 500
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Load",
			"LoadCurrent",
			"Load Current",
			catSort,
			0xEDAD,
			false,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Load",
			"LoadOffsetVoltage",
			"Load Offset Voltage",
			catSort+1,
			0xEDAC,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Load",
			"LoadOutputVoltage",
			"Load Output Voltage",
			catSort+2,
			0xEDA9,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Load",
			"LoadOutputState",
			"Load Output State",
			catSort+3,
			0xEDA8,
			false,
			false,
			veconst.BooleanOffOnFactory,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Load",
			"LoadSwitchHighLevel",
			"Load Switch High Level",
			catSort+4,
			0xED9D,
			false,
			false,
			false,
			100,
			0,
			"v",
		),
		newNumberRegisterStruct(
			"Load",
			"LoadSwitchLowLevel",
			"Load Switch Low Level",
			catSort+5,
			0xED9C,
			false,
			false,
			false,
			100,
			0,
			"v",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Load",
			"LoadOffReason",
			"Load Off Reason",
			catSort+6,
			0xED91,
			false,
			false,
			veconst.SolarChargerLoadOffReasonFactory,
		),
	)
}
