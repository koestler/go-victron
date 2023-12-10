package ble

import (
	"encoding/hex"
	"github.com/koestler/go-iotdevice/victronDefinitions"
	"testing"
)

// this test is based on https://github.com/keshavdv/victron-ble/blob/e28c5f8cc5f9f3062a2f36c2115d38214c07e741/tests/test_solar_charger.py

func TestDecodeSolarChargeRecord(t *testing.T) {
	inp, err := hex.DecodeString("04006c050e000300130000fe409ac069")
	if err != nil {
		t.Fatalf("hex.DecodeString failed: %s", err)
	}

	ret, err := DecodeSolarChargeRecord(inp)
	if err != nil {
		t.Fatalf("DecodeSolarChargeRecord failed: %s", err)
	}

	if expect, got := victronDefinitions.SolarChargerStateAbsorptionCharging, ret.DeviceState; expect != got {
		t.Errorf("ret.DeviceState: expect=%s but got=%s", expect, got)
	}
	if expect, got := victronDefinitions.SolarChargerErrorNoError, ret.ChargerError; expect != got {
		t.Errorf("ret.ChargerError: expect=%s but got=%s", expect, got)
	}
	if expect, got := 13.88, ret.BatteryVoltage; expect != got {
		t.Errorf("ret.BatteryVoltage: expect=%f but got=%f", expect, got)
	}
	if expect, got := 1.4, ret.BatteryCurrent; expect != got {
		t.Errorf("ret.BatteryCurrent: expect=%f but got=%f", expect, got)
	}
	if expect, got := 30.0, ret.YieldToday; expect != got {
		t.Errorf("ret.YieldToday: expect=%f but got=%f", expect, got)
	}
	if expect, got := 19.0, ret.PvPower; expect != got {
		t.Errorf("ret.PvPower: expect=%f but got=%f", expect, got)
	}
	if expect, got := 0.0, ret.LoadCurrent; expect != got {
		t.Errorf("ret.LoadCurrent: expect=%f but got=%f", expect, got)
	}
}
