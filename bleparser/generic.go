package bleparser

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

type Registers interface {
	NumberRegisters() []NumberRegister
	EnumRegisters() []EnumRegister
	FieldListRegisters() []FieldListRegister
}

type Register struct {
	name        string
	description string
}

func (r Register) Name() string {
	return r.name
}

func (r Register) Description() string {
	return r.description
}

type NumberRegister struct {
	Register
	value float64
	unit  string
}

func (r NumberRegister) GenericValue() interface{} {
	return r.value
}

func (r NumberRegister) Value() float64 {
	return r.value
}

func (r NumberRegister) Unit() string {
	return r.unit
}

type EnumRegister struct {
	Register
	value veconst.Enum
}

func (r EnumRegister) GenericValue() interface{} {
	return r.value
}

func (r EnumRegister) Value() veconst.Enum {
	return r.value
}

func (r EnumRegister) Unit() string {
	return ""
}

type FieldListRegister struct {
	Register
	value veconst.FieldList
}

func (r FieldListRegister) GenericValue() interface{} {
	return r.value
}

func (r FieldListRegister) Value() veconst.FieldList {
	return r.value
}

func (r FieldListRegister) Unit() string {
	return ""
}
