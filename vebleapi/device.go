package vebleapi

import (
	"context"
	"fmt"
	"github.com/koestler/go-victron/veconst"
	"tinygo.org/x/bluetooth"
)

// Logger is the interface for a logger. It is implemented by e.g. log.Logger.
type Logger interface {
	Println(v ...any)
}

type Device struct {
	DebugLogger Logger // optional: a logger for debug output; if nil, no debug output is written

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

func (d Device) Consume(ctx context.Context, ba *BleApi, handler func(ctx context.Context, result string)) error {
	return ba.Listen(ctx, d.MacAddress, d.handleResult)
}

func (d Device) handleResult(result bluetooth.ScanResult) {
	if d.DebugLogger != nil {
		d.DebugLogger.Println(fmt.Sprintf("received advertisement: %s", result))
	}

	data := extractData(result.AdvertisementPayload.ManufacturerData())
	if data == nil {
		if d.DebugLogger != nil {
			d.DebugLogger.Println(fmt.Sprintf("no manufacturer data found for CompanyID %d", veconst.BleManufacturerId))
		}
		return
	}
}

func extractData(result bluetooth.ScanResult) []byte {
	mdList := result.AdvertisementPayload.ManufacturerData()
	for _, md := range mdList {
		if md.CompanyID == veconst.BleManufacturerId {
			return md.Data
		}
	}
	return nil
}
