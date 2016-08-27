package cpu

import (
	"time"
	"github.com/mcai/acogo/simutil"
	"os"
	"reflect"
	"fmt"
)

type CPUExperiment struct {
	Config                  *CPUConfig
	Stats                   simutil.Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	CycleAccurateEventQueue *simutil.CycleAccurateEventQueue
	BlockingEventDispatcher *simutil.BlockingEventDispatcher

	Kernel                  *Kernel
	Processor               *Processor
	MemoryHierarchy         *MemoryHierarchy
}

func NewCPUExperiment(config *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		Config:config,
		CycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		BlockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Kernel = NewKernel(experiment)
	experiment.Processor = NewProcessor(experiment)
	experiment.MemoryHierarchy = NewMemoryHierarchy(experiment)

	experiment.BlockingEventDispatcher.AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)

		fmt.Printf("[thread#%d] %s\n", staticInstExecutedEvent.Context.ThreadId, staticInstExecutedEvent.StaticInst.Disassemble(staticInstExecutedEvent.Pc))
	})

	return experiment
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

		experiment.CycleAccurateEventQueue.AdvanceOneCycle()
	}

	experiment.EndTime = time.Now()

	experiment.DumpConfig()

	experiment.DumpStats()
}

func (experiment *CPUExperiment) canAdvanceOneCycle() bool {
	return experiment.Config.MaxDynamicInsts == -1 ||
		experiment.Processor.Cores[0].Threads[0].NumDynamicInsts < experiment.Config.MaxDynamicInsts
}
