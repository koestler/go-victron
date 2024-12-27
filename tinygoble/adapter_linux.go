// Package tinygoble provides a simple API to listen for BLE advertisements from Victron devices.
// It currently only works on Linux since macOS does not expose the MAC address of bluetooth devices.
package tinygoble

import (
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/veble"
	"github.com/koestler/go-victron/veconst"
	"tinygo.org/x/bluetooth"
)

type MacListener func(rssi int, localName string, victronData []byte)

type setMacListenerRequest struct {
	mac      veble.MAC
	listener MacListener
}
type DefaultListener func(mac veble.MAC, rssi int, localName string)

// Adapter is a wrapper around the bluetooth.Adapter type.
// It listens for BLE advertisements, filters out only advertisements from Victron devices by looking at the
// manufacturer data, and then calls a mac specific listener or a default listener.
type Adapter struct {
	logger  log.Logger
	adapter *bluetooth.Adapter

	scanResult chan bluetooth.ScanResult

	macListener        map[veble.MAC]MacListener // map of mac address to handler function
	setMacListener     chan setMacListenerRequest
	defaultListener    DefaultListener
	setDefaultListener chan DefaultListener
}

// NewDefaultAdapter creates a new Adapter with the default bluetooth adapter.
// It can only be called once.
func NewDefaultAdapter(logger log.Logger) (*Adapter, error) {
	a := &Adapter{
		logger:             logger,
		adapter:            bluetooth.DefaultAdapter,
		scanResult:         make(chan bluetooth.ScanResult),
		macListener:        make(map[veble.MAC]MacListener),
		setMacListener:     make(chan setMacListenerRequest),
		setDefaultListener: make(chan DefaultListener),
	}
	err := a.adapter.Enable()
	if err != nil {
		return nil, err
	}

	go a.scan()
	go a.run()

	return a, nil
}

// RegisterDefaultListener registers a listener that is called when no MAC specific listener is found.
// Only one default listener can be registered. If a default listener is already registered, it is replaced.
func (a *Adapter) RegisterDefaultListener(l DefaultListener) {
	a.setDefaultListener <- l
}

// RegisterMacListener registers a listener that is called only for announcements sent to a specific MAC address.
// Only one listener can be registered per MAC address. If it is already registered, it is replaced.
func (a *Adapter) RegisterMacListener(mac veble.MAC, l MacListener) {
	a.setMacListener <- setMacListenerRequest{mac, l}
}

// UnregisterMacListener removes the listener for the given MAC address.
func (a *Adapter) UnregisterMacListener(mac veble.MAC, l MacListener) {
	a.setMacListener <- setMacListenerRequest{mac, l}
}

func (a *Adapter) Close() {
	err := a.adapter.StopScan()
	if err != nil {
		a.logger.Printf("error while stopping scan: %s", err)
	}
}

func (a *Adapter) scan() {
	// when this function returns, the run routine is closed by this channel
	defer close(a.scanResult)

	a.logger.Printf("(*Adapter).scan: starting")
	if err := a.adapter.Scan(func(adapter *bluetooth.Adapter, sr bluetooth.ScanResult) {
		a.scanResult <- sr
	}); err != nil {
		a.logger.Printf("error while scanning %s", err)
	}

	a.logger.Printf("(*Adapter).scan: stopped")
}

func (a *Adapter) run() {
	for {
		select {
		case ml := <-a.setMacListener:
			a.macListener[ml.mac] = ml.listener
		case dl := <-a.setDefaultListener:
			a.defaultListener = dl
		case sr, ok := <-a.scanResult:
			if !ok {
				// terminate
				a.logger.Printf("(*Adapter).run: scanResult channel closed")
				return
			}
			victronData := extractVictronData(sr)
			if victronData == nil {
				// ignore all non-victron devices
				continue
			}

			mac := veble.MAC(sr.Address.MAC) // does not work on darwin
			if ml, ok := a.macListener[mac]; ok {
				ml(int(sr.RSSI), sr.LocalName(), victronData)
			} else if a.defaultListener != nil {
				a.defaultListener(mac, int(sr.RSSI), sr.LocalName())
			}
		}
	}
}

func extractVictronData(sr bluetooth.ScanResult) []byte {
	mdList := sr.AdvertisementPayload.ManufacturerData()
	for _, md := range mdList {
		if md.CompanyID == veconst.BleManufacturerId {
			return md.Data
		}
	}
	return nil
}
