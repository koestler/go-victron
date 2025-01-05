package veble

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/koestler/go-victron/veconst"
	"github.com/koestler/go-victron/velog"
)

type DecryptedFrame struct {
	RecordType     veconst.BleRecordType
	DecryptedBytes []byte
}

var ErrInvalidEncryptionKey = errors.New("invalid encryption key")

func DecryptFrame(ef EncryptedFrame, encryptionKey []byte, l velog.Logger) (df DecryptedFrame, error error) {
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
