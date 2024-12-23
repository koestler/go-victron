package bleapi

import (
	"context"
	"fmt"
)

type Device struct {
	Name          string
	MacAddress    string
	EncryptionKey string
}

func (d Device) String() string {
	return fmt.Sprintf("Device %s %s", d.Name, d.MacAddress)
}

func NewDevice(name, macAddress, encryptionKey string) Device {
	return Device{
		Name:          name,
		MacAddress:    macAddress,
		EncryptionKey: encryptionKey,
	}
}

func (d Device) StreamRegisters(ctx context.Context, ba *BleApi) string {
	return d.Name
}
