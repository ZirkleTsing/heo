package acogo

type Flit struct {
	Packet    Packet
	Num       int
	Head      bool
	Tail      bool
	Node      *Node
	State     FlitState
	Timestamp int64
}

func NewFlit(packet Packet, num int, head bool, tail bool) *Flit {
	var flit = &Flit{
		Packet:packet,
		Num:num,
		Head:head,
		Tail:tail,
		State:FLIT_STATE_UNKNOWN,
		Timestamp: packet.GetNetwork().Experiment.CycleAccurateEventQueue.CurrentCycle,
	}

	packet.SetFlits(append(packet.GetFlits(), flit))

	return flit
}
