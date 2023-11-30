package ble

import (
	"context"
	"fmt"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"log"
	"strings"
)

type Config interface {
	Name() string
	LogDebug() bool
	Devices() []DeviceConfig
}

type DeviceConfig interface {
	MacAddress() string
	EncryptionKey() string
}

type BleStruct struct {
	cfg    Config
	ctx    context.Context
	cancel context.CancelFunc
}

func New(cfg Config) (*BleStruct, error) {
	ctx, cancel := context.WithCancel(context.Background())

	ble := &BleStruct{
		cfg:    cfg,
		ctx:    ctx,
		cancel: cancel,
	}

	adapterID := "hci0"

	go func() {
		//clean up connection on exit
		defer api.Exit()

		if cfg.LogDebug() {
			log.Printf("ble[%s]: get adapterId=%s", ble.Name(), adapterID)
		}
		a, err := adapter.GetAdapter(adapterID)
		if err != nil {
			log.Printf("ble[%s]: error while getting adapter: %s", ble.Name(), err)
			return
		}

		if cfg.LogDebug() {
			log.Printf("ble[%s]: flush devices", ble.Name())
		}
		err = a.FlushDevices()
		if err != nil {
			log.Printf("ble[%s]: error during flush: %s", ble.Name(), err)
			return
		}

		if cfg.LogDebug() {
			log.Printf("ble[%s]: start discovery", ble.Name())
		}

		discovery, cancel, err := api.Discover(a, nil)
		if err != nil {
			log.Printf("ble[%s]: error during discovery: %s", ble.Name(), err)
		}
		defer cancel()

		go func() {
			for ev := range discovery {
				if ev.Type == adapter.DeviceRemoved {
					continue
				}

				dev, err := device.NewDevice1(ev.Path)
				if err != nil {
					log.Printf("ble[%s]: %s: %s", ble.Name(), ev.Path, err)
					continue
				}

				if !ble.getDevice(dev.Properties.Address) {
					continue
				}

				log.Printf("ble[%s] device recovered: %s", ble.Name(), ev.Path)

				if dev == nil {
					log.Printf("ble[%s]: device %s not found", ble.Name(), ev.Path)
					continue
				}

				log.Printf("ble[%s]: name=%s addr=%s rssi=%d", ble.Name(), dev.Properties.Name, dev.Properties.Address, dev.Properties.RSSI)

				log.Printf("ble[%s]: addr=%s manufacturerData=%x", ble.Name(), dev.Properties.Address, dev.Properties.ManufacturerData)

				go func(ev *adapter.DeviceDiscovered) {
					err = connectDevice(dev)
					if err != nil {
						log.Printf("ble[%s]: beacon %s failed: %s", ble.Name(), ev.Path, err)
					}
				}(ev)
			}
		}()

		// wait for cancel
		<-ble.ctx.Done()
	}()

	return ble, nil
}

func connectDevice(dev *device.Device1) error {
	log.Printf("")

	propUpdates, err := dev.WatchProperties()
	if err != nil {
		return fmt.Errorf("cannot watch props: %s", err)
	}

	for pu := range propUpdates {
		log.Printf("dev=%s, received prop update: %#v", dev.Properties.Name, pu)
	}

	if err := dev.UnwatchProperties(propUpdates); err != nil {
		log.Printf("error during unwatch: %s", err)
	}

	return nil
}

func (ble *BleStruct) getDevice(bluezAddr string) bool {
	for _, d := range ble.cfg.Devices() {
		if d.MacAddress() == bluezAddrToOurAddr(bluezAddr) {
			return true
		}
	}
	return false
}

// input: D4:9D:D2:92:62:02
// output d49dd2926202
func bluezAddrToOurAddr(i string) string {
	return strings.ToLower(strings.ReplaceAll(i, ":", ""))
}

func (ble *BleStruct) Name() string {
	return ble.cfg.Name()
}

func (ble *BleStruct) Shutdown() {
	ble.cancel()
}
