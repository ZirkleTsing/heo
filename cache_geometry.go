package acogo

import "math"

type CacheGeometry struct {
	Size           uint32
	Associativity  uint32
	LineSize       uint32
	LineSizeInLog2 uint32
	NumSets        uint32
	NumSetsInLog2  uint32
	NumLines       uint32
}

func NewCacheGeometry(size uint32, associativity uint32, lineSize uint32) *CacheGeometry {
	var cacheGeometry = &CacheGeometry{
		Size:size,
		Associativity:associativity,
		LineSize:lineSize,
		LineSizeInLog2:uint32(math.Log2(float64(lineSize))),
		NumSets:size / associativity / lineSize,
		NumSetsInLog2:uint32(math.Log2(float64(size / associativity / lineSize))),
		NumLines:size / lineSize,
	}

	return cacheGeometry
}

func (cacheGeometry *CacheGeometry) GetDisplacement(address uint32) uint32 {
	return address & (cacheGeometry.LineSize - 1)
}

func (cacheGeometry *CacheGeometry) GetTag(address uint32) uint32 {
	return address & ^(cacheGeometry.LineSize - 1)
}

func (cacheGeometry *CacheGeometry) GetLineId(address uint32) uint32 {
	return address >> uint(cacheGeometry.LineSizeInLog2)
}

func (cacheGeometry *CacheGeometry) GetSet(address uint32) uint32 {
	return cacheGeometry.GetLineId(address) % cacheGeometry.NumSets
}

func (cacheGeometry *CacheGeometry) IsAligned(address uint32) bool {
	return cacheGeometry.GetDisplacement(address) == 0
}


