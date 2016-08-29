package uncore

type SimpleCacheLine struct {
	Cache *SimpleCache
	Set   uint32
	Way   uint32
	Value interface{}
}

func newSimpleCacheLine(cache *SimpleCache, set uint32, way uint32) *SimpleCacheLine {
	var cacheLine = &SimpleCacheLine{
		Cache:cache,
		Set:set,
		Way:way,
	}

	return cacheLine
}

type SimpleCacheSet struct {
	Cache *SimpleCache
	Lines []*SimpleCacheLine
	Num   uint32
}

func newSimpleCacheSet(cache *SimpleCache, assoc uint32, num uint32) *SimpleCacheSet {
	var cacheSet = &SimpleCacheSet{
		Cache:cache,
		Num:num,
	}

	for i := uint32(0); i < assoc; i++ {
		cacheSet.Lines = append(cacheSet.Lines, newSimpleCacheLine(cache, num, i))
	}

	return cacheSet
}

type SimpleCache struct {
	NumSets uint32
	Assoc   uint32
	Sets    []*SimpleCacheSet
}

func NewSimpleCache(numSets uint32, assoc uint32) *SimpleCache {
	var cache = &SimpleCache{
		NumSets:numSets,
		Assoc:assoc,
	}

	for i := uint32(0); i < numSets; i++ {
		cache.Sets = append(cache.Sets, newSimpleCacheSet(cache, assoc, i))
	}

	return cache
}

func NewSimpleCacheFromCache(cache *Cache) *SimpleCache {
	return NewSimpleCache(cache.NumSets(), cache.Assoc())
}
