package cmd

import (
	"encoding/hex"
	"errors"
	"log"
	"regexp"
)

type macKeyPair struct {
	name string
	key  []byte
}

var pairMatcher = regexp.MustCompile("^([0-9a-zA-Z\\s\\-]{1,32})=([0-9a-f]{32})$")

var errInvalidMacKeyPairFormat = errors.New("invalid mac key pair format")

func parseMacKeyPair(arg string) (p macKeyPair, err error) {
	log.Printf("parsing mac key pair: %s", arg)

	m := pairMatcher.FindStringSubmatch(arg)
	if m == nil {
		return p, errInvalidMacKeyPairFormat
	}
	nameStr := m[1]
	keyStr := m[2]

	p.name = nameStr
	p.key, err = hex.DecodeString(keyStr)
	if err != nil {
		return p, err
	}

	return p, nil
}
