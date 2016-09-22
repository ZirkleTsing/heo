package main

import (
	"github.com/mcai/acogo/cpu"
)

var (
	numCores = int32(2)
	numThreadsPerCore = int32(2)

	maxDynamicInsts = int64(1)
	fastForwardDynamicInsts = int64(100000000)
)

func main() {
	mstBaseline()
	mstHelperThreaded()
}

func mstBaseline() {
	var config = cpu.NewCPUConfig("test_results/real/mst_baseline")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/baseline/mst.mips", "400"))

	config.NumCores = numCores
	config.NumThreadsPerCore = numThreadsPerCore
	config.MaxDynamicInsts = maxDynamicInsts
	config.FastForwardDynamicInsts = fastForwardDynamicInsts

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}

func mstHelperThreaded() {
	var config = cpu.NewCPUConfig("test_results/real/mst_ht")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "400"))

	config.NumCores = numCores
	config.NumThreadsPerCore = numThreadsPerCore
	config.MaxDynamicInsts = maxDynamicInsts
	config.FastForwardDynamicInsts = fastForwardDynamicInsts

	var experiment = cpu.NewCPUExperiment(config)

	experiment.Run(false)
}
