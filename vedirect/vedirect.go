// Package vedirect implements the VE.Direct serial protocol.
package vedirect

import (
	"bufio"
	"bytes"
	"encoding/binary"
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

// NewVedirect creates a new Vedirect instance.
func NewVedirect(cfg *Config) (*Vedirect, error) {
	if cfg.IOPort == nil {
		return nil, ErrInvalidConfig
	}

	vd := &Vedirect{
		cfg:            cfg,
		reader:         bufio.NewReader(cfg.IOPort),
		logDebugIndent: 0,
	}
	if vd.cfg.IoLogger != nil {
		vd.resetIoLogBuffers()
	}
	return vd, nil
}

func (vd *Vedirect) resetIoLogBuffers() {
	vd.logIoTxBuff = make([]byte, 0, 32)
	vd.logIoRxBuff = make([]byte, 0, 32)
}

// Ping sends a ping command to the device and waits for a response.
func (vd *Vedirect) Ping() (err error) {
	if vd.cfg.DebugLogger != nil || vd.cfg.IoLogger != nil {
		vd.debugPrintf("Ping() begin")
		defer func() {
			vd.ioLoggerLineEnd("Ping()")
			vd.debugPrintf("Ping end err=%v", err)
		}()
	}

	_, err = vd.sendReceive(VeCommandPing, []byte{})
	return
}

// GetDeviceId fetches what Victron Energy calls the device id.
// It is not a serial number, but it is a product id which can be decoded using veproduct.
func (vd *Vedirect) GetDeviceId() (deviceId uint16, err error) {
	if vd.cfg.DebugLogger != nil || vd.cfg.IoLogger != nil {
		vd.debugPrintf("GetDeviceId() begin")
		defer func() {
			vd.ioLoggerLineEnd("GetDeviceId() = 0x%X", deviceId)
			vd.debugPrintf("GetDeviceId end deviceId=%x err=%v", deviceId, err)
		}()
	}

	rawValue, err := vd.VeCommand(VeCommandDeviceId, 0)
	if err != nil {
		return
	}

	deviceId = binary.LittleEndian.Uint16(rawValue)
	return
}

// GetUint fetches the addressed register assuming it contains an unsigned integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetUint(address uint16) (value uint64, err error) {
	if vd.cfg.DebugLogger != nil || vd.cfg.IoLogger != nil {
		vd.debugPrintf("GetUint(address=0x%X) begin", address)
		defer func() {
			vd.ioLoggerLineEnd("GetUint(0x%X) = %d", address, value)
			vd.debugPrintf("GetUint end value=%d end=%v", value, err)
		}()
	}

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		return
	}

	value = littleEndianBytesToUint(rawValue)
	return
}

// GetInt fetches the addressed register assuming it contains a signed integer of 1, 2, 4 or 8 bytes.
func (vd *Vedirect) GetInt(address uint16) (value int64, err error) {
	if vd.cfg.DebugLogger != nil || vd.cfg.IoLogger != nil {
		vd.debugPrintf("GetInt(address=0x%X) begin", address)
		defer func() {
			vd.ioLoggerLineEnd("GetInt(0x%X) = %d", address, value)
			vd.debugPrintf("GetInt end value=%d err=%v", value, err)
		}()
	}

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		return
	}

	value, err = littleEndianBytesToInt(rawValue)
	return
}

// GetString fetches the addressed register assuming it contains a string of arbitrary length.
func (vd *Vedirect) GetString(address uint16) (value string, err error) {
	if vd.cfg.DebugLogger != nil || vd.cfg.IoLogger != nil {
		vd.debugPrintf("GetString(address=0x%X) begin", address)
		defer func() {
			vd.ioLoggerLineEnd("GetString(0x%X) = %s", address, value)
			vd.debugPrintf("GetString end value=%v err=%v", value, err)
		}()
	}

	rawValue, err := vd.VeCommandGet(address)
	if err != nil {
		return
	}

	value = string(bytes.TrimRightFunc(rawValue, func(r rune) bool { return r == 0 }))
	return
}
