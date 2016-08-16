package acogo

type half uint16
type word uint32
type dword uint64

type sbyte int8
type shalf int16
type sword int32
type sdword int64

func Sext32(x uint32, b uint32) int32 {
	if uint32(x) & (uint32(1) << (b - 1)) != 0 {
		return int32(uint32(x) | ^((uint32(1) << b) - 1))
	} else {
		return int32(x)
	}
}

func Bits32(x uint32, hi uint32, lo uint32) uint32 {
	return (x >> lo) & ((uint32(1) << (hi - lo + 1)) - 1)
}

func Bits64(x uint64, hi uint64, lo uint64) uint64 {
	return (x >> lo) & ((uint64(1) << (hi - lo + 1)) - 1)
}
