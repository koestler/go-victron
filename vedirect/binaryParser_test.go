package vedirect

import "testing"

func TestLittleEndianBytesToUint(t *testing.T) {
	var tests = []struct {
		inp    []byte
		expect uint64
	}{
		{[]byte{0x00, 0x00}, 0},
		{[]byte{0x01, 0x00}, 1},
		{[]byte{0x01, 0x00, 0x00, 0x00}, 1},
		{[]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 1},
		{[]byte{0x00, 0x01}, 256},
		{[]byte{0x01, 0x01}, 257},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, 4294967295},
	}

	for _, tc := range tests {
		got := littleEndianBytesToUint(tc.inp)
		if tc.expect != got {
			t.Errorf("expect %d but got %d for input %x", tc.expect, got, tc.inp)
		}
	}
}

func TestLittleEndianBytesToInt(t *testing.T) {
	var tests = []struct {
		inp    []byte
		expect int64
	}{
		{[]byte{0x00, 0x00}, 0},
		{[]byte{0x01}, 1},
		{[]byte{0x01, 0x00}, 1},
		{[]byte{0x01, 0x00, 0x00, 0x00}, 1},
		{[]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 1},
		{[]byte{0x00, 0x01}, 256},
		{[]byte{0x01, 0x01}, 257},
		{[]byte{0xFF, 0xFF}, -1},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, -1},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, -1},
	}

	for _, tc := range tests {
		got, err := littleEndianBytesToInt(tc.inp)
		if err != nil {
			t.Errorf("littleEndianBytesToInt failed: %s", err)
		}
		if tc.expect != got {
			t.Errorf("expect %d but got %d for input %x", tc.expect, got, tc.inp)
		}
	}

	if _, err := littleEndianBytesToInt([]byte{0x01, 0x00, 0x00}); err == nil {
		t.Errorf("expect error for length=3")
	}
}
