package acogo

import (
	"time"
	"os"
	"fmt"
)

type Experiment struct {
	Config                  *NoCConfig
	Stats                   Stats
	statMap                 map[string]interface{}

	BeginTime, EndTime      time.Time
	CycleAccurateEventQueue *CycleAccurateEventQueue
	Network                 *Network
}

func NewExperiment(config *NoCConfig) *Experiment {
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
	var done = make(chan bool)

	for i, e := range experiments {
		go func(i int, experiment *Experiment, c chan bool) {
			var len = len(experiments)

			fmt.Printf("[%s] Experiment %d/%d started.\n",
				time.Now().Format("2006-01-02 15:04:05"), i + 1, len)

			experiment.Run(skipIfStatsFileExists)

			done <- true

			fmt.Printf("[%s] Experiment %d/%d ended.\n",
				time.Now().Format("2006-01-02 15:04:05"), i + 1, len)
		}(i, e, done)
	}

	for i := 0; i < len(experiments); i++ {
		<-done

		fmt.Printf("[%s] There are %d experiments to be run.\n",
			time.Now().Format("2006-01-02 15:04:05"), len(experiments) - i - 1)
	}
}
