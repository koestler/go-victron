package victronDevice

import (
	"github.com/koestler/go-iotdevice/dataflow"
)

var RegisterListBmvProduct = dataflow.Registers{
	dataflow.CreateNumberRegisterStruct(
		"Product",
		"ProductId",
		"Product id",
		0x0100,
		true,
		false,
		1,
		"",
		200,
	),
	dataflow.CreateNumberRegisterStruct(
		"Product",
		"ProductRevision",
		"Product revision",
		0x0101,
		true,
		false,
		1,
		"",
		201,
	),
	dataflow.CreateTextRegisterStruct(
		"Product",
		"SerialNumber",
		"Serial number",
		0x010A,
		true,
		202,
	),
	dataflow.CreateTextRegisterStruct(
		"Product",
		"ModelName",
		"Model name",
		0x010B,
		true,
		203,
	),
	dataflow.CreateTextRegisterStruct(
		"Product",
		"Description",
		"Description",
		0x010C,
		true,
		204,
	),
	dataflow.CreateNumberRegisterStruct(
		"Product",
		"Uptime",
		"Device uptime",
		0x0120,
		false,
		false,
		1,
		"s",
		205,
	),
	// skipped Bluetooth capabilities
}

var RegisterListBmvMonitor = dataflow.Registers{
	dataflow.CreateNumberRegisterStruct(
		"Essential",
		"Power",
		"Power",
		0xED8E,
		false,
		true,
		1,
		"W",
		0,
	), dataflow.CreateNumberRegisterStruct(
		"Essential",
		"CurrentHighRes",
		"Current",
		0xED8C,
		false,
		true,
		1000,
		"A",
		1,
	),
	dataflow.CreateNumberRegisterStruct(
		"Essential",
		"MainVoltage",
		"Main voltage",
		0xED8D,
		false,
		true,
		100,
		"V",
		2,
	),
	dataflow.CreateNumberRegisterStruct(
		"Monitor",
		"AuxVoltage",
		"Aux (starter) voltage",
		0xED7D,
		false,
		false,
		100,
		"V",
		100,
	),
	dataflow.CreateNumberRegisterStruct(
		"Monitor",
		"Consumed",
		"Consumed",
		0xEEFF,
		false,
		true,
		10,
		"Ah",
		101,
	), dataflow.CreateNumberRegisterStruct(
		"Essential",
		"SOC",
		"State of charge",
		0x0FFF,
		false,
		false,
		100,
		"%",
		3,
	), dataflow.CreateNumberRegisterStruct(
		"Monitor",
		"TTG",
		"Time to go",
		0x0FFE,
		false,
		false,
		1,
		"min",
		102,
	), dataflow.CreateNumberRegisterStruct(
		"Essential",
		"Temperature",
		"Battery Temperature",
		0xEDEC,
		false,
		false,
		100,
		"K",
		4,
	), dataflow.CreateNumberRegisterStruct(
		"Monitor",
		"MidPointVoltage",
		"Mid-point voltage",
		0x0382,
		false,
		false,
		100,
		"V",
		104,
	), dataflow.CreateNumberRegisterStruct(
		"Monitor",
		"MidPointVoltageDeviation",
		"Mid-point voltage deviation",
		0x0383,
		false,
		true,
		10,
		"%",
		105,
	), dataflow.CreateEnumRegisterStruct(
		"Monitor",
		"SynchronizationState",
		"Synchronization state",
		0xEEB6,
		false,
		map[int]string{
			0: "false",
			1: "true",
		},
		106,
	),
}

var RegisterListBmvHistoric = dataflow.Registers{
	dataflow.CreateNumberRegisterStruct(
		"Historic",
		"DepthOfTheDeepestDischarge",
		"Depth of the deepest discharge",
		0x0300,
		false,
		true,
		10,
		"Ah",
		300,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"DepthOfTheLastDischarge",
		"Depth of the last discharge",
		0x0301,
		false,
		true,
		10,
		"Ah",
		301,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"DepthOfTheAverageDischarge",
		"Depth of the average discharge",
		0x0302,
		false,
		true,
		10,
		"Ah",
		302,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"NumberOfCycles",
		"Number of cycles",
		0x0303,
		false,
		false,
		1,
		"",
		303,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"NumberOfFullDischarges",
		"Number of full discharges",
		0x0304,
		false,
		false,
		1,
		"",
		304,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"CumulativeAmpHours",
		"Cumulative amp hours",
		0x0305,
		false,
		true,
		10,
		"Ah",
		305,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"MainVoltageMinimum",
		"Minimum voltage",
		0x0306,
		false,
		false,
		100,
		"V",
		306,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"MainVoltageMaximum",
		"Maximum voltage",
		0x0307,
		false,
		false,
		100,
		"V",
		307,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"TimeSinceFullCharge",
		"Time since full charge",
		0x0308,
		false,
		false,
		1,
		"s",
		308,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"NumberOfAutomaticSynchronizations",
		"Number of automatic synchronizations",
		0x0309,
		false,
		false,
		1,
		"",
		309,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"NumberOfLowMainVoltageAlarms",
		"Number of low voltage alarms",
		0x030A,
		false,
		false,
		1,
		"",
		310,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"NumberOfHighMainVoltageAlarms",
		"Number of high voltage alarms",
		0x030B,
		false,
		false,
		1,
		"",
		311,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"AuxVoltageMinimum",
		"Minimum starter voltage",
		0x030E,
		false,
		true,
		100,
		"V",
		312,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"AuxVoltageMaximum",
		"Maximum starter voltage",
		0x030F,
		false,
		true,
		100,
		"V",
		313,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"AmountOfDischargedEnergy",
		"Amount of discharged energy",
		0x0310,
		false,
		false,
		100,
		"kWh",
		314,
	), dataflow.CreateNumberRegisterStruct(
		"Historic",
		"AmountOfChargedEnergy",
		"Amount of charged energy",
		0x0311,
		false,
		false,
		100,
		"kWh",
		315,
	),
}

var RegisterListBmv = dataflow.MergeRegisters(
	RegisterListBmvMonitor,
	RegisterListBmvProduct,
	RegisterListBmvHistoric,
)

var RegisterListBmv712 = dataflow.FilterRegisters(
	RegisterListBmv,
	[]string{
		"ProductRevision",
		"Description",
	},
	[]string{},
)

var RegisterListBmv702 = dataflow.FilterRegisters(
	RegisterListBmv,
	[]string{
		"ProductRevision",
		"Description",
	},
	[]string{},
)

var RegisterListBmv700 = dataflow.FilterRegisters(
	RegisterListBmv,
	[]string{
		"AuxVoltage",
		"Temperature",
		"MidPointVoltage",
		"MidPointVoltageDeviation",
		"AuxVoltageMinimum",
		"AuxVoltageMaximum",
	},
	[]string{},
)
