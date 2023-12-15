package bleparser

import (
	"encoding/hex"
	"github.com/koestler/go-victron/victronDefinitions"
	"math"
	"testing"
)

// this test is based on https://github.com/keshavdv/victron-ble/blob/e28c5f8cc5f9f3062a2f36c2115d38214c07e741/tests/test_dcdc_converter.py

func TestDecodeDcDcConverterRecord(t *testing.T) {
	cases := map[string]DcDcConverterRecord{
		"00002305ff7f80000000cbdd494cc5d1": {
			DeviceState:   victronDefinitions.DcDcConverterStateNotCharging,
			ChargerError:  victronDefinitions.DcDcConverterErrorNoError,
			InputVoltage:  13.15,
			OutputVoltage: math.NaN(),
			OffReason:     0x00000080,
		},
	}

	for stimuli, expResp := range cases {
		inp, err := hex.DecodeString(stimuli)
		if err != nil {
			t.Fatalf("hex.DecodeString failed: %s", err)
		}

		ret, err := DecodeDcDcConverterRecord(inp)
		if err != nil {
			t.Fatalf("DecodeDcDcConverterRecord failed: %s", err)
		}

		if expect, got := expResp.DeviceState, ret.DeviceState; expect != got {
			t.Errorf("ret.DeviceState: expect=%s but got=%s", expect, got)
		}
		if expect, got := expResp.ChargerError, ret.ChargerError; expect != got {
			t.Errorf("ret.ChargerError: expect=%s but got=%s", expect, got)
		}
		if expect, got := expResp.InputVoltage, ret.InputVoltage; !compF(expect, got) {
			t.Errorf("ret.InputVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.OutputVoltage, ret.OutputVoltage; !compF(expect, got) {
			t.Errorf("ret.OutputVoltage: expect=%f but got=%f", expect, got)
		}
		if expect, got := expResp.OffReason, ret.OffReason; expect != got {
			t.Errorf("ret.OffReason: expect=%d but got=%d", expect, got)
		}
	}
}
