package acogo

import "math"

type CacheGeometry struct {
	Size int
	Associativity int
	LineSize int
	LineSizeInLog2 int
	NumSets int
	NumSetsInLog2 int
	NumLines int
}

func NewCacheGeometry(size int, associativity int, lineSize int) *CacheGeometry {
	var cacheGeometry = &CacheGeometry{
		Size:size,
		Associativity:associativity,
		LineSize:lineSize,
		LineSizeInLog2:int(math.Log2(float64(lineSize))),
		NumSets:size / associativity / lineSize,
		NumSetsInLog2:int(math.Log2(float64(size / associativity / lineSize))),
		NumLines:size / lineSize,
	}

	return cacheGeometry
}

func (cacheGeometry *CacheGeometry) GetDisplacement(address int) int {
	return address & (cacheGeometry.LineSize - 1)
}

func (cacheGeometry *CacheGeometry) GetTag(address int) int {
	return address & ^(cacheGeometry.LineSize - 1)
}

func (cacheGeometry *CacheGeometry) GetLineId(address int) int {
	return address >> uint(cacheGeometry.LineSizeInLog2)
}

func (cacheGeometry *CacheGeometry) GetSet(address int) int {
	return cacheGeometry.GetLineId(address) % cacheGeometry.NumSets
}

func (cacheGeometry *CacheGeometry) IsAligned(address int) bool {
	return cacheGeometry.GetDisplacement(address) == 0
}


