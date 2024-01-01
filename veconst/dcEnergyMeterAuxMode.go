package veconst

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

func NewDcEnergyMeterAuxModeEnum(v int) (Enum, error) {
	s := DcEnergyMeterAuxMode(v)
	if _, ok := dcEnergyMeterAuxModeMap[s]; !ok {
		return nil, ErrInvalidEnumIdx
	}
	return s, nil
}

func DcEnergyMeterAuxModeMap() map[int]string {
	ret := make(map[int]string, len(dcEnergyMeterAuxModeMap))
	for k, v := range dcEnergyMeterAuxModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s DcEnergyMeterAuxMode) Idx() int {
	return int(s)
}

func (s DcEnergyMeterAuxMode) Value() string {
	if v, ok := dcEnergyMeterAuxModeMap[s]; ok {
		return v
	}
	return ""
}
