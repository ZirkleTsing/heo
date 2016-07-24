package acogo

import "fmt"
import "math/rand"

type NoCExperiment struct {
	Config                  *NoCConfig
	Stats                   map[string]string

	BeginTime, EndTime      int
	CycleAccurateEventQueue *CycleAccurateEventQueue
	Network                 *Network
	rand                    *rand.Rand
}

func NewNoCExperiment(config *NoCConfig) *NoCExperiment {
	var cycleAccurateEventQueue = NewCycleAccurateEventQueue()

	var rand = rand.New(rand.NewSource(int64(config.RandSeed)))

	var experiment = &NoCExperiment{
		Config:config,
		CycleAccurateEventQueue:cycleAccurateEventQueue,
		rand: rand,
	}

	experiment.Stats = make(map[string]string)

	var network = NewNetwork(experiment, config.NumNodes)

	experiment.Network = network

	var _ = NewTransposeTrafficGenerator(network, config.DataPacketInjectionRate, config.MaxPackets, func(src int, dest int) Packet {
		return NewDataPacket(network, src, dest, config.DataPacketSize, true, func() {})
	})

	return experiment
}

func (experiment *NoCExperiment) Run() {
	fmt.Printf("[%d] Welcome to ACOGo simulator!\n", experiment.CycleAccurateEventQueue.CurrentCycle)

	for (experiment.Config.MaxCycles == -1 || experiment.CycleAccurateEventQueue.CurrentCycle < experiment.Config.MaxCycles) && (experiment.Config.MaxPackets == -1 || experiment.Network.NumPacketsReceived < experiment.Config.MaxPackets) {
		experiment.CycleAccurateEventQueue.AdvanceOneCycle()
	}

	if !experiment.Config.NoDrain {
		experiment.Network.AcceptPacket = false

		for experiment.Network.NumPacketsReceived != experiment.Network.NumPacketsTransmitted {
			experiment.CycleAccurateEventQueue.AdvanceOneCycle()
		}
	}

	fmt.Printf("[%d] Simulation ended!\n", experiment.CycleAccurateEventQueue.CurrentCycle)
}
