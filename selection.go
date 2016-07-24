package acogo

type SelectionAlgorithm interface {
	Select(src int, dest int, ivc int, directions []Direction) Direction
}

type GeneralSelectionAlgorithm struct {
	Node *Node
}

func NewGeneralSelectionAlgorithm(node *Node) *GeneralSelectionAlgorithm {
	var generalSelectionAlgorithm = &GeneralSelectionAlgorithm{
		Node:node,
	}

	return generalSelectionAlgorithm
}

func (selectionAlgorithm *GeneralSelectionAlgorithm) HandleDestArrived(packet *Packet, inputVirtualChannel *InputVirtualChannel) {
	packet.Memorize(selectionAlgorithm.Node.Id)

	packet.EndCycle = selectionAlgorithm.Node.Network.Experiment.CycleAccurateEventQueue.CurrentCycle

	if packet.OnCompletedCallback != nil {
		packet.OnCompletedCallback()
	}
}

func (selectionAlgorithm *GeneralSelectionAlgorithm) DoRouteComputation(packet *Packet, inputVirtualChannel *InputVirtualChannel) Direction {
	var parent = -1

	if len(packet.Memory) > 0 {
		parent = packet.Memory[len(packet.Memory) - 1].NodeId
	}

	packet.Memorize(selectionAlgorithm.Node.Id)

	var directions = selectionAlgorithm.Node.RoutingAlgorithm.NextHop(packet.Src, packet.Dest, parent)

	return selectionAlgorithm.Select(packet.Src, packet.Dest, inputVirtualChannel.Num, directions)
}

func (selectionAlgorithm *GeneralSelectionAlgorithm) Select(src int, dest int, ivc int, directions []Direction) Direction {
	return Direction(-1)
}

type BufferLevelSelectionAlgorithm struct {
	*GeneralSelectionAlgorithm
}

func NewBufferLevelSelectionAlgorithm(node *Node) *BufferLevelSelectionAlgorithm {
	var bufferLevelSelectionAlgorithm = &BufferLevelSelectionAlgorithm{
		GeneralSelectionAlgorithm:NewGeneralSelectionAlgorithm(node),
	}

	return bufferLevelSelectionAlgorithm
}

func (selectionAlgorithm *BufferLevelSelectionAlgorithm) Select(src int, dest int, ivc int, directions []Direction) Direction {
	var bestDirections []Direction

	var maxFreeSlots = -1

	for i := 0; i < DirectionWest; i++ {
		var direction = Direction(i)
		var neighborRouter = selectionAlgorithm.Node.Network.Nodes[selectionAlgorithm.Node.Neighbors[direction]].Router
		var freeSlots = neighborRouter.FreeSlots(direction.GetReflexDirection(), ivc)

		if (freeSlots > maxFreeSlots) {
			maxFreeSlots = freeSlots;
			bestDirections = []Direction{direction}
		} else if (freeSlots == maxFreeSlots) {
			bestDirections = append(bestDirections, direction)
		}
	}

	if len(bestDirections) > 0 {
		return bestDirections[selectionAlgorithm.Node.Network.Experiment.rand.Intn(len(bestDirections))]
	}

	return selectionAlgorithm.GeneralSelectionAlgorithm.Select(src, dest, ivc, directions)
}