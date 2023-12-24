package bleparser

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

func DecodeDcDcConverterRecord(inp []byte) (ret DcDcConverterRecord, err error) {
	if len(inp) < 10 {
		err = ErrInputTooShort
		return
	}

	ret.DeviceState = veconst.DcDcConverterState(inp[0])
	if !ret.DeviceState.Exists() {
		ret.DeviceState = veconst.DcDcConverterStateUnavailable
	}

	ret.ChargerError = veconst.DcDcConverterError(inp[1])
	if !ret.ChargerError.Exists() {
		ret.ChargerError = veconst.DcDcConverterErrorUnknown
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
