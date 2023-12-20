// Package vedirect implements the VE.Direct serial protocol.
package vedirect

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// IOPort is the interface for the underlying serial port. It is implemented by e.g. tarm/serial.
type IOPort interface {
	io.ReadWriteCloser
	Flush() error
}

// Logger is the interface for a logger. It is implemented by e.g. log.Logger.
type Logger interface {
	Println(v ...any)
}

// Config is the configuration for a Vedirect instance.
type Config struct {
	IOPort      IOPort // mandatory: an implementation of a serial port
	DebugLogger Logger // optional: a logger for debug output; if nil, no debug output is written
	IoLogger    Logger // optional: a logger for all io operations; if nil, no io output is written
}

// Vedirect is the main struct for the VE.Direct serial protocol driver.
// There should be one instance per physical device.
type Vedirect struct {
	cfg      *Config
	reader   *bufio.Reader
	lastSent time.Time

	logDebugIndent int
	logIoTxBuff    []byte
	logIoRxBuff    []byte
}

var ErrNoIOPort = fmt.Errorf("no io port")

// NewVedirect creates a new Vedirect instance.
func NewVedirect(cfg *Config) (*Vedirect, error) {
	if cfg.IOPort == nil {
		return nil, ErrNoIOPort
	}

	vd := &Vedirect{
		cfg:            cfg,
		reader:         bufio.NewReader(cfg.IOPort),
		logDebugIndent: 0,
	}
	vd.clearIoLogBuffers()
	return vd, nil
}

func (vd *Vedirect) clearIoLogBuffers() {
	if vd.cfg.IoLogger != nil {
		vd.logIoTxBuff = make([]byte, 0, 255)
		vd.logIoRxBuff = make([]byte, 0, 255)
	}
}

// Ping sends a ping command to the device and waits for a response.
func (vd *Vedirect) Ping() (err error) {
	vd.debugPrintf("Ping begin")

	_, err = vd.sendReceive(VeCommandPing, []byte{})
	if err != nil {
		vd.debugPrintf("Ping end err=%v", err)
		return err
	}

	vd.ioLoggerLineEnd("Ping()")
	vd.debugPrintf("Ping end")
	return nil
}

// GetDeviceId fetches what Victron Energy calls the device id.
// It is not a serial number, but it is a product id which can be decoded using veproduct.
func (vd *Vedirect) GetDeviceId() (deviceId uint16, err error) {
	vd.debugPrintf("GetDeviceId begin")

	rawValue, err := vd.VeCommand(VeCommandDeviceId, 0)
	if err != nil {
		vd.debugPrintf("GetDeviceId end err=%v", err)
		return 0, err
	}

	deviceId = binary.LittleEndian.Uint16(rawValue)

	vd.ioLoggerLineEnd("GetDeviceId()")
	vd.debugPrintf("GetDeviceId end deviceId=%x", deviceId)
	return deviceId, nil
}

// GetUint fetches the addressed register assuming it contains an unsigned integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetUint(address uint16) (value uint64, err error) {
	vd.debugPrintf("GetUint begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("GetUint end err=%v", err)
		return
	}

	value = littleEndianBytesToUint(rawValue)

	vd.ioLoggerLineEnd("GetUint(0x%X)", address)
	vd.debugPrintf("GetUint end value=%v", value)
	return
}

// GetInt fetches the addressed register assuming it contains a signed integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetInt(address uint16) (value int64, err error) {
	vd.debugPrintf("GetInt begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("GetInt end err=%v", err)
		return
	}
	value, err = littleEndianBytesToInt(rawValue)

	vd.ioLoggerLineEnd("GetInt(0x%X)", address)
	vd.debugPrintf("GetInt end value=%v", value)
	return
}

// GetString fetches the addressed register assuming it contains a string of arbitrary length.
func (vd *Vedirect) GetString(address uint16) (value string, err error) {
	vd.debugPrintf("GetString begin")

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		vd.debugPrintf("GetString end err=%v", err)
		return
	}

	value = string(bytes.TrimRightFunc(rawValue, func(r rune) bool { return r == 0 }))

	vd.ioLoggerLineEnd("GetString(0x%X)", address)
	vd.debugPrintf("GetString end value=%v", value)
	return
}
