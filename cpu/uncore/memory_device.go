package uncore

type MemoryDeviceType string

const (
	MemoryDeviceType_L1I_CONTROLLER = MemoryDeviceType("l1i")
	MemoryDeviceType_L1D_CONTROLLER = MemoryDeviceType("l1d")
	MemoryDeviceType_L2_CONTROLLER = MemoryDeviceType("l2")
	MemoryDeviceType_MEMORY_CONTROLLER = MemoryDeviceType("mem")
)

type MemoryDevice struct {
	MemoryHierarchy *MemoryHierarchy
	Name            string
	DeviceType      MemoryDeviceType
}

func NewMemoryDevice(memoryHierarchy *MemoryHierarchy, name string, deviceType MemoryDeviceType) *MemoryDevice {
	var memoryDevice = &MemoryDevice{
		MemoryHierarchy:memoryHierarchy,
		Name:name,
		DeviceType:deviceType,
	}

	return memoryDevice
}

func (memoryDevice *MemoryDevice) Transfer(to *MemoryDevice, size uint32, onCompletedCallback func()) {
	memoryDevice.MemoryHierarchy.Transfer(memoryDevice, to, size, onCompletedCallback)
}
