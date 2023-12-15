package ble

import (
	"encoding/binary"
	"fmt"
	"github.com/koestler/go-victron/victronDefinitions"
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
	DeviceState   victronDefinitions.DcDcConverterState `Description:"Device state"`
	ChargerError  victronDefinitions.DcDcConverterError `Description:"Charger error"`
	InputVoltage  float64                               `Description:"Input voltage" Unit:"V"`
	OutputVoltage float64                               `Description:"Output voltage" Unit:"V"`
	OffReason     uint32                                `Description:"Off reason"`
}

func DecodeDcDcConverterRecord(inp []byte) (ret DcDcConverterRecord, err error) {
	if len(inp) < 10 {
		err = fmt.Errorf("inp too short")
		return
	}

	ret.DeviceState = victronDefinitions.DcDcConverterState(inp[0])
	{
		sm := victronDefinitions.GetDcDcConverterStateMap()
		if _, ok := sm[ret.DeviceState]; !ok {
			ret.DeviceState = victronDefinitions.DcDcConverterStateUnavailable
		}
	}

	ret.ChargerError = victronDefinitions.DcDcConverterError(inp[1])
	{
		sm := victronDefinitions.GetDcDcConverterErrorMap()
		if _, ok := sm[ret.ChargerError]; !ok {
			ret.ChargerError = victronDefinitions.DcDcConverterErrorUnknown
		}
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
