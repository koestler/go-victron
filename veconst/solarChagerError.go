package veconst

type SolarChargerError uint8

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

var solarChargerErrorMap = map[SolarChargerError]string{
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

// SolarChargerErrorStringMap returns a map of SolarChargerError values to their string representation.
func SolarChargerErrorStringMap() map[SolarChargerError]string {
	ret := make(map[SolarChargerError]string, len(solarChargerErrorMap))
	for k, v := range solarChargerErrorMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the SolarChargerError exists.
func (s SolarChargerError) Exists() bool {
	_, ok := solarChargerErrorMap[s]
	return ok
}

// String returns the string representation of a SolarChargerError.
func (s SolarChargerError) String() string {
	if v, ok := solarChargerErrorMap[s]; ok {
		return v
	}
	return ""
}
