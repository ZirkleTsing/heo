package cpu

import "github.com/mcai/acogo/cpu/uncore"

type MemoryHierarchyThread struct {
	*BaseThread
	LineSizeOfICache uint32
	FetchStalled bool
	LastFetchedCacheLine uint32
}

func NewMemoryHierarchyThread(core Core, num int32) *MemoryHierarchyThread {
	var memoryHierarchyThread = &MemoryHierarchyThread{
		BaseThread:NewBaseThread(core, num),
	}

	return memoryHierarchyThread
}

func (thread *MemoryHierarchyThread) Itlb() *uncore.TranslationLookasideBuffer {
	return thread.Core().Processor().Experiment.MemoryHierarchy.ITlbs[thread.Id()]
}

func (thread *MemoryHierarchyThread) Dtlb() *uncore.TranslationLookasideBuffer {
	return thread.Core().Processor().Experiment.MemoryHierarchy.DTlbs[thread.Id()]
}

func (thread *MemoryHierarchyThread) WarmupOneCycle() {
	panic("unimplemented") //TODO
}