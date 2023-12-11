package ble

import (
	"encoding/binary"
	"fmt"
	"github.com/koestler/go-iotdevice/victronDefinitions"
	"math"
)

type BatteryMonitorRecord struct {
	Ttg            float64                       `Description:"Time to go" Unit:"s"`
	BatteryVoltage float64                       `Description:"Battery voltage" Unit:"V"`
	AlarmReason    uint16                        `Description:"Alarm reason"`
	AuxVoltage     float64                       `Description:"Aux voltage" Unit:"V"`
	MidVoltage     float64                       `Description:"Mid voltage" Unit:"V"`
	Temperature    float64                       `Description:"Temperature" Unit:"K"`
	AuxMode        victronDefinitions.BmvAuxMode `Description:"Aux mode"`
	BatteryCurrent float64                       `Description:"Battery current" Unit:"A"`
	ConsumedAh     float64                       `Description:"Consumed Ah" Unit:"Ah"`
	StateOfCharge  float64                       `Description:"State of charge" Unit:"%"`
}

// Battery Monitor
// Start bits | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 16         | TTG                   | 1m      | 0 .. 45.5 d         | 0xFFFF
// 16         | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 32         | 16         | Alarm reason          |         | 0 .. 0xFFFF         | -
// 48         | 16         | Aux voltage           | 0.01V   | -327.68 .. 327.66 V | -
// 48         | 16         | Mid voltage           | 0.01V   | 0 .. 655.34 V       | -
// 48         | 16         | Temperature           | 0.01K   | 0 .. 655.34 K       | -
// 64         | 2          | Aux input             |         | 0 .. 3              | 0x3
// 66         | 22         | Battery current       | 0.001A  | -4194 .. 4194 A     | 0x3FFFFF
// 88         | 20         | Consumed Ah           | 0.1Ah   | -104.857 .. 0 Ah    | 0x0FFFFF
// 108        | 10         | State of charge       | 0.1%    | 0 .. 100%           | 0x3FF

func DecodeBatteryMonitorRecord(inp []byte) (ret BatteryMonitorRecord, err error) {
	if len(inp) < 15 {
		err = fmt.Errorf("inp too short")
		return
	}

	if v := binary.LittleEndian.Uint16(inp[0:2]); v != 0xFFFF {
		ret.Ttg = float64(v) * 60
	} else {
		ret.Ttg = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.AlarmReason = binary.LittleEndian.Uint16(inp[4:6])

	ret.AuxMode = victronDefinitions.BmvAuxMode(inp[8] & 0x3)
	ret.AuxVoltage = math.NaN()
	ret.MidVoltage = math.NaN()
	ret.Temperature = math.NaN()

	switch ret.AuxMode {
	case 0:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0x7FFF {
			ret.AuxVoltage = float64(int16(v)) / 100
		}
	case 1:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
			ret.MidVoltage = float64(v) / 100
		}
	case 2:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
			ret.Temperature = float64(v) / 100
		}
	}

	// drop first two bits
	if v := binary.LittleEndian.Uint32([]byte{inp[8], inp[9], inp[10], 0x00}) >> 2; v != 0x3FFFFF {
		ret.BatteryCurrent = float64(int32(v)) / 1000
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	// drop last 4 bits
	if v := binary.LittleEndian.Uint32(inp[11:15]) & 0x0FFFFF; v != 0x0FFFFF {
		ret.ConsumedAh = float64(-int32(v)) / 10
	} else {
		ret.ConsumedAh = math.NaN()
	}

	// drop first 4 bits and drop last 2 bits
	if v := (binary.LittleEndian.Uint16([]byte{inp[13], inp[14]}) >> 4) & 0x3FF; v != 0x3FF {
		ret.StateOfCharge = float64(v) / 10
	} else {
		ret.StateOfCharge = math.NaN()
	}

	return
}
