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
}

func NewNetwork(experiment *NoCExperiment, numNodes int) *Network {
	var network = &Network{
		Experiment:experiment,
		NumNodes:numNodes,
		Width:int(math.Sqrt(float64(numNodes))),
		AcceptPacket:true,
	}

	for i := 0; i < numNodes; i++ {
		var node = NewNode(network, i)
		network.Nodes = append(network.Nodes, node)
	}

	network.Experiment.CycleAccurateEventQueue.AddPerCycleEvent(func() {
		for _, node := range network.Nodes {
			node.Router.AdvanceOneCycle()
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