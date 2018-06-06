// +build darwin

typedef signed char BOOL;
typedef unsigned short uint16_t;
typedef unsigned long long uint64_t;

struct Device {
	void *       device;
	BOOL         headless;
	BOOL         lowPower;
	BOOL         removable;
	uint64_t     registryID;
	const char * name;
};

struct Devices {
	struct Device * devices;
	int             length;
};

struct Device CreateSystemDefaultDevice();
struct Devices CopyAllDevices();
BOOL Device_SupportsFeatureSet(void * device, uint16_t featureSet);
void * Device_NewCommandQueue(void * device);
void * CommandQueue_CommandBuffer(void * commandQueue);
