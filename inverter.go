package victronDevice

var RegisterListInverterProduct = []VictronRegister{
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
		"HardwareRevision",
		"Hardware Revision",
		0x0101,
		true,
		false,
		1,
		0,
		"",
		301,
	),
	NewNumberRegisterStruct(
		"Product",
		"Software Revision",
		"Software Revision",
		0x0102,
		true,
		false,
		1,
		0,
		"",
		303,
	),
	NewTextRegisterStruct(
		"Product",
		"SerialNumber",
		"Serial number",
		0x010A,
		true,
		304,
	),
	NewTextRegisterStruct(
		"Product",
		"ModelName",
		"Model name",
		0x010B,
		true,
		305,
	),
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
		306,
	),
	// skipping capabilities
	NewNumberRegisterStruct(
		"Product",
		"ACOutNomVoltage",
		"AC Out Nominal Voltage",
		0x2202,
		false,
		true,
		1,
		0,
		"V",
		307,
	),
}

var RegisterListInverterGeneric = []VictronRegister{
	NewEnumRegisterStruct(
		"Essential",
		"State",
		"Device state",
		0x0201,
		false,
		map[int]string{
			0: "Off",
			1: "Low Power",
			2: "Fault",
			9: "Inverting",
		},
		0,
	),
	// skip bluetooth registers
	NewEnumRegisterStruct(
		"Essential",
		"Mode",
		"Device mode",
		0x0200,
		false,
		map[int]string{
			2:    "Inverter On",
			3:    "Device On",
			4:    "Device Off",
			5:    "Eco mode",
			0xFD: "Hibernate",
		},
		1,
	),
	// todo: add device off reason, device warning reason and alarm reasno (all bit masks)
}

var RegisterListInverter = MergeRegisters(
	RegisterListInverterProduct,
	RegisterListInverterGeneric,
)
