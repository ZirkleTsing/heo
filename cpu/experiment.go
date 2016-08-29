package cpu

import (
	"time"
	"github.com/mcai/acogo/simutil"
	"os"
	"reflect"
	"github.com/mcai/acogo/cpu/uncore"
)

type CPUExperiment struct {
	Config                  *CPUConfig
	Stats                   simutil.Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	cycleAccurateEventQueue *simutil.CycleAccurateEventQueue
	blockingEventDispatcher *simutil.BlockingEventDispatcher

	Kernel                  *Kernel
	Processor               *Processor
	MemoryHierarchy         *uncore.MemoryHierarchy
}

func NewCPUExperiment(config *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		Config:config,
		cycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		blockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Kernel = NewKernel(experiment)
	experiment.Processor = NewProcessor(experiment)

	var memoryHierarchyConfig = uncore.NewMemoryHierarchyConfig() //TODO: to be passed here externally

	experiment.MemoryHierarchy = uncore.NewMemoryHierarchy(experiment, memoryHierarchyConfig)

	experiment.blockingEventDispatcher.AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		//var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)
		//fmt.Printf("[thread#%d] %s\n", staticInstExecutedEvent.Context.ThreadId, staticInstExecutedEvent.StaticInst.Disassemble(staticInstExecutedEvent.Pc))
		//fmt.Printf("#dynamicInsts: %d\n", experiment.Processor.Cores[0].Threads[0].NumDynamicInsts)
		//fmt.Println(staticInstExecutedEvent.Context.Regs.Dump())
	})

	return experiment
}

func (experiment *CPUExperiment) CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue {
	return experiment.cycleAccurateEventQueue
}

func (experiment *CPUExperiment) BlockingEventDispatcher() *simutil.BlockingEventDispatcher {
	return experiment.blockingEventDispatcher
}

func (experiment *CPUExperiment) Run(skipIfStatsFileExists bool) {
	if skipIfStatsFileExists {
		if _, err := os.Stat(experiment.Config.OutputDirectory + "/" + simutil.STATS_JSON_FILE_NAME); err == nil {
			return
		}
	}

	experiment.BeginTime = time.Now()

	for len(experiment.Kernel.Contexts) > 0 && experiment.canAdvanceOneCycle() {
		for _, core := range experiment.Processor.Cores {
			core.AdvanceOneCycle()
		}

		experiment.Kernel.AdvanceOneCycle()
		experiment.Processor.UpdateContextToThreadAssignments()

		experiment.cycleAccurateEventQueue.AdvanceOneCycle()
	}

	experiment.EndTime = time.Now()

	experiment.DumpConfig()

	experiment.DumpStats()
}

func (experiment *CPUExperiment) canAdvanceOneCycle() bool {
	return experiment.Config.MaxDynamicInsts == -1 ||
		experiment.Processor.Cores[0].Threads[0].NumDynamicInsts < experiment.Config.MaxDynamicInsts
}
