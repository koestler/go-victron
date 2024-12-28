package veble

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// DC-DC Converter
// Start bits | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 8          | Charger error         |         | 0 .. 0xFE           | 0xFF
// 16	      | 16         | Input voltage         | 0.01V   | 0 .. 655.34 V       | 0xFFFF
// 32         | 16         | Output voltage        | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 48         | 32         | Off reason            |         | 0 .. 0xFFFFFFFF     |

type DcDcConverterRecord struct {
	DeviceState   veconst.DcDcConverterState `Description:"Device state"`
	ChargerError  veconst.DcDcConverterError `Description:"Charger error"`
	InputVoltage  float64                    `Description:"Input voltage" Unit:"V"`
	OutputVoltage float64                    `Description:"Output voltage" Unit:"V"`
	OffReason     uint32                     `Description:"Off reason"`
}

func DecodeDcDcConverter(inp []byte) (ret DcDcConverterRecord, err error) {
	if len(inp) < 10 {
		err = ErrInputTooShort
		return
	}

	if v, e := veconst.DcDcConverterStateFactory.New(inp[0]); e != nil {
		err = e
		return
	} else {
		ret.DeviceState = v
	}

	if v, e := veconst.DcDcConverterErrorFactory.New(inp[1]); e != nil {
		err = e
		return
	} else {
		ret.ChargerError = v
	}

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0xFFFF {
		ret.InputVoltage = float64(v) / 100
	} else {
		ret.InputVoltage = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[4:6]); v != 0x7FFF {
		ret.OutputVoltage = float64(int16(v)) / 100
	} else {
		ret.OutputVoltage = math.NaN()
	}

	ret.OffReason = binary.LittleEndian.Uint32(inp[6:10])

	return
}

func (r DcDcConverterRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
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

func (r DcDcConverterRecord) EnumRegisters() []EnumRegister {
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
	}
}

func (r DcDcConverterRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
