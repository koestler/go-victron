package veble

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/veconst"
	"github.com/koestler/go-victron/veproduct"
)

var ErrInvalidEncryptionKey = errors.New("invalid encryption key")

type EncryptedFrame struct {
	Product        veproduct.Product
	RecordType     veconst.BleRecordType
	IV             uint16
	EncryptedBytes []byte
}

type DecryptedFrame struct {
	RecordType     veconst.BleRecordType
	DecryptedBytes []byte
}

func DecodeFrame(rawBytes []byte, l log.Logger) (ef EncryptedFrame, error error) {
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

func DecryptFrame(ef EncryptedFrame, encryptionKey []byte, l log.Logger) (df DecryptedFrame, error error) {
	l.Printf("decrypt frame len=%d, encryptedBytes=%x", len(ef.EncryptedBytes), ef.EncryptedBytes)

	// decrypt rawBytes using aes-ctr algorithm
	// encryption key of config is fixed to 32 hex chars, so 16 bytes, so 128-bit AES is used here
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return df, fmt.Errorf("%w: %s", ErrInvalidEncryptionKey, err)
	}

	paddedEncryptedBytes := PKCS7Padding(ef.EncryptedBytes, block.BlockSize())

	l.Printf("paddedEncryptedBytes=%x, len=%d", paddedEncryptedBytes, len(paddedEncryptedBytes))

	df.DecryptedBytes = make([]byte, len(paddedEncryptedBytes))

	// iv needs to be 16 bytes for 128-bit AES, use nonce and pad with 0
	ivBytes := make([]byte, 16)
	binary.LittleEndian.PutUint16(ivBytes, ef.IV)

	l.Printf("iv=%d, ivBytes=%x, len=%d", ef.IV, ivBytes, len(ivBytes))

	ctrStream := cipher.NewCTR(block, ivBytes)
	ctrStream.XORKeyStream(df.DecryptedBytes, paddedEncryptedBytes)

	l.Printf("decryptedBytes=%x, len=%d", df.DecryptedBytes, len(df.DecryptedBytes))

	df.RecordType = ef.RecordType
	return df, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
