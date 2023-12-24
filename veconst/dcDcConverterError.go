package veconst

type DcDcConverterError uint8

const (
	DcDcConverterErrorNoError                           DcDcConverterError = 0
	DcDcConverterErrorBatteryVoltageTooHigh             DcDcConverterError = 2
	DcDcConverterErrorChargerInternalTemperatureTooHigh DcDcConverterError = 17
	DcDcConverterErrorChargerExcessiveOutputCurrent     DcDcConverterError = 18
	DcDcConverterErrorChargerCurrentPolarityReversed    DcDcConverterError = 19
	DcDcConverterErrorChargerBulkTimeExpired            DcDcConverterError = 20
	DcDcConverterErrorChargerCurrentSensorIssue         DcDcConverterError = 21
	DcDcConverterErrorChargerTerminalsOverheated        DcDcConverterError = 26
	DcDcConverterErrorConverterIssue                    DcDcConverterError = 28
	DcDcConverterErrorInputVoltageTooHigh               DcDcConverterError = 33
	DcDcConverterErrorInputExcessiveCurrent             DcDcConverterError = 34
	DcDcConverterErrorInputShutdownBatteryVoltage       DcDcConverterError = 38
	DcDcConverterErrorInputShutdownCurrentFlowing       DcDcConverterError = 39
	DcDcConverterErrorIncompatibleDeviceInTheNetwork    DcDcConverterError = 66
	DcDcConverterErrorBmsConnectionLost                 DcDcConverterError = 67
	DcDcConverterErrorNetworkMisconfigured              DcDcConverterError = 68
	DcDcConverterErrorCalibrationDataLost               DcDcConverterError = 116
	DcDcConverterErrorIncompatibleFirmware              DcDcConverterError = 117
	DcDcConverterErrorSettingsDataInvalid               DcDcConverterError = 119
	DcDcConverterErrorUnknown                           DcDcConverterError = 255
)

var dcDcConverterErrorMap = map[DcDcConverterError]string{
	DcDcConverterErrorNoError:                           "No error",
	DcDcConverterErrorBatteryVoltageTooHigh:             "Battery voltage too high",
	DcDcConverterErrorChargerInternalTemperatureTooHigh: "Charger internal temperature too high",
	DcDcConverterErrorChargerExcessiveOutputCurrent:     "Charger excessive output current",
	DcDcConverterErrorChargerCurrentPolarityReversed:    "Charger current polarity reversed",
	DcDcConverterErrorChargerBulkTimeExpired:            "Charger bulk time expired (when 10 hour bulk time protection active)",
	DcDcConverterErrorChargerCurrentSensorIssue:         "Charger current sensor issue (bias not within expected limits during off state)",
	DcDcConverterErrorChargerTerminalsOverheated:        "Charger terminals overheated",
	DcDcConverterErrorConverterIssue:                    "Converter issue (dual converter models, one of the converters is not working)",
	DcDcConverterErrorInputVoltageTooHigh:               "Input voltage too high",
	DcDcConverterErrorInputExcessiveCurrent:             "Input excessive current",
	DcDcConverterErrorInputShutdownBatteryVoltage:       "Input shutdown (due to excessive battery voltage)",
	DcDcConverterErrorInputShutdownCurrentFlowing:       "Input shutdown (current flowing while the converter is switched off)",
	DcDcConverterErrorIncompatibleDeviceInTheNetwork:    "Incompatible device in the network (for synchronized charging)",
	DcDcConverterErrorBmsConnectionLost:                 "BMS connection lost",
	DcDcConverterErrorNetworkMisconfigured:              "Network misconfigured (e.g. combining ESS with ve.smart networking)",
	DcDcConverterErrorCalibrationDataLost:               "Calibration data lost",
	DcDcConverterErrorIncompatibleFirmware:              "Incompatible firmware (i.e. not for this model)",
	DcDcConverterErrorSettingsDataInvalid:               "Settings data invalid / corrupted (use restore to defaults and set to recover)",
	DcDcConverterErrorUnknown:                           "Unknown error",
}

// DcDcConverterErrorStringMap returns a map of DcDcConverterError values to their string representation.
func DcDcConverterErrorStringMap() map[DcDcConverterError]string {
	ret := make(map[DcDcConverterError]string, len(dcDcConverterErrorMap))
	for k, v := range dcDcConverterErrorMap {
		ret[k] = v
	}
	return ret
}

// Exists returns true if the DcDcConverterError exists.
func (s DcDcConverterError) Exists() bool {
	_, ok := dcDcConverterErrorMap[s]
	return ok
}

// String returns the string representation of a DcDcConverterError.
func (s DcDcConverterError) String() string {
	if v, ok := dcDcConverterErrorMap[s]; ok {
		return v
	}
	return ""
}
