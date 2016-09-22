package main

import (
	"github.com/mcai/acogo/cpu"
)

func main() {
	var config = cpu.NewCPUConfig("test_results/real/mst_ht_100")

	config.ContextMappings = append(config.ContextMappings,
		cpu.NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "100"))

	config.NumCores = 2
	config.NumThreadsPerCore = 2
	config.MaxDynamicInsts = -1
	config.FastForwardDynamicInsts = 1000

	var experiment = cpu.NewCPUExperiment(config)

	//experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*cpu.StaticInstExecutedEvent)(nil)), func(event interface{}) {
	//	var staticInstExecutedEvent = event.(*cpu.StaticInstExecutedEvent)
	//	fmt.Printf("[thread#%d] %s\n", staticInstExecutedEvent.Context.ThreadId, staticInstExecutedEvent.StaticInst.Disassemble(staticInstExecutedEvent.Pc))
	//	fmt.Printf("#dynamicInsts: %d\n", experiment.Processor.Cores[0].Threads()[0].NumDynamicInsts)
	//	fmt.Println(staticInstExecutedEvent.Context.Regs.Dump())
	//})

	experiment.Run(false)
}
