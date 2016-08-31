package uncore

import "github.com/mcai/acogo/cpu/mem"

type DirectoryController struct {
	*BaseController
	CacheControllers         []*CacheController
	NumPendingMemoryAccesses uint32
	FsmFactory *DirectoryControllerFiniteStateMachineFactory
}

func NewDirectoryController(memoryHierarchy *MemoryHierarchy, name string, geometry *mem.Geometry, replacementPolicyType CacheReplacementPolicyType) *DirectoryController {
	var directoryController = &DirectoryController{
	}

	directoryController.BaseController = NewBaseController(
		memoryHierarchy,
		name,
		MemoryDeviceType_L2_CONTROLLER,
		NewEvictableCache(
			geometry,
			func(set uint32, way uint32) CacheLineStateProvider {
				return NewDirectoryControllerFiniteStateMachine(set, way, directoryController)
			},
			replacementPolicyType,
		),
	)

	directoryController.FsmFactory = NewDirectoryControllerFiniteStateMachineFactory()

	return directoryController
}

func (directoryController *DirectoryController) access(producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester CacheController, onReplacementCompletedCallback func(set uint32, way uint32), onReplacementStalledCallback func()) {
	//var set = directoryController.Cache().GetSet(tag)
	//
	//for _, line := range directoryController.Cache().Sets[set].Lines {
	//	var directoryControllerFsm = line.StateProvider.(*DirectoryControllerFiniteStateMachine)
	//	//TODO
	//}
}

func (directoryController *DirectoryController) SendPutAckToRequester(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	directoryController.TransferMessage(requester, 8, NewPutAckMessage(directoryController, producerFlow, producerFlow.Access(), tag))
}