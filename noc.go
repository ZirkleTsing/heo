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
}

type Node struct {
	network   *Network
	id        int
	x, y      int
	neighbors map[Direction]int
	router    *Router
}

type Network struct {
	currentPacketId         int
	cycleAccurateEventQueue *CycleAccurateEventQueue
	numNodes                int
	nodes                   []*Node
	width                   int
	acceptPacket            bool
	numPacketsReceived int
	numPacketsTransmitted int
}

func NewNetwork(numNodes int, cycleAccurateEventQueue *CycleAccurateEventQueue) *Network {
	var network = &Network{
		numNodes:numNodes,
		cycleAccurateEventQueue: cycleAccurateEventQueue,
		width:int(math.Sqrt(float64(numNodes))),
	}

	for i := 0; i < numNodes; i++ {
		var node = &Node{
			network:network,
			id:i,
		}

		network.nodes = append(network.nodes, node)
	}

	return network
}