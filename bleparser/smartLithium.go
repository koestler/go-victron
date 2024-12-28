package bleparser

import (
	"encoding/binary"
	"math"
)

// SmartLithium
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 32         | BMV flags             |         | 0 .. 0xFFFFFFFF     |
// 32         | 4          | 0          | 16         | SmartLithium error    |         | 0 .. 0xFFFF         |
// 48         | 6          | 0          | 7          | Cell 1                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 55         | 6          | 7          | 7          | Cell 2                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 62         | 7          | 6          | 7          | Cell 3                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 69         | 8          | 5          | 7          | Cell 4                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 76         | 9          | 4          | 7          | Cell 5                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 83         | 10         | 3          | 7          | Cell 6                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 90         | 11         | 2          | 7          | Cell 7                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 97         | 12         | 1          | 7          | Cell 8                | 0.01 V  | 2.60 .. 3.86 V      | 0x7F
// 104        | 13         | 0          | 12         | Battery voltage       | 0.01 V  | 0 .. 40.94 V        | 0x0FFF
// 116        | 14         | 4          | 4          | Balancer status       |         | 0 .. 15             | 0x0F
// 120        | 15         | 0          | 7          | Battery temperature   | 1 °C    | -40 .. 86 °C        | 0x7F

type SmartLithiumRecord struct {
	BmvFlags           uint32  `Description:"BMV flags"`          // TODO: decode flags (no doc found)
	SmartLithiumError  uint16  `Description:"SmartLithium error"` // TODO: decode error (no doc found)
	Cell1              float64 `Description:"Cell 1 voltage" Unit:"V"`
	Cell2              float64 `Description:"Cell 2 voltage" Unit:"V"`
	Cell3              float64 `Description:"Cell 3 voltage" Unit:"V"`
	Cell4              float64 `Description:"Cell 4 voltage" Unit:"V"`
	Cell5              float64 `Description:"Cell 5 voltage" Unit:"V"`
	Cell6              float64 `Description:"Cell 6 voltage" Unit:"V"`
	Cell7              float64 `Description:"Cell 7 voltage" Unit:"V"`
	Cell8              float64 `Description:"Cell 8 voltage" Unit:"V"`
	BatteryVoltage     float64 `Description:"Battery voltage" Unit:"V"`
	BalancerStatus     uint8   `Description:"Balancer status"`
	BatteryTemperature float64 `Description:"Battery temperature" Unit:"°C"`
}

func DecodeSmartLithium(inp []byte) (ret SmartLithiumRecord, err error) {
	if len(inp) < 16 {
		err = ErrInputTooShort
		return
	}

	ret.BmvFlags = binary.LittleEndian.Uint32(inp[0:4])
	ret.SmartLithiumError = binary.LittleEndian.Uint16(inp[4:6])

	ret.Cell1 = decodeCellVoltage(inp, 6, 0)
	ret.Cell2 = decodeCellVoltage(inp, 6, 7)
	ret.Cell3 = decodeCellVoltage(inp, 7, 6)
	ret.Cell4 = decodeCellVoltage(inp, 8, 5)
	ret.Cell5 = decodeCellVoltage(inp, 9, 4)
	ret.Cell6 = decodeCellVoltage(inp, 10, 3)
	ret.Cell7 = decodeCellVoltage(inp, 11, 2)
	ret.Cell8 = decodeCellVoltage(inp, 12, 1)

	if v := binary.LittleEndian.Uint16(inp[13:15]); v != 0x0FFF {
		ret.BatteryVoltage = float64(v) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.BalancerStatus = inp[14] >> 4 & 0x0F

	if v := inp[15] & 0x7F; v != 0x7F {
		ret.BatteryTemperature = float64(int8(v))
	} else {
		ret.BatteryTemperature = math.NaN()
	}

	return
}

func decodeCellVoltage(inp []byte, startByte, bitOffset int) float64 {
	if v := (binary.LittleEndian.Uint16(inp[startByte:startByte+2]) >> bitOffset) & 0x7F; v != 0x7F {
		return float64(v)/100 + 2.6
	}
	return math.NaN()
}

func (r SmartLithiumRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "BmvFlags",
				description: "BMV flags",
			},
			value: float64(r.BmvFlags),
		},
		{
			Register: Register{
				name:        "SmartLithiumError",
				description: "SmartLithium error",
			},
			value: float64(r.SmartLithiumError),
		},
		{
			Register: Register{
				name:        "Cell1",
				description: "Cell 1 voltage",
			},
			value: r.Cell1,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell2",
				description: "Cell 2 voltage",
			},
			value: r.Cell2,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell3",
				description: "Cell 3 voltage",
			},
			value: r.Cell3,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell4",
				description: "Cell 4 voltage",
			},
			value: r.Cell4,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell5",
				description: "Cell 5 voltage",
			},
			value: r.Cell5,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell6",
				description: "Cell 6 voltage",
			},
			value: r.Cell6,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell7",
				description: "Cell 7 voltage",
			},
			value: r.Cell7,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "Cell8",
				description: "Cell 8 voltage",
			},
			value: r.Cell8,
			unit:  "V",
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
				name:        "BalancerStatus",
				description: "Balancer status",
			},
			value: float64(r.BalancerStatus),
		},
		{
			Register: Register{
				name:        "BatteryTemperature",
				description: "Battery temperature",
			},
			value: r.BatteryTemperature,
			unit:  "°C",
		},
	}
}

func (r SmartLithiumRecord) EnumRegisters() []EnumRegister {
	return []EnumRegister{}
}

func (r SmartLithiumRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
