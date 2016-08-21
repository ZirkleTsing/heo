package cpu

import "math"

type CacheGeometry struct {
	Size           uint64
	Associativity  uint32
	LineSize       uint64
	LineSizeInLog2 uint32
	NumSets        uint64
	NumSetsInLog2  uint32
	NumLines       uint64
}

func NewCacheGeometry(size uint64, associativity uint32, lineSize uint64) *CacheGeometry {
	var cacheGeometry = &CacheGeometry{
		Size:size,
		Associativity:associativity,
		LineSize:lineSize,
		LineSizeInLog2:uint32(math.Log2(float64(lineSize))),
		NumSets:size / uint64(associativity) / lineSize,
		NumSetsInLog2:uint32(math.Log2(float64(size / uint64(associativity) / lineSize))),
		NumLines:size / lineSize,
	}

	return cacheGeometry
}

func (cacheGeometry *CacheGeometry) GetDisplacement(address uint64) uint64 {
	return address & (uint64(cacheGeometry.LineSize) - 1)
}

func (cacheGeometry *CacheGeometry) GetTag(address uint64) uint64 {
	return address & ^(uint64(cacheGeometry.LineSize) - 1)
}

func (cacheGeometry *CacheGeometry) GetLineId(address uint64) uint64 {
	return address >> uint64(cacheGeometry.LineSizeInLog2)
}

func (cacheGeometry *CacheGeometry) GetSet(address uint64) uint64 {
	return cacheGeometry.GetLineId(address) % cacheGeometry.NumSets
}

func (cacheGeometry *CacheGeometry) IsAligned(address uint64) bool {
	return cacheGeometry.GetDisplacement(address) == 0
}


