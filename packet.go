package acogo

type Packet interface {
	GetNetwork() *Network
	GetId() int
	GetBeginCycle() int
	GetEndCycle() int
	SetEndCycle(endCycle int)
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

type PacketMemoryEntry struct {
	NodeId    int
	Timestamp int
}
