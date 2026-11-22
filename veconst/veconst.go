// Package veconst contains various enum values used by the VE.Direct, the VE.Bus and the BLE protocols.
package veconst

import (
	"errors"
	"math"
)

type Enum interface {
	Idx() int
	String() string
}

type EnumFactory interface {
	NewEnum(v int) (Enum, error)
	IntToStringMap() map[int]string
}

var ErrInvalidEnumIdx = errors.New("enum value does not exist")

type Field interface {
	Idx() int
	String() string
}

type FieldList interface {
	Fields() map[Field]bool
}

type FieldListFactory interface {
	NewFieldList(v uint) (FieldList, error)
	IntToStringMap() map[int]string
}

func enumIdxToUint8(v int) (uint8, error) {
	if v < 0 || v > math.MaxUint8 {
		return 0, ErrInvalidEnumIdx
	}
	return uint8(v), nil
}
