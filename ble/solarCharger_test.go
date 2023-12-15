package ble

import (
	"encoding/hex"
	"github.com/koestler/go-victron/victronDefinitions"
	"math"
	"testing"
)

// this test is based on https://github.com/keshavdv/victron-ble/blob/e28c5f8cc5f9f3062a2f36c2115d38214c07e741/tests/test_solar_charger.py

func TestDecodeSolarChargeRecord(t *testing.T) {
	cases := map[string]SolarChargerRecord{
		"04006c050e000300130000fe409ac069": {
			DeviceState:    victronDefinitions.SolarChargerStateAbsorptionCharging,
			ChargerError:   victronDefinitions.SolarChargerErrorNoError,
			BatteryVoltage: 13.88,
			BatteryCurrent: 1.4,
			YieldToday:     30.0,
			PvPower:        19.0,
			LoadCurrent:    0.0,
		},
		"0300f80402000200030000fe8c9a5572": {
			DeviceState:    victronDefinitions.SolarChargerStateBulkCharging,
			ChargerError:   victronDefinitions.SolarChargerErrorNoError,
			BatteryVoltage: 12.72,
			BatteryCurrent: 0.2,
			YieldToday:     20.0,
			PvPower:        3.0,
			LoadCurrent:    0.0,
		},
		"0300fb09650032000901ffff31bc45ad": {
			DeviceState:    victronDefinitions.SolarChargerStateBulkCharging,
			ChargerError:   victronDefinitions.SolarChargerErrorNoError,
			BatteryVoltage: 25.55,
			BatteryCurrent: 10.1,
			YieldToday:     500.0,
			PvPower:        265.0,
			LoadCurrent:    math.NaN(),
		},
	}

	for stimuli, expResp := range cases {
		inp, err := hex.DecodeString(stimuli)
		if err != nil {
			t.Fatalf("hex.DecodeString failed: %s", err)
		}

		ret, err := DecodeSolarChargeRecord(inp)
		if err != nil {
			t.Fatalf("DecodeSolarChargeRecord failed: %s", err)
		}

		if expect, got := expResp.DeviceState, ret.DeviceState; expect != got {
			t.Errorf("ret.DeviceState: expect=%s but got=%s", expect, got)
		}
		if expect, got := expResp.ChargerError, ret.ChargerError; expect != got {
			t.Errorf("ret.ChargerError: expect=%s but got=%s", expect, got)
		}
		if expect, got := expResp.BatteryVoltage, ret.BatteryVoltage; !compF(expect, got) {
			t.Errorf("ret.BatteryVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.BatteryCurrent, ret.BatteryCurrent; !compF(expect, got) {
			t.Errorf("ret.BatteryCurrent: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.YieldToday, ret.YieldToday; !compF(expect, got) {
			t.Errorf("ret.YieldToday: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.PvPower, ret.PvPower; !compF(expect, got) {
			t.Errorf("ret.PvPower: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.LoadCurrent, ret.LoadCurrent; !compF(expect, got) {
			t.Errorf("ret.LoadCurrent: expect=%f but got=%f", expect, got)
		}
	}
}
