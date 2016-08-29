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

type CacheLine struct {
	Cache        *Cache

	Set          uint32
	Way          uint32

	tag          int32

	Access       *MemoryHierarchyAccess

	InitialState interface{}
	State        interface{}
}

func newCacheLine(cache *Cache, set uint32, way uint32) *CacheLine {
	var cacheLine = &CacheLine{
		Cache:cache,
		Set:set,
		Way:way,
	}

	return cacheLine
}

func (cacheLine *CacheLine) Valid() bool {
	return cacheLine.State != cacheLine.InitialState
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
		cacheSet.Lines = append(cacheSet.Lines, newCacheLine(cache, num, i))
	}

	return cacheSet
}

type Cache struct {
	Geometry     *mem.Geometry
	Sets         []*CacheSet
	NumTagsInUse int32
}

func NewCache(geometry *mem.Geometry) *Cache {
	var cache = &Cache{
		Geometry:geometry,
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

func (cache *Cache) Traverse(callback func(set uint32, way uint32)) {
	for set := uint32(0); set < cache.NumSets(); set++ {
		for way := uint32(0); way < cache.Assoc(); way++ {
			callback(set, way)
		}
	}
}

func (cache *Cache) OccupancyRatio() float32 {
	return float32(cache.NumTagsInUse) / float32(cache.Geometry.NumLines)
}