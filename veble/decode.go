package veble

import (
	"encoding/binary"
	"errors"
	"github.com/koestler/go-victron/veconst"
	"github.com/koestler/go-victron/velog"
	"github.com/koestler/go-victron/veproduct"
)

type EncryptedFrame struct {
	Product        veproduct.Product
	RecordType     veconst.BleRecordType
	IV             uint16
	EncryptedBytes []byte
}

var ErrInputTooShort = errors.New("inp too short")

func DecodeFrame(rawBytes []byte, l velog.Logger) (ef EncryptedFrame, error error) {
	l.Printf("decode frame len=%d, rawBytes=%x", len(rawBytes), rawBytes)

	if len(rawBytes) < 9 {
		return ef, ErrInputTooShort
	}

	// map rawBytes:
	// 00 - 01 : prefix
	// 02 - 03 : product id
	// 04 - 04 : record type
	// 05 - 06 : Nonce/Data counter in LSB order
	// 07 - 07 : first byte of encryption key
	// 08 -    : encrypted data

	prefix := rawBytes[0:2]
	productId := binary.LittleEndian.Uint16(rawBytes[2:4])
	ef.Product = veproduct.Product(productId)
	recordTypeId := rawBytes[4]
	ef.RecordType = veconst.BleRecordType(recordTypeId)
	nonce := rawBytes[5:7] // used ad iv for encryption; is only 16 bits
	ef.IV = binary.LittleEndian.Uint16(nonce)

	firstByteOfEncryptionKey := rawBytes[7]
	ef.EncryptedBytes = rawBytes[8:]

	l.Printf(
		"prefix=%x productId=%x productString=%s recordType=%x, recordTypeString=%s, nonce=%d, firstByteOfEncryptionKey=%x",
		prefix, productId, ef.Product, recordTypeId, ef.RecordType, nonce, firstByteOfEncryptionKey,
	)
	l.Printf("encryptedBytes=%x, len=%d", ef.EncryptedBytes, len(ef.EncryptedBytes))

	return ef, nil
}
