package uncore

type MemoryDeviceType string

const (
	MemoryDeviceType_L1I_CONTROLLER = MemoryDeviceType("l1i")
	MemoryDeviceType_L1D_CONTROLLER = MemoryDeviceType("l1d")
	MemoryDeviceType_L2_CONTROLLER = MemoryDeviceType("l2")
	MemoryDeviceType_MEMORY_CONTROLLER = MemoryDeviceType("mem")
)

type MemoryDevice interface {
	MemoryHierarchy() MemoryHierarchy
	Name() string
	DeviceType() MemoryDeviceType
	Transfer(to MemoryDevice, size uint32, onCompletedCallback func())
}

type BaseMemoryDevice struct {
	memoryHierarchy MemoryHierarchy
	name            string
	deviceType      MemoryDeviceType
}

func NewBaseMemoryDevice(memoryHierarchy MemoryHierarchy, name string, deviceType MemoryDeviceType) *BaseMemoryDevice {
	var baseMemoryDevice = &BaseMemoryDevice{
		memoryHierarchy:memoryHierarchy,
		name:name,
		deviceType:deviceType,
	}

	return baseMemoryDevice
}

func (baseMemoryDevice *BaseMemoryDevice) Transfer(to MemoryDevice, size uint32, onCompletedCallback func()) {
	baseMemoryDevice.memoryHierarchy.Transfer(baseMemoryDevice, to, size, onCompletedCallback)
}

func (baseMemoryDevice *BaseMemoryDevice) MemoryHierarchy() MemoryHierarchy {
	return baseMemoryDevice.memoryHierarchy
}

func (baseMemoryDevice *BaseMemoryDevice) Name() string {
	return baseMemoryDevice.name
}

func (baseMemoryDevice *BaseMemoryDevice) DeviceType() MemoryDeviceType {
	return baseMemoryDevice.deviceType
}
