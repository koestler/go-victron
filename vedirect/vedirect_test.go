package vedirect_test

import (
	"errors"
	"github.com/koestler/go-victron/vedirect"
	"log"
	"testing"
)

func TestVedirect(t *testing.T) {
	io := NewLookupIOPort(t, map[string]string{
		":154\n":       ":51641F9\n",       // Ping() out of docs example
		":7F0ED0071\n": ":7F0ED009600DB\n", // GetUInt out of docs

		":451\n":       ":156A05E\n",                                                                 // GetDeviceId() = 0xA056
		":70001004D\n": ":70001000056A0FF58\n",                                                       // GetUint(0x100) = 4288697856
		":7F2ED006F\n": ":7F2ED00ACF9CA\n",                                                           // GetInt(0xEDF2) = -1620
		":70B010042\n": ":70B0100536D617274536F6C61722043686172676572204D505054203130302F333000BA\n", // GetString(0x10B) = SmartSolar Charger MPPT 100/30
		":7D0ED0091\n": "\r\nPID\t0xA056\r\nFW\t161\r\nSER#\tHQ19499RHC5\r\nV\t11880\r\nI\t0\r\n" +
			"VPV\t10\r\nPPV\t0\r\nCS\t0\r\nMPPT\t0\r\nOR\t0x00000001\r\nERR\t0\r\nLOAD\tON\r\nH19" +
			"\t4803\r\nH20\t0\r\nH21\t0\r\nH22\t0\r\nH23\t0\r\nHSDS\t280\r\nChecksum\t):7D0ED00000091\n", // GetUint(0xEDD0) = 0
		":777770060\n": ":77777015F\n", // Get{Uint, Int, String}(0x7777) -> unknown id
	})
	defer io.CheckEverythingHeard()
	defer io.CheckClosed()

	vd, err := vedirect.NewVedirect(&vedirect.Config{
		IOPort:      io,
		DebugLogger: log.Default(),
		IoLogger:    log.Default(),
	})
	if err != nil {
		t.Fatalf("cannot create vedirect: %v", err)
	}

	defer func() {
		if err := io.Close(); err != nil {
			t.Errorf("cannot close io: %v", err)
		}
	}()

	// Ping()
	if err := vd.Ping(); err != nil {
		t.Errorf("cannot ping: %v", err)
	}

	// GetUInt out of docs (Battery Maximum Current)
	if got, err := vd.GetUint(0xEDF0); err != nil {
		t.Errorf("cannot get 0xEDF0: %v", err)
	} else if expect := uint64(0x96); expect != got {
		t.Errorf("fetching Battery Maximum Current: expected 0x%x but 0x%x", expect, got)
	}

	// GetDeviceId() = 0xA056
	if got, err := vd.GetDeviceId(); err != nil {
		t.Errorf("cannot get device id: %v", err)
	} else if expect := uint16(0xA056); expect != got {
		t.Errorf("fetching device id: expected 0x%x but 0x%x", expect, got)
	}

	// GetUint(0x100) = 4288697856
	if got, err := vd.GetUint(0x100); err != nil {
		t.Errorf("cannot get 0x100: %v", err)
	} else if expect := uint64(4288697856); expect != got {
		t.Errorf("fetching 0x100: expected 0x%x but 0x%x", expect, got)
	}

	// GetInt(0xEDF2) = -1620
	if got, err := vd.GetInt(0xEDF2); err != nil {
		t.Errorf("cannot get 0xEDF2: %v", err)
	} else if expect := int64(-1620); expect != got {
		t.Errorf("fetching 0xEDF2: expected 0x%x but 0x%x", expect, got)
	}

	// GetString(0x10B) = SmartSolar Charger MPPT 100/30
	if got, err := vd.GetString(0x10B); err != nil {
		t.Errorf("cannot get 0x10B: %v", err)
	} else if expect := "SmartSolar Charger MPPT 100/30"; expect != got {
		t.Errorf("fetching 0x10B: expected %q but %q", expect, got)
	}

	// GetUint(0xEDD0) = 0
	if got, err := vd.GetUint(0xEDD0); err != nil {
		t.Errorf("cannot get 0xEDD0: %v", err)
	} else if expect := uint64(0); expect != got {
		t.Errorf("fetching 0xEDD0: expected 0x%x but 0x%x", expect, got)
	}

	// Get{Uint, Int, String}(0x7777) -> unknown id
	if _, err := vd.GetUint(0x7777); !errors.Is(err, vedirect.ErrUnknownId) {
		t.Errorf("GetUInt(invalid address): expected ErrUnknownId but got: %v", err)
	}
	if _, err := vd.GetInt(0x7777); !errors.Is(err, vedirect.ErrUnknownId) {
		t.Errorf("GetInt(invalid address): expected ErrUnknownId but got: %v", err)
	}
	if _, err := vd.GetString(0x7777); !errors.Is(err, vedirect.ErrUnknownId) {
		t.Errorf("GetString(invalid address): expected ErrUnknownId but got: %v", err)
	}
}
