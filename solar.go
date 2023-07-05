package victronDevice

var RegisterListSolarProduct = VictronRegisters{
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateTextRegisterStruct(
		"Product",
		"SerialNumber",
		"Serial number",
		0x010A,
		true,
		302,
	),
	CreateTextRegisterStruct(
		"Product",
		"ModelName",
		"Model name",
		0x010B,
		true,
		303,
	),
	// skipped capabilities
}

var RegisterListSolarGeneric = VictronRegisters{
	CreateEnumRegisterStruct(
		"Generic",
		"DeviceMode",
		"Device mode",
		0x0200,
		false,
		map[int]string{
			0: "Charger off",
			4: "Charger off",
			1: "Charger on",
		},
		400,
	),
	CreateEnumRegisterStruct(
		"Essential",
		"DeviceState",
		"Device state",
		0x0201,
		false,
		map[int]string{
			0:   "Not charging",
			2:   "Fault",
			3:   "Bulk Charging",
			4:   "Absorption Charging",
			5:   "Float Charging",
			7:   "Manual Equalise",
			245: "Wake-Up",
			247: "Auto  Equalise",
			252: "External Control",
			255: "Unavailable",
		},
		6,
	),
	// skipped Remote control used
	// skipped Device off Reason Bitmask
}

var RegisterListSolarSettings = VictronRegisters{
	CreateEnumRegisterStruct(
		"Settings",
		"AdaptiveMode",
		"Adaptive mode",
		0xEDFE,
		true,
		map[int]string{
			0: "off",
			1: "on",
		},
		500,
	),
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateEnumRegisterStruct(
		"Settings",
		"BatteryType",
		"Battery type",
		0xEDF1,
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateEnumRegisterStruct(
		"Settings",
		"BatteryVoltageSetting",
		"Battery voltage setting",
		0xEDEA,
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
	CreateEnumRegisterStruct(
		"Settings",
		"BmsPresent",
		"BMS present",
		0xEDE8,
		false,
		map[int]string{
			0: "no",
			1: "yes",
		},
		512,
	),
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateEnumRegisterStruct(
		"Settings",
		"AutoEqualiseStop",
		"Auto equalise stop on voltage",
		0xEDE5,
		false,
		map[int]string{
			0: "no",
			1: "yes",
		},
		515,
	),
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
}

var RegisterListSolarChargerData = VictronRegisters{
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateEnumRegisterStruct(
		"Generic",
		"ChargerErrorCode",
		"Charger error",
		0xEDDA,
		false,
		map[int]string{
			0:   "No error",
			2:   "Battery voltage too high",
			17:  "Charger internal temperature too high",
			18:  "Charger excessive output current",
			19:  "Charger current polarity reversed",
			20:  "Charger bulk time expired (when 10 hour bulk time protection active)",
			21:  "Charger current sensor issue (bias not within expected limits during off state)",
			26:  "Charger terminals overheated",
			28:  "Converter issue (dual converter models, one of the converters is not working)",
			33:  "Input voltage too high",
			34:  "Input excessive current",
			38:  "Input shutdown (due to excessive battery voltage)",
			39:  "Input shutdown (current flowing while the converter is switched off)",
			66:  "Incompatible device in the network (for synchronized charging)",
			67:  "BMS connection lost",
			68:  "Network misconfigured (e.g. combining ESS with ve.smart networking)",
			116: "Calibration data lost",
			117: "Incompatible firmware (i.e. not for this model)",
			119: "Settings data invalid / corrupted (use restore to defaults and set to recover)",
		},
		401,
	),
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	// skipped voltage setting range
	// skipped history version
	// skipped streetlight version
	// skipped adjustable voltage minimum
	// skipped adjustable voltage maximum
}

var RegisterListSolarPanelData = VictronRegisters{
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateNumberRegisterStruct(
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
	CreateEnumRegisterStruct(
		"Panel",
		"TrackerMode",
		"Tracker mode",
		0xEDB3,
		false,
		map[int]string{
			0: "off",
			1: "voltage/current limited",
			2: "MPP tracker",
		},
		103,
	),
}

var RegisterListSolar = MergeRegisters(
	RegisterListSolarPanelData,
	RegisterListSolarChargerData,
	RegisterListSolarProduct,
	RegisterListSolarGeneric,
	RegisterListSolarSettings,
)
