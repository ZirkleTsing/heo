package acogo

type Packet interface {
	GetNetwork() *Network
	GetId() int64
	GetBeginCycle() int64
	GetEndCycle() int64
	SetEndCycle(endCycle int64)
	GetSrc() int
	GetDest() int
	GetSize() int
	GetOnCompletedCallback() func()
	GetMemory() []*PacketMemoryEntry
	GetFlits() []*Flit
	SetFlits(flits []*Flit)
	GetHasPayload() bool
	HandleDestArrived(inputVirtualChannel *InputVirtualChannel)
	DoRouteComputation(inputVirtualChannel *InputVirtualChannel) Direction
}

func Delay(packet Packet) int {
	if packet.GetEndCycle() == -1 {
		return -1
	} else {
		return int(packet.GetEndCycle() - packet.GetBeginCycle())
	}
}

func Hops(packet Packet) int {
	return len(packet.GetMemory())
}

type PacketMemoryEntry struct {
	NodeId    int
	Timestamp int64
}
