package main

import "github.com/mcai/acogo/cpu"

func main() {
	var config = cpu.NewCPUConfig("test_results/real/mst_ht_100")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "100"))

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}
