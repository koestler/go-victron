package bleparser

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// VE.Bus
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 8          | Device state          |         | 0 .. 0xFE           | 0xFF
// 8          | 1          | 0          | 8          | VE.Bus error          |         | 0 .. 0xFE           | 0xFF
// 16         | 2          | 0          | 16         | Battery current       | 0.1A    | -3276.8 .. 3276.6 A | 0x7FFF
// 32         | 4          | 0          | 14         | Battery voltage       | 0.01V   | 0 .. 163.83 V       | 0x3FFF
// 46         | 5          | 6          | 2          | Active AC in          |         | 0 .. 3              | 0x3
// 48         | 6          | 0          | 19         | Active AC in power    | 1W      | -262144 .. 262142 W | 0x7FFFF
// 67         | 8          | 3          | 19         | AC out power          | 1W      | -262144 .. 262142 W | 0x7FFFF
// 86         | 10         | 6          | 2          | Alarm                 |         | 0 .. 3              | 0x3
// 88         | 11         | 0          | 7		     | Temperature           | 1°C     | -40 .. 86 °C        | 0x7F
// 95         | 11         | 7          | 7          | Soc                   | 1%      | 0 .. 126 %          | 0x7F

type VeBusRecord struct {
	DeviceState     veconst.InverterState      `Description:"Device state"`
	VeBusError      veconst.SolarChargerError  `Description:"VE.Bus error"`
	BatteryCurrent  float64                    `Description:"Battery current" Unit:"A"`
	BatteryVoltage  float64                    `Description:"Battery voltage" Unit:"V"`
	ActiveAcIn      veconst.MultiRsActiveInput `Description:"Active AC in"`
	ActiveAcInPower float64                    `Description:"Active AC in power" Unit:"W"`
	AcOutPower      float64                    `Description:"AC out power" Unit:"W"`
	Alarm           veconst.VeBusAlarm         `Description:"Alarm"`
	Temperature     float64                    `Description:"Temperature" Unit:"°C"`
	Soc             float64                    `Description:"State of charge" Unit:"%"`
}

func DecodeVeBus(inp []byte) (ret VeBusRecord, err error) {
	if len(inp) < 12 {
		err = ErrInputTooShort
		return
	}

	ret.DeviceState = veconst.InverterState(inp[0])
	ret.VeBusError = veconst.SolarChargerError(inp[1])

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryCurrent = float64(int16(v)) / 10
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	if v := binary.LittleEndian.Uint16(inp[4:6]) & 0x3FFF; v != 0x3FFF {
		ret.BatteryVoltage = float64(v) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.ActiveAcIn = veconst.MultiRsActiveInput((inp[5] >> 6) & 0x3)

	if v := binary.LittleEndian.Uint32(inp[6:9]) & 0x7FFFF; v != 0x7FFFF {
		ret.ActiveAcInPower = float64(int32(v))
	} else {
		ret.ActiveAcInPower = math.NaN()
	}

	if v := (binary.LittleEndian.Uint32(inp[8:11]) >> 3) & 0x7FFFF; v != 0x7FFFF {
		ret.AcOutPower = float64(int32(v))
	} else {
		ret.AcOutPower = math.NaN()
	}

	ret.Alarm = veconst.VeBusAlarm((inp[10] >> 6) & 0x3)

	if v := inp[11] & 0x7F; v != 0x7F {
		ret.Temperature = float64(int8(v))
	} else {
		ret.Temperature = math.NaN()
	}

	if v := (binary.LittleEndian.Uint16(inp[11:12]) >> 7) & 0x7F; v != 0x7F {
		ret.Soc = float64(v)
	} else {
		ret.Soc = math.NaN()
	}

	return
}
