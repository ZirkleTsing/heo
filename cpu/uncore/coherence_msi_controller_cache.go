package uncore

import "github.com/mcai/acogo/cpu/mem"

type CacheController struct {
	*BaseController
	NumReadPorts              uint32
	NumWritePorts             uint32
	HitLatency                uint32
	PendingAccesses           map[uint32]*MemoryHierarchyAccess
	NumPendingAccessesPerType map[MemoryHierarchyAccessType]uint32
	FsmFactory *CacheControllerFiniteStateMachineFactory
}

func NewCacheController(memoryHierarchy *MemoryHierarchy, name string, deviceType MemoryDeviceType, geometry *mem.Geometry, replacementPolicyType CacheReplacementPolicyType, numReadPorts uint32, numWritePorts uint32, hitLatency uint32) *CacheController {
	var cacheController = &CacheController{
		BaseController:NewBaseController(memoryHierarchy, name, deviceType, NewEvictableCache(geometry, replacementPolicyType)),
		NumReadPorts:numReadPorts,
		NumWritePorts:numWritePorts,
		HitLatency:hitLatency,
		PendingAccesses:make(map[uint32]*MemoryHierarchyAccess),
		NumPendingAccessesPerType:make(map[MemoryHierarchyAccessType]uint32),
	}

	cacheController.NumPendingAccessesPerType[MemoryHierarchyAccessType_IFETCH] = 0
	cacheController.NumPendingAccessesPerType[MemoryHierarchyAccessType_LOAD] = 0
	cacheController.NumPendingAccessesPerType[MemoryHierarchyAccessType_STORE] = 0

	cacheController.FsmFactory = NewCacheControllerFiniteStateMachineFactory()

	return cacheController
}

func (cacheController *CacheController) FindAccess(physicalTag uint32) *MemoryHierarchyAccess {
	if pendingAccess, ok := cacheController.PendingAccesses[physicalTag]; ok {
		return pendingAccess
	} else {
		return nil
	}
}

func (cacheController *CacheController) CanAccess(accessType MemoryHierarchyAccessType, physicalTag uint32) bool {
	var access = cacheController.FindAccess(physicalTag)

	if access == nil {
		if accessType == MemoryHierarchyAccessType_STORE {
			return cacheController.NumPendingAccessesPerType[accessType] < cacheController.NumWritePorts
		} else {
			return cacheController.NumPendingAccessesPerType[accessType] < cacheController.NumReadPorts
		}
	} else {
		return accessType != MemoryHierarchyAccessType_STORE
	}
}

func (cacheController *CacheController) BeginAccess(accessType MemoryHierarchyAccessType, threadId int32, virtualPc int32, physicalAddress uint32, physicalTag uint32, onCompletedCallback func()) *MemoryHierarchyAccess {
	var newAccess = NewMemoryHierarchyAccess(cacheController.GetMemoryDevice().MemoryHierarchy, accessType, threadId, virtualPc, physicalAddress, physicalTag, onCompletedCallback)

	var access = cacheController.FindAccess(physicalTag)

	if access != nil {
		access.Aliases = append([]*MemoryHierarchyAccess{newAccess}, access.Aliases...)
	} else {
		cacheController.PendingAccesses[physicalTag] = newAccess
		cacheController.NumPendingAccessesPerType[accessType]++
	}

	return newAccess
}

func (cacheController *CacheController) EndAccess(physicalTag uint32) {
	var access = cacheController.FindAccess(physicalTag)

	access.Complete()

	for _, alias := range access.Aliases {
		alias.Complete()
	}

	cacheController.NumPendingAccessesPerType[access.AccessType]--

	delete(cacheController.PendingAccesses, physicalTag)
}