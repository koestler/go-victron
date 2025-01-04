package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMacKeyPair(t *testing.T) {
	p, err := parseMacKeyPair("foobar-test=713f501f0b05beb18ec0947768162d4e")
	assert.NoError(t, err)
	assert.Equal(t, "foobar-test", p.name)
	assert.Equal(t, []byte{0x71, 0x3f, 0x50, 0x1f, 0x0b, 0x05, 0xbe, 0xb1, 0x8e, 0xc0, 0x94, 0x77, 0x68, 0x16, 0x2d, 0x4e}, p.key)
}
