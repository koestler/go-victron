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

type Handler func(rssi int, localName string, record any)

type RecordApi struct {
	adapter Adapter
	mac     mac.MAC
	key     []byte
	logger  log.Logger
}

func NewRecordApi(adapter Adapter, mac mac.MAC, key []byte, logger log.Logger) *RecordApi {
	return &RecordApi{
		adapter: adapter,
		mac:     mac,
		key:     key,
		logger:  logger,
	}
}

func (a *RecordApi) Stream(ctx context.Context, h Handler) {
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

		record, err := veblerecord.Decode(df.RecordType, df.DecryptedBytes)
		if err != nil {
			fmt.Println("error decoding registers:", err)
			return
		}

		h(rssi, localName, record)
	}

	a.adapter.RegisterMacListener(a.mac, listener)
	defer a.adapter.UnregisterMacListener(a.mac)

	// run until context is canceled
	<-ctx.Done()
}
