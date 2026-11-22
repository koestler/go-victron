package veblerecord

import (
	"encoding/binary"
	"math"
)

// (Lynx Smart) BMS
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 8          | Error                 |         |                     | 0x00
// 8          | 1          | 0          | 16         | TTG                   | 1m      | 0 .. 45.5d          | 0xFFFF
// 24         | 3          | 0          | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 40         | 5          | 0          | 16         | Battery current       | 0.1A    | -3276.8 .. 3276.6 A | 0x7FFF
// 56         | 7          | 0          | 16         | IO status             |         |                     | 0x0000
// 72         | 9          | 0          | 18         | Warnings/Alarms       |         |                     | 0x000000
// 90         | 11         | 2          | 10         | SOC                   | 0.1%    | 0 .. 100%           | 0x3FF
// 100        | 12         | 4          | 20         | Consumed Ah           | 0.1Ah   | -104.857 .. 0Ah     | 0xFFFFF
// 120        | 15         | 0          | 7          | Temperature           | 1°C     | -40 .. 86 °C        | 0x7F

type LynxSmartBms struct {
	Error          uint8   `Description:"Error"` // TODO: decode
	TTG            float64 `Description:"Time to go" Unit:"s"`
	BatteryVoltage float64 `Description:"Battery voltage" Unit:"V"`
	BatteryCurrent float64 `Description:"Battery current" Unit:"A"`
	IoStatus       uint16  `Description:"IO status"`       // TODO: decode
	WarningsAlarms uint32  `Description:"Warnings/Alarms"` // TODO: decode
	SOC            float64 `Description:"State of charge" Unit:"%"`
	ConsumedAh     float64 `Description:"Consumed Ah" Unit:"Ah"`
	Temperature    float64 `Description:"Temperature" Unit:"°C"`
}

func DecodeLynxSmartBms(inp []byte) (ret LynxSmartBms, err error) {
	if len(inp) < 16 {
		err = ErrInputTooShort
		return
	}

	ret.Error = inp[0]

	if v := binary.LittleEndian.Uint16(inp[1:3]); v != 0xFFFF {
		ret.TTG = float64(v) * 60
	} else {
		ret.TTG = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[3:5]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[5:7]); v != 0x7FFF {
		ret.BatteryCurrent = float64(int16(v)) / 10
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	ret.IoStatus = binary.LittleEndian.Uint16(inp[7:9])

	ret.WarningsAlarms = binary.LittleEndian.Uint32(inp[9:13]) & 0x3FFFF

	if v := (binary.LittleEndian.Uint16(inp[11:13]) >> 2) & 0x3FF; v != 0x3FF {
		ret.SOC = float64(v) / 10
	} else {
		ret.SOC = math.NaN()
	}

	if v := (binary.LittleEndian.Uint32(inp[12:16]) >> 4) & 0xFFFFF; v != 0xFFFFF {
		ret.ConsumedAh = float64(v) / 10
	} else {
		ret.ConsumedAh = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[15:17]) & 0x7F; v != 0x7F {
		ret.Temperature = float64(int16(v) - 40)
	} else {
		ret.Temperature = math.NaN()
	}

	return
}

func (r LynxSmartBms) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "Error",
				description: "Error",
			},
			value: float64(r.Error),
		},
		{
			Register: Register{
				name:        "TTG",
				description: "Time to go",
			},
			value: r.TTG,
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
				name:        "BatteryCurrent",
				description: "Battery current",
			},
			value: r.BatteryCurrent,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "IoStatus",
				description: "IO status",
			},
			value: float64(r.IoStatus),
		},
		{
			Register: Register{
				name:        "WarningsAlarms",
				description: "Warnings/Alarms",
			},
			value: float64(r.WarningsAlarms),
		},
		{
			Register: Register{
				name:        "SOC",
				description: "State of charge",
			},
			value: r.SOC,
			unit:  "%",
		},
		{
			Register: Register{
				name:        "ConsumedAh",
				description: "Consumed Ah",
			},
			value: r.ConsumedAh,
			unit:  "Ah",
		},
		{
			Register: Register{
				name:        "Temperature",
				description: "Temperature",
			},
			value: r.Temperature,
			unit:  "°C",
		},
	}
}

func (r LynxSmartBms) EnumRegisters() []EnumRegister {
	return []EnumRegister{}
}

func (r LynxSmartBms) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
