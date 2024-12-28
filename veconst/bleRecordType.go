package veconst

type BleRecordType uint8

const (
	BleTestRecord          BleRecordType = 0x00
	BleSolarCharger        BleRecordType = 0x01
	BleBatteryMonitor      BleRecordType = 0x02
	BleInverter            BleRecordType = 0x03
	BleDCDCConverter       BleRecordType = 0x04
	BleSmartLithium        BleRecordType = 0x05
	BleInverterRS          BleRecordType = 0x06
	BleGXDevice            BleRecordType = 0x07
	BleACCharger           BleRecordType = 0x08
	BleSmartBatteryProtect BleRecordType = 0x09
	BleLynxSmartBMS        BleRecordType = 0x0a
	BleMultiRS             BleRecordType = 0x0b
	BleVEBus               BleRecordType = 0x0c
	BleDCEnergyMeter       BleRecordType = 0x0d
)

var bleRecordTypeMap = map[BleRecordType]string{
	BleTestRecord:          "Test record",
	BleSolarCharger:        "Solar charger",
	BleBatteryMonitor:      "Battery monitor",
	BleInverter:            "Inverter",
	BleDCDCConverter:       "DC/DC converter",
	BleSmartLithium:        "SmartLithium",
	BleInverterRS:          "Inverter RS",
	BleGXDevice:            "GX-Device",
	BleACCharger:           "AC charger",
	BleSmartBatteryProtect: "Smart Battery Protect",
	BleLynxSmartBMS:        "Lynx Smart BMS",
	BleMultiRS:             "Multi RS",
	BleVEBus:               "VE.Bus",
	BleDCEnergyMeter:       "DC energy meter",
}

func (rt BleRecordType) Exists() bool {
	_, ok := bleRecordTypeMap[rt]
	return ok
}

func (rt BleRecordType) String() string {
	if v, ok := bleRecordTypeMap[rt]; ok {
		return v
	}
	return ""
}
