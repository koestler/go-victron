package ble

type Config interface {
	Name() string
	LogDebug() bool
}
