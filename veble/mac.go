package veble

import (
	"errors"
	"fmt"
)

var ErrInvalidFormat = errors.New("invalid MAC address format (expected format: E6:75:A3:1E:A8:72)")

type MAC [6]byte

// ParseColonMAC a mac address given in the format E6:75:A3:1E:A8:72.
func ParseColonMAC(s string) (MAC, error) {
	if len(s) != 17 {
		return MAC{}, fmt.Errorf("%w: expect format E6:75:A3:1E:A8:72 but got '%s'", ErrInvalidFormat, s)
	}
	var mac MAC
	_, err := fmt.Sscanf(s, "%02x:%02x:%02x:%02x:%02x:%02x", &mac[0], &mac[1], &mac[2], &mac[3], &mac[4], &mac[5])
	if err != nil {
		return MAC{}, fmt.Errorf("%w: expect format E6:75:A3:1E:A8:72 but got '%s'", ErrInvalidFormat, s)
	}
	return mac, nil
}

// ParseCompactMAC parses a mac address given in the format e675a31ea872 provided by the Victron Energy app.
func ParseCompactMAC(s string) (MAC, error) {
	if len(s) != 12 {
		return MAC{}, fmt.Errorf("%w: expect format e675a31ea872 but got '%s'", ErrInvalidFormat, s)
	}
	var mac MAC
	_, err := fmt.Sscanf(s, "%02x%02x%02x%02x%02x%02x", &mac[0], &mac[1], &mac[2], &mac[3], &mac[4], &mac[5])
	if err != nil {
		return MAC{}, fmt.Errorf("%w: expect format e675a31ea872 but got '%s'", ErrInvalidFormat, s)
	}
	return mac, nil
}

// As capitalized string with : separators. E.g. E6:75:A3:1E:A8:72.
func (m MAC) String() string {
	return fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X", m[0], m[1], m[2], m[3], m[4], m[5])
}
