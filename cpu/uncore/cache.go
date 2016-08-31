package uncore

import (
	"github.com/mcai/acogo/cpu/mem"
)

const (
	INVALID_TAG = -1
)

const (
	INVALID_WAY = -1
)

type CacheLineStateProvider interface {
	InitialState() interface{}
	State() interface{}
}

type BaseCacheLineStateProvider struct {
	initialState interface{}
	state        interface{}
}

func NewBaseCacheLineStateProvider(initialState bool) *BaseCacheLineStateProvider {
	var boolCacheLineStateProvider = &BaseCacheLineStateProvider{
		initialState:initialState,
		state:initialState,
	}

	return boolCacheLineStateProvider
}

func (baseCacheLineStateProvider *BaseCacheLineStateProvider) InitialState() interface{} {
	return baseCacheLineStateProvider.initialState
}

func (baseCacheLineStateProvider *BaseCacheLineStateProvider) State() interface{} {
	return baseCacheLineStateProvider.state
}

func (baseCacheLineStateProvider *BaseCacheLineStateProvider) SetState(state interface{}) {
	baseCacheLineStateProvider.state = state
}

type CacheLine struct {
	Cache         *Cache

	Set           uint32
	Way           uint32

	tag           int32

	Access        *MemoryHierarchyAccess

	StateProvider CacheLineStateProvider
}

func newCacheLine(cache *Cache, set uint32, way uint32, stateProvider CacheLineStateProvider) *CacheLine {
	var cacheLine = &CacheLine{
		Cache:cache,
		Set:set,
		Way:way,
		StateProvider:stateProvider,
	}

	return cacheLine
}

func (cacheLine *CacheLine) Valid() bool {
	return cacheLine.State() != cacheLine.InitialState()
}

func (cacheLine *CacheLine) Tag() int32 {
	return cacheLine.tag
}

func (cacheLine *CacheLine) SetTag(tag int32) {
	if tag != INVALID_TAG {
		for _, line := range cacheLine.Cache.Sets[cacheLine.Set].Lines {
			if line.Tag() == tag {
				panic("Impossible")
			}
		}
	}

	if cacheLine.Tag() == INVALID_TAG && tag != INVALID_TAG {
		cacheLine.Cache.NumTagsInUse++
	} else if cacheLine.tag != INVALID_TAG && tag == INVALID_TAG {
		cacheLine.Cache.NumTagsInUse--
	}

	cacheLine.tag = tag
}

func (cacheLine *CacheLine) InitialState() interface{} {
	return cacheLine.StateProvider.InitialState()
}

func (cacheLine *CacheLine) State() interface{} {
	return cacheLine.StateProvider.State()
}

type CacheSet struct {
	Cache *Cache
	Lines []*CacheLine
	Num   uint32
}

func newCacheSet(cache *Cache, assoc uint32, num uint32) *CacheSet {
	var cacheSet = &CacheSet{
		Cache:cache,
		Num:num,
	}

	for i := uint32(0); i < assoc; i++ {
		cacheSet.Lines = append(cacheSet.Lines,
			newCacheLine(
				cache,
				num,
				i,
				cache.LineStateProviderFactory(num, i),
			),
		)
	}

	return cacheSet
}

type Cache struct {
	Geometry                 *mem.Geometry
	Sets                     []*CacheSet
	NumTagsInUse             int32
	LineStateProviderFactory func(set uint32, way uint32) CacheLineStateProvider
}

func NewCache(geometry *mem.Geometry, lineStateProviderFactory func(set uint32, way uint32) CacheLineStateProvider) *Cache {
	var cache = &Cache{
		Geometry:geometry,
		LineStateProviderFactory:lineStateProviderFactory,
	}

	for i := uint32(0); i < geometry.NumSets; i++ {
		cache.Sets = append(cache.Sets, newCacheSet(cache, geometry.Assoc, i))
	}

	return cache
}

func (cache *Cache) GetTag(address uint32) uint32 {
	return cache.Geometry.GetTag(address)
}

func (cache *Cache) GetSet(address uint32) uint32 {
	return cache.Geometry.GetSet(address)
}

func (cache *Cache) NumSets() uint32 {
	return cache.Geometry.NumSets
}

func (cache *Cache) Assoc() uint32 {
	return cache.Geometry.Assoc
}

func (cache *Cache) LineSize() uint32 {
	return cache.Geometry.LineSize
}

func (cache *Cache) FindWay(address uint32) int32 {
	var tag = cache.GetTag(address)
	var set = cache.GetSet(address)

	for _, line := range cache.Sets[set].Lines {
		if line.Valid() && line.Tag() == int32(tag) {
			return int32(line.Way)
		}
	}

	return INVALID_WAY
}

func (cache *Cache) FindLine(address uint32) *CacheLine {
	var set = cache.GetSet(address)
	var way = cache.FindWay(address)

	if way != INVALID_WAY {
		return cache.Sets[set].Lines[way]
	} else {
		return nil
	}
}

func (cache *Cache) OccupancyRatio() float32 {
	return float32(cache.NumTagsInUse) / float32(cache.Geometry.NumLines)
}