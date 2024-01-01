package veconst

type DcEnergyMeterAuxMode uint8
type DcEnergyMeterAuxModeFactoryType struct{}

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
var DcEnergyMeterAuxModeFactory DcEnergyMeterAuxModeFactoryType

func (f DcEnergyMeterAuxModeFactoryType) New(v uint8) (DcEnergyMeterAuxMode, error) {
	s := DcEnergyMeterAuxMode(v)
	if _, ok := dcEnergyMeterAuxModeMap[s]; !ok {
		return DcEnergyMeterAuxModeDisabled, ErrInvalidEnumIdx
	}
	return s, nil
}

func (f DcEnergyMeterAuxModeFactoryType) NewEnum(v int) (Enum, error) {
	return f.New(uint8(v))
}

func (f DcEnergyMeterAuxModeFactoryType) IntToStringMap() map[int]string {
	ret := make(map[int]string, len(dcEnergyMeterAuxModeMap))
	for k, v := range dcEnergyMeterAuxModeMap {
		ret[int(k)] = v
	}
	return ret
}

func (s DcEnergyMeterAuxMode) Idx() int {
	return int(s)
}

func (s DcEnergyMeterAuxMode) String() string {
	if v, ok := dcEnergyMeterAuxModeMap[s]; ok {
		return v
	}
	return ""
}
