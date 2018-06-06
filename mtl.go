// +build darwin

// Package mtl provides access to Apple's Metal API (https://developer.apple.com/documentation/metal).
//
// This package is in very early stages of development.
// Less than 5% of the Metal API surface is implemented.
// Current functionality is sufficient to list Metal-capable devices
// on the system and query basic information about them.
package mtl

import (
	"errors"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Metal
#include <stdlib.h>
#include "mtl.h"
*/
import "C"

// Device is abstract representation of the GPU that
// serves as the primary interface for a Metal app.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice.
type Device struct {
	device unsafe.Pointer

	// Headless indicates whether a device is configured as headless.
	Headless bool

	// LowPower indicates whether a device is low-power.
	LowPower bool

	// Removable determines whether or not a GPU is removable.
	Removable bool

	// RegistryID is the registry ID value for the device.
	RegistryID uint64

	// Name is the name of the device.
	Name string
}

// CreateSystemDefaultDevice returns the preferred system default Metal device.
//
// Reference: https://developer.apple.com/documentation/metal/1433401-mtlcreatesystemdefaultdevice.
func CreateSystemDefaultDevice() (Device, error) {
	d := C.CreateSystemDefaultDevice()
	if d.device == nil {
		return Device{}, errors.New("Metal is not supported on this system")
	}

	return Device{
		device:     d.device,
		Headless:   d.headless != 0,
		LowPower:   d.lowPower != 0,
		Removable:  d.removable != 0,
		RegistryID: uint64(d.registryID),
		Name:       C.GoString(d.name),
	}, nil
}

// CopyAllDevices returns all Metal devices in the system.
//
// Reference: https://developer.apple.com/documentation/metal/1433367-mtlcopyalldevices.
func CopyAllDevices() []Device {
	d := C.CopyAllDevices()
	defer C.free(unsafe.Pointer(d.devices))

	ds := make([]Device, d.length)
	for i := 0; i < len(ds); i++ {
		d := (*C.struct_Device)(unsafe.Pointer(uintptr(unsafe.Pointer(d.devices)) + uintptr(i)*C.sizeof_struct_Device))

		ds[i].device = d.device
		ds[i].Headless = d.headless != 0
		ds[i].LowPower = d.lowPower != 0
		ds[i].Removable = d.removable != 0
		ds[i].RegistryID = uint64(d.registryID)
		ds[i].Name = C.GoString(d.name)
	}
	return ds
}

// SupportsFeatureSet reports whether device d supports feature set fs.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433418-supportsfeatureset.
func (d Device) SupportsFeatureSet(fs FeatureSet) bool {
	return C.Device_SupportsFeatureSet(d.device, C.uint16_t(fs)) != 0
}

// NewCommandQueue creates a new command queue object.
//
// Reference: https://developer.apple.com/documentation/metal/mtldevice/1433388-newcommandqueue.
func (d Device) NewCommandQueue() CommandQueue {
	return CommandQueue{C.Device_NewCommandQueue(d.device)}
}

// FeatureSet defines a specific platform, hardware, and software configuration.
//
// Reference: https://developer.apple.com/documentation/metal/mtlfeatureset.
type FeatureSet uint16

const (
	MacOSGPUFamily1V1 FeatureSet = 10000 // The GPU family 1, version 1 feature set for macOS.
	MacOSGPUFamily1V2 FeatureSet = 10001 // The GPU family 1, version 2 feature set for macOS.
	MacOSGPUFamily1V3 FeatureSet = 10003 // The GPU family 1, version 3 feature set for macOS.
	MacOSGPUFamily1V4 FeatureSet = 10004 // The GPU family 1, version 4 feature set for macOS.
	MacOSGPUFamily2V1 FeatureSet = 10005 // The GPU family 2, version 1 feature set for macOS.
)

// CommandQueue is a queue that organizes the order
// in which command buffers are executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue.
type CommandQueue struct {
	commandQueue unsafe.Pointer
}

// CommandBuffer creates a command buffer.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508686-commandbuffer.
func (cq CommandQueue) CommandBuffer() CommandBuffer {
	return CommandBuffer{C.CommandQueue_CommandBuffer(cq.commandQueue)}
}

// CommandBuffer is a container that stores encoded commands
// that are committed to and executed by the GPU.
//
// Reference: https://developer.apple.com/documentation/metal/mtlcommandbuffer.
type CommandBuffer struct {
	commandBuffer unsafe.Pointer
}
