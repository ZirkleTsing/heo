package uncore

import "github.com/mcai/acogo/cpu/mem"

type DirectoryController struct {
	*BaseController
	CacheControllers         []*CacheController
	NumPendingMemoryAccesses uint32
}

func NewDirectoryController(memoryHierarchy *MemoryHierarchy, name string, geometry *mem.Geometry, replacementPolicyType CacheReplacementPolicyType) *DirectoryController {
	var directoryController = &DirectoryController{
		BaseController:NewBaseController(memoryHierarchy, name, MemoryDeviceType_L2_CONTROLLER, NewEvictableCache(geometry, replacementPolicyType)),
	}

	return directoryController
}