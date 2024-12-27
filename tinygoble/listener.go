package tinygoble

import (
	"github.com/koestler/go-victron/veconst"
	"tinygo.org/x/bluetooth"
)

type Listener struct {
	a                *Adapter
	macAddress       string
	manufacturerData chan []byte
}

func (l *Listener) Close() {
	l.a.logger.Printf("closing listener for %s", l.macAddress)
	l.a.removeListener(l.macAddress)
	close(l.manufacturerData)
}

func (l *Listener) handle(sr bluetooth.ScanResult) {
	l.a.logger.Printf("advertisement received: %s, %d, %s", sr.Address.String(), sr.RSSI, sr.LocalName())
	d := extractData(sr)
	l.manufacturerData <- d
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
