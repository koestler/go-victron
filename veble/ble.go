package veble

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/koestler/go-victron/bleparser"
	"github.com/koestler/go-victron/veconst"
	"github.com/koestler/go-victron/veproduct"
	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	"log"
	"runtime/debug"
	"strings"
)

func (ble *BleStruct) connectDevice(dev *device.Device1, deviceConfig DeviceConfig) error {
	propUpdates, err := dev.WatchProperties()
	if err != nil {
		return fmt.Errorf("cannot watch props: %s", err)
	}

	defer func() {
		// unwatch properties calls recover
		if r := recover(); r != nil {
			panic(r)
		}

		if err := dev.UnwatchProperties(propUpdates); err != nil {
			log.Printf("error during unwatch: %s", err)
		} else {
			log.Printf("unwatch ok")
		}
	}()

	var lastRawBytes []uint8

	for pu := range propUpdates {
		if pu.Name == "ManufacturerData" {
			data := dev.Properties.ManufacturerData

			var rawBytes []uint8
			if md, ok := data[veconst.BleManufacturerId]; !ok {
				log.Printf("ble[%s]->%s: invalid manufacturer data record",
					ble.cfg.Name(), deviceConfig.Name(),
				)
				continue
			} else {
				rawBytes = md.([]uint8)
			}

			if bytes.Equal(rawBytes, lastRawBytes) {
				continue
			}

			ble.handleNewManufacturerData(deviceConfig, rawBytes)

			lastRawBytes = rawBytes
		}
	}

	return nil
}

func (ble *BleStruct) handleNewManufacturerData(deviceConfig DeviceConfig, rawBytes []uint8) {
	if ble.cfg.LogDebug() {
		log.Printf("ble[%s]->%s: handle len=%d, rawBytes=%x",
			ble.cfg.Name(), deviceConfig.Name(), len(rawBytes), rawBytes,
		)
	}

	if len(rawBytes) < 9 {
		log.Printf("ble[%s]->%s: len(rawBytes) is to low",
			ble.cfg.Name(), deviceConfig.Name(),
		)
		return
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
	product := veproduct.Product(productId)
	recordType := rawBytes[4]
	nonce := rawBytes[5:7] // used ad iv for encryption; is only 16 bits
	iv := binary.LittleEndian.Uint16(nonce)

	firstByteOfEncryptionKey := rawBytes[7]
	encryptedBytes := rawBytes[8:]

	log.Printf("ble[%s]->%s: prefix=%x productId=%x productString=%s recordType=%x, nonce=%d, firstByteOfEncryptionKey=%x",
		ble.cfg.Name(), deviceConfig.Name(), prefix, productId, product.String(), recordType, nonce, firstByteOfEncryptionKey,
	)

	log.Printf("ble[%s]->%s: encryptedBytes=%x, len=%d",
		ble.cfg.Name(), deviceConfig.Name(), encryptedBytes, len(encryptedBytes),
	)

	// decrypt rawBytes using aes-ctr algorithm
	// encryption key of config is fixed to 32 hex chars, so 16 bytes, so 128-bit AES is used here
	block, err := aes.NewCipher(deviceConfig.EncryptionKey())
	if err != nil {
		log.Printf("ble[%s]->%s: cannot create aes cipher: %s",
			ble.cfg.Name(), deviceConfig.Name(), err,
		)
		return
	}

	paddedEncryptedBytes := PKCS7Padding(encryptedBytes, block.BlockSize())

	log.Printf("ble[%s]->%s: paddedEncryptedBytes=%x, len=%d",
		ble.cfg.Name(), deviceConfig.Name(), paddedEncryptedBytes, len(paddedEncryptedBytes),
	)

	decryptedBytes := make([]byte, len(paddedEncryptedBytes))

	// iv needs to be 16 bytes for 128-bit AES, use nonce and pad with 0
	ivBytes := make([]byte, 16)
	binary.LittleEndian.PutUint16(ivBytes, iv)

	log.Printf("ble[%s]->%s: iv=%d ivBytes=%x, len=%d",
		ble.cfg.Name(), deviceConfig.Name(),
		iv, ivBytes, len(ivBytes),
	)

	ctrStream := cipher.NewCTR(block, ivBytes)
	ctrStream.XORKeyStream(decryptedBytes, paddedEncryptedBytes)

	log.Printf("ble[%s]->%s: decryptedBytes=%x, len=%d",
		ble.cfg.Name(), deviceConfig.Name(), decryptedBytes, len(decryptedBytes),
	)

	// handle decryptedBytes
	switch recordType {
	case 0x01:
		// solar charger
		record, err := bleparser.DecodeSolarChargeRecord(decryptedBytes)
		if err != nil {
			log.Printf("ble[%s]->%s: cannot decode solar charger record: %s",
				ble.cfg.Name(), deviceConfig.Name(), err,
			)
			return
		}
		log.Printf("ble[%s]->%s: solar charger record=%#v", ble.cfg.Name(), deviceConfig.Name(), record)
	}
}

func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (ble *BleStruct) getDeviceConfig(bluezAddr string) DeviceConfig {
	for _, d := range ble.cfg.Devices() {
		if bytes.Equal(d.MacAddress(), bluezAddrBytes(bluezAddr)) {
			return d
		}
	}
	return nil
}

// input: D4:9D:D2:92:62:02
// output []byte{0x, 0xd4, 0x9d, 0xd2, 0x92, 0x62, 0x02}
func bluezAddrBytes(i string) []byte {
	b, err := hex.DecodeString(strings.ReplaceAll(i, ":", ""))
	if err != nil {
		log.Printf("cannot decode %s, got: %s", i, err)
	}
	return b
}

func (ble *BleStruct) Name() string {
	return ble.cfg.Name()
}

func (ble *BleStruct) Shutdown() {
	ble.cancel()
}
