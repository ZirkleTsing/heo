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
}

func NewMemoryHierarchy(experiment MemoryHierarchyDriver) *MemoryHierarchy {
	var memoryHierarchy = &MemoryHierarchy{
		Driver:experiment,
	}

	return memoryHierarchy
}
