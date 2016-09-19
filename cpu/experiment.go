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

	MemoryHierarchy         uncore.MemoryHierarchy
	OoO                     *OoO
}

func NewCPUExperiment(config *CPUConfig) *CPUExperiment {
	var experiment = &CPUExperiment{
		CPUConfig:config,
		cycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		blockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Processor = NewProcessor(experiment)
	experiment.Kernel = NewKernel(experiment)

	var uncoreConfig = uncore.NewUncoreConfig()
	uncoreConfig.NumCores = config.NumCores
	uncoreConfig.NumThreadsPerCore = config.NumThreadsPerCore

	var nocConfig = noc.NewNoCConfig(config.OutputDirectory, -1, -1, -1, false)

	experiment.MemoryHierarchy = uncore.NewBaseMemoryHierarchy(experiment, uncoreConfig, nocConfig)

	experiment.OoO = NewOoO(experiment)

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

	//TODO: to be called based on config
	//experiment.DoFastForward()
	//experiment.DoWarmup()
	experiment.DoMeasurement()

	experiment.EndTime = time.Now()

	experiment.CPUConfig.Dump(experiment.CPUConfig.OutputDirectory)

	experiment.MemoryHierarchy.Config().Dump(experiment.CPUConfig.OutputDirectory)

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

func (experiment *CPUExperiment) DoWarmup() {
	for len(experiment.Kernel.Contexts) > 0 && experiment.canAdvanceOneCycle() {
		for _, core := range experiment.Processor.Cores {
			core.WarmupOneCycle()
		}

		experiment.advanceOneCycle()
	}
}

func (experiment *CPUExperiment) DoMeasurement() {
	for len(experiment.Kernel.Contexts) > 0 && experiment.canAdvanceOneCycle() {
		for _, core := range experiment.Processor.Cores {
			core.(*OoOCore).MeasurementOneCycle()
		}

		experiment.advanceOneCycle()
	}
}

func (experiment *CPUExperiment) SimulationTime() time.Duration {
	return experiment.EndTime.Sub(experiment.BeginTime)
}

func (experiment *CPUExperiment) CyclesPerSecond() float64 {
	return float64(experiment.CycleAccurateEventQueue().CurrentCycle) / experiment.EndTime.Sub(experiment.BeginTime).Seconds()
}

func (experiment *CPUExperiment) InstructionsPerSecond() float64 {
	return float64(experiment.Processor.NumDynamicInsts()) / experiment.EndTime.Sub(experiment.BeginTime).Seconds()
}