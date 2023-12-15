package ble

import (
	"encoding/binary"
	"math"
)

// GX-Device
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 16         | Battery voltage       | 0.01V   | 0 .. 655.34 V       | 0xFFFF
// 16         | 2          | 0          | 20         | PV power              | 1W      | 0 .. 1 MW           | 0xFFFFF
// 36         | 4          | 4          | 7          | SOC                   | 1%      | 0 .. 100%           | 0x7F
// 43         | 5          | 3          | 21         | Battery power         | 1W      | -1 .. 1 MW          | 0x0FFFFF
// 64         | 8          | 0          | 21         | DC power              | 1W      | -1 .. 1 MW          | 0x0FFFFF

type GxDeviceRecord struct {
	BatteryVoltage float64 `Description:"Battery voltage" Unit:"V"`
	PvPower        float64 `Description:"Solar power" Unit:"W"`
	Soc            float64 `Description:"State of charge" Unit:"%"`
	BatteryPower   float64 `Description:"Battery power" Unit:"W"`
	DcPower        float64 `Description:"DC power" Unit:"W"`
}

func DecodeGxDeviceRecord(inp []byte) (ret GxDeviceRecord, err error) {
	if len(inp) < 11 {
		err = ErrInputTooShort
		return
	}

	if v := binary.LittleEndian.Uint16(inp[0:2]); v != 0xFFFF {
		ret.BatteryVoltage = float64(v) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}
	if v := binary.LittleEndian.Uint32(inp[2:6]) & 0x0FFFFF; v != 0x0FFFFF {
		ret.PvPower = float64(v)
	} else {
		ret.PvPower = math.NaN()
	}
	if v := (binary.LittleEndian.Uint16(inp[4:6]) >> 4) & 0x7F; v != 0x7F {
		ret.Soc = float64(v)
	} else {
		ret.Soc = math.NaN()
	}
	if v := (binary.LittleEndian.Uint32(inp[5:9]) >> 3) & 0x0FFFFF; v != 0x0FFFFF {
		ret.BatteryPower = float64(int32(v))
	} else {
		ret.BatteryPower = math.NaN()
	}
	if v := binary.LittleEndian.Uint32(inp[12:15]); v != 0x0FFFFF {
		ret.DcPower = float64(int32(v))
	} else {
		ret.DcPower = math.NaN()
	}
	return
}
