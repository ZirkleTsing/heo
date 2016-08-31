package uncore

import (
	"github.com/mcai/acogo/simutil"
)

type MemoryHierarchyDriver interface {
	CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue
	BlockingEventDispatcher() *simutil.BlockingEventDispatcher
}

type MemoryHierarchy struct {
	Driver MemoryHierarchyDriver
	Config *MemoryHierarchyConfig

	CurrentMemoryHierarchyAccessId int32
	CurrentCacheCoherenceFlowId int32

	PendingFlows []CacheCoherenceFlow
}

func NewMemoryHierarchy(driver MemoryHierarchyDriver, config *MemoryHierarchyConfig) *MemoryHierarchy {
	var memoryHierarchy = &MemoryHierarchy{
		Driver:driver,
		Config:config,
	}

	return memoryHierarchy
}

func (memoryHierarchy *MemoryHierarchy) Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func()) {
	panic("Unimplemented") //TODO
}

func (memoryHierarchy *MemoryHierarchy) TransferMessage(from Controller, to Controller, size uint32, message CoherenceMessage) {
	panic("Unimplemented") //TODO
}
