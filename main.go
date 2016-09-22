package main

import (
	"github.com/mcai/acogo/cpu"
)

func main() {
	mstBaseline()
	mstHelperThreaded()
}

func mstBaseline() {
	var config = cpu.NewCPUConfig("test_results/real/mst_baseline_100")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/baseline/mst.mips", "100"))

	config.NumCores = 2
	config.NumThreadsPerCore = 2
	config.MaxDynamicInsts = 1
	config.FastForwardDynamicInsts = 100000000

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}

func mstHelperThreaded() {
	var config = cpu.NewCPUConfig("test_results/real/mst_ht_100")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "100"))

	config.NumCores = 2
	config.NumThreadsPerCore = 2
	config.MaxDynamicInsts = 1
	config.FastForwardDynamicInsts = 100000000

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}
