package noc

import (
	"fmt"
	"math"
)

type DataPacket struct {
	Network              *Network
	Id                   int64
	BeginCycle, EndCycle int64
	Src, Dest            int
	Size                 int
	OnCompletedCallback  func()
	Memory               []*PacketMemoryEntry
	Flits                []*Flit
	HasPayload           bool
}

func NewDataPacket(network *Network, src int, dest int, size int, hasPayload bool, onCompletedCallback func()) *DataPacket {
	var packet = &DataPacket{
		Network:network,
		Id:network.CurrentPacketId,
		BeginCycle:network.Experiment.CycleAccurateEventQueue.CurrentCycle,
		EndCycle:-1,
		Src:src,
		Dest:dest,
		Size:size,
		OnCompletedCallback:onCompletedCallback,
		HasPayload:hasPayload,
	}

	network.CurrentPacketId++

	var numFlits = int(math.Ceil(float64(packet.Size) / float64(network.Experiment.Config.LinkWidth)))
	if numFlits > network.Experiment.Config.MaxInputBufferSize {
		panic(fmt.Sprintf("Number of flits (%d) in a packet cannot be greater than max input buffer size (%d)", numFlits, network.Experiment.Config.MaxInputBufferSize))
	}

	return packet
}

func (packet *DataPacket) GetNetwork() *Network {
	return packet.Network
}

func (packet *DataPacket) GetId() int64 {
	return packet.Id
}

func (packet *DataPacket) GetBeginCycle() int64 {
	return packet.BeginCycle
}

func (packet *DataPacket) GetEndCycle() int64 {
	return packet.EndCycle
}

func (packet *DataPacket) SetEndCycle(endCycle int64) {
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

func (packet *DataPacket) HandleDestArrived(inputVirtualChannel *InputVirtualChannel) {
	packet.Memorize(inputVirtualChannel.InputPort.Router.Node)

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

	packet.Memorize(inputVirtualChannel.InputPort.Router.Node)

	var directions = inputVirtualChannel.InputPort.Router.Node.RoutingAlgorithm.NextHop(packet, parent)

	return inputVirtualChannel.InputPort.Router.Node.SelectionAlgorithm.Select(packet, inputVirtualChannel.Num, directions)
}

func (packet *DataPacket) Memorize(node *Node) {
	for _, entry := range packet.Memory {
		if entry.NodeId == node.Id {
			panic(fmt.Sprintf("packet#%d(src=%d, dest=%d): %d", packet.Id, packet.Src, packet.Dest, node.Id))
		}
	}

	packet.Memory = append(packet.Memory, &PacketMemoryEntry{
		NodeId:node.Id,
		Timestamp:packet.Network.Experiment.CycleAccurateEventQueue.CurrentCycle,
	})
}

func (packet *DataPacket) DumpMemory() {
	for i, entry := range packet.Memory {
		fmt.Printf("packet#%d.memory[%d]=%d\n", packet.Id, i, entry.NodeId)
	}

	fmt.Println()
}
