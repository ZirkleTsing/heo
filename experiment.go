package acogo

import (
	"time"
	"os"
)

type Experiment struct {
	Config                  *Config
	Stats                   Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	CycleAccurateEventQueue *CycleAccurateEventQueue
	Network                 *Network
}

func NewExperiment(config *Config) *Experiment {
	var experiment = &Experiment{
		Config:config,
		CycleAccurateEventQueue:NewCycleAccurateEventQueue(),
	}

	var network = NewNetwork(experiment)

	experiment.Network = network

	return experiment
}

func (experiment *Experiment) Run(skipIfStatsFileExists bool) {
	if skipIfStatsFileExists {
		if _, err := os.Stat(experiment.Config.OutputDirectory + "/" + STATS_JSON_FILE_NAME); err == nil {
			return
		}
	}

	experiment.BeginTime = time.Now()

	for (experiment.Config.MaxCycles == -1 || experiment.CycleAccurateEventQueue.CurrentCycle < experiment.Config.MaxCycles) && (experiment.Config.MaxPackets == -1 || experiment.Network.NumPacketsReceived < experiment.Config.MaxPackets) {
		experiment.CycleAccurateEventQueue.AdvanceOneCycle()
	}

	if experiment.Config.DrainPackets {
		experiment.Network.AcceptPacket = false

		for experiment.Network.NumPacketsReceived != experiment.Network.NumPacketsTransmitted {
			experiment.CycleAccurateEventQueue.AdvanceOneCycle()
		}
	}

	experiment.EndTime = time.Now()

	experiment.DumpConfig()

	experiment.DumpStats()
}

func RunExperiments(experiments []*Experiment, skipIfStatsFileExists bool) {
	var tasks []func()

	for _, experiment := range experiments {
		var e = experiment

		tasks = append(tasks, func() {
			e.Run(skipIfStatsFileExists)
		})
	}

	RunInParallel(tasks)
}

func AnalyzeExperiments(outputDirectory string, outputCSVFileName string, experiments []*Experiment) {
	for _, experiment := range experiments {
		experiment.LoadStats()
	}

	WriteCSVFile(outputDirectory, outputCSVFileName, experiments, GetCSVFields())
}
