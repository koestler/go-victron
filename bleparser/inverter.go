package bleparser

import (
	"encoding/binary"
	"github.com/koestler/go-victron/victronDefinitions"
	"math"
)

// Inverter
// Start bits | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 16         | Alarm Reason          |         | 0 .. 0xFFFF         |
// 24         | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 40         | 16         | AC Apparent power     | 1VA     | 0 .. 65534 VA       | 0xFFFF
// 56         | 15         | AC voltage            | 0.01V   | 0 .. 327.66 V       | 0x7FFF
// 71         | 11         | AC current            | 0.1A    | 0 .. 204.6 A        | 0x7FF

type InverterRecord struct {
	DeviceState     victronDefinitions.InverterState `Description:"Device state"`
	AlarmReason     uint16                           `Description:"Alarm reason"`
	BatteryVoltage  float64                          `Description:"Battery voltage" Unit:"V"`
	AcApparentPower float64                          `Description:"AC Apparent power" Unit:"VA"`
	AcVoltage       float64                          `Description:"AC voltage" Unit:"V"`
	AcCurrent       float64                          `Description:"AC current" Unit:"A"`
}

func DecodeInverterRecord(inp []byte) (ret InverterRecord, err error) {
	if len(inp) < 11 {
		err = ErrInputTooShort
		return
	}

	ret.DeviceState = victronDefinitions.InverterState(inp[0])
	ret.AlarmReason = binary.LittleEndian.Uint16(inp[1:3])
	if v := binary.LittleEndian.Uint16(inp[3:5]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[5:7]); v != 0xFFFF {
		ret.AcApparentPower = float64(v)
	} else {
		ret.AcApparentPower = math.NaN()
	}

	// drop last bit
	if v := binary.LittleEndian.Uint16(inp[7:9]) & 0x7FFF; v != 0x7FFF {
		ret.AcVoltage = float64(v) / 100
	} else {
		ret.AcVoltage = math.NaN()
	}

	// drop first 7 bits
	if v := (binary.LittleEndian.Uint32([]byte{inp[8], inp[9], inp[10], 0x00}) >> 7) & 0x07FF; v != 0x7FF {
		ret.AcCurrent = float64(v) / 10
	} else {
		ret.AcCurrent = math.NaN()
	}

	return
}
