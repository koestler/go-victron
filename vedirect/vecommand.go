package vedirect

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"
)

func (vd *Vedirect) VeCommand(command VeCommand, address uint16) (values []byte, err error) {
	vd.debugPrintf("VeCommand begin command=%v, address=%x", command, address)

	var param []byte
	if command == VeCommandGet || command == VeCommandSet {
		id := []byte{byte(address), byte(address >> 8)}
		param = append(id, 0x00)
	}

	var responseData []byte
	responseData, err = vd.sendReceive(command, param)
	if err != nil {
		vd.debugPrintf("VeCommand end err=%v", err)
		return
	}

	if len(responseData) < 7 {
		err = fmt.Errorf("responseData too short, len(responseData)=%v", len(responseData))
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	}

	// extract and check command
	var response VeResponse
	s := string(responseData[0])
	if i, err := strconv.ParseUint(s, 16, 8); err != nil {
		err = fmt.Errorf("cannot parse response, address=%x, s=%v, err=%v", address, s, err)
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	} else {
		response = VeResponse(i)
	}

	expectedResponse := ResponseForCommand(command)
	if expectedResponse != response {
		err = fmt.Errorf("expectedResponse != response, expectedResponse=%v, response=%v",
			expectedResponse, response)
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	}

	// extract data
	hexData := responseData[1:]
	if len(hexData)%2 != 0 {
		err = fmt.Errorf("received an odd number of hex bytes, len(hexData)=%v", len(hexData))
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	}

	numbBytes := len(hexData) / 2
	binData := make([]byte, numbBytes)

	if n, err := hex.Decode(binData, hexData); err != nil || n != numbBytes {
		err = fmt.Errorf("hex to bin conversion failed: n=%v, err=%v", n, err)
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	}

	// extract and check checksum
	values = binData[:len(binData)-1]
	responseChecksum := binData[len(binData)-1]

	checksum := computeChecksum(byte(response), values)
	if checksum != responseChecksum {
		err = fmt.Errorf("checksum != responseChecksum, checksum=%X, responseChecksum=%X", checksum, responseChecksum)
		vd.debugPrintf("VeCommand end err=%v", err)
		return nil, err
	}

	vd.debugPrintf("VeCommand end")
	return
}

// VeCommandGet fetches the addressed register and returns its raw value or an error.
func (vd *Vedirect) VeCommandGet(address uint16) (value []byte, err error) {
	vd.debugPrintf("VeCommandGet begin address=%x", address)

	// fetch response using multiple tries to
	// deal with old data in the tx buffer of the veproduct device and our rx buffer
	const numbTries = 8
	for try := 0; try < numbTries; try++ {
		var rawValues []byte
		rawValues, err = vd.VeCommand(VeCommandGet, address)
		if err != nil {
			if try > 0 {
				log.Printf("VeCommandGet(address=%x) retry try=%v err=%v", address, try, err)
			}
			continue
		}

		// check address
		responseAddress := uint16(littleEndianBytesToUint(rawValues[0:2]))
		if address != responseAddress {
			err = fmt.Errorf("address != responseAddress, address=%x, responseAddress=%x", address, responseAddress)
			if try > 0 {
				log.Printf("VeCommandGet(address=%x) retry try=%v err=%v", address, try, err)
			}
			continue
		}

		// check flag
		responseFlag := VeResponseFlag(littleEndianBytesToUint(rawValues[2:3]))
		if responseErr := responseError(responseFlag); responseErr != nil {
			err = responseErr
			if try > 0 {
				log.Printf("VeCommandGet(address=%x) retry try=%v err=%v", address, try, err)
			}
			continue
		}

		// extract value
		vd.debugPrintf("VeCommandGet end")
		return rawValues[3:], nil
	}

	vd.debugPrintf("VeCommandGet(address=%x) end tries=%v last err=%v", address, numbTries, err)
	err = fmt.Errorf("gave up after %v tries, last err=%v", numbTries, err)
	return nil, err
}

func (vd *Vedirect) sendReceive(cmd VeCommand, data []byte) (response []byte, err error) {
	vd.debugPrintf("sendReceive begin")
	defer vd.debugPrintf("sendReceive end err=%v", err)

	now := time.Now()
	if now.Sub(vd.lastSent) > 200*time.Millisecond {
		// after a while, the BMV starts sending asynchronous messages
		// flush the receiver to get rid of them
		vd.flushReceiver()
	}
	vd.lastSent = now

	if err = vd.sendCommand(cmd, data); err != nil {
		return
	}

	response, err = vd.receiveResponse()
	return
}

func (vd *Vedirect) sendCommand(cmd VeCommand, data []byte) (err error) {
	vd.debugPrintf("sendCommand begin")

	checksum := computeChecksum(byte(cmd), data)
	str := fmt.Sprintf(":%X%X%X\n", cmd, data, checksum)

	_, err = vd.write([]byte(str))

	vd.debugPrintf("sendCommand end")
	return
}

func (vd *Vedirect) receiveResponse() (data []byte, err error) {
	vd.debugPrintf("receiveResponse begin")

	for {
		// search start marker
		_, err = vd.recvUntil(':')
		if err != nil {
			vd.debugPrintf("receiveResponse end err=%v", err)
			return nil, err
		}

		// search end marker
		data, err = vd.recvUntil('\n')
		if err != nil {
			vd.debugPrintf("receiveResponse end err=%v", err)
			return nil, err
		}

		if len(data) > 0 && data[0] == 'A' {
			vd.debugPrintf("receiveResponse async message received; ignore and read next response")
		} else {
			break
		}
	}

	vd.debugPrintf("receiveResponse end data=%s", data)
	return
}
