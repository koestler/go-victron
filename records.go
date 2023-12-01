package ble

import (
	"encoding/binary"
	"fmt"
)

type SolarChargerRecord struct {
	DeviceState    uint
	ChargerError   uint
	BatteryVoltage float64 // in V
	BatteryCurrent float64 // in A
	YieldToday     float64 // in kWh
	PvPower        float64 // in W
	LoadCurrent    float64 // in A
}

// Solar Charger
// Start bits | Nr of bits | Meaning               | Units   | NA value
// 0          | 8          | Device state          |         | 0xFF
// 8          | 8          | Charger error         |         | 0xFF
// 16		  | 16         | Battery voltage       | 0.01V   | 0x7FFF
// 32         | 16         | Battery current       | 0.1A    | 0x7FFF
// 48         | 16         | Yield today           | 0.01kWh | 0xFFFF
// 64         | 16         | PV power              | 1W      | 0xFFFF
// 80         | 9          | Load current          | 0.1A    | 0x1FF

func DecodeSolarChargeRecord(inp []byte) (ret SolarChargerRecord, err error) {
	if len(inp) < 12 {
		err = fmt.Errorf("inp too short")
		return
	}

	ret.DeviceState = uint(inp[0])
	ret.ChargerError = uint(inp[1])
	if v := binary.LittleEndian.Uint16(inp[2:4]); v != 0x7FFF {
		ret.BatteryVoltage = float64(v) / 100
	}
	if v := binary.LittleEndian.Uint16(inp[4:6]); v != 0x7FFF {
		ret.BatteryCurrent = float64(v) / 10
	}
	if v := binary.LittleEndian.Uint16(inp[6:8]); v != 0xFFFF {
		ret.YieldToday = float64(v) / 100
	}
	if v := binary.LittleEndian.Uint16(inp[8:10]); v != 0xFFFF {
		ret.PvPower = float64(v)
	}
	if v := binary.LittleEndian.Uint16(inp[10:12]); v != 0x1FF {
		ret.LoadCurrent = float64(v) / 10
	}

	return
}
