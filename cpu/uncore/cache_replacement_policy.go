package uncore

type CacheReplacementPolicy interface {
	Cache() *EvictableCache
	HandleReplacement(access *MemoryHierarchyAccess, set uint32, tag uint32) *CacheAccess
	HandlePromotionOnHit(access *MemoryHierarchyAccess, set uint32, way uint32)
	HandleInsertionOnMiss(access *MemoryHierarchyAccess, set uint32, way uint32)
}

type BaseCacheReplacementPolicy struct {
	cache *EvictableCache
}

func NewBaseCacheReplacementPolicy(cache *EvictableCache) *BaseCacheReplacementPolicy {
	var baseCacheReplacementPolicy = &BaseCacheReplacementPolicy{
		cache:cache,
	}

	return baseCacheReplacementPolicy
}

func (baseCacheReplacementPolicy *BaseCacheReplacementPolicy) Cache() *EvictableCache {
	return baseCacheReplacementPolicy.cache
}

func NewMiss(cacheReplacementPolicy CacheReplacementPolicy, access *MemoryHierarchyAccess, set uint32, address uint32) *CacheAccess {
	var tag = cacheReplacementPolicy.Cache().GetTag(address)

	for way := uint32(0); way < cacheReplacementPolicy.Cache().Assoc(); way++ {
		var line = cacheReplacementPolicy.Cache().Sets[set].Lines[way]
		if !line.Valid() {
			return NewCacheAccess(cacheReplacementPolicy.Cache(), access, set, way, tag)
		}
	}

	return cacheReplacementPolicy.HandleReplacement(access, set, tag)
}