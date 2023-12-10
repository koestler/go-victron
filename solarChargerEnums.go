package victronDefinitions

type SolarChargerState uint16

const (
	SolarChargerStateNotCharging        = 0
	SolarChargerStateFault              = 2
	SolarChargerStateBulkCharging       = 3
	SolarChargerStateAbsorptionCharging = 4
	SolarChargerStateFloatCharging      = 5
	SolarChargerStateManualEqualise     = 7
	SolarChargerStateWakeUp             = 245
	SolarChargerStateAutoEqualise       = 247
	SolarChargerStateExternalControl    = 252
	SolarChargerStateUnavailable        = 255
)

func GetSolarChargerStateMap() map[SolarChargerState]string {
	return map[SolarChargerState]string{
		SolarChargerStateNotCharging:        "Not charging",
		SolarChargerStateFault:              "Fault",
		SolarChargerStateBulkCharging:       "Bulk Charging",
		SolarChargerStateAbsorptionCharging: "Absorption Charging",
		SolarChargerStateFloatCharging:      "Float Charging",
		SolarChargerStateManualEqualise:     "Manual Equalise",
		SolarChargerStateWakeUp:             "Wake-Up",
		SolarChargerStateAutoEqualise:       "Auto Equalise",
		SolarChargerStateExternalControl:    "External Control",
		SolarChargerStateUnavailable:        "Unavailable",
	}
}

func (s SolarChargerState) String() string {
	m := GetSolarChargerStateMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
