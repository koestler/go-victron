package vedirect

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

func littleEndianBytesToInt(input []byte) (res int64, err error) {
	length := len(input)
	buf := bytes.NewReader(input)

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
		return 0, fmt.Errorf("vecommand: littleEndianBytesToInt: unhandled length=%v, input=%x", length, input)
	}

	if err != nil {
		return 0, fmt.Errorf("vecommand: littleEndianBytesToInt: binary.Read failed: %w", err)
	}

	return
}
