package acogo

import (
	"math/rand"
	"fmt"
	"sort"
)

type NoCExperiment struct {
	Config                  *NoCConfig
	Stats                   map[string]string

	BeginTime, EndTime      int64 //TODO
	CycleAccurateEventQueue *CycleAccurateEventQueue
	Network                 *Network
	Rand                    *rand.Rand
}

func NewNoCExperiment(config *NoCConfig) *NoCExperiment {
	var experiment = &NoCExperiment{
		Config:config,
		Stats: make(map[string]string),
		CycleAccurateEventQueue:NewCycleAccurateEventQueue(),
		Rand: rand.New(rand.NewSource(config.RandSeed)),
	}

	var network = NewNetwork(experiment, config.NumNodes)

	experiment.Network = network

	var _ = NewTransposeTrafficGenerator(network, config.DataPacketInjectionRate, config.MaxPackets, func(src int, dest int) Packet {
		return NewDataPacket(network, src, dest, config.DataPacketSize, true, func() {})
	})

	return experiment
}

func (experiment *NoCExperiment) Run() {
	fmt.Printf("[%d] Welcome to ACOGo simulator!\n", experiment.CycleAccurateEventQueue.CurrentCycle)

	// TODO: dump config

	for (experiment.Config.MaxCycles == -1 || experiment.CycleAccurateEventQueue.CurrentCycle < experiment.Config.MaxCycles) && (experiment.Config.MaxPackets == -1 || experiment.Network.NumPacketsReceived < experiment.Config.MaxPackets) {
		experiment.CycleAccurateEventQueue.AdvanceOneCycle()
	}

	if experiment.Config.DrainPackets {
		experiment.Network.AcceptPacket = false

		for experiment.Network.NumPacketsReceived != experiment.Network.NumPacketsTransmitted {
			experiment.CycleAccurateEventQueue.AdvanceOneCycle()
		}
	}

	experiment.CollectStats()

	fmt.Printf("[%d] Simulation ended!\n", experiment.CycleAccurateEventQueue.CurrentCycle)

	var keys []string

	for k, _ := range experiment.Stats {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %s\n", k, experiment.Stats[k])
	}
}

func (experiment *NoCExperiment) CollectStats() {
	experiment.Stats["TotalCycles"] = fmt.Sprintf("%d", experiment.CycleAccurateEventQueue.CurrentCycle)
	experiment.Stats["NumPacketsReceived"] = fmt.Sprintf("%d", experiment.Network.NumPacketsReceived)
	experiment.Stats["NumPacketsTransmitted"] = fmt.Sprintf("%d", experiment.Network.NumPacketsTransmitted)
	//TODO
}
