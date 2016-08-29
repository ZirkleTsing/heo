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
}

func NewMemoryHierarchy(driver MemoryHierarchyDriver, config *MemoryHierarchyConfig) *MemoryHierarchy {
	var memoryHierarchy = &MemoryHierarchy{
		Driver:driver,
		Config:config,
	}

	return memoryHierarchy
}
