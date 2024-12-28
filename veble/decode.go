package veble

import (
	"errors"
	"github.com/koestler/go-victron/veconst"
)

var ErrUnsupportedRecordType = errors.New("unsupported record type")

func Decode(recordType veconst.BleRecordType, decryptedBytes []byte) (Registers, error) {
	switch recordType {
	case veconst.BleSolarCharger:
		return DecodeSolarCharge(decryptedBytes)
	case veconst.BleBatteryMonitor:
		return DecodeBatteryMonitor(decryptedBytes)
	case veconst.BleInverter:
		return DecodeInverter(decryptedBytes)
	case veconst.BleDCDCConverter:
		return DecodeDcDcConverter(decryptedBytes)
	case veconst.BleSmartLithium:
		return DecodeSmartLithium(decryptedBytes)
	case veconst.BleInverterRS:
		return DecodeInverterRs(decryptedBytes)
	case veconst.BleGXDevice:
		return DecodeGxDevice(decryptedBytes)
	case veconst.BleACCharger:
		return DecodeAcCharger(decryptedBytes)
	case veconst.BleSmartBatteryProtect:
		return DecodeSmartBatteryProtect(decryptedBytes)
	case veconst.BleLynxSmartBMS:
		return DecodeLynxSmartBms(decryptedBytes)
	case veconst.BleMultiRS:
		return DecodeMultiRs(decryptedBytes)
	case veconst.BleVEBus:
		return DecodeVeBus(decryptedBytes)
	case veconst.BleDCEnergyMeter:
		return DecodeDcEnergyMeter(decryptedBytes)
	default:
		return nil, ErrUnsupportedRecordType
	}
}
