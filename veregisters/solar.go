package veregisters

import "github.com/koestler/go-victron/victronDefinitions"

func appendSolarProduct(rl *RegisterList) {
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
			300,
		),
		NewNumberRegisterStruct(
			"Product",
			"GroupId",
			"Group id",
			0x0104,
			true,
			false,
			1,
			0,
			"",
			301,
		),
	)
	rl.appendTextRegisterStruct(
		NewTextRegisterStruct(
			"Product",
			"SerialNumber",
			"Serial number",
			0x010A,
			true,
			302,
		),
		NewTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			0x010B,
			true,
			303,
		),
	)
	// skipped capabilities
}

func appendSolarGeneric(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Generic",
			"DeviceMode",
			"Device mode",
			0x0200, -1,
			false,
			map[int]string{
				0: "Charger off",
				4: "Charger off",
				1: "Charger on",
			},
			400,
		),
		NewEnumRegisterStruct(
			"Essential",
			"State",
			"Device state",
			0x0201, -1,
			false,
			victronDefinitions.GetSolarChargerStateMap(),
			6,
		),
		// skipped Remote control used
		// skipped Device off Reason Bitmask
	)
}

func appendSolarSettings(rl *RegisterList) {
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Settings",
			"AdaptiveMode",
			"Adaptive mode",
			0xEDFE, -1,
			true,
			map[int]string{
				0: "off",
				1: "on",
			},
			500,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Settings",
			"AutomaticEqualisationMode",
			"Automatic equalisation mode",
			0xEDFD,
			true,
			false,
			1,
			0,
			"",
			501,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryAbsorptionTimeLimit",
			"Battery absorption time limit",
			0xEDFB,
			true,
			false,
			100,
			0,
			"h",
			502,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryAbsorptionVoltage",
			"Battery absorption voltage",
			0xEDF7,
			true,
			false,
			100,
			0,
			"V",
			503,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryFloatVoltage",
			"Battery float voltage",
			0xEDF6,
			true,
			false,
			100,
			0,
			"V",
			504,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryEqualisationVoltage",
			"Battery equalisation voltage",
			0xEDF4,
			true,
			false,
			100,
			0,
			"V",
			505,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryTempCompensation",
			"Battery temperature compensation",
			0xEDF2,
			true,
			true,
			100,
			0,
			"mV/K",
			506,
		),
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Settings",
			"BatteryType",
			"Battery type",
			0xEDF1, -1,
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
			507,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Settings",
			"BatteryMaximumCurrent",
			"Battery maximum current",
			0xEDF0,
			true,
			false,
			10,
			0,
			"A",
			508,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryVoltage",
			"Battery voltage",
			0xEDEF,
			true,
			false,
			1,
			0,
			"V",
			509,
		),
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Settings",
			"BatteryVoltageSetting",
			"Battery voltage setting",
			0xEDEA, -1,
			false,
			map[int]string{
				0:  "Auto detection at startup",
				12: "12V battery",
				24: "24V battery",
				36: "36V battery",
				48: "48V battery",
			},
			511,
		),
		NewEnumRegisterStruct(
			"Settings",
			"BmsPresent",
			"BMS present",
			0xEDE8, -1,
			false,
			map[int]string{
				0: "no",
				1: "yes",
			},
			512,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Settings",
			"TailCurrent",
			"Tail current",
			0xEDE7,
			true,
			false,
			10,
			0,
			"A",
			513,
		),
		NewNumberRegisterStruct(
			"Settings",
			"LowTempCurrent",
			"Low temperature charge current",
			0xEDE6,
			true,
			false,
			10,
			0,
			"A",
			514,
		),
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Settings",
			"AutoEqualiseStop",
			"Auto equalise stop on voltage",
			0xEDE5, -1,
			false,
			map[int]string{
				0: "no",
				1: "yes",
			},
			515,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Settings",
			"EqualisationCurrentLevel",
			"Equalisation current level",
			0xEDE4,
			true,
			false,
			1,
			0,
			"%",
			516,
		),
		NewNumberRegisterStruct(
			"Settings",
			"EqualisationDuration",
			"Equalisation duration",
			0xEDE3,
			true,
			false,
			100,
			0,
			"h",
			517,
		),
		NewNumberRegisterStruct(
			"Settings",
			"ReBulkVoltageOffset",
			"Re-bulk voltage offset",
			0xED2E,
			true,
			false,
			100,
			0,
			"V",
			518,
		),
		NewNumberRegisterStruct(
			"Settings",
			"BatteryLowTemperatureLevel",
			"Battery low temperature level",
			0xEDE0,
			true,
			true,
			100,
			0,
			"°C",
			519,
		),
	)
}

func appendSolarChargerData(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Essential",
			"BatteryTemperature",
			"Battery temperature",
			0xEDEC,
			false,
			false,
			100,
			-273.15, // unit outputs temp in K
			"°C",
			5,
		),
		NewNumberRegisterStruct(
			"Charger",
			"ChargerMaximumCurrent",
			"Charger maximum current",
			0xEDDF,
			false,
			false,
			10,
			0,
			"A",
			201,
		),
		NewNumberRegisterStruct(
			"Charger",
			"SystemYield",
			"System yield",
			0xEDDD,
			false,
			false,
			100,
			0,
			"kWh",
			202,
		),
		NewNumberRegisterStruct(
			"Charger",
			"SystemYieldResettable",
			"System yield (resettable)",
			0xEDDC,
			false,
			false,
			100,
			0,
			"kWh",
			203,
		),
		NewNumberRegisterStruct(
			"Charger",
			"ChargerInternalTemperature",
			"Charger internal temperature",
			0xEDDB,
			false,
			true,
			100,
			0,
			"°C",
			204,
		),
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Generic",
			"ChargerErrorCode",
			"Charger error",
			0xEDDA, -1,
			false,
			victronDefinitions.GetSolarChargerErrorMap(),
			401,
		),
	)
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Essential",
			"ChargerCurrent",
			"Charger current",
			0xEDD7,
			false,
			false,
			10,
			0,
			"A",
			1,
		),
		NewNumberRegisterStruct(
			"Essential",
			"ChargerVoltage",
			"Charger voltage",
			0xEDD5,
			false,
			false,
			100,
			0,
			"V",
			2,
		),
		// skipped Additional charger state info (bitmask)
		NewNumberRegisterStruct(
			"Essential",
			"YieldToday",
			"Yield today",
			0xEDD3,
			false,
			false,
			100,
			0,
			"kWh",
			3,
		),
		NewNumberRegisterStruct(
			"Charger",
			"MaximumPowerToday",
			"Maximum power today",
			0xEDD2,
			false,
			false,
			1,
			0,
			"W",
			205,
		),
		NewNumberRegisterStruct(
			"Essential",
			"YieldYesterday",
			"Yield yesterday",
			0xEDD1,
			false,
			false,
			100,
			0,
			"kWh",
			4,
		),
		NewNumberRegisterStruct(
			"Charger",
			"MaximumPowerYesterday",
			"Maximum power yesterday",
			0xEDD0,
			false,
			false,
			1,
			0,
			"W",
			207,
		),
	)
	// skipped voltage setting range
	// skipped history version
	// skipped streetlight version
	// skipped adjustable voltage minimum
	// skipped adjustable voltage maximum
}

func appendSolarPanelData(rl *RegisterList) {
	rl.appendNumberRegisterStruct(
		NewNumberRegisterStruct(
			"Essential",
			"PanelPower",
			"Panel power",
			0xEDBC,
			false,
			false,
			100,
			0,
			"W",
			0,
		),
		NewNumberRegisterStruct(
			"Panel",
			"PanelVoltage",
			"Panel voltage",
			0xEDBB,
			false,
			false,
			100,
			0,
			"V",
			100,
		),
		NewNumberRegisterStruct(
			"Panel",
			"PanelCurrent",
			"Panel current",
			0xEDBD,
			false,
			false,
			10,
			0,
			"A",
			101,
		),
		NewNumberRegisterStruct(
			"Panel",
			"PanelMaximumVoltage",
			"Panel maximum voltage",
			0xEDB8,
			true,
			false,
			100,
			0,
			"V",
			102,
		),
	)
	rl.appendEnumRegisterStruct(
		NewEnumRegisterStruct(
			"Panel",
			"TrackerMode",
			"Tracker mode",
			0xEDB3, -1,
			false,
			map[int]string{
				0: "off",
				1: "voltage/current limited",
				2: "MPP tracker",
			},
			103,
		),
	)
}

func appendSolar(rl *RegisterList) {
	appendSolarProduct(rl)
	appendSolarGeneric(rl)
	appendSolarSettings(rl)
	appendSolarChargerData(rl)
	appendSolarPanelData(rl)
}
