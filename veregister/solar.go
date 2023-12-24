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
}

// AppendSolarProduct appends all registers of the Product category to the given RegisterList.
func AppendSolarProduct(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"ProductId",
			"Product id",
			300,
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
			301,
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
			302,
			0x010A,
			true,
			false,
		),
		newTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			303,
			0x010B,
			true,
			false,
		),
	)
	// skipped capabilities
}

// AppendSolarGeneric appends all registers of the Generic category to the given RegisterList.
func AppendSolarGeneric(rl *RegisterList) {
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Generic",
			"DeviceMode",
			"Device mode",
			400,
			0x0200, -1,
			false,
			false,
			veconst.GetSolarChargerDeviceModeStringMap(),
		),
		newEnumRegisterStruct(
			"Essential",
			"State",
			"Device state",
			6,
			0x0201,
			-1,
			false,
			false,
			veconst.GetSolarChargerStateStringMap(),
		),
		// skipped Remote control used
		// skipped Device off Reason Bitmask
	)
}

// AppendSolarSettings appends all registers of the Settings category to the given RegisterList.
func AppendSolarSettings(rl *RegisterList) {
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Settings",
			"AdaptiveMode",
			"Adaptive mode",
			500,
			0xEDFE, -1,
			true,
			false,
			veconst.GetBooleanOffOnStringMap(),
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"AutomaticEqualisationMode",
			"Automatic equalisation mode",
			501,
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
			502,
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
			503,
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
			504,
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
			505,
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
			506,
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
			507,
			0xEDF1, -1,
			false,
			false,
			veconst.GetSolarChargerBatteryTypeStringMap(),
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"BatteryMaximumCurrent",
			"Battery maximum current",
			508,
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
			509,
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
			511,
			0xEDEA, -1,
			false,
			false,
			veconst.GetSolarChargerBatteryVoltageStringMap(),
		),
		newEnumRegisterStruct(
			"Settings",
			"BmsPresent",
			"BMS present",
			512,
			0xEDE8, -1,
			false,
			false,
			veconst.GetBooleanNoYesStringMap(),
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"TailCurrent",
			"Tail current",
			513,
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
			514,
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
			515,
			0xEDE5, -1,
			false,
			false,
			veconst.GetBooleanNoYesStringMap(),
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Settings",
			"EqualisationCurrentLevel",
			"Equalisation current level",
			516,
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
			517,
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
			518,
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
			519,
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
			201,
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
			202,
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
			203,
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
			204,
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
			0xEDDA, -1,
			false,
			false,
			veconst.GetSolarChargerErrorStringMap(),
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
			205,
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
			207,
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
			100,
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
			101,
			0xEDBD,
			false,
			false,
			false,
			10,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Panel",
			"PanelMaximumVoltage",
			"Panel maximum voltage",
			102,
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
			103,
			0xEDB3, -1,
			false,
			false,
			veconst.GetSolarChargerTrackerModeStringMap(),
		),
	)
}
