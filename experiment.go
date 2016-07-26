package acogo

import (
	"time"
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

func (experiment *Experiment) Run() {
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

func RunExperiments(experiments []*Experiment) {
	var tasks []func()

	for _, experiment := range experiments {
		tasks = append(tasks, experiment.Run)
	}

	RunInParallel(tasks)
}
