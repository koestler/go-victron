package vebleapi

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"tinygo.org/x/bluetooth"
)

var ErrMacAlreadyListening = errors.New("mac address already listening")

type Config interface {
	LogDebug() bool
}

type HandlerFunc func(bluetooth.ScanResult)

type BleApi struct {
	cfg     Config
	adapter *bluetooth.Adapter

	listener     map[string]HandlerFunc // map of mac address to handler function
	listenerLock sync.Mutex

	close chan struct{}
}

func NewBleApi(cfg Config) (*BleApi, error) {
	ba := &BleApi{
		cfg:      cfg,
		adapter:  bluetooth.DefaultAdapter,
		listener: make(map[string]HandlerFunc),
		close:    make(chan struct{}),
	}

	srChan := make(chan bluetooth.ScanResult)
	if err := ba.adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		srChan <- result
	}); err != nil {
		if ba.cfg.LogDebug() {
			fmt.Printf("bleapi: error while scanning %s", err)
		}
		return nil, err
	}
	go ba.run(srChan)

	return ba, nil
}

func (ba *BleApi) Close() {
	close(ba.close)
}

func (ba *BleApi) run(srChan chan bluetooth.ScanResult) {
	defer func() {
		err := ba.adapter.StopScan()
		if err != nil {
			fmt.Printf("bleapi: error while stopping scan: %s", err)
		}
	}()

	for {
		select {
		case <-ba.close:
			if ba.cfg.LogDebug() {
				fmt.Printf("bleapi: terminating")
			}
			return
		case sr := <-srChan:
			if ba.cfg.LogDebug() {
				fmt.Printf("bleapi: received result from address: %s", sr.Address)
			}
			ba.listenerLock.Lock()
			l, ok := ba.listener[sr.Address.String()]
			ba.listenerLock.Unlock()
			if ok {
				l(sr)
			} else if ba.cfg.LogDebug() {
				fmt.Printf("bleapi: no listener for address: %s", sr.Address)
			}
		}
	}
}

func (ba *BleApi) addListener(mac string, handler HandlerFunc) error {
	ba.listenerLock.Lock()
	defer ba.listenerLock.Unlock()
	if _, ok := ba.listener[mac]; ok {
		return ErrMacAlreadyListening
	}

	ba.listener[mac] = handler
	return nil
}

func (ba *BleApi) removeListener(mac string) {
	ba.listenerLock.Lock()
	defer ba.listenerLock.Unlock()
	delete(ba.listener, mac)
}

// Listen listens for scan results for a given mac. The handler function is called with the scan result.
// The method blocks until the context is canceled.
func (ba *BleApi) Listen(ctx context.Context, mac string, handler HandlerFunc) error {
	if err := ba.addListener(mac, handler); err != nil {
		return err
	}
	defer ba.removeListener(mac)

	<-ctx.Done()
	return nil
}
