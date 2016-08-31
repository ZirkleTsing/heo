package uncore

type Controller interface {
	MemoryDevice
	Cache() *EvictableCache
	Next() MemoryDevice
	SetNext(next MemoryDevice)
	ReceiveMessage(message CoherenceMessage)
	TransferMessage(to Controller, size uint32, message CoherenceMessage)
}

type BaseController struct {
	*BaseMemoryDevice
	cache                  *EvictableCache
	next                   MemoryDevice
	NumDownwardReadHits    int32
	NumDownwardReadMisses  int32
	NumDownwardWriteHits   int32
	NumDownwardWriteMisses int32
	NumEvictions           int32
}

func NewBaseController(memoryHierarchy *MemoryHierarchy, name string, deviceType MemoryDeviceType, cache *EvictableCache) *BaseController {
	var baseController = &BaseController{
		BaseMemoryDevice:NewBaseMemoryDevice(memoryHierarchy, name, deviceType),
		cache:cache,
	}

	return baseController
}

func (baseController *BaseController) ReceiveMessage(message CoherenceMessage) {
	panic("Impossible")
}

func (baseController *BaseController) TransferMessage(to Controller, size uint32, message CoherenceMessage) {
	baseController.MemoryHierarchy().TransferMessage(baseController, to, size, message)
}

func (baseController *BaseController) UpdateStats(write bool, hitInCache bool) {
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

func (baseController *BaseController) Cache() *EvictableCache {
	return baseController.cache
}

func (baseController *BaseController) Next() MemoryDevice {
	return baseController.next
}

func (baseController *BaseController) SetNext(next MemoryDevice) {
	baseController.next = next
}

