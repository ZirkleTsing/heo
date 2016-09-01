package cpu

import "github.com/mcai/acogo/cpu/uncore"

func (thread *Thread) Itlb() *uncore.TranslationLookasideBuffer {
	return thread.Core.Processor.Experiment.MemoryHierarchy.ITlbs[thread.Id]
}

func (thread *Thread) Dtlb() *uncore.TranslationLookasideBuffer {
	return thread.Core.Processor.Experiment.MemoryHierarchy.DTlbs[thread.Id]
}
