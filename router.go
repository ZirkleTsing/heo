package acogo

import (
	"math"
	"fmt"
)

type Router struct {
	Node                    *Node
	InjectionBuffer         *InjectionBuffer
	InputPorts              map[Direction]*InputPort
	OutputPorts             map[Direction]*OutputPort
	NumInflightHeadFlits    map[FlitState]int
	NumInflightNonHeadFlits map[FlitState]int
}

func NewRouter(node *Node) *Router {
	var router = &Router{
		Node:node,
		InputPorts:make(map[Direction]*InputPort),
		OutputPorts:make(map[Direction]*OutputPort),
		NumInflightHeadFlits:make(map[FlitState]int),
		NumInflightNonHeadFlits:make(map[FlitState]int),
	}

	router.InjectionBuffer = NewInjectionBuffer(router)

	router.InputPorts[DIRECTION_LOCAL] = NewInputPort(router, DIRECTION_LOCAL)
	router.OutputPorts[DIRECTION_LOCAL] = NewOutputPort(router, DIRECTION_LOCAL)

	for direction := range node.Neighbors {
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
				if flit != nil && flit.State == FLIT_STATE_SWITCH_TRAVERSAL {
					if outputPort.Direction != DIRECTION_LOCAL {
						flit.State = FLIT_STATE_LINK_TRAVERSAL

						var nextHop = router.Node.Neighbors[outputPort.Direction]
						var ip = outputPort.Direction.GetReflexDirection()
						var ivc = outputVirtualChannel.Num

						if flit.Head {
							fmt.Printf("NextHopArrived(packet#%d, src=%d, dest=%d, current=%d, nextHop=%d, op=%s)\n", flit.Packet.GetId(), flit.Packet.GetSrc(), flit.Packet.GetDest(), router.Node.Id, nextHop, outputPort.Direction)
							router.Node.DumpNeighbors()
						}

						router.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
							router.NextHopArrived(flit, nextHop, ip, ivc)
						}, router.Node.Network.Experiment.Config.LinkDelay)
					}

					inputVirtualChannel.InputBuffer.Pop()

					if outputPort.Direction != DIRECTION_LOCAL {
						outputVirtualChannel.Credits--
					} else {
						flit.State = FLIT_STATE_DESTINATION_ARRIVED
					}

					if flit.Tail {
						inputVirtualChannel.OutputVirtualChannel = nil
						outputVirtualChannel.InputVirtualChannel = nil

						if outputPort.Direction == DIRECTION_LOCAL {
							flit.Packet.HandleDestArrived(inputVirtualChannel)
						}
					}
				}
			}
		}
	}
}

func (router *Router) NextHopArrived(flit *Flit, nextHop int, ip Direction, ivc int) {
	var inputBuffer = router.Node.Network.Nodes[nextHop].Router.InputPorts[ip].VirtualChannels[ivc].InputBuffer

	if !inputBuffer.Full() {
		router.Node.Network.Nodes[nextHop].Router.InsertFlit(flit, ip, ivc)
	} else {
		router.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
			router.NextHopArrived(flit, nextHop, ip, ivc)
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
					if flit != nil && flit.State == FLIT_STATE_SWITCH_ALLOCATION {
						flit.State = FLIT_STATE_SWITCH_TRAVERSAL

						if inputPort.Direction != DIRECTION_LOCAL {
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
			flit.State = FLIT_STATE_SWITCH_ALLOCATION
		}
	}
}

func (router *Router) findWinnerForSwitchAllocation(outputPort *OutputPort) *InputVirtualChannel {
	for _, inputPort := range router.InputPorts {
		for _, inputVirtualChannel := range inputPort.VirtualChannels {
			if inputVirtualChannel.OutputVirtualChannel != nil && inputVirtualChannel.OutputVirtualChannel.OutputPort == outputPort {
				var flit = inputVirtualChannel.InputBuffer.Peek()
				if flit != nil && (flit.Head && flit.State == FLIT_STATE_VIRTUAL_CHANNEL_ALLOCATION || !flit.Head && flit.State == FLIT_STATE_INPUT_BUFFER) {
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
					flit.State = FLIT_STATE_VIRTUAL_CHANNEL_ALLOCATION

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
				if flit != nil && flit.Head && flit.State == FLIT_STATE_ROUTE_COMPUTATION {
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

			if flit != nil && flit.Head && flit.State == FLIT_STATE_INPUT_BUFFER {
				if flit.Packet.GetDest() == router.Node.Id {
					inputVirtualChannel.Route = DIRECTION_LOCAL
				} else {
					inputVirtualChannel.Route = flit.Packet.DoRouteComputation(inputVirtualChannel)
				}

				flit.State = FLIT_STATE_ROUTE_COMPUTATION
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

			var inputBuffer = router.InputPorts[DIRECTION_LOCAL].VirtualChannels[ivc].InputBuffer

			if inputBuffer.Count() + numFlits <= inputBuffer.Size() {
				for i := 0; i < numFlits; i++ {
					var flit = NewFlit(packet, i, i == 0, i == numFlits - 1)
					router.InsertFlit(flit, DIRECTION_LOCAL, ivc)
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
	flit.State = FLIT_STATE_INPUT_BUFFER
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