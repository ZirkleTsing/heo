package cpu

import (
	"time"
	"github.com/mcai/acogo/simutil"
	"os"
)

type CPUExperiment struct {
	Config                  *CPUConfig
	Stats                   simutil.Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	CycleAccurateEventQueue *simutil.CycleAccurateEventQueue

	Kernel                  *Kernel
	Processor               *Processor
	MemoryHierarchy         *MemoryHierarchy
}

func NewCPUExperiment(config *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		Config:config,
		CycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
	}

	experiment.Kernel = NewKernel(experiment)
	experiment.Processor = NewProcessor(experiment)
	experiment.MemoryHierarchy = NewMemoryHierarchy(experiment)

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
	return experiment.Config.MaxInstructions == -1 ||
		experiment.Processor.Cores[0].Threads[0].NumInstructions < uint64(experiment.Config.MaxInstructions)
}
