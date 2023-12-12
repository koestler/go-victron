package victronDefinitions

type DcDcConverterState uint16

const (
	DcDcConverterStateNotCharging        DcDcConverterState = 0
	DcDcConverterStateFault              DcDcConverterState = 2
	DcDcConverterStateBulkCharging       DcDcConverterState = 3
	DcDcConverterStateAbsorptionCharging DcDcConverterState = 4
	DcDcConverterStateFloatCharging      DcDcConverterState = 5
	DcDcConverterStateManualEqualise     DcDcConverterState = 7
	DcDcConverterStateWakeUp             DcDcConverterState = 245
	DcDcConverterStateAutoEqualise       DcDcConverterState = 247
	DcDcConverterStateExternalControl    DcDcConverterState = 252
	DcDcConverterStateUnavailable        DcDcConverterState = 255
)

func GetDcDcConverterStateMap() map[DcDcConverterState]string {
	return map[DcDcConverterState]string{
		DcDcConverterStateNotCharging:        "Not charging",
		DcDcConverterStateFault:              "Fault",
		DcDcConverterStateBulkCharging:       "Bulk Charging",
		DcDcConverterStateAbsorptionCharging: "Absorption Charging",
		DcDcConverterStateFloatCharging:      "Float Charging",
		DcDcConverterStateManualEqualise:     "Manual Equalise",
		DcDcConverterStateWakeUp:             "Wake-Up",
		DcDcConverterStateAutoEqualise:       "Auto Equalise",
		DcDcConverterStateExternalControl:    "External Control",
		DcDcConverterStateUnavailable:        "Unavailable",
	}
}

func (s DcDcConverterState) String() string {
	m := GetDcDcConverterStateMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}

type DcDcConverterError uint16

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

func GetDcDcConverterErrorMap() map[DcDcConverterError]string {
	return map[DcDcConverterError]string{
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
}

func (s DcDcConverterError) String() string {
	m := GetDcDcConverterErrorMap()
	if v, ok := m[s]; ok {
		return v
	}
	return ""
}
