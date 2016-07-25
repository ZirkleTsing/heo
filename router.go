package acogo

import (
	"math"
	"fmt"
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

type Flit struct {
	Packet    Packet
	Num       int
	Head      bool
	Tail      bool
	Node      *Node
	State     FlitState
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
		Route:Direction(-1),
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
}

func NewOutputVirtualChannel(outputPort *OutputPort, num int) *OutputVirtualChannel {
	var outputVirtualChannel = &OutputVirtualChannel{
		OutputPort:outputPort,
		Num: num,
		Credits:10,
	}

	return outputVirtualChannel
}

type OutputPort struct {
	Router          *Router
	Direction       Direction
	VirtualChannels []*OutputVirtualChannel
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

	return outputPort
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

	router.InputPorts[DirectionLocal] = NewInputPort(router, DirectionLocal)
	router.OutputPorts[DirectionLocal] = NewOutputPort(router, DirectionLocal)

	for direction:= range node.Neighbors {
		router.InputPorts[direction] = NewInputPort(router, direction)
		router.OutputPorts[direction] = NewOutputPort(router, direction)
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
				if flit != nil && flit.State == FlitStateSwitchTraversal {
					if outputPort.Direction != DirectionLocal {
						flit.State = FlitStateLinkTraversal

						var nextHop = router.Node.Neighbors[outputPort.Direction]
						var ip = outputPort.Direction.GetReflexDirection()
						var ivc = outputVirtualChannel.Num

						if flit.Head {
							fmt.Printf("Pre::nextHopArrived: packet#%d, src=%d, dest=%d, current=%d, nextHop=%d, outputPort.direction=%d\n", flit.Packet.GetId(), flit.Packet.GetSrc(), flit.Packet.GetDest(), router.Node.Id, nextHop, outputPort.Direction)
							router.Node.DumpNeighbors()
						}

						router.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
							router.nextHopArrived(flit, nextHop, ip, ivc)
						}, router.Node.Network.Experiment.Config.LinkDelay)
					}

					inputVirtualChannel.InputBuffer.Pop()

					if outputPort.Direction != DirectionLocal {
						outputVirtualChannel.Credits--
					} else {
						flit.State = FlitStateDestinationArrived
					}

					if flit.Tail {
						inputVirtualChannel.OutputVirtualChannel = nil
						outputVirtualChannel.InputVirtualChannel = nil

						if outputPort.Direction == DirectionLocal {
							flit.Packet.HandleDestArrived(inputVirtualChannel)
						}
					}
				}
			}
		}
	}
}

func (router *Router) nextHopArrived(flit *Flit, nextHop int, ip Direction, ivc int) {
	var inputBuffer = router.Node.Network.Nodes[nextHop].Router.InputPorts[ip].VirtualChannels[ivc].InputBuffer

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
		var winnerInputVirtualChannel = router.findWinnerForSwitchAllocation(outputPort)

		if winnerInputVirtualChannel != nil {
			var flit = winnerInputVirtualChannel.InputBuffer.Peek()
			flit.State = FlitStateSwitchAllocation
		}
	}
}

func (router *Router) findWinnerForSwitchAllocation(outputPort *OutputPort) *InputVirtualChannel {
	for _, inputPort := range router.InputPorts {
		for _, inputVirtualChannel := range inputPort.VirtualChannels {
			if inputVirtualChannel.OutputVirtualChannel != nil && inputVirtualChannel.OutputVirtualChannel.OutputPort == outputPort {
				var flit = inputVirtualChannel.InputBuffer.Peek()
				if flit != nil && (flit.Head && flit.State == FlitStateVirtualChannelAllocation || !flit.Head && flit.State == FlitStateInputBuffer) {
					return inputVirtualChannel
				}
			}
		}
	}

	return nil
}

func (router *Router) stageVirtualChannelAllocation() {
	for _, outputPort := range router.OutputPorts {
		for _, outputVirtualChannel := range outputPort.VirtualChannels {
			if outputVirtualChannel.InputVirtualChannel == nil {
				var winnerInputVirtualChannel = router.findWinnerForVirtualChannelAllocation(outputVirtualChannel)

				if winnerInputVirtualChannel != nil {
					var flit = winnerInputVirtualChannel.InputBuffer.Peek()
					flit.State = FlitStateVirtualChannelAllocation

					winnerInputVirtualChannel.OutputVirtualChannel = outputVirtualChannel
					outputVirtualChannel.InputVirtualChannel = winnerInputVirtualChannel
				}
			}
		}
	}
}

func (router *Router) findWinnerForVirtualChannelAllocation(outputVirtualChannel *OutputVirtualChannel) *InputVirtualChannel {
	for _, inputPort := range outputVirtualChannel.OutputPort.Router.InputPorts {
		for _, inputVirtualChannel := range inputPort.VirtualChannels {
			if inputVirtualChannel.Route == outputVirtualChannel.OutputPort.Direction {
				var flit = inputVirtualChannel.InputBuffer.Peek()
				if flit != nil && flit.Head && flit.State == FlitStateRouteComputation {
					return inputVirtualChannel
				}
			}
		}
	}

	return nil
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
				return
			}

			var packet = router.InjectionBuffer.Peek()

			var numFlits = int(math.Ceil(float64(packet.GetSize()) / float64(router.Node.Network.Experiment.Config.LinkWidth)))

			var inputBuffer = router.InputPorts[DirectionLocal].VirtualChannels[ivc].InputBuffer

			if inputBuffer.Count() + numFlits <= inputBuffer.Size() {
				for i := 0; i < numFlits; i++ {
					var flit = NewFlit(packet, i, i == 0, i == numFlits - 1)
					router.InsertFlit(flit, DirectionLocal, ivc)
				}

				router.InjectionBuffer.Pop()
				requestInserted = true;
				break
			}
		}

		if !requestInserted {
			break
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
		for _, inputVirtualChannel := range inputPort.VirtualChannels {
			inputVirtualChannels = append(inputVirtualChannels, inputVirtualChannel)
		}
	}

	return inputVirtualChannels
}

func (router *Router) FreeSlots(ip Direction, ivc int) int {
	return router.InputPorts[ip].VirtualChannels[ivc].InputBuffer.FreeSlots()
}