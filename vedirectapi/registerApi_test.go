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
			":154\n": "\x00\r\nSER#\tHQ19499RHC5\r\nV\t11880\r\nI\t0\r\nVPV\t10\r\nPPV\t0\r\nCS\t0\r\n" +
				"MPPT\t0\r\nOR\t0x00000001\r\nERR\t0\r\nLOAD\tON\r\nH19\t4803\r\nH20\t0\r\nH21\t0\r\nH22\t" +
				"0\r\nH23\t0\r\nHSDS\t298\r\nChecksum\t :56141AE\n", // Ping()
			":451\n":       ":156A05E\n",                                                                 // GetDeviceId() = 0xA056
			":70001004D\n": ":70001000056A0FF58\n",                                                       // GetUint(0x100) = 4288697856
			":704010049\n": ":70401000A3F\n",                                                             // GetUint(0x104) = 10
			":7FDED0064\n": ":7FDED000064\n",                                                             // GetUint(0xEDFD) = 0
			":7FBED0066\n": ":7FBED0058020C\n",                                                           // GetUint(0xEDFB) = 600
			":7F7ED006A\n": ":7F7ED00A005C5\n",                                                           // GetUint(0xEDF7) = 1440
			":7F6ED006B\n": ":7F6ED00640502\n",                                                           // GetUint(0xEDF6) = 1380
			":7F4ED006D\n": ":7F4ED00540613\n",                                                           // GetUint(0xEDF4) = 1620
			":7F2ED006F\n": ":7F2ED00ACF9CA\n",                                                           // GetInt(0xEDF2) = -1620
			":7F0ED0071\n": ":7F0ED002C0144\n",                                                           // GetUint(0xEDF0) = 300
			":7EFED0072\n": ":7EFED000C66\n",                                                             // GetUint(0xEDEF) = 12
			":7E7ED007A\n": ":7E7ED00140066\n",                                                           // GetUint(0xEDE7) = 20
			":7E6ED007B\n": ":7E6ED00FFFF7D\n",                                                           // GetUint(0xEDE6) = 65535
			":7E4ED007D\n": ":7E4ED000875\n",                                                             // GetUint(0xEDE4) = 8
			":7E3ED007E\n": ":7E3ED0064001A\n",                                                           // GetUint(0xEDE3) = 100
			":72EED0033\n": ":72EED0028000B\n",                                                           // GetUint(0xED2E) = 40
			":7E0ED0081\n": ":7E0ED00F4018C\n",                                                           // GetInt(0xEDE0) = 500
			":7ECED0075\n": ":7ECED00FFFF77\n",                                                           // GetUint(0xEDEC) = 65535
			":7DFED0082\n": ":7DFED002C0155\n",                                                           // GetUint(0xEDDF) = 300
			":7DDED0084\n": ":7DDED00D734000079\n",                                                       // GetUint(0xEDDD) = 13527
			":7DCED0085\n": ":7DCED00C3120000B0\n",                                                       // GetUint(0xEDDC) = 4803
			":7DBED0086\n": ":7DBED00680B13\n",                                                           // GetInt(0xEDDB) = 2920
			":7D7ED008A\n": ":7D7ED0000008A\n",                                                           // GetUint(0xEDD7) = 0
			":7D5ED008C\n": ":7D5ED00A404E4\n",                                                           // GetUint(0xEDD5) = 1188
			":7D3ED008E\n": ":7D3ED0000008E\n",                                                           // GetUint(0xEDD3) = 0
			":7D2ED008F\n": ":7D2ED0000008F\n",                                                           // GetUint(0xEDD2) = 0
			":7D1ED0090\n": ":7D1ED00000090\n",                                                           // GetUint(0xEDD1) = 0
			":7D0ED0091\n": ":7D0ED00000091\n",                                                           // GetUint(0xEDD0) = 0
			":7BCED00A5\n": ":7BCED0000000000A5\n",                                                       // GetUint(0xEDBC) = 0
			":7BBED00A6\n": ":7BBED000100A5\n",                                                           // GetUint(0xEDBB) = 1
			":7BDED00A4\n": ":7BDED000000A4\n",                                                           // GetUint(0xEDBD) = 0
			":7B8ED00A9\n": ":7B8ED0074270E\n",                                                           // GetUint(0xEDB8) = 10100
			":70A010043\n": ":70A01004851313934393952484335000000000088\n",                               // GetString(0x10A) = HQ19499RHC5
			":70B010042\n": ":70B0100536D617274536F6C61722043686172676572204D505054203130302F333000BA\n", // GetString(0x10B) = SmartSolar Charger MPPT 100/30
			":70002004C\n": ":7000200014B\n",                                                             // GetUint(0x200) = 1
			":70102004B\n": ":7010200004B\n",                                                             // GetUint(0x201) = 0
			":7FEED0063\n": ":7FEED000162\n",                                                             // GetUint(0xEDFE) = 1
			":7F1ED0070\n": ":7F1ED00036D\n",                                                             // GetUint(0xEDF1) = 3
			":7EAED0077\n": ":7EAED000C6B\n",                                                             // GetUint(0xEDEA) = 12
			":7E8ED0079\n": ":7E8ED000079\n",                                                             // GetUint(0xEDE8) = 0
			":7E5ED007C\n": ":7E5ED00017B\n",                                                             // GetUint(0xEDE5) = 1
			":7DAED0087\n": ":7DAED000087\n",                                                             // GetUint(0xEDDA) = 0
			":7B3ED00AE\n": ":7B3ED0000AE\n",                                                             // GetUint(0xEDB3) = 0
			":707020045\n": ":70702000100000044\n",                                                       // GetUint(0x207) = 1
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

	t.Run("Solar Piegn Main", func(t *testing.T) {
		ioMap := map[string]string{
			":451\n":       ":169A04B\n",           // GetDeviceId() = 0xA069
			":154\n":       ":56141AE\n",           // Ping()
			":70001004D\n": ":70001000069A0FF45\n", // GetUint(0x100) = 4288702720
			":704010049\n": ":70401000A3F\n",       // GetUint(0x104) = 10
			":7FDED0064\n": ":7FDED000064\n",       // GetUint(0xEDFD) = 0
			":7FBED0066\n": ":7FBED00C8009E\n",     // GetUint(0xEDFB) = 200
			":7F7ED006A\n": ":7F7ED00180B47\n",     // GetUint(0xEDF7) = 2840
			":7F6ED006B\n": ":7F6ED008C0AD5\n",     // GetUint(0xEDF6) = 2700
			":7F4ED006D\n": ":7F4ED0000006D\n",     // GetUint(0xEDF4) = 0
			":7F2ED006F\n": ":7F2ED0000006F\n",     // GetInt(0xEDF2) = 0
			":7F0ED0071\n": ":7F0ED00BC02B3\n",     // GetUint(0xEDF0) = 700
			":7EFED0072\n": ":7EFED00185A\n",       // GetUint(0xEDEF) = 24
			":7E7ED007A\n": ":7E7ED0000007A\n",     // GetUint(0xEDE7) = 0
			":7E6ED007B\n": "\r\nPID\t0xA069\r\nFW\t161\r\nSER#\tHQ2104VESU7\r\nV\t27350\r\nI\t5100\r\nVPV\t" +
				"106630\r\nPPV\t142\r\nCS\t3\r\nMPPT\t2\r\nOR\t0x00000000\r\nERR\t0\r\nLOAD\tOFF\r\nRelay\t" +
				"OFF\r\nH19\t62753\r\nH20\t20\r\nH21\t141\r\nH22\t67\r\nH23\t322\r\nHSDS\t259\r\nChecksum\t" +
				"\x87:7E6ED0000007B\n", // GetUint(0xEDE6) = 0
			":7E4ED007D\n": ":7E4ED00007D\n",                                                                       // GetUint(0xEDE4) = 0
			":7E3ED007E\n": ":7E3ED0000007E\n",                                                                     // GetUint(0xEDE3) = 0
			":72EED0033\n": ":72EED0014001F\n",                                                                     // GetUint(0xED2E) = 20
			":7E0ED0081\n": ":7E0ED00F4018C\n",                                                                     // GetInt(0xEDE0) = 500
			":7ECED0075\n": ":7ECED00316FD5\n",                                                                     // GetUint(0xEDEC) = 28465
			":7DFED0082\n": ":7DFED00BC02C4\n",                                                                     // GetUint(0xEDDF) = 700
			":7DDED0084\n": ":7DDED0021F500006E\n",                                                                 // GetUint(0xEDDD) = 62753
			":7DCED0085\n": ":7DCED0021F500006F\n",                                                                 // GetUint(0xEDDC) = 62753
			":7DBED0086\n": ":7DBED00EE0791\n",                                                                     // GetInt(0xEDDB) = 2030
			":7D7ED008A\n": ":7D7ED002F005B\n",                                                                     // GetUint(0xEDD7) = 47
			":7D5ED008C\n": ":7D5ED00AB0AD7\n",                                                                     // GetUint(0xEDD5) = 2731
			":7D3ED008E\n": ":7D3ED0014007A\n",                                                                     // GetUint(0xEDD3) = 20
			":7D2ED008F\n": ":7D2ED008D0002\n",                                                                     // GetUint(0xEDD2) = 141
			":7D1ED0090\n": ":A032000FF7FAA\n:7D1ED0043004D\n",                                                     // GetUint(0xEDD1) = 67
			":7D0ED0091\n": ":A2720007837000055\n:7D0ED0042014E\n",                                                 // GetUint(0xEDD0) = 322
			":7BCED00A5\n": ":A132000FFFFFF7F9C\n:7BCED0095360000DA\n",                                             // GetUint(0xEDBC) = 13973
			":7BBED00A6\n": ":AD5ED00AF0AD0\n:7BBED00C52AB7\n",                                                     // GetUint(0xEDBB) = 10949
			":7BDED00A4\n": ":ABBED00C52AB4\n:ABCED008B360000E1\n:AD7ED00310056\n:7BDED000C0098\n",                 // GetUint(0xEDBD) = 12
			":7B8ED00A9\n": ":ABDED000C0095\n:7B8ED00A861A0\n",                                                     // GetUint(0xEDB8) = 25000
			":70A010043\n": ":70A01004851323130345645535537000000000069\n",                                         // GetString(0x10A) = HQ2104VESU7
			":70B010042\n": ":70B0100536D617274536F6C61722043686172676572204D505054203235302F373020726576320011\n", // GetString(0x10B) = SmartSolar Charger MPPT 250/70 rev2
			":70002004C\n": ":7000200014B\n",                                                                       // GetUint(0x200) = 1
			":70102004B\n": ":70102000348\n",                                                                       // GetUint(0x201) = 3
			":7FEED0063\n": ":7FEED000063\n",                                                                       // GetUint(0xEDFE) = 0
			":7F1ED0070\n": ":7F1ED000868\n",                                                                       // GetUint(0xEDF1) = 8
			":7EAED0077\n": ":7EAED00185F\n",                                                                       // GetUint(0xEDEA) = 24
			":7E8ED0079\n": ":7E8ED000079\n",                                                                       // GetUint(0xEDE8) = 0
			":7E5ED007C\n": ":7E5ED00017B\n",                                                                       // GetUint(0xEDE5) = 1
			":7DAED0087\n": ":7DAED000087\n",                                                                       // GetUint(0xEDDA) = 0
			":7B3ED00AE\n": ":7B3ED0002AC\n",                                                                       // GetUint(0xEDB3) = 2
			":707020045\n": ":70702000000000045\n",                                                                 // GetUint(0x207) = 0// GetUint(0x207) = 0
		}

		expected := strings.Join([]string{
			"AdaptiveMode=0:Off",
			"AutoEqualiseStop=1:Yes",
			"AutomaticEqualisationMode=0.000000",
			"BatteryAbsorptionTimeLimit=2.000000h",
			"BatteryAbsorptionVoltage=28.400000V",
			"BatteryEqualisationVoltage=0.000000V",
			"BatteryFloatVoltage=27.000000V",
			"BatteryLowTemperatureLevel=5.000000°C",
			"BatteryMaximumCurrent=70.000000A",
			"BatteryTempCompensation=0.000000mV/K",
			"BatteryTemperature=11.500000°C",
			"BatteryType=8:LiFEPO4 (14.2V)",
			"BatteryVoltage=24.000000V",
			"BatteryVoltageSetting=24:24V battery",
			"BmsPresent=0:No",
			"ChargerCurrent=4.700000A",
			"ChargerErrorCode=0:No error",
			"ChargerInternalTemperature=20.300000°C",
			"ChargerMaximumCurrent=70.000000A",
			"ChargerVoltage=27.310000V",
			"DeviceMode=1:Charger On",
			"EqualisationCurrentLevel=0.000000%",
			"EqualisationDuration=0.000000h",
			"GroupId=10.000000",
			"LowTempCurrent=0.000000A",
			"MaximumPowerToday=141.000000W",
			"MaximumPowerYesterday=322.000000W",
			"ModelName=SmartSolar Charger MPPT 250/70 rev2",
			"OffReason=",
			"PanelCurrent=1.200000A",
			"PanelMaximumVoltage=250.000000V",
			"PanelPower=139.730000W",
			"PanelVoltage=109.490000V",
			"ProductId=4288702720.000000",
			"ReBulkVoltageOffset=0.200000V",
			"SerialNumber=HQ2104VESU7",
			"State=3:Bulk Charging",
			"SystemYield=627.530000kWh",
			"SystemYieldResettable=627.530000kWh",
			"TailCurrent=0.000000A",
			"TrackerMode=2:MPP tracker",
			"YieldToday=0.200000kWh",
			"YieldYesterday=0.670000kWh",
		}, "\n")

		runNewRegisterApiTest(t, ioMap, expected)
	})

	t.Run("Solar Piegn Aux", func(t *testing.T) {
		ioMap := map[string]string{
			":451\n":       ":14BA069\n",           // GetDeviceId() = 0xA04B
			":154\n":       ":56141AE\n",           // Ping()
			":70001004D\n": ":7000100004BA0FF63\n", // GetUint(0x100) = 4288695040
			":704010049\n": ":70401000A3F\n",       // GetUint(0x104) = 10
			":7FDED0064\n": ":7FDED000064\n",       // GetUint(0xEDFD) = 0
			":7FBED0066\n": ":7FBED0058020C\n",     // GetUint(0xEDFB) = 600
			":7F7ED006A\n": ":7F7ED00400B1F\n",     // GetUint(0xEDF7) = 2880
			":7F6ED006B\n": ":7F6ED00C80A99\n",     // GetUint(0xEDF6) = 2760
			":7F4ED006D\n": ":7F4ED00A80CB9\n",     // GetUint(0xEDF4) = 3240
			":7F2ED006F\n": ":7F2ED0058F324\n",     // GetInt(0xEDF2) = -3240
			":7F0ED0071\n": ":7F0ED005E0112\n",     // GetUint(0xEDF0) = 350
			":7EFED0072\n": ":7EFED00185A\n",       // GetUint(0xEDEF) = 24
			":7E7ED007A\n": ":7E7ED00140066\n",     // GetUint(0xEDE7) = 20
			":7E6ED007B\n": ":7E6ED00FFFF7D\n",     // GetUint(0xEDE6) = 65535
			":7E4ED007D\n": ":7E4ED000875\n",       // GetUint(0xEDE4) = 8
			":7E3ED007E\n": ":7E3ED0064001A\n",     // GetUint(0xEDE3) = 100
			":72EED0033\n": ":72EED005000E3\n",     // GetUint(0xED2E) = 80
			":7E0ED0081\n": ":7E0ED00F4018C\n",     // GetInt(0xEDE0) = 500
			":7ECED0075\n": ":7ECED00FFFF77\n",     // GetUint(0xEDEC) = 65535
			":7DFED0082\n": ":7DFED005E0123\n",     // GetUint(0xEDDF) = 350
			":7DDED0084\n": ":7DDED00F427060063\n", // GetUint(0xEDDD) = 403444
			":7DCED0085\n": "\r\nPID\t0xA04B\r\nFW\t161\r\nSER#\tHQ1546WBRK9\r\nV\t24510\r\nI\t300\r\nVPV\t50230\r\n" +
				"PPV\t8\r\nCS\t3\r\nMPPT\t2\r\nOR\t0x00000000\r\nERR\t0\r\nLOAD\tOFF\r\nH19\t39080\r\nH20\t29\r\n" +
				"H21\t75\r\nH22\t33\r\nH23\t189\r\nHSDS\t28\r\nChecksum\t\x9d:7DCED00A898000045\n", // GetUint(0xEDDC) = 39080
			":7DBED0086\n": "30084\n:AD5ED009309ED\n:7DBED00B608C8\n",                                            // GetInt(0xEDDB) = 2230
			":7D7ED008A\n": ":ABCED00540300004B\n:ABBED006B1325\n:ABDED0002009F\n:7D7ED00040086\n",               // GetUint(0xEDD7) = 4
			":7D5ED008C\n": ":7D5ED009309F0\n",                                                                   // GetUint(0xEDD5) = 2451
			":7D3ED008E\n": ":7D3ED001D0071\n",                                                                   // GetUint(0xEDD3) = 29
			":7D2ED008F\n": ":7D2ED004B0044\n",                                                                   // GetUint(0xEDD2) = 75
			":7D1ED0090\n": ":7D1ED0021006F\n",                                                                   // GetUint(0xEDD1) = 33
			":7D0ED0091\n": ":7D0ED00BD00D4\n",                                                                   // GetUint(0xEDD0) = 189
			":7BCED00A5\n": ":7BCED00A003000002\n",                                                               // GetUint(0xEDBC) = 928
			":7BBED00A6\n": ":7BBED00F110A5\n",                                                                   // GetUint(0xEDBB) = 4337
			":7BDED00A4\n": ":7BDED000200A2\n",                                                                   // GetUint(0xEDBD) = 2
			":7B8ED00A9\n": ":7B8ED00983AD7\n",                                                                   // GetUint(0xEDB8) = 15000
			":70A010043\n": ":70A01004851313534365742524B3900000000006B\n",                                       // GetString(0x10A) = HQ1546WBRK9
			":70B010042\n": ":70B0100426C7565536F6C61722043686172676572204D505054203135302F333520726576320090\n", // GetString(0x10B) = BlueSolar Charger MPPT 150/35 rev2
			":70002004C\n": ":7000200014B\n",                                                                     // GetUint(0x200) = 1
			":70102004B\n": ":70102000348\n",                                                                     // GetUint(0x201) = 3
			":7FEED0063\n": ":7FEED000162\n",                                                                     // GetUint(0xEDFE) = 1
			":7F1ED0070\n": ":7F1ED00036D\n",                                                                     // GetUint(0xEDF1) = 3
			":7EAED0077\n": ":7EAED00185F\n",                                                                     // GetUint(0xEDEA) = 24
			":7E8ED0079\n": ":7E8ED000079\n",                                                                     // GetUint(0xEDE8) = 0
			":7E5ED007C\n": ":7E5ED00017B\n",                                                                     // GetUint(0xEDE5) = 1
			":7DAED0087\n": ":7DAED000087\n",                                                                     // GetUint(0xEDDA) = 0
			":7B3ED00AE\n": ":7B3ED0002AC\n",                                                                     // GetUint(0xEDB3) = 2
			":707020045\n": ":70702000000000045\n",                                                               // GetUint(0x207) = 0
		}

		expected := strings.Join([]string{
			"AdaptiveMode=1:On",
			"AutoEqualiseStop=1:Yes",
			"AutomaticEqualisationMode=0.000000",
			"BatteryAbsorptionTimeLimit=6.000000h",
			"BatteryAbsorptionVoltage=28.800000V",
			"BatteryEqualisationVoltage=32.400000V",
			"BatteryFloatVoltage=27.600000V",
			"BatteryLowTemperatureLevel=5.000000°C",
			"BatteryMaximumCurrent=35.000000A",
			"BatteryTempCompensation=-32.400000mV/K",
			"BatteryTemperature=382.200000°C",
			"BatteryType=3:Gel Victron Deep discharge (14.4V)",
			"BatteryVoltage=24.000000V",
			"BatteryVoltageSetting=24:24V battery",
			"BmsPresent=0:No",
			"ChargerCurrent=0.400000A",
			"ChargerErrorCode=0:No error",
			"ChargerInternalTemperature=22.300000°C",
			"ChargerMaximumCurrent=35.000000A",
			"ChargerVoltage=24.510000V",
			"DeviceMode=1:Charger On",
			"EqualisationCurrentLevel=8.000000%",
			"EqualisationDuration=1.000000h",
			"GroupId=10.000000",
			"LowTempCurrent=6553.500000A",
			"MaximumPowerToday=75.000000W",
			"MaximumPowerYesterday=189.000000W",
			"ModelName=BlueSolar Charger MPPT 150/35 rev2",
			"OffReason=",
			"PanelCurrent=0.200000A",
			"PanelMaximumVoltage=150.000000V",
			"PanelPower=9.280000W",
			"PanelVoltage=43.370000V",
			"ProductId=4288695040.000000",
			"ReBulkVoltageOffset=0.800000V",
			"SerialNumber=HQ1546WBRK9",
			"State=3:Bulk Charging",
			"SystemYield=4034.440000kWh",
			"SystemYieldResettable=390.800000kWh",
			"TailCurrent=2.000000A",
			"TrackerMode=2:MPP tracker",
			"YieldToday=0.290000kWh",
			"YieldYesterday=0.330000kWh",
		}, "\n")

		runNewRegisterApiTest(t, ioMap, expected)
	})

	t.Run("BMV Piegn Main", func(t *testing.T) {
		ioMap := map[string]string{
			":451\n":       ":189A328\n",                                                     // GetDeviceId() = 0xA389
			":154\n":       ":51444F8\n",                                                     // Ping()
			":70001004D\n": ":70001000089A3FE23\n",                                           // GetUint(0x100) = 4272130304
			":72001002D\n": ":7200100732D3E004F\n",                                           // GetUint(0x120) = 4074867
			":78EED00D3\n": ":78EED00FCFFD8\n",                                               // GetInt(0xED8E) = -4
			":78CED00D5\n": ":78CED006AFFFFFF6E\n",                                           // GetInt(0xED8C) = -150
			":78DED00D4\n": ":78DED006D0A5D\n",                                               // GetInt(0xED8D) = 2669
			":77DED00E4\n": ":77DED00FF7F66\n",                                               // GetUint(0xED7D) = 32767
			":7FFEE0061\n": ":7FFEE00FAFAFFFF6F\n",                                           // GetInt(0xEEFF) = -1286
			":7FF0F0040\n": ":7FF0F004F1BD6\n",                                               // GetUint(0xFFF) = 6991
			":7FE0F0041\n": ":7FE0F004038C9\n",                                               // GetUint(0xFFE) = 14400
			":7ECED0075\n": ":7ECED00FFFF77\n",                                               // GetUint(0xEDEC) = 65535
			":7820300C9\n": ":7820300FFFFCB\n",                                               // GetUint(0x382) = 65535
			":7830300C8\n": ":7830300FF7F4A\n",                                               // GetInt(0x383) = 32767
			":70003004B\n": ":70003003EF0FFFF1F\n",                                           // GetInt(0x300) = -4034
			":70103004A\n": ":70103006CF6FFFFEA\n",                                           // GetInt(0x301) = -2452
			":702030049\n": ":7020300BDF6FFFF98\n",                                           // GetInt(0x302) = -2371
			":703030048\n": ":70303000C0000003C\n",                                           // GetUint(0x303) = 12
			":704030047\n": ":70403000100000046\n",                                           // GetUint(0x304) = 1
			":705030046\n": ":70503004F11FDFFEA\n",                                           // GetInt(0x305) = -192177
			":706030045\n": ":7060300C20800007B\n",                                           // GetUint(0x306) = 2242
			":707030044\n": ":70703003F0B0000FA\n",                                           // GetUint(0x307) = 2879
			":708030043\n": ":70803008FF70300BA\n",                                           // GetUint(0x308) = 259983
			":709030042\n": ":70903005F000000E3\n",                                           // GetUint(0x309) = 95
			":70A030041\n": ":70A03000000000041\n",                                           // GetUint(0x30A) = 0
			":70B030040\n": ":70B03000000000040\n",                                           // GetUint(0x30B) = 0
			":70E03003D\n": ":70E0300000000003D\n",                                           // GetInt(0x30E) = 0
			":70F03003C\n": ":70F0300000000003C\n",                                           // GetInt(0x30F) = 0
			":71003003B\n": ":710030048C600002D\n",                                           // GetUint(0x310) = 50760
			":71103003A\n": ":711030050C9000021\n",                                           // GetUint(0x311) = 51536
			":70A010043\n": ":70A01004851323131374E54565834000000000000000000000000005B\n",   // GetString(0x10A) = HQ2117NTVX4
			":70B010042\n": ":70B0100424D562D536D6172745368756E7420353030412F35306D5600CA\n", // GetString(0x10B) = BMV-SmartShunt 500A/50mV
			":7B6EE00AA\n": ":7B6EE0001A9\n",                                                 // GetUint(0xEEB6) = 1
		}

		expected := strings.Join([]string{
			"AmountOfChargedEnergy=515.360000kWh",
			"AmountOfDischargedEnergy=507.600000kWh",
			"AuxVoltage=327.670000V",
			"AuxVoltageMaximum=0.000000V",
			"AuxVoltageMinimum=0.000000V",
			"BatteryTemperature=382.200000°C",
			"Consumed=-128.600000Ah",
			"CumulativeAmpHours=-19217.700000Ah",
			"CurrentHighRes=-0.150000A",
			"DepthOfTheAverageDischarge=-237.100000Ah",
			"DepthOfTheDeepestDischarge=-403.400000Ah",
			"DepthOfTheLastDischarge=-245.200000Ah",
			"MainVoltage=26.690000V",
			"MainVoltageMaximum=28.790000V",
			"MainVoltageMinimum=22.420000V",
			"MidPointVoltage=655.350000V",
			"MidPointVoltageDeviation=3276.700000%",
			"ModelName=BMV-SmartShunt 500A/50mV",
			"NumberOfAutomaticSynchronizations=95.000000",
			"NumberOfCycles=12.000000",
			"NumberOfFullDischarges=1.000000",
			"NumberOfHighMainVoltageAlarms=0.000000",
			"NumberOfLowMainVoltageAlarms=0.000000",
			"Power=-4.000000W",
			"ProductId=4272130304.000000",
			"SOC=69.910000%",
			"SerialNumber=HQ2117NTVX4",
			"SynchronizationState=1:True",
			"TTG=14400.000000min",
			"TimeSinceFullCharge=259983.000000s",
			"Uptime=4074867.000000s",
		}, "\n")

		runNewRegisterApiTest(t, ioMap, expected)
	})

	t.Run("BMV Piegn Aux", func(t *testing.T) {
		ioMap := map[string]string{
			":451\n":       ":189A328\n",                                                     // GetDeviceId() = 0xA389
			":154\n":       ":51444F8\n",                                                     // Ping()
			":70001004D\n": ":70001000089A3FE23\n",                                           // GetUint(0x100) = 4272130304
			":72001002D\n": ":720010078464E0120\n",                                           // GetUint(0x120) = 21907064
			":78EED00D3\n": ":78EED00D6FFFE\n",                                               // GetInt(0xED8E) = -42
			":78CED00D5\n": ":78CED004AF9FFFF94\n",                                           // GetInt(0xED8C) = -1718
			":78DED00D4\n": ":78DED008F093C\n",                                               // GetInt(0xED8D) = 2447
			":77DED00E4\n": ":77DED00FF7F66\n",                                               // GetUint(0xED7D) = 32767
			":7FFEE0061\n": ":7FFEE00FEFEFFFF67\n",                                           // GetInt(0xEEFF) = -258
			":7FF0F0040\n": ":7FF0F001A2105\n",                                               // GetUint(0xFFF) = 8474
			":7FE0F0041\n": ":7FE0F007907C1\n",                                               // GetUint(0xFFE) = 1913
			":7ECED0075\n": ":7ECED00546FB2\n",                                               // GetUint(0xEDEC) = 28500
			":7820300C9\n": ":7820300FFFFCB\n",                                               // GetUint(0x382) = 65535
			":7830300C8\n": ":7830300FF7F4A\n",                                               // GetInt(0x383) = 32767
			":70003004B\n": ":70003003BFBFFFF17\n",                                           // GetInt(0x300) = -1221
			":70103004A\n": ":70103009CFDFFFFB3\n",                                           // GetInt(0x301) = -612
			":702030049\n": ":702030068FEFFFFE5\n",                                           // GetInt(0x302) = -408
			":703030048\n": ":7030300190000002F\n",                                           // GetUint(0x303) = 25
			":704030047\n": ":70403000000000047\n",                                           // GetUint(0x304) = 0
			":705030046\n": ":70503004644FEFFBF\n",                                           // GetInt(0x305) = -113594
			":706030045\n": ":70603002A08000013\n",                                           // GetUint(0x306) = 2090
			":707030044\n": ":7070300BB0B00007E\n",                                           // GetUint(0x307) = 3003
			":708030043\n": ":7080300CFF603007B\n",                                           // GetUint(0x308) = 259791
			":709030042\n": ":7090300F70000004B\n",                                           // GetUint(0x309) = 247
			":70A030041\n": ":70A03000000000041\n",                                           // GetUint(0x30A) = 0
			":70B030040\n": ":70B03000000000040\n",                                           // GetUint(0x30B) = 0
			":70E03003D\n": ":70E030047040000F2\n",                                           // GetInt(0x30E) = 1095
			":70F03003C\n": ":70F0300C70600006F\n",                                           // GetInt(0x30F) = 1735
			":71003003B\n": ":7100300006D0000CE\n",                                           // GetUint(0x310) = 27904
			":71103003A\n": ":7110300CA760000FA\n",                                           // GetUint(0x311) = 30410
			":70A010043\n": ":70A0100485132323039464D57524D0000000000000000000000000054\n",   // GetString(0x10A) = HQ2209FMWRM
			":70B010042\n": ":70B0100424D562D536D6172745368756E7420353030412F35306D5600CA\n", // GetString(0x10B) = BMV-SmartShunt 500A/50mV
			":7B6EE00AA\n": ":7B6EE0001A9\n",                                                 // GetUint(0xEEB6) = 1
		}

		expected := strings.Join([]string{
			"AmountOfChargedEnergy=304.100000kWh",
			"AmountOfDischargedEnergy=279.040000kWh",
			"AuxVoltage=327.670000V",
			"AuxVoltageMaximum=17.350000V",
			"AuxVoltageMinimum=10.950000V",
			"BatteryTemperature=11.850000°C",
			"Consumed=-25.800000Ah",
			"CumulativeAmpHours=-11359.400000Ah",
			"CurrentHighRes=-1.718000A",
			"DepthOfTheAverageDischarge=-40.800000Ah",
			"DepthOfTheDeepestDischarge=-122.100000Ah",
			"DepthOfTheLastDischarge=-61.200000Ah",
			"MainVoltage=24.470000V",
			"MainVoltageMaximum=30.030000V",
			"MainVoltageMinimum=20.900000V",
			"MidPointVoltage=655.350000V",
			"MidPointVoltageDeviation=3276.700000%",
			"ModelName=BMV-SmartShunt 500A/50mV",
			"NumberOfAutomaticSynchronizations=247.000000",
			"NumberOfCycles=25.000000",
			"NumberOfFullDischarges=0.000000",
			"NumberOfHighMainVoltageAlarms=0.000000",
			"NumberOfLowMainVoltageAlarms=0.000000",
			"Power=-42.000000W",
			"ProductId=4272130304.000000",
			"SOC=84.740000%",
			"SerialNumber=HQ2209FMWRM",
			"SynchronizationState=1:True",
			"TTG=1913.000000min",
			"TimeSinceFullCharge=259791.000000s",
			"Uptime=21907064.000000s",
		}, "\n")

		runNewRegisterApiTest(t, ioMap, expected)
	})
}
