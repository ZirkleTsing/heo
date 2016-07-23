package acogo

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