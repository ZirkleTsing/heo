package main

import (
	"fmt"
	"container/heap"
)

type Direction int

const (
	DirectionLocal = 0
	DirectionNorth = 1
	DirectionEast = 2
	DirectionSouth = 3
	DirectionWest = 4
)

type FlitState int

const (
	FlitStateInputBuffer = 0
	FlitStateRouteComputation = 1
	FlitStateVirtualChannelAllocation = 2
	FlitStateSwitchAllocation = 3
	FlitStateSwitchTraversal = 4
	FlitStateLinkTraversal = 5
	FlitStateDestinationArrived = 6
)

type CycleAccurateEvent struct {
	eventQueue *CycleAccurateEventQueue
	when       int
	action     func()
	id         int
}

type CycleAccurateEventQueue struct {
	events         []*CycleAccurateEvent
	currentCycle   int
	currentEventId int
}

func (eventQueue CycleAccurateEventQueue) AdvanceOneCycle() {
	fmt.Println("Welcome to ACOGo, haha!")
}

func GetReflexDirection(direction Direction) int {
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

}

type Flit struct {
	packet                        *Packet
	num                           int
	head                          bool
	tail                          bool
	node                          *Node
	state                         FlitState
	prevStateTimestamp, timestamp int
}

type InputBuffer struct {
	inputVirtualChannel *InputVirtualChannel
	flits               []*Flit
}

type InputVirtualChannel struct {
	inputPort            *InputPort
	id                   int
	inputBuffer          *InputBuffer
	route                Direction
	outputVirtualChannel *OutputVirtualChannel
}

type InputPort struct {
	router          *Router
	direction       *Direction
	virtualChannels []*InputVirtualChannel
}

type OutputVirtualChannel struct {
	outputPort          *OutputPort
	id                  int
	inputVirtualChannel *InputVirtualChannel
	credits             int
	arbiter             *VirtualChannelArbiter
}

type OutputPort struct {
	router          *Router
	direction       Direction
	virtualChannels []*OutputVirtualChannel
	arbiter         *SwitchArbiter
}

type VirtualChannelArbiter struct {
	inputVirtualChannels []*InputVirtualChannel
}

type SwitchArbiter struct {
	inputVirtualChannels []*InputVirtualChannel
}

type Router struct {
	node            *Node
	injectionBuffer []*Packet
	inputPorts      map[Direction]*InputPort
	outputPorts     map[Direction]*OutputPort
}

func main() {
	var eventQueue CycleAccurateEventQueue
	eventQueue.AdvanceOneCycle()
}
