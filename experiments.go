package acogo

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

	var _ = NewTransposeTrafficGenerator(network, config.DataPacketInjectionRate, config.MaxPackets, func(src int, dest int) *Packet {
		return NewPacket(network, src, dest, config.DataPacketSize, func() {})
	})

	return experiment
}

func (experiment *NoCExperiment) Run() {
	for (experiment.Config.MaxCycles == -1 || experiment.CycleAccurateEventQueue.CurrentCycle < experiment.Config.MaxCycles) && (experiment.Config.MaxPackets == -1 || experiment.Network.NumPacketsReceived < experiment.Config.MaxPackets) {
		experiment.CycleAccurateEventQueue.AdvanceOneCycle()
	}

	if !experiment.Config.NoDrain {
		experiment.Network.AcceptPacket = false

		for experiment.Network.NumPacketsReceived != experiment.Network.NumPacketsTransmitted {
			experiment.CycleAccurateEventQueue.AdvanceOneCycle()
		}
	}
}
