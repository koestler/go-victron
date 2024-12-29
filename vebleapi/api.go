package vebleapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/mac"
	"github.com/koestler/go-victron/veble"
	"github.com/koestler/go-victron/veblerecord"
)

type Adapter interface {
	RegisterMacListener(mac mac.MAC, l func(rssi int, localName string, data []byte))
	UnregisterMacListener(mac mac.MAC)
}

type Api struct {
	adapter Adapter
	mac     mac.MAC
	key     []byte
	logger  log.Logger
}

func NewApi(adapter Adapter, mac mac.MAC, key []byte, logger log.Logger) *Api {
	return &Api{
		adapter: adapter,
		mac:     mac,
		key:     key,
		logger:  logger,
	}
}

func (a *Api) StreamRegisters(ctx context.Context, h func(rssi int, localName string, registers veblerecord.Registers)) {
	var lastData []byte
	listener := func(rssi int, localName string, victronData []byte) {
		// ignore duplicate packets
		if bytes.Equal(lastData, victronData) {
			return
		}
		lastData = victronData

		ef, err := veble.DecodeFrame(victronData, a.logger)
		if err != nil {
			fmt.Println("error decoding frame:", err)
			return
		}

		fmt.Printf("received packet MAC=%s, RSSI=%d, name=%s, IV=%d\n", a.mac, rssi, localName, ef.IV)

		df, err := veble.DecryptFrame(ef, a.key, a.logger)
		if err != nil {
			fmt.Println("error decrypting frame:", err)
			return
		}

		registers, err := veblerecord.Decode(df.RecordType, df.DecryptedBytes)
		if err != nil {
			fmt.Println("error decoding registers:", err)
			return
		}

		h(rssi, localName, registers)
	}

	a.adapter.RegisterMacListener(a.mac, listener)
	defer a.adapter.UnregisterMacListener(a.mac)

	// run until context is canceled
	<-ctx.Done()
}
