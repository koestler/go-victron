package ble

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
	BmvFlags           uint32  `Description:"BMV flags"`
	SmartLithiumError  uint16  `Description:"SmartLithium error"`
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

func DecodeSmartLithiumRecord(inp []byte) (ret SmartLithiumRecord, err error) {
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
