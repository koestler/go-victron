// Package vedirect implements the VE.Direct serial protocol.
package vedirect

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/koestler/go-victron/veproduct"
	"io"
)

type IoPort interface {
	io.ReadWriteCloser
	Flush() error
}

type Config struct {
	ioPort      IoPort
	debugLogger io.Writer
	ioLogger    io.Writer
}

type Vedirect struct {
	cfg            *Config
	reader         *bufio.Reader
	logDebugIndent int
	lastWritten    []byte
}

// NewVedirect creates a new Vedirect instance.
func NewVedirect(cfg *Config) *Vedirect {
	return &Vedirect{
		cfg,
		bufio.NewReader(cfg.ioPort),
		0,
		nil,
	}
}

// FlushReceiver flushes the underlying receiver buffer.
// Todo: remove this function; do it automatically after some inactivity.
func (vd *Vedirect) FlushReceiver() {
	vd.debugPrintf("vedirect: FlushReceiver begin")
	vd.flushReceiver()
	vd.debugPrintf("vedirect: FlushReceiver end")
}

// Ping sends a ping command to the device and waits for a response.
func (vd *Vedirect) Ping() (err error) {
	vd.debugPrintf("vedirect: VeCommandPing begin")

	err = vd.sendVeCommand(VeCommandPing, []byte{})
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandPing end err=%v", err)
		return err
	}

	_, err = vd.recvVeResponse()
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandPing end err=%v", err)
		return err
	}

	vd.debugPrintf("vedirect: VeCommandPing end")
	return nil
}

// GetDeviceId fetches what Victron Energy calls the device id.
// It is not a serial number, but it is a product id which can be decoded using veproduct.
func (vd *Vedirect) GetDeviceId() (deviceId veproduct.Product, err error) {
	vd.debugPrintf("vedirect: VeCommandDeviceId begin")

	rawValue, err := vd.VeCommand(VeCommandDeviceId, 0)
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandDeviceId end err=%v", err)
		return 0, err
	}

	deviceId = veproduct.Product(littleEndianBytesToUint(rawValue))
	if len(deviceId.String()) < 1 {
		vd.debugPrintf("vedirect: VeCommandDeviceId end unknown deviceId=%x", rawValue)
		return 0, fmt.Errorf("unknownw deviceId=%x", rawValue)
	}

	vd.debugPrintf("vedirect: VeCommandDeviceId end deviceId=%x", deviceId)
	return deviceId, nil
}

// GetUint fetches the addressed register assuming it contains an unsigned integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetUint(address uint16) (value uint64, err error) {
	vd.debugPrintf("vedirect: VeCommandGetUint begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandGetUint end err=%v", err)
		return
	}

	value = littleEndianBytesToUint(rawValue)
	vd.debugPrintf("vedirect: VeCommandGetUint end value=%v", value)
	return
}

// GetInt fetches the addressed register assuming it contains a signed integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetInt(address uint16) (value int64, err error) {
	vd.debugPrintf("vedirect: VeCommandGetInt begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandGetInt end err=%v", err)
		return
	}
	value = littleEndianBytesToInt(rawValue)

	vd.debugPrintf("vedirect: VeCommandGetInt end value=%v", value)
	return
}

// GetString fetches the addressed register assuming it contains a string of arbitrary length.
func (vd *Vedirect) GetString(address uint16) (value string, err error) {
	vd.debugPrintf("vedirect: VeCommandGetString begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("vedirect: VeCommandGetString end err=%v", err)
		return
	}

	value = string(bytes.TrimRightFunc(rawValue, func(r rune) bool { return r == 0 }))

	vd.debugPrintf("vedirect: VeCommandGetString end value=%v", value)
	return
}
