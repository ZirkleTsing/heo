package acogo

import (
	"math"
)

type Direction int

const (
	DirectionLocal = 0
	DirectionNorth = 1
	DirectionEast = 2
	DirectionSouth = 3
	DirectionWest = 4
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
	nodeId    int
	timestamp int
}

type Packet struct {
	network              *Network
	id                   int
	beginCycle, endCycle int
	src, dest            int
	size                 int
	onCompletedCallback  func()
	memory               []*PacketMemoryEntry
	flits                []*Flit
	hasPayload           bool
}

func NewPacket(network *Network, src int, dest int, size int, onCompletedCallback func()) *Packet {
	var packet = &Packet{
		network:network,
		src:src,
		dest:dest,
		size:size,
		onCompletedCallback:onCompletedCallback,
		hasPayload:true,
	}

	return packet
}

type AntPacket struct {
	forward bool
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
	network   *Network
	id        int
	x, y      int
	neighbors map[Direction]int
	router    *Router
}

func NewNode(network *Network, id int) *Node {
	var node = &Node{
		network:network,
		id:id,
		x:network.GetX(id),
		y:network.GetY(id),
	}

	node.neighbors = make(map[Direction]int)

	if (id / network.width > 0) {
		node.neighbors[DirectionNorth] = id - network.width
	}

	if ( (id % network.width) != network.width - 1) {
		node.neighbors[DirectionEast] = id + 1
	}

	if (id / network.width < network.width - 1) {
		node.neighbors[DirectionSouth] = id + network.width
	}

	if (id % network.width != 0) {
		node.neighbors[DirectionWest] = id - 1
	}

	node.router = &Router{
		node:node,
	}

	return node
}

type Network struct {
	experiment              *NoCExperiment
	currentPacketId         int
	cycleAccurateEventQueue *CycleAccurateEventQueue
	numNodes                int
	nodes                   []*Node
	width                   int
	acceptPacket            bool
	numPacketsReceived      int
	numPacketsTransmitted   int
}

func NewNetwork(experiment *NoCExperiment, numNodes int, cycleAccurateEventQueue *CycleAccurateEventQueue) *Network {
	var network = &Network{
		experiment:experiment,
		numNodes:numNodes,
		cycleAccurateEventQueue: cycleAccurateEventQueue,
		width:int(math.Sqrt(float64(numNodes))),
	}

	for i := 0; i < numNodes; i++ {
		var node = NewNode(network, i)
		network.nodes = append(network.nodes, node)
	}

	return network
}

func (network *Network) GetX(id int) int {
	return id % network.width
}

func (network *Network) GetY(id int) int {
	return (id - id % network.width) / network.width
}

func (network *Network) Receive(packet *Packet) bool {
	if !network.nodes[packet.src].router.InjectPacket(packet) {
		network.cycleAccurateEventQueue.Schedule(func() {
			network.Receive(packet)
		}, 1)
		return false
	}

	return true
}