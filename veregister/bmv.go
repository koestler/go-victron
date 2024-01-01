package veregister

import "github.com/koestler/go-victron/veconst"

// AppendBmv appends all registers of BMV units to the given RegisterList.
// The list is based on:
// https://www.victronenergy.com/upload/documents/BMV-7xx-HEX-Protocol.pdf
func AppendBmv(rl *RegisterList) {
	AppendBmvProduct(rl)
	AppendBmvMonitor(rl)
	AppendBmvHistoric(rl)
}

// AppendBmvProduct appends all registers of the Product category to the given RegisterList.
func AppendBmvProduct(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"ProductId",
			"Product id",
			200,
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
			"Product revision",
			201,
			0x0101,
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
			202,
			0x010A,
			true,
			false,
		),
		newTextRegisterStruct(
			"Product",
			"ModelName",
			"Model name",
			203,
			0x010B,
			true,
			false,
		),
		newTextRegisterStruct(
			"Product",
			"Description",
			"Description",
			204,
			0x010C,
			true,
			false,
		),
	)
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Product",
			"Uptime",
			"Device uptime",
			205,
			0x0120,
			false,
			false,
			false,
			1,
			0,
			"s",
		),
	)
	// skipped Bluetooth capabilities
}

// AppendBmvMonitor appends all registers of the Monitor category to the given RegisterList.
func AppendBmvMonitor(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Essential",
			"Power",
			"Power",
			0,
			0xED8E,
			false,
			false,
			true,
			1,
			0,
			"W",
		),
		newNumberRegisterStruct(
			"Essential",
			"CurrentHighRes",
			"Current",
			1,
			0xED8C,
			false,
			false,
			true,
			1000,
			0,
			"A",
		),
		newNumberRegisterStruct(
			"Essential",
			"MainVoltage",
			"Main voltage",
			2,
			0xED8D,
			false,
			false,
			true,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Monitor",
			"AuxVoltage",
			"Aux (starter) voltage",
			100,
			0xED7D,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Monitor",
			"Consumed",
			"Consumed",
			101,
			0xEEFF,
			false,
			false,
			true,
			10,
			0,
			"Ah",
		),
		newNumberRegisterStruct(
			"Essential",
			"SOC",
			"State of charge",
			3,
			0x0FFF,
			false,
			false,
			false,
			100,
			0,
			"%",
		),
		newNumberRegisterStruct(
			"Monitor",
			"TTG",
			"Time to go",
			102,
			0x0FFE,
			false,
			false,
			false,
			1,
			0,
			"min",
		),
		newNumberRegisterStruct(
			"Essential",
			"BatteryTemperature",
			"Battery Temperature",
			4,
			0xEDEC,
			false,
			false,
			false,
			100,
			-273.15, // unit outputs temp in K
			"°C",
		),
		newNumberRegisterStruct(
			"Monitor",
			"MidPointVoltage",
			"Mid-point voltage",
			104,
			0x0382,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Monitor",
			"MidPointVoltageDeviation",
			"Mid-point voltage deviation",
			105,
			0x0383,
			false,
			false,
			true,
			10,
			0,
			"%",
		),
	)
	rl.AppendEnumRegisterStruct(
		newEnumRegisterStruct(
			"Monitor",
			"SynchronizationState",
			"Synchronization state",
			106,
			0xEEB6, -1,
			false,
			false,
			veconst.BooleanFalseTrueFactory,
		),
	)
}

// AppendBmvHistoric appends all registers of the Historic category to the given RegisterList.
func AppendBmvHistoric(rl *RegisterList) {
	rl.AppendNumberRegisterStruct(
		newNumberRegisterStruct(
			"Historic",
			"DepthOfTheDeepestDischarge",
			"Depth of the deepest discharge",
			300,
			0x0300,
			false,
			false,
			true,
			10,
			0,
			"Ah",
		),
		newNumberRegisterStruct(
			"Historic",
			"DepthOfTheLastDischarge",
			"Depth of the last discharge",
			301,
			0x0301,
			false,
			false,
			true,
			10,
			0,
			"Ah",
		),
		newNumberRegisterStruct(
			"Historic",
			"DepthOfTheAverageDischarge",
			"Depth of the average discharge",
			302,
			0x0302,
			false,
			false,
			true,
			10,
			0,
			"Ah",
		),
		newNumberRegisterStruct(
			"Historic",
			"NumberOfCycles",
			"Number of cycles",
			303,
			0x0303,
			false,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Historic",
			"NumberOfFullDischarges",
			"Number of full discharges",
			304,
			0x0304,
			false,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Historic",
			"CumulativeAmpHours",
			"Cumulative amp hours",
			305,
			0x0305,
			false,
			false,
			true,
			10,
			0,
			"Ah",
		),
		newNumberRegisterStruct(
			"Historic",
			"MainVoltageMinimum",
			"Minimum voltage",
			306,
			0x0306,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Historic",
			"MainVoltageMaximum",
			"Maximum voltage",
			307,
			0x0307,
			false,
			false,
			false,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Historic",
			"TimeSinceFullCharge",
			"Time since full charge",
			308,
			0x0308,
			false,
			false,
			false,
			1,
			0,
			"s",
		),
		newNumberRegisterStruct(
			"Historic",
			"NumberOfAutomaticSynchronizations",
			"Number of automatic synchronizations",
			309,
			0x0309,
			false,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Historic",
			"NumberOfLowMainVoltageAlarms",
			"Number of low voltage alarms",
			310,
			0x030A,
			false,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Historic",
			"NumberOfHighMainVoltageAlarms",
			"Number of high voltage alarms",
			311,
			0x030B,
			false,
			false,
			false,
			1,
			0,
			"",
		),
		newNumberRegisterStruct(
			"Historic",
			"AuxVoltageMinimum",
			"Minimum starter voltage",
			312,
			0x030E,
			false,
			false,
			true,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Historic",
			"AuxVoltageMaximum",
			"Maximum starter voltage",
			313,
			0x030F,
			false,
			false,
			true,
			100,
			0,
			"V",
		),
		newNumberRegisterStruct(
			"Historic",
			"AmountOfDischargedEnergy",
			"Amount of discharged energy",
			314,
			0x0310,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
		newNumberRegisterStruct(
			"Historic",
			"AmountOfChargedEnergy",
			"Amount of charged energy",
			315,
			0x0311,
			false,
			false,
			false,
			100,
			0,
			"kWh",
		),
	)
}
