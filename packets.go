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

type DataPacket struct {
	Network              *Network
	Id                   int
	BeginCycle, EndCycle int
	Src, Dest            int
	Size                 int
	OnCompletedCallback  func()
	Memory               []*PacketMemoryEntry
	Flits                []*Flit
	HasPayload           bool
}

func (packet *DataPacket) GetNetwork() *Network {
	return packet.Network
}

func (packet *DataPacket) GetId() int {
	return packet.Id
}

func (packet *DataPacket) GetBeginCycle() int {
	return packet.BeginCycle
}

func (packet *DataPacket) GetEndCycle() int {
	return packet.EndCycle
}

func (packet *DataPacket) SetEndCycle(endCycle int) {
	packet.EndCycle = endCycle
}

func (packet *DataPacket) GetSrc() int {
	return packet.Src
}

func (packet *DataPacket) GetDest() int {
	return packet.Dest
}

func (packet *DataPacket) GetSize() int {
	return packet.Size
}

func (packet *DataPacket) GetOnCompletedCallback() func() {
	return packet.OnCompletedCallback
}

func (packet *DataPacket) GetMemory() []*PacketMemoryEntry {
	return packet.Memory
}

func (packet *DataPacket) GetFlits() []*Flit {
	return packet.Flits
}

func (packet *DataPacket) SetFlits(flits []*Flit) {
	packet.Flits = flits
}

func (packet *DataPacket) GetHasPayload() bool {
	return packet.HasPayload
}

func NewDataPacket(network *Network, src int, dest int, size int, hasPayload bool, onCompletedCallback func()) *DataPacket {
	var packet = &DataPacket{
		Network:network,
		Src:src,
		Dest:dest,
		Size:size,
		OnCompletedCallback:onCompletedCallback,
		HasPayload:hasPayload,
	}

	return packet
}

func (packet *DataPacket) HandleDestArrived(inputVirtualChannel *InputVirtualChannel) {
	packet.Memorize(inputVirtualChannel.InputPort.Router.Node.Id)

	packet.EndCycle = inputVirtualChannel.InputPort.Router.Node.Network.Experiment.CycleAccurateEventQueue.CurrentCycle

	inputVirtualChannel.InputPort.Router.Node.Network.LogPacketTransmitted(packet)

	if packet.OnCompletedCallback != nil {
		packet.OnCompletedCallback()
	}
}

func (packet *DataPacket) DoRouteComputation(inputVirtualChannel *InputVirtualChannel) Direction {
	var parent = -1

	if len(packet.Memory) > 0 {
		parent = packet.Memory[len(packet.Memory) - 1].NodeId
	}

	packet.Memorize(inputVirtualChannel.InputPort.Router.Node.Id)

	var directions = inputVirtualChannel.InputPort.Router.Node.RoutingAlgorithm.NextHop(packet.Src, packet.Dest, parent)

	return inputVirtualChannel.InputPort.Router.Node.SelectionAlgorithm.Select(packet.Src, packet.Dest, inputVirtualChannel.Num, directions)
}

func (packet *DataPacket) Memorize(currentNodeId int) {
	packet.Memory = append(packet.Memory, &PacketMemoryEntry{
		NodeId:currentNodeId,
		Timestamp:packet.Network.Experiment.CycleAccurateEventQueue.CurrentCycle,
	})
}