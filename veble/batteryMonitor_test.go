package veble

import (
	"encoding/hex"
	"github.com/koestler/go-victron/veconst"
	"math"
	"testing"
)

// this test is based on https://github.com/keshavdv/victron-ble/blob/e28c5f8cc5f9f3062a2f36c2115d38214c07e741/tests/test_battery_monitor.py

func TestDecodeBatteryMonitor(t *testing.T) {
	cases := map[string]BatteryMonitorRecord{
		"ffffe50400000000030000f40140df03": {
			TTG:            math.NaN(),
			BatteryVoltage: 12.53,
			AlarmReason:    0,
			AuxVoltage:     math.NaN(),
			MidVoltage:     math.NaN(),
			Temperature:    math.NaN(),
			AuxMode:        veconst.BmvAuxModeDisabled,
			BatteryCurrent: 0.0,
			ConsumedAh:     -50.0,
			StateOfCharge:  50.0,
		},
		"ffffe6040000feff000000000080feac": {
			TTG:            math.NaN(),
			BatteryVoltage: 12.54,
			AlarmReason:    0,
			AuxVoltage:     -0.02,
			MidVoltage:     math.NaN(),
			Temperature:    math.NaN(),
			AuxMode:        veconst.BmvAuxModeStarterVoltage,
			BatteryCurrent: 0.0,
			ConsumedAh:     0.0,
			StateOfCharge:  100.0,
		},
		"ffffe6040000feff010000000080fe0c": {
			TTG:            math.NaN(),
			BatteryVoltage: 12.54,
			AlarmReason:    0,
			AuxVoltage:     math.NaN(),
			MidVoltage:     655.34,
			Temperature:    math.NaN(),
			AuxMode:        veconst.BmvAuxModeMidpointVoltage,
			BatteryCurrent: 0.0,
			ConsumedAh:     0.0,
			StateOfCharge:  100.0,
		},
		"ffffc60400007d73feff7fffffffff12": {
			TTG:            math.NaN(),
			BatteryVoltage: 12.22,
			AlarmReason:    0,
			AuxVoltage:     math.NaN(),
			MidVoltage:     math.NaN(),
			Temperature:    22.5,
			AuxMode:        veconst.BmvAuxModeTemperature,
			BatteryCurrent: math.NaN(),
			ConsumedAh:     math.NaN(),
			StateOfCharge:  math.NaN(),
		},
		"fffff80400008971feff7fffffffff5c": {
			TTG:            math.NaN(),
			BatteryVoltage: 12.72,
			AlarmReason:    0,
			AuxVoltage:     math.NaN(),
			MidVoltage:     math.NaN(),
			Temperature:    17.5,
			AuxMode:        veconst.BmvAuxModeTemperature,
			BatteryCurrent: math.NaN(),
			ConsumedAh:     math.NaN(),
			StateOfCharge:  math.NaN(),
		},
	}

	for stimuli, expResp := range cases {
		inp, err := hex.DecodeString(stimuli)
		if err != nil {
			t.Fatalf("hex.DecodeString failed: %s", err)
		}

		ret, err := DecodeBatteryMonitor(inp)
		if err != nil {
			t.Fatalf("DecodeBatteryMonitor failed: %s", err)
		}

		if expect, got := expResp.TTG, ret.TTG; !compF(expect, got) {
			t.Errorf("ret.TTG: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.BatteryVoltage, ret.BatteryVoltage; !compF(expect, got) {
			t.Errorf("ret.BatteryVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.AlarmReason, ret.AlarmReason; expect != got {
			t.Errorf("ret.AlarmReason: expect=%d but got=%d", expect, got)
		}
		if expect, got := expResp.AuxVoltage, ret.AuxVoltage; !compF(expect, got) {
			t.Errorf("ret.AuxVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.MidVoltage, ret.MidVoltage; !compF(expect, got) {
			t.Errorf("ret.MidVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.Temperature, ret.Temperature; !compF(expect, got) {
			t.Errorf("ret.Temperature: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.AuxMode, ret.AuxMode; expect != got {
			t.Errorf("ret.AuxMode: expect=%s but got=%s", expect, got)
		}
		if expect, got := expResp.BatteryCurrent, ret.BatteryCurrent; !compF(expect, got) {
			t.Errorf("ret.BatteryCurrent: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.ConsumedAh, ret.ConsumedAh; !compF(expect, got) {
			t.Errorf("ret.ConsumedAh: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.StateOfCharge, ret.StateOfCharge; !compF(expect, got) {
			t.Errorf("ret.StateOfCharge: expect=%f but got=%f", expect, got)
		}
	}
}
