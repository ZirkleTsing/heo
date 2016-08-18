package acogo

const (
	CLOCKS_PER_SEC = 1000000
	CPU_FREQUENCY = 300000
)

func Clock(numCycles uint64) uint64 {
	return CLOCKS_PER_SEC * numCycles / CPU_FREQUENCY
}
