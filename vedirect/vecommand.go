package vedirect

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// VeCommandGet fetches the addressed register and returns its raw value or an error.
func (vd *Vedirect) VeCommandGet(address uint16) (value []byte, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("VeCommandGet(address=0x%X) begin", address)
		defer func() {
			vd.debugPrintf("VeCommandGet(address=0x%X) end value=%q, err=%v", address, value, err)
		}()
	}

	// fetch response using multiple tries to
	// deal with old data in the tx buffer of the veproduct device and our rx buffer
	const numbTries = 8
	for try := 0; try < numbTries; try++ {
		var rawValues []byte
		rawValues, err = vd.VeCommand(VeCommandGet, address)
		if err != nil {
			if try > 0 {
				vd.debugPrintf("retry try=%d err=%v", try, err)
			}
			continue
		}

		// check address
		responseAddress := uint16(littleEndianBytesToUint(rawValues[0:2]))
		if address != responseAddress {
			err = fmt.Errorf("address != responseAddress, 0x%X != 0x%X", address, responseAddress)
			if try > 0 {
				vd.debugPrintf("retry try=%d err=%v", try, err)
			}
			continue
		}

		// check flag
		responseFlag := VeResponseFlag(littleEndianBytesToUint(rawValues[2:3]))
		if e := responseError(responseFlag); e != nil {
			// do not retry an error returned by the device for the correct address as it will not change
			err = e
			return
		}

		// extract value
		value = rawValues[3:]
		return
	}

	err = fmt.Errorf("gave up after %d tries, last err=%v", numbTries, err)
	return
}

// VeCommand executes, sends the given command returns the raw response or an error.
// address is only used for VeCommandGet and VeCommandSet commands.
func (vd *Vedirect) VeCommand(command VeCommand, address uint16) (values []byte, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("VeCommand(command=0x%X, address=0x%X) begin", command, address)
		defer func() {
			vd.debugPrintf("VeCommand end values=%q err=%v", values, err)
		}()
	}

	var param []byte
	if command == VeCommandGet || command == VeCommandSet {
		id := []byte{byte(address), byte(address >> 8)}
		param = append(id, 0x00)
	}

	var responseData []byte
	responseData, err = vd.sendReceive(command, param)
	if err != nil {
		return
	}

	if len(responseData) < 7 {
		err = fmt.Errorf("responseData too short, len(responseData)=%d", len(responseData))
		return
	}

	// extract and check command
	var response VeResponse
	s := string(responseData[0])
	if i, e := strconv.ParseUint(s, 16, 8); err != nil {
		err = fmt.Errorf("cannot parse response, address=0x%X, s=%q: %s", address, s, e)
		return
	} else {
		response = VeResponse(i)
	}

	expectedResponse := ResponseForCommand(command)
	if expectedResponse != response {
		err = fmt.Errorf("expectedResponse != response, 0x%X != 0x%X", expectedResponse, response)
		return
	}

	// extract data
	hexData := responseData[1:]
	if len(hexData)%2 != 0 {
		err = fmt.Errorf("received an odd number of hex bytes, len(hexData)=%d", len(hexData))
		return
	}

	numbBytes := len(hexData) / 2
	binData := make([]byte, numbBytes)

	if n, e := hex.Decode(binData, hexData); e != nil || n != numbBytes {
		err = fmt.Errorf("hex to bin conversion failed: n=%d, e=%s", n, e)
		return
	}

	// extract and check checksum
	values = binData[:len(binData)-1]
	responseChecksum := binData[len(binData)-1]

	checksum := computeChecksum(byte(response), values)
	if checksum != responseChecksum {
		err = fmt.Errorf("checksum != responseChecksum, 0x%X != 0x%X", checksum, responseChecksum)
		return
	}

	return
}

func (vd *Vedirect) sendReceive(cmd VeCommand, data []byte) (response []byte, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("sendReceive(cmd=0x%X, data=%q) begin", cmd, data)
		defer func() {
			vd.debugPrintf("sendReceive end response=%q err=%v", response, err)
		}()
	}

	now := time.Now()
	if now.Sub(vd.lastSent) > 100*time.Millisecond {
		// after a while, the BMV starts sending asynchronous messages
		// flush the receiver to get rid of them before sending a command
		// otherwise we might use up all tries in VeCommandGet
		vd.flushReceiver()
	}
	vd.lastSent = now

	err = vd.sendCommand(cmd, data)
	if err != nil {
		return
	}

	response, err = vd.receiveResponse()
	return
}

func (vd *Vedirect) sendCommand(cmd VeCommand, data []byte) (err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("sendCommand(cmd=0x%X, data=%q) begin", cmd, data)
		defer func() {
			vd.debugPrintf("sendCommand end err=%v", err)
		}()
	}

	checksum := computeChecksum(byte(cmd), data)
	str := fmt.Sprintf(":%X%X%X\n", cmd, data, checksum)
	_, err = vd.write([]byte(str))
	return
}

func (vd *Vedirect) receiveResponse() (data []byte, err error) {
	if vd.cfg.DebugLogger != nil {
		vd.debugPrintf("receiveResponse() begin")
		defer func() {
			vd.debugPrintf("receiveResponse end data=%q, err=%v", data, err)
		}()
	}

	var received []byte
	for {
		// search start marker
		_, err = vd.recvUntil(':')
		if err != nil {
			return
		}

		// search end marker
		received, err = vd.recvUntil('\n')
		if err != nil {
			return
		}

		if len(received) > 0 && received[0] == 'A' {
			vd.debugPrintf("async message received; ignore and read next response")
		} else {
			data = received
			return
		}
	}
}
