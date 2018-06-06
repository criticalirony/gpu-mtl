// +build darwin

#include <stdlib.h>
#import <Metal/Metal.h>
#include "mtl.h"

struct Device CreateSystemDefaultDevice() {
	id<MTLDevice> device = MTLCreateSystemDefaultDevice();
	if (!device) {
		struct Device d;
		d.device = NULL;
		return d;
	}

	struct Device d;
	d.device = device;
	d.headless = device.headless;
	d.lowPower = device.lowPower;
	d.removable = device.removable;
	d.registryID = device.registryID;
	d.name = device.name.UTF8String;
	return d;
}

// Caller must call free(d.devices).
struct Devices CopyAllDevices() {
	NSArray<id<MTLDevice>> * devices = MTLCopyAllDevices();

	struct Devices d;
	d.devices = malloc(devices.count * sizeof(struct Device));
	for (int i = 0; i < devices.count; i++) {
		d.devices[i].device = devices[i];
		d.devices[i].headless = devices[i].headless;
		d.devices[i].lowPower = devices[i].lowPower;
		d.devices[i].removable = devices[i].removable;
		d.devices[i].registryID = devices[i].registryID;
		d.devices[i].name = devices[i].name.UTF8String;
	}
	d.length = devices.count;
	return d;
}

BOOL Device_SupportsFeatureSet(void * device, uint16_t featureSet) {
	return [(id<MTLDevice>)device supportsFeatureSet:featureSet];
}

void * Device_NewCommandQueue(void * device) {
	return [(id<MTLDevice>)device newCommandQueue];
};

void * CommandQueue_CommandBuffer(void * commandQueue) {
	return [(id<MTLCommandQueue>)commandQueue commandBuffer];
};
