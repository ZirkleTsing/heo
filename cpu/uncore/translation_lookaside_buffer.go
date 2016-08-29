package uncore

import "github.com/mcai/acogo/cpu/mem"

type TranslationLookasideBuffer struct {
	MemoryHierarchy *MemoryHierarchy
	Name            string
	Cache           *EvictableCache
	NumHits         int32
	NumMisses       int32
	NumEvictions    int32
}

func NewTranslationLookasideBuffer(memoryHierarchy *MemoryHierarchy, name string) *TranslationLookasideBuffer {
	var tlb = &TranslationLookasideBuffer{
		MemoryHierarchy:memoryHierarchy,
		Name:name,
		Cache:NewEvictableCache(
			mem.NewGeometry(
				memoryHierarchy.Config.TlbSize,
				memoryHierarchy.Config.TlbAssoc,
				memoryHierarchy.Config.TlbLineSize,
			),
			CacheReplacementPolicyType_LRU,
		),
	}

	tlb.Cache.Traverse(func(set uint32, way uint32) {
		tlb.Cache.Sets[set].Lines[way].InitialState = false
		tlb.Cache.Sets[set].Lines[way].State = false
	})

	return tlb
}

func (tlb *TranslationLookasideBuffer) NumAccesses() int32 {
	return tlb.NumHits + tlb.NumMisses
}

func (tlb *TranslationLookasideBuffer) HitRatio() float32 {
	if tlb.NumAccesses() == 0 {
		return 0
	} else {
		return float32(tlb.NumHits) / float32(tlb.NumAccesses())
	}
}

func (tlb *TranslationLookasideBuffer) OccupancyRatio() float32 {
	return tlb.Cache.OccupancyRatio()
}

func (tlb *TranslationLookasideBuffer) HitLatency() uint32 {
	return tlb.MemoryHierarchy.Config.TlbHitLatency
}

func (tlb *TranslationLookasideBuffer) MissLatency() uint32 {
	return tlb.MemoryHierarchy.Config.TlbMissLatency
}

func (tlb *TranslationLookasideBuffer) Access(access *MemoryHierarchyAccess, onCompletedCallback func()) {
	var set = tlb.Cache.GetSet(access.PhysicalAddress)
	var cacheAccess = tlb.Cache.NewAccess(access, access.PhysicalAddress)

	if cacheAccess.HitInCache {
		tlb.Cache.ReplacementPolicy.HandlePromotionOnHit(access, set, cacheAccess.Way)

		tlb.NumHits++
	} else {
		if cacheAccess.Replacement {
			tlb.NumEvictions++
		}

		var line = tlb.Cache.Sets[set].Lines[cacheAccess.Way]
		line.State = true
		line.Access = access
		line.SetTag(int32(access.PhysicalTag))
		tlb.Cache.ReplacementPolicy.HandleInsertionOnMiss(access, set, cacheAccess.Way)

		tlb.NumMisses++
	}

	var delay uint32

	if cacheAccess.HitInCache {
		delay = tlb.HitLatency()
	} else {
		delay = tlb.MissLatency()
	}

	tlb.MemoryHierarchy.Driver.CycleAccurateEventQueue().Schedule(onCompletedCallback, int(delay))
}
