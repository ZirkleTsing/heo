package acogo

func ContainsBit(value uint, i uint) bool {
	return (value & (1 << i)) != 0
}

func SetBit(value uint, i uint) uint {
	return value | (1 << i)
}

func ClearBit(value uint, i uint) uint {
	return value & ^(1 << i)
}

func Mask(numBits uint) uint32 {
	return (1 << numBits) - 1
}

func Bits(value uint32, first uint, last uint) uint32 {
	return (value >> last) & Mask(first - last + 1)
}

func Mbits(value uint32, first uint, last uint) uint32 {
	return value & (Mask(first + 1) & ^Mask(last))
}

func SignExtend(value uint32) int32 {
	return int32((value << 16) >> 16)
}

func ZeroExtend(value int32) uint32 {
	return uint32(value) & 0xffff
}
