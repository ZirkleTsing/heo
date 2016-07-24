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

func (direction Direction) GetReflexDirection() Direction {
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

type PacketHandler interface {
	HandleDestArrived(packet *Packet, inputVirtualChannel *InputVirtualChannel)

	DoRouteComputation(packet *Packet, inputVirtualChannel *InputVirtualChannel) Direction
}

func NewPacket(network *Network, src int, dest int, size int, hasPayload bool, onCompletedCallback func()) *Packet {
	var packet = &Packet{
		Network:network,
		Src:src,
		Dest:dest,
		Size:size,
		OnCompletedCallback:onCompletedCallback,
		HasPayload:hasPayload,
	}

	return packet
}

func (packet *Packet) HandleDestArrived(inputVirtualChannel *InputVirtualChannel) {
	packet.Memorize(inputVirtualChannel.InputPort.Router.Node.Id)

	packet.EndCycle = inputVirtualChannel.InputPort.Router.Node.Network.Experiment.CycleAccurateEventQueue.CurrentCycle

	if packet.OnCompletedCallback != nil {
		packet.OnCompletedCallback()
	}
}

func (packet *Packet) DoRouteComputation(inputVirtualChannel *InputVirtualChannel) Direction {
	var parent = -1

	if len(packet.Memory) > 0 {
		parent = packet.Memory[len(packet.Memory) - 1].NodeId
	}

	packet.Memorize(inputVirtualChannel.InputPort.Router.Node.Id)

	var directions = inputVirtualChannel.InputPort.Router.Node.RoutingAlgorithm.NextHop(packet.Src, packet.Dest, parent)

	return inputVirtualChannel.InputPort.Router.Node.SelectionAlgorithm.Select(packet.Src, packet.Dest, inputVirtualChannel.Num, directions)
}

func (packet *Packet) Memorize(currentNodeId int) {
	packet.Memory = append(packet.Memory, &PacketMemoryEntry{
		NodeId:currentNodeId,
		Timestamp:packet.Network.Experiment.CycleAccurateEventQueue.CurrentCycle,
	})
}

type Node struct {
	Network   *Network
	Id        int
	X, Y      int
	Neighbors map[Direction]int
	Router    *Router
	RoutingAlgorithm RoutingAlgorithm //TODO
	SelectionAlgorithm SelectionAlgorithm //TODO
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

	network.Experiment.CycleAccurateEventQueue.AddPerCycleEvent(func() {
		for i := 0; i < numNodes; i++ {
			network.Nodes[i].Router.AdvanceOneCycle()
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

func (network *Network) Receive(packet *Packet) bool {
	if !network.Nodes[packet.Src].Router.InjectPacket(packet) {
		network.Experiment.CycleAccurateEventQueue.Schedule(func() {
			network.Receive(packet)
		}, 1)
		return false
	}

	return true
}