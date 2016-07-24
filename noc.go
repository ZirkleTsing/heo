package acogo

import (
	"fmt"
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

func (direction Direction) GetReflexDirection() Direction {
	switch direction {
	case DirectionLocal:
		return DirectionLocal
	case DirectionNorth:
		return DirectionSouth
	case DirectionEast:
		return DirectionWest
	case DirectionSouth:
		return DirectionNorth
	case DirectionWest:
		return DirectionEast
	default:
		panic(fmt.Sprintf("%d", direction))
	}
}

type Node struct {
	Network            *Network
	Id                 int
	X, Y               int
	Neighbors          map[Direction]int
	Router             *Router
	RoutingAlgorithm   RoutingAlgorithm
	SelectionAlgorithm SelectionAlgorithm
}

func NewNode(network *Network, id int) *Node {
	var node = &Node{
		Network:network,
		Id:id,
		X:network.GetX(id),
		Y:network.GetY(id),
		Neighbors:make(map[Direction]int),
	}

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

	node.RoutingAlgorithm = NewOddEvenRoutingAlgorithm(node)
	//node.RoutingAlgorithm = NewXYRoutingAlgorithm(node)

	//node.SelectionAlgorithm = NewACOSelectionAlgorithm(node)
	node.SelectionAlgorithm = NewBufferLevelSelectionAlgorithm(node)

	return node
}

func (node *Node) DumpNeighbors() {
	for direction, neighbor := range node.Neighbors {
		fmt.Printf("node#%d.neighbors[%d]=%d\n", node.Id, direction, neighbor)
	}

	fmt.Println()
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