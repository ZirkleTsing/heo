package uncore

type Controller interface {
	GetMemoryDevice() *MemoryDevice
	Cache() *EvictableCache
	Next() *MemoryDevice
}

type BaseController struct {
	*MemoryDevice
	cache                  *EvictableCache
	next                   *MemoryDevice
	NumDownwardReadHits    int32
	NumDownwardReadMisses  int32
	NumDownwardWriteHits   int32
	NumDownwardWriteMisses int32
	NumEvictions           int32
}

func NewBaseController(memoryHierarchy *MemoryHierarchy, name string, deviceType MemoryDeviceType, cache *EvictableCache) *BaseController {
	var baseController = &BaseController{
		MemoryDevice:NewMemoryDevice(memoryHierarchy, name, deviceType),
		cache:cache,
	}

	return baseController
}

func (baseController *BaseController) updateStats(write bool, hitInCache bool) {
	if write {
		if hitInCache {
			baseController.NumDownwardWriteHits++
		} else {
			baseController.NumDownwardWriteMisses++
		}
	} else {
		if hitInCache {
			baseController.NumDownwardReadHits++
		} else {
			baseController.NumDownwardReadMisses++
		}
	}
}

func (baseController *BaseController) NumDownwardHits() int32 {
	return baseController.NumDownwardReadHits + baseController.NumDownwardWriteHits
}

func (baseController *BaseController) NumDownwardMisses() int32 {
	return baseController.NumDownwardReadMisses + baseController.NumDownwardWriteMisses
}

func (baseController *BaseController) NumDownwardAccesses() int32 {
	return baseController.NumDownwardHits() + baseController.NumDownwardMisses()
}

func (baseController *BaseController) HitRatio() float32 {
	if baseController.NumDownwardAccesses() == 0 {
		return 0
	} else {
		return float32(baseController.NumDownwardHits()) / float32(baseController.NumDownwardAccesses())
	}
}

func (baseController *BaseController) OccupancyRatio() float32 {
	return baseController.cache.Cache.OccupancyRatio()
}

func (baseController *BaseController) GetMemoryDevice() *MemoryDevice {
	return baseController.MemoryDevice
}

func (baseController *BaseController) Name() string {
	return baseController.MemoryDevice.Name
}

func (baseController *BaseController) DeviceType() MemoryDeviceType {
	return baseController.MemoryDevice.DeviceType
}

func (baseController *BaseController) Cache() *EvictableCache {
	return baseController.cache
}

func (baseController *BaseController) Next() *MemoryDevice {
	return baseController.MemoryDevice
}

