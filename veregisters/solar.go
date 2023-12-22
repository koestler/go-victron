package veregisters

import "github.com/koestler/go-victron/victronDefinitions"

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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewTextRegisterStruct(
			"Product",
			"SerialNumber",
			"Serial number",
			302,
			0x010A,
			true,
			false,
		),
		NewTextRegisterStruct(
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
		NewEnumRegisterStruct(
			"Generic",
			"DeviceMode",
			"Device mode",
			400,
			0x0200, -1,
			false,
			false,
			map[int]string{
				0: "Charger off",
				4: "Charger off",
				1: "Charger on",
			},
		),
		NewEnumRegisterStruct(
			"Essential",
			"State",
			"Device state",
			6,
			0x0201,
			-1,
			false,
			false,
			victronDefinitions.GetSolarChargerStateMap(),
		),
		// skipped Remote control used
		// skipped Device off Reason Bitmask
	)
}

// AppendSolarSettings appends all registers of the Settings category to the given RegisterList.
func AppendSolarSettings(rl *RegisterList) {
	rl.AppendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Settings",
			"AdaptiveMode",
			"Adaptive mode",
			500,
			0xEDFE, -1,
			true,
			false,
			map[int]string{
				0: "off",
				1: "on",
			},
		),
	)
	rl.AppendNumberRegisterStruct(
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewEnumRegisterStruct(
			"Settings",
			"BatteryType",
			"Battery type",
			507,
			0xEDF1, -1,
			false,
			false,
			map[int]string{
				1:   "Gel Victron Long Life (14.1V)",
				2:   "Gel Victron Deep discharge (14.3V)",
				3:   "Gel Victron Deep discharge (14.4V)",
				4:   "AGM Victron Deep discharge (14.7V)",
				5:   "Tubular plate cyclic mode 1 (14.9V)",
				6:   "Tubular plate cyclic mode 2 (15.1V)",
				7:   "Tubular plate cyclic mode 3 (15.3V)",
				8:   "LiFEPO4 (14.2V)",
				255: "User defined",
			},
		),
	)
	rl.AppendNumberRegisterStruct(
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewEnumRegisterStruct(
			"Settings",
			"BatteryVoltageSetting",
			"Battery voltage setting",
			511,
			0xEDEA, -1,
			false,
			false,
			map[int]string{
				0:  "Auto detection at startup",
				12: "12V battery",
				24: "24V battery",
				36: "36V battery",
				48: "48V battery",
			},
		),
		NewEnumRegisterStruct(
			"Settings",
			"BmsPresent",
			"BMS present",
			512,
			0xEDE8, -1,
			false,
			false,
			map[int]string{
				0: "no",
				1: "yes",
			},
		),
	)
	rl.AppendNumberRegisterStruct(
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewEnumRegisterStruct(
			"Settings",
			"AutoEqualiseStop",
			"Auto equalise stop on voltage",
			515,
			0xEDE5, -1,
			false,
			false,
			map[int]string{
				0: "no",
				1: "yes",
			},
		),
	)
	rl.AppendNumberRegisterStruct(
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewEnumRegisterStruct(
			"Generic",
			"ChargerErrorCode",
			"Charger error",
			401,
			0xEDDA, -1,
			false,
			false,
			victronDefinitions.GetSolarChargerErrorMap(),
		),
	)
	rl.AppendNumberRegisterStruct(
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewNumberRegisterStruct(
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
		NewEnumRegisterStruct(
			"Panel",
			"TrackerMode",
			"Tracker mode",
			103,
			0xEDB3, -1,
			false,
			false,
			map[int]string{
				0: "off",
				1: "voltage/current limited",
				2: "MPP tracker",
			},
		),
	)
}
