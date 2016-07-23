package acogo

import "math/rand"

type NoCExperiment struct {
	config                  *NoCConfig
	stats                   map[string]string

	beginTime, endTime      int
	cycleAccurateEventQueue *CycleAccurateEventQueue
	network                 *Network
	rand                    *rand.Rand
}

func NewNoCExperiment(config *NoCConfig) *NoCExperiment {
	var cycleAccurateEventQueue = NewCycleAccurateEventQueue()

	var rand = rand.New(rand.NewSource(int64(config.randSeed)))

	var experiment = &NoCExperiment{
		config:config,
		cycleAccurateEventQueue:cycleAccurateEventQueue,
		rand: rand,
	}

	experiment.stats = make(map[string]string)

	var network = NewNetwork(experiment, config.numNodes, cycleAccurateEventQueue)

	experiment.network = network

	var _ = NewTransposeTrafficGenerator(network, config.dataPacketInjectionRate, config.maxPackets, func(src int, dest int) *Packet {
		return NewPacket(network, src, dest, config.dataPacketSize, func() {})
	})

	return experiment
}

func (experiment *NoCExperiment) run() {
	for (experiment.config.maxCycles == -1 || experiment.cycleAccurateEventQueue.currentCycle < experiment.config.maxCycles) && (experiment.config.maxPackets == -1 || experiment.network.numPacketsReceived < experiment.config.maxPackets) {
		experiment.cycleAccurateEventQueue.AdvanceOneCycle()
	}

	if !experiment.config.noDrain {
		experiment.network.acceptPacket = false

		for experiment.network.numPacketsReceived != experiment.network.numPacketsTransmitted {
			experiment.cycleAccurateEventQueue.AdvanceOneCycle()
		}
	}
}
