package vedirect

import "testing"

func TestChecksum(t *testing.T) {
	var tests = []struct {
		cmd    byte
		data   []byte
		expect byte
	}{
		{0x7, []byte{0x00, 0x10, 0x00}, 0x3E},
		{0x7, []byte{0x00, 0x10, 0x00, 0xC8, 0x00}, 0x76},
	}

	for _, tc := range tests {
		got := computeChecksum(tc.cmd, tc.data)
		if tc.expect != got {
			t.Errorf("expect %d but got %d for cmd=%x and data=%x", tc.expect, got, tc.cmd, tc.data)
		}
	}

}
