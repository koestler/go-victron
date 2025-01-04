package vebleapi

import (
	"github.com/koestler/go-list"
	"github.com/koestler/go-victron/log"
	"sync"
	"tinygo.org/x/bluetooth"
)

type Packet struct {
	rssi             int
	address          string
	name             string
	manufacturerData []byte
}

func (p Packet) RSSI() int {
	return p.rssi
}

func (p Packet) Address() string {
	return p.address
}

func (p Packet) Name() string {
	return p.name
}

func (p Packet) ManufacturerData() []byte {
	return p.manufacturerData
}

// Adapter is a wrapper around the bluetooth.Adapter type.
// It listens for BLE advertisements, filters out only advertisements from Victron devices by looking at the
// manufacturer data, and then calls a mac specific listener or a default listener.
type Adapter struct {
	manufacturerId uint16
	logger         log.Logger
	adapter        *bluetooth.Adapter

	scanResult chan bluetooth.ScanResult

	listener     map[string]*list.List[*Listener] // map of local names to listeners
	listenerLock sync.Mutex
}

// NewDefaultAdapter creates a new Adapter with the default bluetooth adapter.
// It can only be called once.
func NewDefaultAdapter(manufacturerId uint16, logger log.Logger) (*Adapter, error) {
	a := &Adapter{
		manufacturerId: manufacturerId,
		logger:         logger,
		adapter:        bluetooth.DefaultAdapter,
		scanResult:     make(chan bluetooth.ScanResult),
		listener:       make(map[string]*list.List[*Listener]),
	}
	err := a.adapter.Enable()
	if err != nil {
		return nil, err
	}

	go a.scan()
	go a.run()

	return a, nil
}

func (a *Adapter) RegisterNameListener(name string) *Listener {
	a.listenerLock.Lock()
	defer a.listenerLock.Unlock()

	l, ok := a.listener[name]
	if !ok {
		l = list.New[*Listener]()
		a.listener[name] = l
	}

	listener := NewListener()
	e := l.PushBack(listener)
	e.Value.unsubscribe = func() {
		a.listenerLock.Lock()
		defer a.listenerLock.Unlock()
		l.Remove(e)
	}

	return listener
}

func (a *Adapter) RegisterDefaultListener() *Listener {
	return a.RegisterNameListener("")
}

func (a *Adapter) Close() {
	err := a.adapter.StopScan()
	if err != nil {
		a.logger.Printf("error while stopping scan: %s", err)
	}

	// close all listener channels
	a.listenerLock.Lock()
	defer a.listenerLock.Unlock()
	for _, l := range a.listener {
		for e := l.Front(); e != nil; e = e.Next() {
			e.Value.close()
		}
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
		case sr, ok := <-a.scanResult:
			if !ok {
				// terminate
				a.logger.Printf("(*Adapter).run: scanResult channel closed")
				return
			}
			manufacturerData := extractManufacturerData(sr, a.manufacturerId)
			if manufacturerData == nil {
				// ignore all non-victron devices
				continue
			}

			a.logger.Printf("recieved scan result from %s, RSSI=%d, name=%s", sr.Address.String(), sr.RSSI, sr.LocalName())

			name := sr.LocalName()
			if name == "" {
				a.logger.Printf("(*Adapter).run: ignoring device with empty name")
				continue
			}

			a.notifyListeners(sr, manufacturerData)
		}
	}
}

func (a *Adapter) notifyListeners(sr bluetooth.ScanResult, manufacturerData []byte) {
	a.listenerLock.Lock()
	defer a.listenerLock.Unlock()

	ok := a.notifyListenersList(sr.LocalName(), sr, manufacturerData)
	if !ok {
		// fallback to default listener
		a.notifyListenersList("", sr, manufacturerData)
	}
}

// notifyListenersList notifies all listeners that are listening for the given key.
// It returns true, if at lest one listener was notified.
// This function assumes, that the caller holds the listenerLock.
func (a *Adapter) notifyListenersList(key string, sr bluetooth.ScanResult, manufacturerData []byte) bool {
	l, ok := a.listener[key]
	if !ok {
		return false
	}

	p := Packet{
		rssi:             int(sr.RSSI),
		address:          sr.Address.String(),
		name:             sr.LocalName(),
		manufacturerData: manufacturerData,
	}

	for e := l.Front(); e != nil; e = e.Next() {
		e.Value.packets <- p
	}

	return true

}

func extractManufacturerData(sr bluetooth.ScanResult, manufacturerId uint16) []byte {
	mdList := sr.AdvertisementPayload.ManufacturerData()
	for _, md := range mdList {
		if md.CompanyID == manufacturerId {
			return md.Data
		}
	}
	return nil
}
