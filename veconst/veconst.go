// Package veconst contains various enum values used by the VE.Direct, the VE.Bus and the BLE protocols.
package veconst

import "errors"

var ErrInvalidEnumIdx = errors.New("enum value does not exist")

type Enum interface {
	Idx() int
	Value() string
}

type EnumFactory func(v int) (Enum, error)
