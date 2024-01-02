// Package veconst contains various enum values used by the VE.Direct, the VE.Bus and the BLE protocols.
package veconst

import "errors"

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
