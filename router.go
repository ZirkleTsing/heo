package acogo

import "math"

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
	Packet                        Packet
	Num                           int
	Head                          bool
	Tail                          bool
	Node                          *Node
	State                         FlitState
	Timestamp int
}

func NewFlit(packet Packet, num int, head bool, tail bool) *Flit {
	var flit = &Flit{
		Packet:packet,
		Num:num,
		Head:head,
		Tail:tail,
		Timestamp: packet.GetNetwork().Experiment.CycleAccurateEventQueue.CurrentCycle,
	}

	packet.SetFlits(append(packet.GetFlits(), flit))

	return flit
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

func (arbiter *VirtualChannelArbiter) Next() *InputVirtualChannel  {
	if arbiter.OutputVirtualChannel.InputVirtualChannel != nil {
		return nil
	}

	for _, inputVirtualChannel := range arbiter.InputVirtualChannels {
		if inputVirtualChannel.Route == arbiter.OutputVirtualChannel.OutputPort.Direction {
			var flit = inputVirtualChannel.InputBuffer.Peek()
			if flit != nil && flit.Head && flit.State == FlitStateRouteComputation {
				return inputVirtualChannel
			}
		}
	}

	return nil
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

func (arbiter *SwitchArbiter) Next() *InputVirtualChannel  {
	for _, inputVirtualChannel := range arbiter.InputVirtualChannels {
		if inputVirtualChannel.OutputVirtualChannel != nil && inputVirtualChannel.OutputVirtualChannel.OutputPort == arbiter.OutputPort {
			var flit = inputVirtualChannel.InputBuffer.Peek()
			if flit != nil && (flit.Head && flit.State == FlitStateVirtualChannelAllocation || !flit.Head && flit.State == FlitStateInputBuffer) {
				return inputVirtualChannel
			}
		}
	}

	return nil
}

type Router struct {
	Node            *Node
	InjectionBuffer *InjectionBuffer
	InputPorts      map[Direction]*InputPort
	OutputPorts     map[Direction]*OutputPort
}

func NewRouter(node *Node) *Router {
	var router = &Router{
		Node:node,
		InputPorts:make(map[Direction]*InputPort),
		OutputPorts:make(map[Direction]*OutputPort),
	}

	router.InjectionBuffer = NewInjectionBuffer(router)

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
	for _, outputPort := range router.OutputPorts {
		for _, outputVirtualChannel := range outputPort.VirtualChannels {
			var inputVirtualChannel = outputVirtualChannel.InputVirtualChannel
			if inputVirtualChannel != nil && outputVirtualChannel.Credits > 0 {
				var flit = inputVirtualChannel.InputBuffer.Peek()
				if (flit != nil && flit.State == FlitStateSwitchTraversal) {
					if (outputPort.Direction != DirectionLocal) {
						flit.Node = router.Node
						flit.State = FlitStateLinkTraversal

						var nextHop = router.Node.Neighbors[outputPort.Direction]
						var ip = outputPort.Direction.GetReflexDirection()
						var ivc = outputVirtualChannel.Num

						router.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
							router.nextHopArrived(flit, nextHop, ip, ivc)
						}, router.Node.Network.Experiment.Config.LinkDelay)
					}

					inputVirtualChannel.InputBuffer.Pop()

					if outputPort.Direction != DirectionLocal {
						outputVirtualChannel.Credits--
					} else {
						flit.Node = router.Node
						flit.State = FlitStateDestinationArrived
					}

					if flit.Tail {
						inputVirtualChannel.OutputVirtualChannel = nil
						outputVirtualChannel.InputVirtualChannel = nil

						if (outputPort.Direction == DirectionLocal) {
							flit.Packet.HandleDestArrived(inputVirtualChannel)
						}
					}
				}
			}
		}
	}
}

func (router *Router) nextHopArrived(flit *Flit, nextHop int, ip Direction, ivc int) {
	var inputBuffer =
		router.Node.Network.Nodes[nextHop].Router.InputPorts[ip].VirtualChannels[ivc].InputBuffer

	if !inputBuffer.Full() {
		router.Node.Network.Nodes[nextHop].Router.InsertFlit(flit, ip, ivc)
	} else {
		router.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
			router.nextHopArrived(flit, nextHop, ip, ivc)
		}, 1)
	}
}

func (router *Router) stageSwitchTraversal() {
	for _, outputPort := range router.OutputPorts {
		for _, inputPort := range router.InputPorts {
			if outputPort.Direction == inputPort.Direction {
				continue;
			}

			for _, inputVirtualChannel := range inputPort.VirtualChannels {
				if inputVirtualChannel.OutputVirtualChannel != nil && inputVirtualChannel.OutputVirtualChannel.OutputPort == outputPort {
					var flit = inputVirtualChannel.InputBuffer.Peek()
					if flit != nil && flit.State == FlitStateSwitchAllocation {
						flit.Node = router.Node
						flit.State = FlitStateSwitchTraversal

						if inputPort.Direction != DirectionLocal {
							var parent = router.Node.Network.Nodes[router.Node.Neighbors[inputPort.Direction]]

							var parentOutputVirtualChannel = parent.Router.OutputPorts[inputPort.Direction.GetReflexDirection()].VirtualChannels[inputVirtualChannel.Num]

							parentOutputVirtualChannel.Credits++
						}
					}
				}
			}
		}
	}
}

func (router *Router) stageSwitchAllocation() {
	for _, outputPort := range router.OutputPorts {
		var winnerInputVirtualChannel = outputPort.Arbiter.Next()

		if winnerInputVirtualChannel != nil {
			var flit = winnerInputVirtualChannel.InputBuffer.Peek()
			flit.Node = router.Node
			flit.State = FlitStateSwitchAllocation
		}
	}
}

func (router *Router) stageVirtualChannelAllocation() {
	for _, outputPort := range router.OutputPorts {
		for _, outputVirtualChannel := range outputPort.VirtualChannels {
			if outputVirtualChannel.InputVirtualChannel == nil {
				var winnerInputVirtualChannel = outputVirtualChannel.Arbiter.Next()

				if winnerInputVirtualChannel != nil {
					var flit = winnerInputVirtualChannel.InputBuffer.Peek()
					flit.Node = router.Node
					flit.State = FlitStateVirtualChannelAllocation

					winnerInputVirtualChannel.OutputVirtualChannel = outputVirtualChannel
					outputVirtualChannel.InputVirtualChannel = winnerInputVirtualChannel
				}
			}
		}
	}
}

func (router *Router) stageRouteComputation() {
	for _, inputPort := range router.InputPorts {
		for _, inputVirtualChannel := range inputPort.VirtualChannels {
			var flit = inputVirtualChannel.InputBuffer.Peek()

			if flit != nil && flit.Head && flit.State == FlitStateInputBuffer {
				if flit.Packet.GetDest() == router.Node.Id {
					inputVirtualChannel.Route = DirectionLocal
				} else {
					inputVirtualChannel.Route = flit.Packet.DoRouteComputation(inputVirtualChannel)
				}

				flit.Node = router.Node
				flit.State = FlitStateRouteComputation
			}
		}
	}
}

func (router *Router) localPacketInjection() {
	for {
		var requestInserted = false;

		for ivc := 0; ivc < router.Node.Network.Experiment.Config.NumVirtualChannels; ivc++ {
			if router.InjectionBuffer.Count() == 0 {
				return;
			}

			var packet = router.InjectionBuffer.Peek().(Packet)

			var numFlits = int(math.Ceil(float64(packet.GetSize()) / float64(router.Node.Network.Experiment.Config.LinkWidth)))

			var inputBuffer = router.InputPorts[DirectionLocal].VirtualChannels[ivc].InputBuffer

			if inputBuffer.Count() + numFlits <= inputBuffer.Size() {
				for i:= 0; i < numFlits; i++ {
					var flit = NewFlit(packet, i, i == 0, i == numFlits - 1)
					router.InsertFlit(flit, DirectionLocal, ivc)
				}

				router.InjectionBuffer.Pop()
				requestInserted = true;
				break;
			}
		}

		if(!requestInserted) {
			break;
		}
	}
}

func (router *Router) InjectPacket(packet Packet) bool {
	if !router.InjectionBuffer.Full() {
		router.InjectionBuffer.Push(packet)
		return true
	}

	return false
}

func (router *Router) InsertFlit(flit *Flit, ip Direction, ivc int) {
	router.InputPorts[ip].VirtualChannels[ivc].InputBuffer.Push(flit)
	flit.Node = router.Node
	flit.State = FlitStateInputBuffer
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