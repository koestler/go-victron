package veble

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// AC Charger
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 1          | 0          | 8          | Charger Error         |         | 0 .. 0xFE           | 0xFF
// 16         | 2          | 0          | 13         | Battery voltage 1     | 0.01V   | 0 .. 81.90 V        | 0x1FFF
// 29         | 3          | 5          | 11         | Battery current 1     | 0.1A    | 0 .. 204.6 A        | 0x7FF
// 40         | 5          | 0          | 13         | Battery voltage 2     | 0.01V   | 0 .. 81.90 V        | 0x1FFF
// 53         | 6          | 5          | 11         | Battery current 2     | 0.1A    | 0 .. 204.6 A        | 0x7FF
// 64         | 8          | 0          | 13         | Battery voltage 3     | 0.01V   | 0 .. 81.90 V        | 0x1FFF
// 77         | 9          | 5          | 11         | Battery current 3     | 0.1A    | 0 .. 204.6 A        | 0x7FF
// 88         | 11         | 0          | 7          | Temperature           | 1°C     | -40 .. 86 °C        | 0x7F
// 95         | 11         | 7          | 9          | AC Current            | 0.1A    | 0 .. 51.0 A         | 0x1FF

type AcChargerRecord struct {
	DeviceState     veconst.SolarChargerState `Description:"Device state"`
	ChargerError    veconst.SolarChargerError `Description:"Charger error"`
	BatteryVoltage1 float64                   `Description:"Battery voltage 1" Unit:"V"`
	BatteryCurrent1 float64                   `Description:"Battery current 1" Unit:"A"`
	BatteryVoltage2 float64                   `Description:"Battery voltage 2" Unit:"V"`
	BatteryCurrent2 float64                   `Description:"Battery current 2" Unit:"A"`
	BatteryVoltage3 float64                   `Description:"Battery voltage 3" Unit:"V"`
	BatteryCurrent3 float64                   `Description:"Battery current 3" Unit:"A"`
	Temperature     float64                   `Description:"Temperature" Unit:"°C"`
	AcCurrent       float64                   `Description:"AC Current" Unit:"A"`
}

func DecodeAcCharger(inp []byte) (ret AcChargerRecord, err error) {
	if len(inp) < 12 {
		err = ErrInputTooShort
		return
	}

	if v, e := veconst.SolarChargerStateFactory.New(inp[0]); e != nil {
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

	if v := binary.LittleEndian.Uint16(inp[2:4]) & 0x1FFF; v != 0x1FFF {
		ret.BatteryVoltage1 = float64(v) / 100
	} else {
		ret.BatteryVoltage1 = math.NaN()
	}
	if v := (binary.LittleEndian.Uint16(inp[3:5]) << 5) & 0x7FF; v != 0x7FF {
		ret.BatteryCurrent1 = float64(v) / 10
	} else {
		ret.BatteryCurrent1 = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[5:7]); v != 0x1FFF {
		ret.BatteryVoltage2 = float64(v) / 100
	} else {
		ret.BatteryVoltage2 = math.NaN()
	}
	if v := (binary.LittleEndian.Uint16(inp[6:8]) << 5) & 0x7FF; v != 0x7FF {
		ret.BatteryCurrent2 = float64(v) / 10
	} else {
		ret.BatteryCurrent2 = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[8:10]); v != 0x1FFF {
		ret.BatteryVoltage3 = float64(v) / 100
	} else {
		ret.BatteryVoltage3 = math.NaN()
	}
	if v := (binary.LittleEndian.Uint16(inp[9:11]) << 5) & 0x7FF; v != 0x7FF {
		ret.BatteryCurrent3 = float64(v) / 10
	} else {
		ret.BatteryCurrent3 = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[11:13]) & 0x7F; v != 0x7F {
		ret.Temperature = float64(int16(v) - 40)
	} else {
		ret.Temperature = math.NaN()
	}
	if v := (binary.LittleEndian.Uint16(inp[11:13]) << 7) & 0x1FF; v != 0x1FF {
		ret.AcCurrent = float64(v) / 10
	} else {
		ret.AcCurrent = math.NaN()
	}
	return
}

func (r AcChargerRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "BatteryVoltage1",
				description: "Battery voltage 1",
			},
			value: r.BatteryVoltage1,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "BatteryCurrent1",
				description: "Battery current 1",
			},
			value: r.BatteryCurrent1,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "BatteryVoltage2",
				description: "Battery voltage 2",
			},
			value: r.BatteryVoltage2,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "BatteryCurrent2",
				description: "Battery current 2",
			},
			value: r.BatteryCurrent2,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "BatteryVoltage3",
				description: "Battery voltage 3",
			},
			value: r.BatteryVoltage3,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "BatteryCurrent3",
				description: "Battery current 3",
			},
			value: r.BatteryCurrent3,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "Temperature",
				description: "Temperature",
			},
			value: r.Temperature,
			unit:  "°C",
		},
		{
			Register: Register{
				name:        "AcCurrent",
				description: "AC Current",
			},
			value: r.AcCurrent,
			unit:  "A",
		},
	}
}

func (r AcChargerRecord) EnumRegisters() []EnumRegister {
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

func (r AcChargerRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
