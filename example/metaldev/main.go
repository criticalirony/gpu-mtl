// +build ignore
// +build darwin

// metaldev is a test program for developing the Metal API library.
package main

import (
	"log"

	"dmitri.shuralyov.com/gpu/mtl"
	"github.com/shurcooL/go-goon"
)

func main() {
	goon.DumpExpr(mtl.CopyAllDevices())

	device, err := mtl.CreateSystemDefaultDevice()
	if err != nil {
		log.Fatalln(err)
	}
	goon.DumpExpr(device)

	goon.DumpExpr(device.SupportsFeatureSet(mtl.MacOSGPUFamily1V1))
	goon.DumpExpr(device.SupportsFeatureSet(mtl.MacOSGPUFamily1V2))
	goon.DumpExpr(device.SupportsFeatureSet(mtl.MacOSGPUFamily1V3))
	goon.DumpExpr(device.SupportsFeatureSet(mtl.MacOSGPUFamily1V4))
	goon.DumpExpr(device.SupportsFeatureSet(mtl.MacOSGPUFamily2V1))

	commandQueue := device.NewCommandQueue()
	commandBuffer := commandQueue.CommandBuffer()
	_ = commandBuffer
}
