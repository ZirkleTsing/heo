package cpu

import (
	"testing"
	"reflect"
	"fmt"
)

func TestCPUExperiment(t *testing.T) {
	var config = NewCPUConfig("test_results/real/mst_ht_100")

	config.ContextMappings = append(config.ContextMappings,
		NewContextMapping(0, "/home/itecgo/Projects/Archimulator/benchmarks/Olden_Custom1/mst/ht/mst.mips", "100"))

	config.MaxDynamicInsts = 1000000

	var experiment = NewCPUExperiment(config)

	experiment.blockingEventDispatcher.AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)
		fmt.Printf("[thread#%d] %s\n", staticInstExecutedEvent.Context.ThreadId, staticInstExecutedEvent.StaticInst.Disassemble(staticInstExecutedEvent.Pc))
		//fmt.Printf("#dynamicInsts: %d\n", experiment.Processor.Cores[0].Threads[0].NumDynamicInsts)
		//fmt.Println(staticInstExecutedEvent.Context.Regs.Dump())
	})

	experiment.Run(false)
}