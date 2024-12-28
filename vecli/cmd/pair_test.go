package cmd

import (
	"github.com/koestler/go-victron/veble"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMacKeyPair(t *testing.T) {
	p, err := parseMacKeyPair("e675a31ea872=713f501f0b05beb18ec0947768162d4e")
	assert.NoError(t, err)
	assert.Equal(t, veble.MAC{0x72, 0xa8, 0x1e, 0xa3, 0x75, 0xe6}, p.mac)
	assert.Equal(t, []byte{0x71, 0x3f, 0x50, 0x1f, 0x0b, 0x05, 0xbe, 0xb1, 0x8e, 0xc0, 0x94, 0x77, 0x68, 0x16, 0x2d, 0x4e}, p.key)
}
