package vebleapi

import "github.com/koestler/go-victron/mac"

type Adapter interface {
	RegisterMacListener(mac mac.MAC, l func(rssi int, localName string, data []byte))
	UnregisterMacListener(mac mac.MAC)
}
