package tinygoble

import (
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/veconst"

	"tinygo.org/x/bluetooth"
)

type MacListener func(rssi int, localName string, victronData []byte)
type DefaultListener func(mac string, rssi int, localName string)

// Adapter is a wrapper around the bluetooth.Adapter type.
// It listens for BLE advertisements, filters out only advertisements from Victron devices by looking at the
// manufacturer data, and then calls a mac specific listener or a default listener.
type Adapter struct {
	logger  log.Logger
	adapter *bluetooth.Adapter

	macListener     map[string]MacListener // map of mac address to handler function
	defaultListener DefaultListener
}

// NewDefaultAdapter creates a new Adapter with the default bluetooth adapter.
// It can only be called once.
func NewDefaultAdapter(logger log.Logger) (*Adapter, error) {
	a := &Adapter{
		logger:      logger,
		adapter:     bluetooth.DefaultAdapter,
		macListener: make(map[string]MacListener),
	}
	err := a.adapter.Enable()
	if err != nil {
		return nil, err
	}
	return a, nil
}

// RegisterDefaultListener registers a listener that is called when no MAC specific listener is found.
// Only one default listener can be registered. If a default listener is already registered, it is replaced.
// Execute this function before calling Run.
func (a *Adapter) RegisterDefaultListener(l DefaultListener) {
	a.defaultListener = l
}

// RegisterMacListener registers a listener that is called only for announcements sent to a specific MAC address.
// Only one listener can be registered per MAC address. If it is already registered, it is replaced.
// Execute this function before calling Run.
func (a *Adapter) RegisterMacListener(mac string, l MacListener) {
	a.macListener[mac] = l
}

func (a *Adapter) Close() {
	err := a.adapter.StopScan()
	if err != nil {
		a.logger.Printf("error while stopping scan: %s", err)
	}
}

func (a *Adapter) Run() {
	go func() {
		a.logger.Printf("(*Adapter).Run: starting scan")
		if err := a.adapter.Scan(func(adapter *bluetooth.Adapter, sr bluetooth.ScanResult) {

			macStr := sr.Address.String()
			if ml, ok := a.macListener[macStr]; ok {
				ml(int(sr.RSSI), sr.LocalName(), extractVictronMD(sr))
			} else if a.defaultListener != nil {
				a.defaultListener(macStr, int(sr.RSSI), sr.LocalName())
			}
		}); err != nil {
			a.logger.Printf("error while scanning %s", err)
			a.Close()
		}

		a.logger.Printf("scan stopped")
	}()
}

func extractVictronMD(sr bluetooth.ScanResult) []byte {
	mdList := sr.AdvertisementPayload.ManufacturerData()
	for _, md := range mdList {
		if md.CompanyID == veconst.BleManufacturerId {
			return md.Data
		}
	}
	return nil
}
