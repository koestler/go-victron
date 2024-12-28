package veble

import (
	"encoding/binary"
	"math"
)

// Smart Battery Protect
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 1          | 0          | 8          | Output state          |         | 0 .. 0xFE           | 0xFF
// 16         | 2          | 0          | 8          | Error code            |         | 0 .. 0xFE           | 0xFF
// 24         | 3          | 0          | 16         | Alarm reason          |         | 0 .. 0xFFFF         |
// 40         | 5          | 0          | 16         | Warning reason        |         | 0 .. 0xFFFF         |
// 56         | 7          | 0          | 16         | Input voltage         | 0.01 V  | -327.68 .. 327.66 V | 0x7FFF
// 72         | 9          | 0          | 16         | Output voltage        | 0.01 V  | 0 .. 655.34 V       | 0xFFFF
// 88         | 11         | 0          | 32         | Off reason            |         | 0 .. 0xFFFFFFFF     |

type SmartBatteryProtectRecord struct {
	DeviceState   uint8   `Description:"Device state"`   // TODO: decode state
	OutputState   uint8   `Description:"Output state"`   // TODO: decode state
	ErrorCode     uint8   `Description:"Error code"`     // TODO: decode
	AlarmReason   uint16  `Description:"Alarm reason"`   // TODO: decode
	WarningReason uint16  `Description:"Warning reason"` // TODO: decode
	InputVoltage  float64 `Description:"Input voltage" Unit:"V"`
	OutputVoltage float64 `Description:"Output voltage" Unit:"V"`
	OffReason     uint32  `Description:"Off reason"` // TODO: decode
}

func DecodeSmartBatteryProtect(inp []byte) (ret SmartBatteryProtectRecord, err error) {
	if len(inp) < 15 {
		err = ErrInputTooShort
		return
	}

	ret.DeviceState = inp[0]
	ret.OutputState = inp[1]
	ret.ErrorCode = inp[2]
	ret.AlarmReason = binary.LittleEndian.Uint16(inp[3:5])
	ret.WarningReason = binary.LittleEndian.Uint16(inp[5:7])

	if v := binary.LittleEndian.Uint16(inp[7:9]); v != 0x7FFF {
		ret.InputVoltage = float64(int16(v)) / 100
	} else {
		ret.InputVoltage = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[9:11]); v != 0xFFFF {
		ret.OutputVoltage = float64(v) / 100
	} else {
		ret.OutputVoltage = math.NaN()
	}

	ret.OffReason = binary.LittleEndian.Uint32(inp[11:15])
	return
}

func (r SmartBatteryProtectRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "DeviceState",
				description: "Device state",
			},
			value: float64(r.DeviceState),
		},
		{
			Register: Register{
				name:        "OutputState",
				description: "Output state",
			},
			value: float64(r.OutputState),
		},
		{
			Register: Register{
				name:        "ErrorCode",
				description: "Error code",
			},
			value: float64(r.ErrorCode),
		},
		{
			Register: Register{
				name:        "AlarmReason",
				description: "Alarm reason",
			},
			value: float64(r.AlarmReason),
		},
		{
			Register: Register{
				name:        "WarningReason",
				description: "Warning reason",
			},
			value: float64(r.WarningReason),
		},
		{
			Register: Register{
				name:        "InputVoltage",
				description: "Input voltage",
			},
			value: r.InputVoltage,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "OutputVoltage",
				description: "Output voltage",
			},
			value: r.OutputVoltage,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "OffReason",
				description: "Off reason",
			},
			value: float64(r.OffReason),
		},
	}
}

func (r SmartBatteryProtectRecord) EnumRegisters() []EnumRegister {
	return []EnumRegister{}
}

func (r SmartBatteryProtectRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
