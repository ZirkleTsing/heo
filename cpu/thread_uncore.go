package cpu

import "github.com/mcai/acogo/cpu/uncore"

type MemoryHierarchyThread struct {
	*BaseThread
	FetchStalled bool
	LastFetchedCacheLine int32
	NextDynamicInstInWarmup *DynamicInst
}

func NewMemoryHierarchyThread(core Core, num int32) *MemoryHierarchyThread {
	var memoryHierarchyThread = &MemoryHierarchyThread{
		BaseThread:NewBaseThread(core, num),
		LastFetchedCacheLine:-1,
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
	if thread.Context() != nil && thread.Context().State == ContextState_RUNNING && !thread.FetchStalled {
		if thread.NextDynamicInstInWarmup == nil {
			var staticInst *StaticInst

			for {
				staticInst = thread.Context().DecodeNextStaticInst()

				thread.NextDynamicInstInWarmup = NewDynamicInst(
					thread,
					thread.Context().Regs.Pc,
					staticInst,
				)

				staticInst.Execute(thread.Context())

				if staticInst.Mnemonic.StaticInstType != StaticInstType_NOP {
					thread.numDynamicInsts++
				}

				if !(thread.Context() != nil &&
					thread.Context().State == ContextState_RUNNING &&
					!thread.FetchStalled &&
					staticInst.Mnemonic.Name == Mnemonic_NOP) {
					break
				}
			}
		}

		var pc = thread.NextDynamicInstInWarmup.Pc

		var cacheLineToFetch = thread.Core().L1IController().Cache.GetTag(pc)

		if int32(cacheLineToFetch) != thread.LastFetchedCacheLine {
			if thread.Core().CanIfetch(thread, pc) {
				thread.Core().Ifetch(thread, pc, pc, func() {
					thread.FetchStalled = false
				})
			}

			thread.FetchStalled = true
			thread.LastFetchedCacheLine = int32(cacheLineToFetch)
		} else {
			return
		}

		var effectiveAddress = thread.NextDynamicInstInWarmup.EffectiveAddress

		if thread.NextDynamicInstInWarmup.StaticInst.Mnemonic.StaticInstType == StaticInstType_LOAD {
			if thread.Core().CanLoad(thread, uint32(effectiveAddress)) {
				thread.Core().Load(thread, uint32(effectiveAddress), pc, func(){
				})

				thread.NextDynamicInstInWarmup = nil
			}
		} else if thread.NextDynamicInstInWarmup.StaticInst.Mnemonic.StaticInstType == StaticInstType_STORE {
			if thread.Core().CanStore(thread, uint32(effectiveAddress)) {
				thread.Core().Store(thread, uint32(effectiveAddress), pc, func() {
				})

				thread.NextDynamicInstInWarmup = nil
			}
		} else {
			thread.NextDynamicInstInWarmup = nil
		}
	}
}