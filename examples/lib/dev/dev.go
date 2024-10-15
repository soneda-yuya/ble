package dev

import (
	"github.com/go-ble/ble"
)

// NewDevice ...
func NewDevice(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return DefaultDevice(opts...)
}

// NewBLE5Device ...
func NewBLE5Device(impl string, opts ...ble.Option) (d ble.Device, err error) {
	return BLE5Device(opts...)
}
