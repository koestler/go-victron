package veble

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// Multi RS
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 8		     | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 1          | 0          | 8          | Charger error         |         | 0 .. 0xFE           | 0xFF
// 16         | 2          | 0          | 16         | Battery current       | 0.1A    | -3276.8 .. 3276.6 A | 0x7FFF
// 32         | 4          | 0          | 14         | Battery voltage       | 0.1V    | 0 .. 163.83 V       | 0x3FFF
// 46         | 5          | 6          | 2          | Active AC in          |         | 0 .. 3              | 0x3
// 48         | 6          | 0          | 16         | Active AC in power    | 1W      | -32768 .. 32766 W   | 0x7FFF
// 64         | 8          | 0          | 16         | AC out power          | 1W      | -32768 .. 32766 W   | 0x7FFF
// 80         | 10         | 0          | 16         | PV power              | 1W      | 0 .. 65534 W        | 0xFFFF
// 96         | 12         | 0          | 16         | Yield today           | 0.01kWh | 0 .. 655.34 kWh     | 0xFFFF

type MultiRsRecord struct {
	DeviceState     veconst.InverterState      `Description:"Device state"`
	ChargerError    veconst.SolarChargerError  `Description:"Charger error"`
	BatteryCurrent  float64                    `Description:"Battery current" Unit:"A"`
	BatteryVoltage  float64                    `Description:"Battery voltage" Unit:"V"`
	ActiveAcIn      veconst.MultiRsActiveInput `Description:"Active AC in"`
	ActiveAcInPower float64                    `Description:"Active AC in power" Unit:"W"`
	AcOutPower      float64                    `Description:"AC out power" Unit:"W"`
	PvPower         float64                    `Description:"PV power" Unit:"W"`
	YieldToday      float64                    `Description:"Yield today" Unit:"Wh"`
}

func DecodeMultiRs(inp []byte) (ret MultiRsRecord, err error) {
	if len(inp) < 13 {
		err = ErrInputTooShort
		return
	}

	if v, e := veconst.InverterStateFactory.New(inp[0]); e != nil {
		err = e
		return
	} else {
		ret.DeviceState = v
	}

	if v, e := veconst.SolarChargerErrorFactory.New(inp[1]); e != nil {
		err = e
		return
	} else {
		ret.ChargerError = v
	}

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryCurrent = float64(int16(v)) / 10
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[4:6]) & 0x3FFF; v != 0x3FFF {
		ret.BatteryVoltage = float64(v) / 10
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.ActiveAcIn = veconst.MultiRsActiveInput((inp[5] >> 6) & 0x3)

	if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0x7FFF {
		ret.ActiveAcInPower = float64(int16(v))
	} else {
		ret.ActiveAcInPower = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[8:10]); v != 0x7FFF {
		ret.AcOutPower = float64(int16(v))
	} else {
		ret.AcOutPower = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[10:12]); v != 0xFFFF {
		ret.PvPower = float64(v)
	} else {
		ret.PvPower = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[12:14]); v != 0xFFFF {
		ret.YieldToday = float64(v) * 10
	} else {
		ret.YieldToday = math.NaN()
	}

	return
}

func (r MultiRsRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "BatteryCurrent",
				description: "Battery current",
			},
			value: r.BatteryCurrent,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "BatteryVoltage",
				description: "Battery voltage",
			},
			value: r.BatteryVoltage,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "ActiveAcInPower",
				description: "Active AC in power",
			},
			value: r.ActiveAcInPower,
			unit:  "W",
		},
		{
			Register: Register{
				name:        "AcOutPower",
				description: "AC out power",
			},
			value: r.AcOutPower,
			unit:  "W",
		},
		{
			Register: Register{
				name:        "PvPower",
				description: "PV power",
			},
			value: r.PvPower,
			unit:  "W",
		},
		{
			Register: Register{
				name:        "YieldToday",
				description: "Yield today",
			},
			value: r.YieldToday,
			unit:  "Wh",
		},
	}
}

func (r MultiRsRecord) EnumRegisters() []EnumRegister {
	return []EnumRegister{
		{
			Register: Register{
				name:        "DeviceState",
				description: "Device state",
			},
			value: r.DeviceState,
		},
		{
			Register: Register{
				name:        "ChargerError",
				description: "Charger error",
			},
			value: r.ChargerError,
		},
		{
			Register: Register{
				name:        "ActiveAcIn",
				description: "Active AC in",
			},
			value: r.ActiveAcIn,
		},
	}
}

func (r MultiRsRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
