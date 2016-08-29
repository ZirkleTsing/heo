package uncore

type CacheReplacementPolicy interface {
	Cache() *Cache
	NewMiss(access *MemoryHierarchyAccess, set uint32, address uint32) *CacheAccess
	HandleReplacement(access *MemoryHierarchyAccess, set uint32, tag uint32) *CacheAccess
	HandlePromotionOnHit(access *MemoryHierarchyAccess, set uint32, way uint32)
	HandleInsertionOnMiss(access *MemoryHierarchyAccess, set uint32, way uint32)
}

type BaseCacheReplacementPolicy struct {
	cache *Cache
}

func NewBaseCacheReplacementPolicy(cache *Cache) *BaseCacheReplacementPolicy {
	var baseCacheReplacementPolicy = &BaseCacheReplacementPolicy{
		cache:cache,
	}

	return baseCacheReplacementPolicy
}

func (baseCacheReplacementPolicy *BaseCacheReplacementPolicy) Cache() *Cache {
	return baseCacheReplacementPolicy.cache
}

func NewMiss(baseCacheReplacementPolicy CacheReplacementPolicy, access *MemoryHierarchyAccess, set uint32, address uint32) *CacheAccess {
	var tag = baseCacheReplacementPolicy.Cache().GetTag(address)

	for way := uint32(0); way < baseCacheReplacementPolicy.Cache().Assoc(); way++ {
		var line = baseCacheReplacementPolicy.Cache().Sets[set].Lines[way]
		if !line.Valid() {
			return NewCacheAccess(baseCacheReplacementPolicy.Cache(), access, set, way, tag)
		}
	}

	return baseCacheReplacementPolicy.HandleReplacement(access, set, tag)
}