package cpu

import (
	"time"
	"github.com/mcai/acogo/simutil"
	"os"
	"github.com/mcai/acogo/cpu/uncore"
	"github.com/mcai/acogo/noc"
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
	MemoryHierarchy         *uncore.NoCMemoryHierarchy
}

func NewCPUExperiment(cpuConfig *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		CPUConfig:cpuConfig,
		cycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		blockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Processor = NewProcessor(experiment)
	experiment.Kernel = NewKernel(experiment)

	var uncoreConfig = uncore.NewUncoreConfig()
	uncoreConfig.NumCores = cpuConfig.NumCores
	uncoreConfig.NumThreadsPerCore = cpuConfig.NumThreadsPerCore

	var nocConfig = noc.NewNoCConfig(cpuConfig.OutputDirectory, -1, -1, -1, false)

	experiment.MemoryHierarchy = uncore.NewNocMemoryHierarchy(experiment, uncoreConfig, nocConfig)

	experiment.Processor.UpdateContextToThreadAssignments()

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

	experiment.DoFastForward()

	experiment.EndTime = time.Now()

	experiment.CPUConfig.Dump(experiment.CPUConfig.OutputDirectory)

	experiment.DumpStats()
}

func (experiment *CPUExperiment) canAdvanceOneCycle() bool {
	return experiment.CPUConfig.MaxDynamicInsts == -1 ||
		experiment.Processor.Cores[0].Threads()[0].NumDynamicInsts() < experiment.CPUConfig.MaxDynamicInsts
}

func (experiment *CPUExperiment) advanceOneCycle() {
	experiment.Kernel.AdvanceOneCycle()
	experiment.Processor.UpdateContextToThreadAssignments()

	experiment.cycleAccurateEventQueue.AdvanceOneCycle()
}

func (experiment *CPUExperiment) DoFastForward() {
	for len(experiment.Kernel.Contexts) > 0 && experiment.canAdvanceOneCycle() {
		for _, core := range experiment.Processor.Cores {
			core.FastForwardOneCycle()
		}

		experiment.advanceOneCycle()
	}
}

//TODO: to be called based on config
func (experiment *CPUExperiment) DoWarmup() {
	for len(experiment.Kernel.Contexts) > 0 && experiment.canAdvanceOneCycle() {
		for _, core := range experiment.Processor.Cores {
			core.WarmupOneCycle()
		}

		experiment.advanceOneCycle()
	}
}
