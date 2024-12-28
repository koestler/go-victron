package veble

type RecordType uint8

const (
	BleTestRecord          RecordType = 0x00
	BleSolarCharger        RecordType = 0x01
	BleBatteryMonitor      RecordType = 0x02
	BleInverter            RecordType = 0x03
	BleDCDCConverter       RecordType = 0x04
	BleSmartLithium        RecordType = 0x05
	BleInverterRS          RecordType = 0x06
	BleGXDevice            RecordType = 0x07
	BleACCharger           RecordType = 0x08
	BleSmartBatteryProtect RecordType = 0x09
	BleLynxSmartBMS        RecordType = 0x0a
	BleMultiRS             RecordType = 0x0b
	BleVEBus               RecordType = 0x0c
	BleDCEnergyMeter       RecordType = 0x0d
)

var recordTypeMap = map[RecordType]string{
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

func (rt RecordType) Exists() bool {
	_, ok := recordTypeMap[rt]
	return ok
}

func (rt RecordType) String() string {
	if v, ok := recordTypeMap[rt]; ok {
		return v
	}
	return ""
}
