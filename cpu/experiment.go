package cpu

import (
	"time"
	"github.com/mcai/acogo/simutil"
	"os"
	"reflect"
	"github.com/mcai/acogo/cpu/uncore"
)

type CPUExperiment struct {
	CPUConfig               *CPUConfig
	Stats                   simutil.Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	cycleAccurateEventQueue *simutil.CycleAccurateEventQueue
	blockingEventDispatcher *simutil.BlockingEventDispatcher

	Kernel                  *Kernel
	Processor               *Processor
	MemoryHierarchy         *uncore.MemoryHierarchy
}

func NewCPUExperiment(cpuConfig *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		CPUConfig:cpuConfig,
		cycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		blockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Kernel = NewKernel(experiment)
	experiment.Processor = NewProcessor(experiment)

	var uncoreConfig = uncore.NewUncoreConfig()
	uncoreConfig.NumCores = cpuConfig.NumCores
	uncoreConfig.NumThreadsPerCore = cpuConfig.NumThreadsPerCore

	experiment.MemoryHierarchy = uncore.NewMemoryHierarchy(experiment, uncoreConfig)

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
		if _, err := os.Stat(experiment.CPUConfig.OutputDirectory + "/" + simutil.STATS_JSON_FILE_NAME); err == nil {
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

	experiment.CPUConfig.Dump(experiment.CPUConfig.OutputDirectory)

	experiment.DumpStats()
}

func (experiment *CPUExperiment) canAdvanceOneCycle() bool {
	return experiment.CPUConfig.MaxDynamicInsts == -1 ||
		experiment.Processor.Cores[0].Threads[0].NumDynamicInsts < experiment.CPUConfig.MaxDynamicInsts
}
