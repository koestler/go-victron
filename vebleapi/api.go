package vebleapi

import (
	"bytes"
	"context"
	"fmt"
	"github.com/koestler/go-victron/velog"
	"github.com/koestler/go-victron/veble"
	"github.com/koestler/go-victron/veblerecord"
)

type Api struct {
	adapter       *Adapter
	name          string
	encryptionKey []byte
	logger        velog.Logger
}

func NewApi(adapter *Adapter, name string, encryptionKey []byte, logger velog.Logger) *Api {
	return &Api{
		adapter:       adapter,
		name:          name,
		encryptionKey: encryptionKey,
		logger:        logger,
	}
}

func (a *Api) StreamRegisters(ctx context.Context, h func(rssi int, registers veblerecord.Registers)) {
	var lastMD []byte

	listener := a.adapter.RegisterNameListener(a.name)
	defer listener.End()

	// terminate when ctx is done
	go func() {
		<-ctx.Done()
		listener.End()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case p := <-listener.Drain():
			// ignore duplicate packets
			md := p.ManufacturerData()

			if bytes.Equal(lastMD, md) {
				continue
			}
			lastMD = md

			ef, err := veble.DecodeFrame(md, a.logger)
			if err != nil {
				fmt.Println("error decoding frame:", err)
				return
			}

			fmt.Printf("received packet from address=%s, RSSI=%d, name=%s, IV=%d\n", p.Name(), p.RSSI(), p.Name(), ef.IV)

			df, err := veble.DecryptFrame(ef, a.encryptionKey, a.logger)
			if err != nil {
				fmt.Println("error decrypting frame:", err)
				return
			}

			registers, err := veblerecord.Decode(df.RecordType, df.DecryptedBytes)
			if err != nil {
				fmt.Println("error decoding registers:", err)
				return
			}

			h(p.RSSI(), registers)
		}

	}
}
