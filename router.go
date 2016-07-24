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
	Packet                        *Packet
	Num                           int
	Head                          bool
	Tail                          bool
	Node                          *Node
	State                         FlitState
	prevStateTimestamp, Timestamp int
}

type InputBuffer struct {
	InputVirtualChannel *InputVirtualChannel
	Flits               *Queue
}

func NewInputBuffer(inputVirtualChannel *InputVirtualChannel) *InputBuffer {
	var inputBuffer = &InputBuffer{
		InputVirtualChannel:inputVirtualChannel,
		Flits:NewQueue(inputVirtualChannel.InputPort.Router.Node.Network.Experiment.Config.MaxInputBufferSize),
	}

	return inputBuffer
}

func (inputBuffer *InputBuffer) Push(flit *Flit) {
	inputBuffer.Flits.Push(flit)
}

func (inputBuffer *InputBuffer) Peek() *Flit {
	if inputBuffer.Flits.Count > 0 {
		return inputBuffer.Flits.Peek().(*Flit)
	} else {
		return nil
	}
}

func (inputBuffer *InputBuffer) Pop() {
	inputBuffer.Flits.Pop()
}

func (inputBuffer *InputBuffer) Full() bool {
	return inputBuffer.Flits.Size == inputBuffer.Flits.Count
}

func (inputBuffer *InputBuffer) Count() int {
	return inputBuffer.Flits.Count
}

func (inputBuffer *InputBuffer) FreeSlots() int  {
	return inputBuffer.Flits.Size - inputBuffer.Flits.Count
}

type InputVirtualChannel struct {
	InputPort            *InputPort
	Num                  int
	InputBuffer          *InputBuffer
	Route                Direction
	OutputVirtualChannel *OutputVirtualChannel
}

func NewInputVirtualChannel(inputPort *InputPort, num int) *InputVirtualChannel {
	var inputVirtualChannel = &InputVirtualChannel{
		InputPort:inputPort,
		Num:num,
	}

	inputVirtualChannel.InputBuffer = NewInputBuffer(inputVirtualChannel)

	return inputVirtualChannel
}

type InputPort struct {
	Router          *Router
	Direction       Direction
	VirtualChannels []*InputVirtualChannel
}

func NewInputPort(router *Router, direction Direction) *InputPort {
	var inputPort = &InputPort{
		Router:router,
		Direction:direction,
	}

	for i := 0; i < router.Node.Network.Experiment.Config.NumVirtualChannels; i++ {
		inputPort.VirtualChannels = append(inputPort.VirtualChannels, NewInputVirtualChannel(inputPort, i))
	}

	return inputPort
}

type OutputVirtualChannel struct {
	OutputPort          *OutputPort
	Num                 int
	InputVirtualChannel *InputVirtualChannel
	Credits             int
	Arbiter             *VirtualChannelArbiter
}

func NewOutputVirtualChannel(outputPort *OutputPort, num int) *OutputVirtualChannel {
	var outputVirtualChannel = &OutputVirtualChannel{
		OutputPort:outputPort,
		Num: num,
		Credits:10,
	}

	outputVirtualChannel.Arbiter = NewVirtualChannelArbiter(outputVirtualChannel)

	return outputVirtualChannel
}

type OutputPort struct {
	Router          *Router
	Direction       Direction
	VirtualChannels []*OutputVirtualChannel
	Arbiter         *SwitchArbiter
}

func NewOutputPort(router *Router, direction Direction) *OutputPort {
	var outputPort = &OutputPort{
		Router:router,
		Direction:direction,
	}

	for i := 0; i < router.Node.Network.Experiment.Config.NumVirtualChannels; i++ {
		var outputVirtualChannel = NewOutputVirtualChannel(outputPort, i)
		outputPort.VirtualChannels = append(outputPort.VirtualChannels, outputVirtualChannel)
	}

	outputPort.Arbiter = NewSwitchArbiter(outputPort)

	return outputPort
}

type VirtualChannelArbiter struct {
	OutputVirtualChannel *OutputVirtualChannel
	InputVirtualChannels []*InputVirtualChannel
}

func NewVirtualChannelArbiter(outputVirtualChannel *OutputVirtualChannel) *VirtualChannelArbiter {
	var virtualChannelArbiter = &VirtualChannelArbiter{
		OutputVirtualChannel:outputVirtualChannel,
		InputVirtualChannels:outputVirtualChannel.OutputPort.Router.GetInputVirtualChannels(),
	}

	return virtualChannelArbiter
}

type SwitchArbiter struct {
	OutputPort           *OutputPort
	InputVirtualChannels []*InputVirtualChannel
}

func NewSwitchArbiter(outputPort *OutputPort) *SwitchArbiter {
	var switchArbiter = &SwitchArbiter{
		OutputPort:outputPort,
		InputVirtualChannels:outputPort.Router.GetInputVirtualChannels(),
	}

	return switchArbiter
}

type Router struct {
	Node            *Node
	InjectionBuffer []*Packet
	InputPorts      map[Direction]*InputPort
	OutputPorts     map[Direction]*OutputPort
}

func NewRouter(node *Node) *Router {
	var router = &Router{
		Node:node,
		InputPorts:make(map[Direction]*InputPort),
		OutputPorts:make(map[Direction]*OutputPort),
	}

	for i := 0; i < DirectionWest; i++ {
		router.InputPorts[Direction(i)] = NewInputPort(router, Direction(i))
		router.OutputPorts[Direction(i)] = NewOutputPort(router, Direction(i))
	}

	return router
}

func (router *Router) AdvanceOneCycle() {
	router.stageLinkTraversal()
	router.stageSwitchTraversal()
	router.stageSwitchAllocation()
	router.stageVirtualChannelAllocation()
	router.stageRouteComputation()
	router.localPacketInjection()
}

func (router *Router) stageLinkTraversal() {
	//TODO
}

func (router *Router) stageSwitchTraversal() {
	//TODO
}

func (router *Router) stageSwitchAllocation() {
	//TODO
}

func (router *Router) stageVirtualChannelAllocation() {
	//TODO
}

func (router *Router) stageRouteComputation() {
	//TODO
}

func (router *Router) localPacketInjection() {
	//TODO
}

func (router *Router) InjectPacket(packet *Packet) bool {
	return true //TODO
}

func (router *Router) GetInputVirtualChannels() []*InputVirtualChannel {
	var inputVirtualChannels []*InputVirtualChannel

	for _, inputPort := range router.InputPorts {
		for i := range inputPort.VirtualChannels {
			inputVirtualChannels = append(inputVirtualChannels, inputPort.VirtualChannels[i])
		}
	}

	return inputVirtualChannels
}

func (router *Router) FreeSlots(ip Direction, ivc int) int {
	return router.InputPorts[ip].VirtualChannels[ivc].InputBuffer.FreeSlots()
}