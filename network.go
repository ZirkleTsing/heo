package acogo

import (
	"math"
)

type Network struct {
	Experiment            *NoCExperiment
	CurrentPacketId       int64
	NumNodes              int
	Nodes                 []*Node
	Width                 int
	AcceptPacket          bool
	NumPacketsReceived    int64
	NumPacketsTransmitted int64
	TrafficGenerators     []TrafficGenerator
}

func NewNetwork(experiment *NoCExperiment) *Network {
	var network = &Network{
		Experiment:experiment,
		NumNodes:experiment.Config.NumNodes,
		Width:int(math.Sqrt(float64(experiment.Config.NumNodes))),
		AcceptPacket:true,
	}

	for i := 0; i < network.NumNodes; i++ {
		var node = NewNode(network, i)
		network.Nodes = append(network.Nodes, node)
	}

	switch dataPacketTraffic := experiment.Config.DataPacketTraffic; dataPacketTraffic {
	case "uniform":
		network.TrafficGenerators = append(network.TrafficGenerators, NewUniformTrafficGenerator(network, experiment.Config.DataPacketInjectionRate, experiment.Config.MaxPackets, func(src int, dest int) Packet {
			return NewDataPacket(network, src, dest, experiment.Config.DataPacketSize, true, func() {})
		}))
	case "transpose":
		network.TrafficGenerators = append(network.TrafficGenerators, NewTransposeTrafficGenerator(network, experiment.Config.DataPacketInjectionRate, experiment.Config.MaxPackets, func(src int, dest int) Packet {
			return NewDataPacket(network, src, dest, experiment.Config.DataPacketSize, true, func() {})
		}))
	}

	switch selection := experiment.Config.Selection; selection {
	case "aco":
		switch antPacketTraffic := experiment.Config.AntPacketTraffic; antPacketTraffic {
		case "uniform":
			network.TrafficGenerators = append(network.TrafficGenerators, NewUniformTrafficGenerator(network, experiment.Config.AntPacketInjectionRate, int64(-1), func(src int, dest int) Packet {
				return NewAntPacket(network, src, dest, experiment.Config.AntPacketSize, func() {}, true)
			}))
		case "transpose":
			network.TrafficGenerators = append(network.TrafficGenerators, NewTransposeTrafficGenerator(network, experiment.Config.AntPacketInjectionRate, int64(-1), func(src int, dest int) Packet {
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
	//TODO
}

func (network *Network) LogPacketTransmitted(packet Packet) {
	network.NumPacketsTransmitted++
	//TODO
}