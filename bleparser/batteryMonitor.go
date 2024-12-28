package bleparser

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// Battery Monitor
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 16         | TTG                   | 1m      | 0 .. 45.5 d         | 0xFFFF
// 16         | 2          | 0          | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 32         | 4          | 0          | 16         | Alarm reason          |         | 0 .. 0xFFFF         | -
// 48         | 6          | 0          | 16         | Aux voltage           | 0.01V   | -327.68 .. 327.66 V | -
// 48         | 6          | 0          | 16         | Mid voltage           | 0.01V   | 0 .. 655.34 V       | -
// 48         | 6          | 0          | 16         | Temperature           | 0.01K   | 0 .. 655.34 K       | -
// 64         | 8          | 0          | 2          | Aux input             |         | 0 .. 3              | 0x3
// 66         | 8          | 2          | 22         | Battery current       | 0.001A  | -4194 .. 4194 A     | 0x3FFFFF | 0x1FFFFF
// 88         | 11         | 0          | 20         | Consumed Ah           | 0.1Ah   | -104.857 .. 0 Ah    | 0x0FFFFF
// 108        | 13         | 4          | 10         | State of charge       | 0.1%    | 0 .. 100%           | 0x3FF

type BatteryMonitorRecord struct {
	TTG            float64            `Description:"Time to go" Unit:"s"`
	BatteryVoltage float64            `Description:"Battery voltage" Unit:"V"`
	AlarmReason    uint16             `Description:"Alarm reason"`
	AuxVoltage     float64            `Description:"Aux voltage" Unit:"V"`
	MidVoltage     float64            `Description:"Mid voltage" Unit:"V"`
	Temperature    float64            `Description:"Temperature" Unit:"°C"`
	AuxMode        veconst.BmvAuxMode `Description:"Aux mode"`
	BatteryCurrent float64            `Description:"Battery current" Unit:"A"`
	ConsumedAh     float64            `Description:"Consumed Ah" Unit:"Ah"`
	StateOfCharge  float64            `Description:"State of charge" Unit:"%"`
}

func DecodeBatteryMonitor(inp []byte) (ret BatteryMonitorRecord, err error) {
	if len(inp) < 15 {
		err = ErrInputTooShort
		return
	}

	if v := binary.LittleEndian.Uint16(inp[0:2]); v != 0xFFFF {
		ret.TTG = float64(v) * 60
	} else {
		ret.TTG = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.AlarmReason = binary.LittleEndian.Uint16(inp[4:6])

	ret.AuxMode = veconst.BmvAuxMode(inp[8] & 0x3)
	ret.AuxVoltage = math.NaN()
	ret.MidVoltage = math.NaN()
	ret.Temperature = math.NaN()

	switch ret.AuxMode {
	case veconst.BmvAuxModeStarterVoltage:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0x7FFF {
			ret.AuxVoltage = float64(int16(v)) / 100
		}
	case veconst.BmvAuxModeMidpointVoltage:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
			ret.MidVoltage = float64(v) / 100
		}
	case veconst.BmvAuxModeTemperature:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
			ret.Temperature = float64(v)/100 - 273.15
		}
	}

	if v := (binary.LittleEndian.Uint32(inp[8:12]) >> 2) & 0x3FFFFF; v != 0x3FFFFF && v != 0x1FFFFF {
		ret.BatteryCurrent = float64(int32(v)) / 1000
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	if v := binary.LittleEndian.Uint32(inp[11:15]) & 0x0FFFFF; v != 0x0FFFFF {
		ret.ConsumedAh = float64(-int32(v)) / 10
	} else {
		ret.ConsumedAh = math.NaN()
	}

	if v := (binary.LittleEndian.Uint16(inp[13:15]) >> 4) & 0x3FF; v != 0x3FF {
		ret.StateOfCharge = float64(v) / 10
	} else {
		ret.StateOfCharge = math.NaN()
	}

	return
}

func (r BatteryMonitorRecord) NumberRegisters() []NumberRegister {
	return []NumberRegister{
		{
			Register: Register{
				name:        "TTG",
				description: "Time to go",
			},
			value: r.TTG,
			unit:  "s",
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
				name:        "AlarmReason",
				description: "Alarm reason",
			},
			value: float64(r.AlarmReason),
		},
		{
			Register: Register{
				name:        "AuxVoltage",
				description: "Aux voltage",
			},
			value: r.AuxVoltage,
			unit:  "V",
		},
		{
			Register: Register{
				name:        "MidVoltage",
				description: "Mid voltage",
			},
			value: r.MidVoltage,
			unit:  "V",
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
				name:        "BatteryCurrent",
				description: "Battery current",
			},
			value: r.BatteryCurrent,
			unit:  "A",
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
				name:        "StateOfCharge",
				description: "State of charge",
			},
			value: r.StateOfCharge,
			unit:  "%",
		},
	}
}

func (r BatteryMonitorRecord) EnumRegisters() []EnumRegister {
	return []EnumRegister{
		{
			Register: Register{
				name:        "AuxMode",
				description: "Aux mode",
			},
			value: r.AuxMode,
		},
	}
}

func (r BatteryMonitorRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
