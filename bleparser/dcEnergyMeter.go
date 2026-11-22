package bleparser

import (
	"encoding/binary"
	"github.com/koestler/go-victron/veconst"
	"math"
)

// DC Energy Meter
// Start bits | Start byte | Bit offset | Nr of bits | Meaning               | Units   | Range               | NA value
// 0          | 0          | 0          | 16         | BMV monitor mode      |         | -32768 .. 32767     |
// 16         | 2          | 0          | 16         | Battery voltage       | 0.01V   | -327.68 .. 327.66 V | 0x7FFF
// 32         | 4          | 0          | 16         | Alarm reason          |         | 0 .. 0xFFFF         |
// 48         | 6          | 0          | 16         | Aux voltage           | 0.01V   | -327.68 .. 327.64 V |
// 48         | 6          | 0          | 16         | Temperature           | 0.01K   | 0 .. 655.34 K       |
// 64         | 8          | 0          | 2          | Aux input             |         | 0 .. 3              | 0x03
// 66         | 8          | 2          | 22         | Battery current       | 0.001A  | -4194 .. 4194 A     | 0x3FFFFF

type DcEnergyMeterRecord struct {
	BmvMonitorMode int16                        `Description:"BMV monitor mode"`
	BatteryVoltage float64                      `Description:"Battery voltage" Unit:"V"`
	AlarmReason    uint16                       `Description:"Alarm reason"`
	AuxVoltage     float64                      `Description:"Aux voltage" Unit:"V"`
	Temperature    float64                      `Description:"Temperature" Unit:"K"`
	AuxMode        veconst.DcEnergyMeterAuxMode `Description:"Aux mode"`
	BatteryCurrent float64                      `Description:"Battery current" Unit:"A"`
}

func DecodeDcEnergyMeterRecord(inp []byte) (ret DcEnergyMeterRecord, err error) {
	if len(inp) < 12 {
		err = ErrInputTooShort
		return
	}

	ret.BmvMonitorMode = int16(binary.LittleEndian.Uint16(inp[0:2]))

	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryVoltage = float64(int16(v)) / 100
	} else {
		ret.BatteryVoltage = math.NaN()
	}

	ret.AlarmReason = binary.LittleEndian.Uint16(inp[4:6])

	ret.AuxMode = veconst.DcEnergyMeterAuxMode(inp[8] & 0x3)
	switch ret.AuxMode {
	case veconst.DcEnergyMeterAuxModeAuxVoltage:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0x7FFF {
			ret.AuxVoltage = float64(int16(v)) / 100
		} else {
			ret.AuxVoltage = math.NaN()
		}
	case veconst.DcEnergyMeterAuxModeTemperature:
		if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
			ret.Temperature = float64(v) / 100
		} else {
			ret.Temperature = math.NaN()
		}
	}

	if v := (binary.LittleEndian.Uint32(inp[8:12]) >> 2) & 0x3FFFFF; v != 0x3FFFFF {
		ret.BatteryCurrent = float64(int32(v)) / 1000
	} else {
		ret.BatteryCurrent = math.NaN()
	}

	return
}
