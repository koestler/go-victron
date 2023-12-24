package veconsts

type DcEnergyMeterAuxMode uint8

const (
	DcEnergyMeterAuxModeAuxVoltage  DcEnergyMeterAuxMode = 0
	DcEnergyMeterAuxModeTemperature DcEnergyMeterAuxMode = 2
	DcEnergyMeterAuxModeDisabled    DcEnergyMeterAuxMode = 3
)

var dcEnergyMeterAuxModeMap = map[DcEnergyMeterAuxMode]string{
	DcEnergyMeterAuxModeAuxVoltage:  "Aux voltage",
	DcEnergyMeterAuxModeTemperature: "Temperature",
	DcEnergyMeterAuxModeDisabled:    "Disabled",
}

// GetDcEnergyMeterAuxModeStringMap returns a map of DcEnergyMeterAuxMode values to their string representation.
func GetDcEnergyMeterAuxModeStringMap() map[DcEnergyMeterAuxMode]string {
	ret := make(map[DcEnergyMeterAuxMode]string, len(dcEnergyMeterAuxModeMap))
	for k, v := range dcEnergyMeterAuxModeMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the DcEnergyMeterAuxMode exists.
func (s DcEnergyMeterAuxMode) Exists() bool {
	_, ok := dcEnergyMeterAuxModeMap[s]
	return ok
}

// String returns the string representation of a DcEnergyMeterAuxMode.
func (s DcEnergyMeterAuxMode) String() string {
	if v, ok := dcEnergyMeterAuxModeMap[s]; ok {
		return v
	}
	return ""
}
