package tinygoble

import (
	"errors"
	"github.com/koestler/go-victron/log"
	"sync"

	"tinygo.org/x/bluetooth"
)

var ErrMacAlreadyListening = errors.New("mac address already listening")

type Adapter struct {
	logger  log.Logger
	adapter *bluetooth.Adapter

	listener     map[string]*Listener // map of mac address to handler function
	listenerLock sync.Mutex

	close chan struct{}
}

func NewDefaultAdapter(logger log.Logger) (*Adapter, error) {
	a := &Adapter{
		logger:   logger,
		adapter:  bluetooth.DefaultAdapter,
		listener: make(map[string]*Listener),
		close:    make(chan struct{}),
	}

	srChan := make(chan bluetooth.ScanResult)
	if err := a.adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		srChan <- result
	}); err != nil {
		a.logger.Printf("error while scanning %s", err)
		return nil, err
	}
	go a.run(srChan)

	return a, nil
}

func (a *Adapter) Close() {
	close(a.close)
}

func (a *Adapter) run(srChan chan bluetooth.ScanResult) {
	defer func() {
		err := a.adapter.StopScan()
		if err != nil {
			a.logger.Printf("error while stopping scan: %s", err)
		}
	}()

	for {
		select {
		case <-a.close:
			a.logger.Printf("adapter terminating")
			return
		case sr := <-srChan:
			a.logger.Printf("received result from address: %s", sr.Address)
			a.listenerLock.Lock()
			l, ok := a.listener[sr.Address.String()]
			a.listenerLock.Unlock()
			if ok {
				l.handle(sr)
			} else {
				a.logger.Printf("no listener for address: %s", sr.Address)
			}
		}
	}
}

func (a *Adapter) addListener(macAddress string, l *Listener) error {
	a.listenerLock.Lock()
	defer a.listenerLock.Unlock()
	if _, ok := a.listener[macAddress]; ok {
		return ErrMacAlreadyListening
	}

	a.listener[macAddress] = l
	return nil
}

func (a *Adapter) removeListener(macAddress string) {
	a.listenerLock.Lock()
	defer a.listenerLock.Unlock()
	delete(a.listener, macAddress)
}

// Listen listens for scan results for a given mac. The handler function is called with the scan result.
// The method blocks until the context is canceled.
func (a *Adapter) Listen(macAddress string) (close func(), md <-chan []byte, err error) {
	l := &Listener{
		a:                a,
		macAddress:       macAddress,
		manufacturerData: make(chan []byte),
	}

	if err := a.addListener(macAddress, l); err != nil {
		return nil, nil, err
	}

	a.logger.Printf("listening for %s", macAddress)

	return l.Close, l.manufacturerData, nil
}
