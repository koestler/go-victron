package bleparser

import (
	"errors"
	"github.com/koestler/go-victron/veconst"
)

var ErrUnsupportedRecordType = errors.New("unsupported record type")

func Decode(recordType veconst.BleRecordType, decryptedBytes []byte) (any, error) {
	switch recordType {
	case veconst.BleSolarCharger:
		return DecodeSolarChargeRecord(decryptedBytes)
	case veconst.BleBatteryMonitor:
		return DecodeBatteryMonitorRecord(decryptedBytes)
	case veconst.BleInverter:
		return DecodeInverterRecord(decryptedBytes)
	case veconst.BleDCDCConverter:
		return DecodeDcDcConverterRecord(decryptedBytes)
	case veconst.BleSmartLithium:
		return DecodeSmartLithiumRecord(decryptedBytes)
	case veconst.BleInverterRS:
		return DecodeInverterRsRecord(decryptedBytes)
	case veconst.BleGXDevice:
		return DecodeGxDeviceRecord(decryptedBytes)
	case veconst.BleACCharger:
		return DecodeAcChargerRecord(decryptedBytes)
	case veconst.BleSmartBatteryProtect:
		return DecodeSmartBatteryProtectRecord(decryptedBytes)
	case veconst.BleLynxSmartBMS:
		return DecodeLynxSmartBms(decryptedBytes)
	case veconst.BleMultiRS:
		return DecodeMultiRsRecord(decryptedBytes)
	case veconst.BleVEBus:
		return DecodeVeBusRecord(decryptedBytes)
	case veconst.BleDCEnergyMeter:
		return DecodeDcEnergyMeterRecord(decryptedBytes)
	default:
		return nil, ErrUnsupportedRecordType
	}
}
