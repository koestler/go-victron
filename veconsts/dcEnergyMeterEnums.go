package veconsts

type DcEnergyMeterAuxMode uint8

const (
	DcEnergyMeterAuxModeAuxVoltage  DcEnergyMeterAuxMode = 0
	DcEnergyMeterAuxModeTemperature DcEnergyMeterAuxMode = 2
	DcEnergyMeterAuxModeDisabled    DcEnergyMeterAuxMode = 3
)

func GetDcEnergyMeterAuxModeMap() map[DcEnergyMeterAuxMode]string {
	return map[DcEnergyMeterAuxMode]string{
		DcEnergyMeterAuxModeAuxVoltage:  "Aux voltage",
		DcEnergyMeterAuxModeTemperature: "Temperature",
		DcEnergyMeterAuxModeDisabled:    "Disabled",
	}
}

func (s DcEnergyMeterAuxMode) String() string {
	m := GetDcEnergyMeterAuxModeMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
