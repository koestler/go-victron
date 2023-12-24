package veconsts

type SolarChargerState uint16

const (
	SolarChargerStateNotCharging        SolarChargerState = 0
	SolarChargerStateFault              SolarChargerState = 2
	SolarChargerStateBulkCharging       SolarChargerState = 3
	SolarChargerStateAbsorptionCharging SolarChargerState = 4
	SolarChargerStateFloatCharging      SolarChargerState = 5
	SolarChargerStateManualEqualise     SolarChargerState = 7
	SolarChargerStateWakeUp             SolarChargerState = 245
	SolarChargerStateAutoEqualise       SolarChargerState = 247
	SolarChargerStateExternalControl    SolarChargerState = 252
	SolarChargerStateUnavailable        SolarChargerState = 255
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

type SolarChargerError uint16

const (
	SolarChargerErrorNoError                           SolarChargerError = 0
	SolarChargerErrorBatteryVoltageTooHigh             SolarChargerError = 2
	SolarChargerErrorChargerInternalTemperatureTooHigh SolarChargerError = 17
	SolarChargerErrorChargerExcessiveOutputCurrent     SolarChargerError = 18
	SolarChargerErrorChargerCurrentPolarityReversed    SolarChargerError = 19
	SolarChargerErrorChargerBulkTimeExpired            SolarChargerError = 20
	SolarChargerErrorChargerCurrentSensorIssue         SolarChargerError = 21
	SolarChargerErrorChargerTerminalsOverheated        SolarChargerError = 26
	SolarChargerErrorConverterIssue                    SolarChargerError = 28
	SolarChargerErrorInputVoltageTooHigh               SolarChargerError = 33
	SolarChargerErrorInputExcessiveCurrent             SolarChargerError = 34
	SolarChargerErrorInputShutdownBatteryVoltage       SolarChargerError = 38
	SolarChargerErrorInputShutdownCurrentFlowing       SolarChargerError = 39
	SolarChargerErrorIncompatibleDeviceInTheNetwork    SolarChargerError = 66
	SolarChargerErrorBmsConnectionLost                 SolarChargerError = 67
	SolarChargerErrorNetworkMisconfigured              SolarChargerError = 68
	SolarChargerErrorCalibrationDataLost               SolarChargerError = 116
	SolarChargerErrorIncompatibleFirmware              SolarChargerError = 117
	SolarChargerErrorSettingsDataInvalid               SolarChargerError = 119
	SolarChargerErrorUnknown                           SolarChargerError = 255
)

func GetSolarChargerErrorMap() map[SolarChargerError]string {
	return map[SolarChargerError]string{
		SolarChargerErrorNoError:                           "No error",
		SolarChargerErrorBatteryVoltageTooHigh:             "Battery voltage too high",
		SolarChargerErrorChargerInternalTemperatureTooHigh: "Charger internal temperature too high",
		SolarChargerErrorChargerExcessiveOutputCurrent:     "Charger excessive output current",
		SolarChargerErrorChargerCurrentPolarityReversed:    "Charger current polarity reversed",
		SolarChargerErrorChargerBulkTimeExpired:            "Charger bulk time expired (when 10 hour bulk time protection active)",
		SolarChargerErrorChargerCurrentSensorIssue:         "Charger current sensor issue (bias not within expected limits during off state)",
		SolarChargerErrorChargerTerminalsOverheated:        "Charger terminals overheated",
		SolarChargerErrorConverterIssue:                    "Converter issue (dual converter models, one of the converters is not working)",
		SolarChargerErrorInputVoltageTooHigh:               "Input voltage too high",
		SolarChargerErrorInputExcessiveCurrent:             "Input excessive current",
		SolarChargerErrorInputShutdownBatteryVoltage:       "Input shutdown (due to excessive battery voltage)",
		SolarChargerErrorInputShutdownCurrentFlowing:       "Input shutdown (current flowing while the converter is switched off)",
		SolarChargerErrorIncompatibleDeviceInTheNetwork:    "Incompatible device in the network (for synchronized charging)",
		SolarChargerErrorBmsConnectionLost:                 "BMS connection lost",
		SolarChargerErrorNetworkMisconfigured:              "Network misconfigured (e.g. combining ESS with ve.smart networking)",
		SolarChargerErrorCalibrationDataLost:               "Calibration data lost",
		SolarChargerErrorIncompatibleFirmware:              "Incompatible firmware (i.e. not for this model)",
		SolarChargerErrorSettingsDataInvalid:               "Settings data invalid / corrupted (use restore to defaults and set to recover)",
		SolarChargerErrorUnknown:                           "Unknown error",
	}
}

func (s SolarChargerError) String() string {
	m := GetSolarChargerErrorMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
