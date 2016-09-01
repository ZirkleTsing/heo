package main

import "github.com/mcai/acogo/cpu"

func main() {
	var config = cpu.NewCPUConfig("test_results/real/mst_ht_1000")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "1000"))

	config.MaxDynamicInsts = 1000000

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}
