package acogo

import (
	"math"
)

type Direction int

const (
	DirectionLocal = 1
	DirectionNorth = 2
	DirectionEast = 3
	DirectionSouth = 4
	DirectionWest = 5
)

func (direction Direction) GetReflexDirection() int {
	switch direction {
	case DirectionLocal:
		return DirectionLocal
	case DirectionNorth:
		return DirectionSouth
	case DirectionEast:
		return DirectionWest
	default:
		return -1
	}
}

type PacketMemoryEntry struct {
	NodeId    int
	Timestamp int
}

type Packet struct {
	Network              *Network
	Id                   int
	BeginCycle, EndCycle int
	Src, Dest            int
	Size                 int
	OnCompletedCallback  func()
	Memory               []*PacketMemoryEntry
	Flits                []*Flit
	HasPayload           bool
}

func NewPacket(network *Network, src int, dest int, size int, onCompletedCallback func()) *Packet {
	var packet = &Packet{
		Network:network,
		Src:src,
		Dest:dest,
		Size:size,
		OnCompletedCallback:onCompletedCallback,
		HasPayload:true,
	}

	return packet
}

type AntPacket struct {
	Forward bool
	*Packet
}

func NewAntPacket(network *Network, src int, dest int, size int, onCompletedCallback func(), forward bool) *AntPacket {
	var packet = &AntPacket{
		forward,
		NewPacket(network, src, dest, size, onCompletedCallback),
	}

	return packet
}

type Node struct {
	Network   *Network
	Id        int
	X, Y      int
	Neighbors map[Direction]int
	Router    *Router
}

func NewNode(network *Network, id int) *Node {
	var node = &Node{
		Network:network,
		Id:id,
		X:network.GetX(id),
		Y:network.GetY(id),
	}

	node.Neighbors = make(map[Direction]int)

	if (id / network.Width > 0) {
		node.Neighbors[DirectionNorth] = id - network.Width
	}

	if ( (id % network.Width) != network.Width - 1) {
		node.Neighbors[DirectionEast] = id + 1
	}

	if (id / network.Width < network.Width - 1) {
		node.Neighbors[DirectionSouth] = id + network.Width
	}

	if (id % network.Width != 0) {
		node.Neighbors[DirectionWest] = id - 1
	}

	node.Router = NewRouter(node)

	return node
}

type Network struct {
	Experiment            *NoCExperiment
	currentPacketId       int
	NumNodes              int
	Nodes                 []*Node
	Width                 int
	AcceptPacket          bool
	NumPacketsReceived    int
	NumPacketsTransmitted int
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

	return network
}

func (network *Network) GetX(id int) int {
	return id % network.Width
}

func (network *Network) GetY(id int) int {
	return (id - id % network.Width) / network.Width
}

func (network *Network) Receive(packet *Packet) bool {
	if !network.Nodes[packet.Src].Router.InjectPacket(packet) {
		network.Experiment.CycleAccurateEventQueue.Schedule(func() {
			network.Receive(packet)
		}, 1)
		return false
	}

	return true
}