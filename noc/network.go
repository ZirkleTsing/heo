package noc

import (
	"math"
	"github.com/mcai/acogo/simutil"
)

type NetworkDriver interface {
	Config() *NoCConfig
	CycleAccurateEventQueue() *simutil.CycleAccurateEventQueue
	BlockingEventDispatcher() *simutil.BlockingEventDispatcher
}

type Network struct {
	Experiment                   *NoCExperiment
	CurrentPacketId              int64
	NumNodes                     int
	Nodes                        []*Node
	Width                        int
	AcceptPacket                 bool
	TrafficGenerators            []TrafficGenerator

	NumPacketsReceived           int64
	NumPacketsTransmitted        int64

	totalPacketDelays            int64
	MaxPacketDelay               int

	totalPacketHops              int64
	MaxPacketHops                int

	NumPayloadPacketsReceived    int64
	NumPayloadPacketsTransmitted int64

	totalPayloadPacketDelays     int64
	MaxPayloadPacketDelay        int

	totalPayloadPacketHops       int64
	MaxPayloadPacketHops         int

	numFlitPerStateDelaySamples  map[FlitState]int64
	totalFlitPerStateDelays      map[FlitState]int64
	MaxFlitPerStateDelay         map[FlitState]int
}

func NewNetwork(experiment *NoCExperiment) *Network {
	var network = &Network{
		Experiment:experiment,
		NumNodes:experiment.Config.NumNodes,
		Width:int(math.Sqrt(float64(experiment.Config.NumNodes))),
		AcceptPacket:true,

		numFlitPerStateDelaySamples:make(map[FlitState]int64),
		totalFlitPerStateDelays:make(map[FlitState]int64),
		MaxFlitPerStateDelay:make(map[FlitState]int),
	}

	for i := 0; i < network.NumNodes; i++ {
		var node = NewNode(network, i)
		network.Nodes = append(network.Nodes, node)
	}

	switch dataPacketTraffic := experiment.Config.DataPacketTraffic; dataPacketTraffic {
	case TRAFFIC_UNIFORM:
		network.TrafficGenerators = append(network.TrafficGenerators, NewUniformTrafficGenerator(network, experiment.Config.DataPacketInjectionRate, experiment.Config.MaxPackets, func(src int, dest int) Packet {
			return NewDataPacket(network, src, dest, experiment.Config.DataPacketSize, true, func() {})
		}))
	case TRAFFIC_TRANSPOSE1:
		network.TrafficGenerators = append(network.TrafficGenerators, NewTranspose1TrafficGenerator(network, experiment.Config.DataPacketInjectionRate, experiment.Config.MaxPackets, func(src int, dest int) Packet {
			return NewDataPacket(network, src, dest, experiment.Config.DataPacketSize, true, func() {})
		}))
	case TRAFFIC_TRANSPOSE2:
		network.TrafficGenerators = append(network.TrafficGenerators, NewTranspose2TrafficGenerator(network, experiment.Config.DataPacketInjectionRate, experiment.Config.MaxPackets, func(src int, dest int) Packet {
			return NewDataPacket(network, src, dest, experiment.Config.DataPacketSize, true, func() {})
		}))
	}

	switch selection := experiment.Config.Selection; selection {
	case SELECTION_ACO:
		switch antPacketTraffic := experiment.Config.AntPacketTraffic; antPacketTraffic {
		case TRAFFIC_UNIFORM:
			network.TrafficGenerators = append(network.TrafficGenerators, NewUniformTrafficGenerator(network, experiment.Config.AntPacketInjectionRate, int64(-1), func(src int, dest int) Packet {
				return NewAntPacket(network, src, dest, experiment.Config.AntPacketSize, func() {}, true)
			}))
		case TRAFFIC_TRANSPOSE1:
			network.TrafficGenerators = append(network.TrafficGenerators, NewTranspose1TrafficGenerator(network, experiment.Config.AntPacketInjectionRate, int64(-1), func(src int, dest int) Packet {
				return NewAntPacket(network, src, dest, experiment.Config.AntPacketSize, func() {}, true)
			}))
		case TRAFFIC_TRANSPOSE2:
			network.TrafficGenerators = append(network.TrafficGenerators, NewTranspose2TrafficGenerator(network, experiment.Config.AntPacketInjectionRate, int64(-1), func(src int, dest int) Packet {
				return NewAntPacket(network, src, dest, experiment.Config.AntPacketSize, func() {}, true)
			}))
		}
	}

	experiment.CycleAccurateEventQueue.AddPerCycleEvent(func() {
		for _, node := range network.Nodes {
			node.Router.AdvanceOneCycle()
		}

		for _, trafficGenerator := range network.TrafficGenerators {
			trafficGenerator.AdvanceOneCycle()
		}
	})

	return network
}

func (network *Network) GetX(id int) int {
	return id % network.Width
}

func (network *Network) GetY(id int) int {
	return (id - id % network.Width) / network.Width
}

func (network *Network) Receive(packet Packet) bool {
	if !network.Nodes[packet.GetSrc()].Router.InjectPacket(packet) {
		network.Experiment.CycleAccurateEventQueue.Schedule(func() {
			network.Receive(packet)
		}, 1)
		return false
	}

	network.LogPacketReceived(packet)

	return true
}

func (network *Network) LogPacketReceived(packet Packet) {
	network.NumPacketsReceived++

	if packet.GetHasPayload() {
		network.NumPayloadPacketsReceived++
	}
}

func (network *Network) LogPacketTransmitted(packet Packet) {
	network.NumPacketsTransmitted++

	if packet.GetHasPayload() {
		network.NumPayloadPacketsTransmitted++
	}

	network.totalPacketDelays += int64(Delay(packet))
	network.totalPacketHops += int64(Hops(packet))

	if packet.GetHasPayload() {
		network.totalPayloadPacketDelays += int64(Delay(packet))
		network.totalPayloadPacketHops += int64(Hops(packet))
	}

	network.MaxPacketDelay = int(math.Max(float64(network.MaxPacketDelay), float64(Delay(packet))))
	network.MaxPacketHops = int(math.Max(float64(network.MaxPacketHops), float64(Hops(packet))))

	if packet.GetHasPayload() {
		network.MaxPayloadPacketDelay = int(math.Max(float64(network.MaxPayloadPacketDelay), float64(Delay(packet))))
		network.MaxPayloadPacketHops = int(math.Max(float64(network.MaxPayloadPacketHops), float64(Hops(packet))))
	}
}

func (network *Network) LogFlitPerStateDelay(state FlitState, delay int) {
	if _, exists := network.numFlitPerStateDelaySamples[state]; !exists {
		network.numFlitPerStateDelaySamples[state] = int64(0)
	}

	network.numFlitPerStateDelaySamples[state]++

	if _, exists := network.totalFlitPerStateDelays[state]; !exists {
		network.totalFlitPerStateDelays[state] = int64(0)
	}

	network.totalFlitPerStateDelays[state] += int64(delay)

	if _, exists := network.MaxFlitPerStateDelay[state]; !exists {
		network.MaxFlitPerStateDelay[state] = 0
	}

	network.MaxFlitPerStateDelay[state] = int(math.Max(float64(network.MaxFlitPerStateDelay[state]), float64(delay)))
}

func (network *Network) Throughput() float64 {
	return float64(network.NumPacketsTransmitted) / float64(network.Experiment.CycleAccurateEventQueue.CurrentCycle) / float64(network.NumNodes)
}

func (network *Network) AveragePacketDelay() float64 {
	if network.NumPacketsTransmitted > 0 {
		return float64(network.totalPacketDelays) / float64(network.NumPacketsTransmitted)
	} else {
		return 0.0
	}
}

func (network *Network) AveragePacketHops() float64 {
	if network.NumPacketsTransmitted > 0 {
		return float64(network.totalPacketHops) / float64(network.NumPacketsTransmitted)
	} else {
		return 0.0
	}
}

func (network *Network) PayloadThroughput() float64 {
	return float64(network.NumPayloadPacketsTransmitted) / float64(network.Experiment.CycleAccurateEventQueue.CurrentCycle) / float64(network.NumNodes)
}

func (network *Network) AveragePayloadPacketDelay() float64 {
	if network.NumPayloadPacketsTransmitted > 0 {
		return float64(network.totalPayloadPacketDelays) / float64(network.NumPayloadPacketsTransmitted)
	} else {
		return 0.0
	}
}

func (network *Network) AveragePayloadPacketHops() float64 {
	if network.NumPayloadPacketsTransmitted > 0 {
		return float64(network.totalPayloadPacketHops) / float64(network.NumPayloadPacketsTransmitted)
	} else {
		return 0.0
	}
}

func (network *Network) AverageFlitPerStateDelay(state FlitState) float64 {
	if _, exists := network.numFlitPerStateDelaySamples[state]; exists {
		if network.numFlitPerStateDelaySamples[state] > 0 {
			return float64(network.totalFlitPerStateDelays[state]) / float64(network.numFlitPerStateDelaySamples[state])
		}
	}

	return 0.0
}