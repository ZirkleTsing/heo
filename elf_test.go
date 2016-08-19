package acogo

import (
	"testing"
	"fmt"
)

func TestElfFile(t *testing.T) {
	var elfFile = NewElfFile(
		"/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/baseline/mst.mips")

	fmt.Printf("Clz: %s, data: %s\n", elfFile.Identification.Clz, elfFile.Identification.Data)
}
