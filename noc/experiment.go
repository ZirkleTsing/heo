package noc

import (
	"time"
	"os"
	"github.com/mcai/acogo/simutil"
)

type NoCExperiment struct {
	cycleAccurateEventQueue *simutil.CycleAccurateEventQueue
	blockingEventDispatcher *simutil.BlockingEventDispatcher

	Network                 *Network

	BeginTime, EndTime      time.Time

	Stats                   simutil.Stats
	statMap                 map[string]interface{}
}

func NewNoCExperiment(config *NoCConfig) *NoCExperiment {
	var experiment = &NoCExperiment{
		cycleAccurateEventQueue:simutil.NewCycleAccurateEventQueue(),
		blockingEventDispatcher:simutil.NewBlockingEventDispatcher(),
	}

	experiment.Network = NewNetwork(experiment, config)

	return experiment
}

func (experiment *NoCExperiment) CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue {
	return experiment.cycleAccurateEventQueue
}

func (experiment *NoCExperiment) BlockingEventDispatcher() *simutil.BlockingEventDispatcher {
	return experiment.blockingEventDispatcher
}

func (experiment *NoCExperiment) Run(skipIfStatsFileExists bool) {
	if skipIfStatsFileExists {
		if _, err := os.Stat(experiment.Network.Config.OutputDirectory + "/" + simutil.STATS_JSON_FILE_NAME); err == nil {
			return
		}
	}

	experiment.BeginTime = time.Now()

	for (experiment.Network.Config.MaxCycles == -1 || experiment.CycleAccurateEventQueue().CurrentCycle < experiment.Network.Config.MaxCycles) && (experiment.Network.Config.MaxPackets == -1 || experiment.Network.NumPacketsReceived < experiment.Network.Config.MaxPackets) {
		experiment.CycleAccurateEventQueue().AdvanceOneCycle()
	}

	if experiment.Network.Config.DrainPackets {
		experiment.Network.AcceptPacket = false

		for experiment.Network.NumPacketsReceived != experiment.Network.NumPacketsTransmitted {
			experiment.CycleAccurateEventQueue().AdvanceOneCycle()
		}
	}

	experiment.EndTime = time.Now()

	experiment.Network.Config.Dump(experiment.Network.Config.OutputDirectory)

	experiment.DumpStats()
}
