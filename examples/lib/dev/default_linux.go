package dev

import (
	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}

// BLE5Device ...
func BLE5Device(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewBLE5Device(opts...)
}
