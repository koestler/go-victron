package vedirect

import (
	"bytes"
	"encoding/binary"
	"log"
)

func littleEndianBytesToUint(bytes []byte) (res uint64) {
	for i, b := range bytes {
		res |= uint64(b) << uint(i*8)
		if i >= 7 {
			break
		}
	}
	return
}

func littleEndianBytesToInt(input []byte) (res int64) {
	length := len(input)
	buf := bytes.NewReader(input)
	var err error

	switch length {
	case 1:
		var v int8
		err = binary.Read(buf, binary.LittleEndian, &v)
		res = int64(v)
	case 2:
		var v int16
		err = binary.Read(buf, binary.LittleEndian, &v)
		res = int64(v)
	case 4:
		var v int32
		err = binary.Read(buf, binary.LittleEndian, &v)
		res = int64(v)
	case 8:
		var v int64
		err = binary.Read(buf, binary.LittleEndian, &v)
		res = v
	default:
		log.Printf("vecommand: littleEndianBytesToInt: unhandled length=%v, input=%x", length, input)
		return 0
	}

	if err != nil {
		log.Printf("vecommand: littleEndianBytesToInt: binary.Read failed: %v", err)
		return 0
	}

	return
}
