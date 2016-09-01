package noc

import (
	"time"
	"os"
	"github.com/mcai/acogo/simutil"
)

type NoCExperiment struct {
	Config                  *NoCConfig
	Stats                   simutil.Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	CycleAccurateEventQueue *simutil.CycleAccurateEventQueue
	Network                 *Network
}

func NewNoCExperiment(config *NoCConfig) *NoCExperiment {
	var experiment = &NoCExperiment{
		Config:config,
		CycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
	}

	var network = NewNetwork(experiment)

	experiment.Network = network

	return experiment
}

func (experiment *NoCExperiment) Run(skipIfStatsFileExists bool) {
	if skipIfStatsFileExists {
		if _, err := os.Stat(experiment.Config.OutputDirectory + "/" + simutil.STATS_JSON_FILE_NAME); err == nil {
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

	experiment.Config.Dump(experiment.Config.OutputDirectory)

	experiment.DumpStats()
}
