package veblerecord

import (
	"encoding/hex"
	"github.com/koestler/go-victron/veconst"
	"math"
	"testing"
)

// this test is based on https://github.com/keshavdv/victron-ble/blob/e28c5f8cc5f9f3062a2f36c2115d38214c07e741/tests/test_dcdc_converter.py

func TestDecodeDcDcConverter(t *testing.T) {
	cases := map[string]DcDcConverterRecord{
		"00002305ff7f80000000cbdd494cc5d1": {
			DeviceState:   veconst.DcDcConverterStateNotCharging,
			ChargerError:  veconst.DcDcConverterErrorNoError,
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

		ret, err := DecodeDcDcConverter(inp)
		if err != nil {
			t.Fatalf("DecodeDcDcConverter failed: %s", err)
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
