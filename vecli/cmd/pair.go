package cmd

import (
	"encoding/hex"
	"errors"
	"github.com/koestler/go-victron/veble"
	"regexp"
)

type macKeyPair struct {
	mac veble.MAC
	key []byte
}

var pairMatcher = regexp.MustCompile("^([0-9a-f]{12})=([0-9a-f]{32})$")

var errInvalidMacKeyPairFormat = errors.New("invalid mac key pair format")

func parseMacKeyPair(arg string) (p macKeyPair, err error) {
	m := pairMatcher.FindStringSubmatch(arg)
	if m == nil {
		return p, errInvalidMacKeyPairFormat
	}
	macStr := m[1]
	keyStr := m[2]

	p.mac, err = veble.ParseCompactMAC(macStr)
	if err != nil {
		return p, err
	}

	p.key, err = hex.DecodeString(keyStr)
	if err != nil {
		return p, err
	}

	return p, nil
}
