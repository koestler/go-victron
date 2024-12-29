package veblerecord

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
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
	DeviceState     veconst.InverterState          `Description:"Device state"`
	AlarmReason     veconst.InverterWarningReasons `Description:"Alarm reason"`
	BatteryVoltage  float64                        `Description:"Battery voltage" Unit:"V"`
	AcApparentPower float64                        `Description:"AC Apparent power" Unit:"VA"`
	AcVoltage       float64                        `Description:"AC voltage" Unit:"V"`
	AcCurrent       float64                        `Description:"AC current" Unit:"A"`
}

func DecodeInverter(inp []byte) (ret InverterRecord, err error) {
	if len(inp) < 11 {
		err = ErrInputTooShort
		return
	}

	if v, e := veconst.InverterStateFactory.New(inp[0]); e != nil {
		err = e
		return
	} else {
		ret.DeviceState = v
	}

	if v, e := veconst.InverterWarningReasonFactory.New(binary.LittleEndian.Uint16(inp[1:3])); e != nil {
		err = e
		return
	} else {
		ret.AlarmReason = v
	}

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

func (r InverterRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
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
				name:        "AcApparentPower",
				description: "AC Apparent power",
			},
			value: r.AcApparentPower,
			unit:  "VA",
		},
		{
			Register: Register{
				name:        "AcVoltage",
				description: "AC voltage",
			},
			value: r.AcVoltage,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "AcCurrent",
				description: "AC current",
			},
			value: r.AcCurrent,
			unit:  "A",
		},
	}
}

func (r InverterRecord) EnumRegisters() []EnumRegister {
	return []EnumRegister{
		{
			Register: Register{
				name:        "DeviceState",
				description: "Device state",
			},
			value: r.DeviceState,
		},
	}
}

func (r InverterRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{
		{
			Register: Register{
				name:        "AlarmReason",
				description: "Alarm reason",
			},
			value: r.AlarmReason,
		},
	}
}
