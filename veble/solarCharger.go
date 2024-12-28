package veble

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// Solar Charger
// Start bits | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 8          | Charger error         |         | 0 .. 0xFE           | 0xFF
// 16	      | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 32         | 16         | Battery current       | 0.1A    | -3276.8 .. 3276.6 A | 0x7FFF
// 48         | 16         | Yield today           | 0.01kWh | 0 .. 65534 W        | 0xFFFF
// 64         | 16         | PV power              | 1W      | 0 .. 65534 W        | 0xFFFF
// 80         | 9          | Load current          | 0.1A    | 0 .. 51.0 A         | 0x1FF

type SolarChargerRecord struct {
	DeviceState    veconst.SolarChargerState `Description:"Device state"`
	ChargerError   veconst.SolarChargerError `Description:"Charger error"`
	BatteryVoltage float64                   `Description:"Battery voltage" Unit:"V"`
	BatteryCurrent float64                   `Description:"Battery current" Unit:"A"`
	YieldToday     float64                   `Description:"Yield today" Unit:"Wh"`
	PvPower        float64                   `Description:"PV power" Unit:"W"`
	LoadCurrent    float64                   `Description:"Load current" Unit:"A"`
}

func DecodeSolarCharge(inp []byte) (ret SolarChargerRecord, err error) {
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

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[4:6]); v != 0x7FFF {
		ret.BatteryCurrent = float64(int16(v)) / 10
	} else {
		ret.BatteryCurrent = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
		ret.YieldToday = float64(v) * 10
	} else {
		ret.YieldToday = math.NaN()
	}
	if v := binary.LittleEndian.Uint16(inp[8:10]); v != 0xFFFF {
		ret.PvPower = float64(v)
	} else {
		ret.PvPower = math.NaN()
	}
	if v := binary.LittleEndian.Uint16([]byte{inp[10], inp[11]}) & 0x1FF; v != 0x1FF {
		ret.LoadCurrent = float64(v) / 10
	} else {
		ret.LoadCurrent = math.NaN()
	}

	return
}

func (r SolarChargerRecord) NumberRegisters() []NumberRegister {
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
				name:        "BatteryCurrent",
				description: "Battery current",
			},
			value: r.BatteryCurrent,
			unit:  "A",
		},
		{
			Register: Register{
				name:        "YieldToday",
				description: "Yield today",
			},
			value: r.YieldToday,
			unit:  "Wh",
		},
		{
			Register: Register{
				name:        "PvPower",
				description: "PV power",
			},
			value: r.PvPower,
			unit:  "W",
		},
		{
			Register: Register{
				name:        "LoadCurrent",
				description: "Load current",
			},
			value: r.LoadCurrent,
			unit:  "A",
		},
	}
}

func (r SolarChargerRecord) EnumRegisters() []EnumRegister {
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

func (r SolarChargerRecord) FieldListRegisters() []FieldListRegister {
	return []FieldListRegister{}
}
