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
