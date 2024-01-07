package vedirectapi

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/koestler/go-victron/internal/vetest"
	"github.com/koestler/go-victron/vedirect"
	"sort"
	"strings"
	"testing"
	"time"
)

func runNewRegisterApiTest(t *testing.T, ioMap map[string]string, expected string) {
	io := vetest.NewLookupIOPort(t, ioMap)
	defer io.CheckEverythingHeard()
	defer io.CheckClosed()

	api, err := NewRegisterApi(io, vedirect.Config{})
	if err != nil {
		t.Errorf("cannot create api: %v", err)
		return
	}
	defer api.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	regs, err := api.ReadAllRegisters(ctx)
	if err != nil {
		t.Errorf("cannot read all registers: %v", err)
		return
	}

	// order list to make list deterministic
	list := make([]string, 0)
	for _, r := range regs.GetList() {
		list = append(list, r.String())
	}
	sort.Strings(list)
	got := strings.Join(list, "\n")

	t.Logf("got: %s", got)

	if diff := cmp.Diff(got, expected); len(diff) > 0 {
		t.Errorf("values list does not match :\nexpected:\n%s\n\ngot:\n%s\n\ndiff:\n%s",
			expected,
			got,
			diff,
		)
	}
}

func TestRegisterApi_ReadAllRegisters(t *testing.T) {
	t.Run("Solar Lab", func(t *testing.T) {
		ioMap := map[string]string{
			":154\n":       "\x00\r\nSER#\tHQ19499RHC5\r\nV\t11880\r\nI\t0\r\nVPV\t10\r\nPPV\t0\r\nCS\t0\r\nMPPT\t0\r\nOR\t0x00000001\r\nERR\t0\r\nLOAD\tON\r\nH19\t4803\r\nH20\t0\r\nH21\t0\r\nH22\t0\r\nH23\t0\r\nHSDS\t298\r\nChecksum\t :56141AE\n", // Ping()
			":451\n":       ":156A05E\n",                                                                                                                                                                                                                // GetDeviceId() = 0xA056
			":70001004D\n": ":70001000056A0FF58\n",                                                                                                                                                                                                      // GetUint(0x100) = 4288697856
			":704010049\n": ":70401000A3F\n",                                                                                                                                                                                                            // GetUint(0x104) = 10
			":7FDED0064\n": ":7FDED000064\n",                                                                                                                                                                                                            // GetUint(0xEDFD) = 0
			":7FBED0066\n": ":7FBED0058020C\n",                                                                                                                                                                                                          // GetUint(0xEDFB) = 600
			":7F7ED006A\n": ":7F7ED00A005C5\n",                                                                                                                                                                                                          // GetUint(0xEDF7) = 1440
			":7F6ED006B\n": ":7F6ED00640502\n",                                                                                                                                                                                                          // GetUint(0xEDF6) = 1380
			":7F4ED006D\n": ":7F4ED00540613\n",                                                                                                                                                                                                          // GetUint(0xEDF4) = 1620
			":7F2ED006F\n": ":7F2ED00ACF9CA\n",                                                                                                                                                                                                          // GetInt(0xEDF2) = -1620
			":7F0ED0071\n": ":7F0ED002C0144\n",                                                                                                                                                                                                          // GetUint(0xEDF0) = 300
			":7EFED0072\n": ":7EFED000C66\n",                                                                                                                                                                                                            // GetUint(0xEDEF) = 12
			":7E7ED007A\n": ":7E7ED00140066\n",                                                                                                                                                                                                          // GetUint(0xEDE7) = 20
			":7E6ED007B\n": ":7E6ED00FFFF7D\n",                                                                                                                                                                                                          // GetUint(0xEDE6) = 65535
			":7E4ED007D\n": ":7E4ED000875\n",                                                                                                                                                                                                            // GetUint(0xEDE4) = 8
			":7E3ED007E\n": ":7E3ED0064001A\n",                                                                                                                                                                                                          // GetUint(0xEDE3) = 100
			":72EED0033\n": ":72EED0028000B\n",                                                                                                                                                                                                          // GetUint(0xED2E) = 40
			":7E0ED0081\n": ":7E0ED00F4018C\n",                                                                                                                                                                                                          // GetInt(0xEDE0) = 500
			":7ECED0075\n": ":7ECED00FFFF77\n",                                                                                                                                                                                                          // GetUint(0xEDEC) = 65535
			":7DFED0082\n": ":7DFED002C0155\n",                                                                                                                                                                                                          // GetUint(0xEDDF) = 300
			":7DDED0084\n": ":7DDED00D734000079\n",                                                                                                                                                                                                      // GetUint(0xEDDD) = 13527
			":7DCED0085\n": ":7DCED00C3120000B0\n",                                                                                                                                                                                                      // GetUint(0xEDDC) = 4803
			":7DBED0086\n": ":7DBED00680B13\n",                                                                                                                                                                                                          // GetInt(0xEDDB) = 2920
			":7D7ED008A\n": ":7D7ED0000008A\n",                                                                                                                                                                                                          // GetUint(0xEDD7) = 0
			":7D5ED008C\n": ":7D5ED00A404E4\n",                                                                                                                                                                                                          // GetUint(0xEDD5) = 1188
			":7D3ED008E\n": ":7D3ED0000008E\n",                                                                                                                                                                                                          // GetUint(0xEDD3) = 0
			":7D2ED008F\n": ":7D2ED0000008F\n",                                                                                                                                                                                                          // GetUint(0xEDD2) = 0
			":7D1ED0090\n": ":7D1ED00000090\n",                                                                                                                                                                                                          // GetUint(0xEDD1) = 0
			":7D0ED0091\n": ":7D0ED00000091\n",                                                                                                                                                                                                          // GetUint(0xEDD0) = 0
			":7BCED00A5\n": ":7BCED0000000000A5\n",                                                                                                                                                                                                      // GetUint(0xEDBC) = 0
			":7BBED00A6\n": ":7BBED000100A5\n",                                                                                                                                                                                                          // GetUint(0xEDBB) = 1
			":7BDED00A4\n": ":7BDED000000A4\n",                                                                                                                                                                                                          // GetUint(0xEDBD) = 0
			":7B8ED00A9\n": ":7B8ED0074270E\n",                                                                                                                                                                                                          // GetUint(0xEDB8) = 10100
			":70A010043\n": ":70A01004851313934393952484335000000000088\n",                                                                                                                                                                              // GetString(0x10A) = HQ19499RHC5
			":70B010042\n": ":70B0100536D617274536F6C61722043686172676572204D505054203130302F333000BA\n",                                                                                                                                                // GetString(0x10B) = SmartSolar Charger MPPT 100/30
			":70002004C\n": ":7000200014B\n",                                                                                                                                                                                                            // GetUint(0x200) = 1
			":70102004B\n": ":7010200004B\n",                                                                                                                                                                                                            // GetUint(0x201) = 0
			":7FEED0063\n": ":7FEED000162\n",                                                                                                                                                                                                            // GetUint(0xEDFE) = 1
			":7F1ED0070\n": ":7F1ED00036D\n",                                                                                                                                                                                                            // GetUint(0xEDF1) = 3
			":7EAED0077\n": ":7EAED000C6B\n",                                                                                                                                                                                                            // GetUint(0xEDEA) = 12
			":7E8ED0079\n": ":7E8ED000079\n",                                                                                                                                                                                                            // GetUint(0xEDE8) = 0
			":7E5ED007C\n": ":7E5ED00017B\n",                                                                                                                                                                                                            // GetUint(0xEDE5) = 1
			":7DAED0087\n": ":7DAED000087\n",                                                                                                                                                                                                            // GetUint(0xEDDA) = 0
			":7B3ED00AE\n": ":7B3ED0000AE\n",                                                                                                                                                                                                            // GetUint(0xEDB3) = 0
			":707020045\n": ":70702000100000044\n",                                                                                                                                                                                                      // GetUint(0x207) = 1
		}

		expected := strings.Join([]string{
			"AdaptiveMode=1:On",
			"AutoEqualiseStop=1:Yes",
			"AutomaticEqualisationMode=0.000000",
			"BatteryAbsorptionTimeLimit=6.000000h",
			"BatteryAbsorptionVoltage=14.400000V",
			"BatteryEqualisationVoltage=16.200000V",
			"BatteryFloatVoltage=13.800000V",
			"BatteryLowTemperatureLevel=5.000000°C",
			"BatteryMaximumCurrent=30.000000A",
			"BatteryTempCompensation=-16.200000mV/K",
			"BatteryTemperature=382.200000°C",
			"BatteryType=3:Gel Victron Deep discharge (14.4V)",
			"BatteryVoltage=12.000000V",
			"BatteryVoltageSetting=12:12V battery",
			"BmsPresent=0:No",
			"ChargerCurrent=0.000000A",
			"ChargerErrorCode=0:No error",
			"ChargerInternalTemperature=29.200000°C",
			"ChargerMaximumCurrent=30.000000A",
			"ChargerVoltage=11.880000V",
			"DeviceMode=1:Charger On",
			"EqualisationCurrentLevel=8.000000%",
			"EqualisationDuration=1.000000h",
			"GroupId=10.000000",
			"LowTempCurrent=6553.500000A",
			"MaximumPowerToday=0.000000W",
			"MaximumPowerYesterday=0.000000W",
			"ModelName=SmartSolar Charger MPPT 100/30",
			"OffReason=No input power",
			"PanelCurrent=0.000000A",
			"PanelMaximumVoltage=101.000000V",
			"PanelPower=0.000000W",
			"PanelVoltage=0.010000V",
			"ProductId=4288697856.000000",
			"ReBulkVoltageOffset=0.400000V",
			"SerialNumber=HQ19499RHC5",
			"State=0:Not charging",
			"SystemYield=135.270000kWh",
			"SystemYieldResettable=48.030000kWh",
			"TailCurrent=2.000000A",
			"TrackerMode=0:Off",
			"YieldToday=0.000000kWh",
			"YieldYesterday=0.000000kWh",
		}, "\n")

		runNewRegisterApiTest(t, ioMap, expected)
	})

}
