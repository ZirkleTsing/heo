package mem

import "math"

type Geometry struct {
	Size           uint64
	Associativity  uint32
	LineSize       uint64
	LineSizeInLog2 uint32
	NumSets        uint64
	NumSetsInLog2  uint32
	NumLines       uint64
}

func NewGeometry(size uint64, associativity uint32, lineSize uint64) *Geometry {
	var geometry = &Geometry{
		Size:size,
		Associativity:associativity,
		LineSize:lineSize,
		LineSizeInLog2:uint32(math.Log2(float64(lineSize))),
		NumSets:size / uint64(associativity) / lineSize,
		NumSetsInLog2:uint32(math.Log2(float64(size / uint64(associativity) / lineSize))),
		NumLines:size / lineSize,
	}

	return geometry
}

func (geometry *Geometry) GetDisplacement(address uint64) uint64 {
	return address & (uint64(geometry.LineSize) - 1)
}

func (geometry *Geometry) GetTag(address uint64) uint64 {
	return address & ^(uint64(geometry.LineSize) - 1)
}

func (geometry *Geometry) GetLineId(address uint64) uint64 {
	return address >> uint64(geometry.LineSizeInLog2)
}

func (geometry *Geometry) GetSet(address uint64) uint64 {
	return geometry.GetLineId(address) % geometry.NumSets
}

func (geometry *Geometry) IsAligned(address uint64) bool {
	return geometry.GetDisplacement(address) == 0
}


