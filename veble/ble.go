package veble

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/koestler/go-victron/log"
	"github.com/koestler/go-victron/veproduct"
)

var (
	ErrInputTooShort        = errors.New("input too short")
	ErrInvalidEncryptionKey = errors.New("invalid encryption key")
)

type DecryptedFrame struct {
	Product        veproduct.Product
	RecordType     uint8
	DecryptedBytes []byte
}

func DecryptFrame(rawBytes []byte, encryptionKey []byte, log log.Logger) (df DecryptedFrame, error error) {
	log.Printf("handle len=%d, rawBytes=%x", len(rawBytes), rawBytes)

	if len(rawBytes) < 9 {
		return df, ErrInputTooShort
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
	df.Product = veproduct.Product(productId)
	df.RecordType = rawBytes[4]
	nonce := rawBytes[5:7] // used ad iv for encryption; is only 16 bits
	iv := binary.LittleEndian.Uint16(nonce)

	firstByteOfEncryptionKey := rawBytes[7]
	encryptedBytes := rawBytes[8:]

	log.Printf(
		"prefix=%x productId=%x productString=%s recordType=%x, nonce=%d, firstByteOfEncryptionKey=%x",
		prefix, productId, df.Product, df.RecordType, nonce, firstByteOfEncryptionKey,
	)
	log.Printf("encryptedBytes=%x, len=%d", encryptedBytes, len(encryptedBytes))

	// decrypt rawBytes using aes-ctr algorithm
	// encryption key of config is fixed to 32 hex chars, so 16 bytes, so 128-bit AES is used here
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return df, fmt.Errorf("%w: %s", ErrInvalidEncryptionKey, err)
	}

	paddedEncryptedBytes := PKCS7Padding(encryptedBytes, block.BlockSize())

	log.Printf("paddedEncryptedBytes=%x, len=%d", paddedEncryptedBytes, len(paddedEncryptedBytes))

	df.DecryptedBytes = make([]byte, len(paddedEncryptedBytes))

	// iv needs to be 16 bytes for 128-bit AES, use nonce and pad with 0
	ivBytes := make([]byte, 16)
	binary.LittleEndian.PutUint16(ivBytes, iv)

	log.Printf("iv=%d, ivBytes=%x, len=%d", iv, ivBytes, len(ivBytes))

	ctrStream := cipher.NewCTR(block, ivBytes)
	ctrStream.XORKeyStream(df.DecryptedBytes, paddedEncryptedBytes)

	log.Printf("decryptedBytes=%x, len=%d", df.DecryptedBytes, len(df.DecryptedBytes))

	return df, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}
