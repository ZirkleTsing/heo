package acogo

type FlitState int

const (
	FlitStateInputBuffer = 1
	FlitStateRouteComputation = 2
	FlitStateVirtualChannelAllocation = 3
	FlitStateSwitchAllocation = 4
	FlitStateSwitchTraversal = 5
	FlitStateLinkTraversal = 6
	FlitStateDestinationArrived = 7
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
	flits               *Queue
}

func NewInputBuffer(inputVirtualChannel *InputVirtualChannel) *InputBuffer {
	var inputBuffer = &InputBuffer{
		inputVirtualChannel:inputVirtualChannel,
		flits:NewQueue(inputVirtualChannel.inputPort.router.node.network.experiment.config.maxInputBufferSize),
	}

	return inputBuffer
}

func (inputBuffer *InputBuffer) Push(flit *Flit) {
	inputBuffer.flits.Push(flit)
}

func (inputBuffer *InputBuffer) Peek() *Flit {
	if inputBuffer.flits.count > 0 {
		return inputBuffer.flits.Peek().(*Flit)
	} else {
		return nil
	}
}

func (inputBuffer *InputBuffer) Pop() {
	inputBuffer.flits.Pop()
}

func (inputBuffer *InputBuffer) Full() bool {
	return inputBuffer.flits.size == inputBuffer.flits.count
}

func (inputBuffer *InputBuffer) Count() int {
	return inputBuffer.flits.count
}

type InputVirtualChannel struct {
	inputPort            *InputPort
	id                   int
	inputBuffer          *InputBuffer
	route                Direction
	outputVirtualChannel *OutputVirtualChannel
}

func NewInputVirtualChannel(inputPort *InputPort, id int) *InputVirtualChannel {
	var inputVirtualChannel = &InputVirtualChannel{
		inputPort:inputPort,
		id:id,
	}

	inputVirtualChannel.inputBuffer = NewInputBuffer(inputVirtualChannel)

	return inputVirtualChannel
}

type InputPort struct {
	router          *Router
	direction       Direction
	virtualChannels []*InputVirtualChannel
}

func NewInputPort(router *Router, direction Direction) *InputPort {
	var inputPort = &InputPort{
		router:router,
		direction:direction,
	}

	for i := 0; i < router.node.network.experiment.config.numVirtualChannels; i++ {
		inputPort.virtualChannels = append(inputPort.virtualChannels, NewInputVirtualChannel(inputPort, i))
	}

	return inputPort
}

type OutputVirtualChannel struct {
	outputPort          *OutputPort
	id                  int
	inputVirtualChannel *InputVirtualChannel
	credits             int
	arbiter             *VirtualChannelArbiter
}

func NewOutputVirtualChannel(outputPort *OutputPort, id int) *OutputVirtualChannel {
	var outputVirtualChannel = &OutputVirtualChannel{
		outputPort:outputPort,
		id: id,
		credits:10,
	}

	outputVirtualChannel.arbiter = NewVirtualChannelArbiter(outputVirtualChannel)

	return outputVirtualChannel
}

type OutputPort struct {
	router          *Router
	direction       Direction
	virtualChannels []*OutputVirtualChannel
	arbiter         *SwitchArbiter
}

func NewOutputPort(router *Router, direction Direction) *OutputPort {
	var outputPort = &OutputPort{
		router:router,
		direction:direction,
	}

	for i := 0; i < router.node.network.experiment.config.numVirtualChannels; i++ {
		var outputVirtualChannel = NewOutputVirtualChannel(outputPort, i)
		outputPort.virtualChannels = append(outputPort.virtualChannels, outputVirtualChannel)
	}

	outputPort.arbiter = NewSwitchArbiter(outputPort)

	return outputPort
}

type VirtualChannelArbiter struct {
	outputVirtualChannel *OutputVirtualChannel
	inputVirtualChannels []*InputVirtualChannel
}

func NewVirtualChannelArbiter(outputVirtualChannel *OutputVirtualChannel) *VirtualChannelArbiter {
	var virtualChannelArbiter = &VirtualChannelArbiter{
		outputVirtualChannel:outputVirtualChannel,
		inputVirtualChannels:outputVirtualChannel.outputPort.router.GetInputVirtualChannels(),
	}

	return virtualChannelArbiter
}

type SwitchArbiter struct {
	outputPort           *OutputPort
	inputVirtualChannels []*InputVirtualChannel
}

func NewSwitchArbiter(outputPort *OutputPort) *SwitchArbiter {
	var switchArbiter = &SwitchArbiter{
		outputPort:outputPort,
		inputVirtualChannels:outputPort.router.GetInputVirtualChannels(),
	}

	return switchArbiter
}

type Router struct {
	node            *Node
	injectionBuffer []*Packet
	inputPorts      map[Direction]*InputPort
	outputPorts     map[Direction]*OutputPort
}

func NewRouter(node *Node) *Router {
	var router = &Router{
		node:node,
		inputPorts:make(map[Direction]*InputPort),
		outputPorts:make(map[Direction]*OutputPort),
	}

	for i := 0; i < DirectionWest; i++ {
		router.inputPorts[Direction(i)] = NewInputPort(router, Direction(i))
		router.outputPorts[Direction(i)] = NewOutputPort(router, Direction(i))
	}

	router.node.network.experiment.cycleAccurateEventQueue.AddPerCycleEvent(router.advanceOneCycle)

	return router
}

func (router *Router) advanceOneCycle() {
	router.stageLinkTraversal()
	router.stageSwitchTraversal()
	router.stageSwitchAllocation()
	router.stageVirtualChannelAllocation()
	router.stageRouteComputation()
	router.localPacketInjection()
}

func (router *Router) stageLinkTraversal() {
}

func (router *Router) stageSwitchTraversal() {
}

func (router *Router) stageSwitchAllocation() {
}

func (router *Router) stageVirtualChannelAllocation() {
}

func (router *Router) stageRouteComputation() {
}

func (router *Router) localPacketInjection() {
}

func (router *Router) InjectPacket(packet *Packet) bool {
	return false //TODO
}

func (router *Router) GetInputVirtualChannels() []*InputVirtualChannel {
	var inputVirtualChannels []*InputVirtualChannel

	for _, inputPort := range router.inputPorts {
		for i := range inputPort.virtualChannels {
			inputVirtualChannels = append(inputVirtualChannels, inputPort.virtualChannels[i])
		}
	}

	return inputVirtualChannels
}