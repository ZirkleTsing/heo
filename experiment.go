package acogo

import (
	"math/rand"
	"fmt"
	"sort"
	"time"
)

type NoCExperiment struct {
	Config                  *NoCConfig
	Stats                   map[string]string

	BeginTime, EndTime      time.Time
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

	var network = NewNetwork(experiment)

	experiment.Network = network

	return experiment
}

func (experiment *NoCExperiment) Run() {
	fmt.Printf("[%d] Welcome to ACOGo simulator!\n", experiment.CycleAccurateEventQueue.CurrentCycle)

	// TODO: dump config

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

	experiment.Stats["SimulationTime"] = fmt.Sprintf("%v", experiment.EndTime.Sub(experiment.BeginTime))
	experiment.Stats["CyclesPerSecond"] = fmt.Sprintf("%f", float64(experiment.CycleAccurateEventQueue.CurrentCycle) / experiment.EndTime.Sub(experiment.BeginTime).Seconds())
	experiment.Stats["PacketsPerSecond"] = fmt.Sprintf("%f", float64(experiment.Network.NumPacketsTransmitted) / experiment.EndTime.Sub(experiment.BeginTime).Seconds())

	experiment.Stats["NumPacketsReceived"] = fmt.Sprintf("%d", experiment.Network.NumPacketsReceived)
	experiment.Stats["NumPacketsTransmitted"] = fmt.Sprintf("%d", experiment.Network.NumPacketsTransmitted)
	experiment.Stats["Throughput"] = fmt.Sprintf("%f", experiment.Network.Throughput())
	experiment.Stats["AveragePacketDelay"] = fmt.Sprintf("%f", experiment.Network.AveragePacketDelay())
	experiment.Stats["AveragePacketHops"] = fmt.Sprintf("%f", experiment.Network.AveragePacketHops())
	experiment.Stats["MaxPacketDelay"] = fmt.Sprintf("%d", experiment.Network.MaxPacketDelay)
	experiment.Stats["MaxPacketHops"] = fmt.Sprintf("%d", experiment.Network.MaxPacketHops)

	experiment.Stats["NumPayloadPacketsReceived"] = fmt.Sprintf("%d", experiment.Network.NumPayloadPacketsReceived)
	experiment.Stats["NumPayloadPacketsTransmitted"] = fmt.Sprintf("%d", experiment.Network.NumPayloadPacketsTransmitted)
	experiment.Stats["PayloadThroughput"] = fmt.Sprintf("%f", experiment.Network.PayloadThroughput())
	experiment.Stats["AveragePayloadPacketDelay"] = fmt.Sprintf("%f", experiment.Network.AveragePayloadPacketDelay())
	experiment.Stats["AveragePayloadPacketHops"] = fmt.Sprintf("%f", experiment.Network.AveragePayloadPacketHops())
	experiment.Stats["MaxPayloadPacketDelay"] = fmt.Sprintf("%d", experiment.Network.MaxPayloadPacketDelay)
	experiment.Stats["MaxPayloadPacketHops"] = fmt.Sprintf("%d", experiment.Network.MaxPayloadPacketHops)
}
