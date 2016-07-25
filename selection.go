package acogo

type SelectionAlgorithm interface {
	Select(packet Packet, ivc int, directions []Direction) Direction
}

type BufferLevelSelectionAlgorithm struct {
	Node *Node
}

func NewBufferLevelSelectionAlgorithm(node *Node) *BufferLevelSelectionAlgorithm {
	var selectionAlgorithm = &BufferLevelSelectionAlgorithm{
		Node:node,
	}

	return selectionAlgorithm
}

func (selectionAlgorithm *BufferLevelSelectionAlgorithm) Select(packet Packet, ivc int, directions []Direction) Direction {
	var bestDirections []Direction

	var maxFreeSlots = -1

	for direction, neighbor := range selectionAlgorithm.Node.Neighbors {
		var neighborRouter = selectionAlgorithm.Node.Network.Nodes[neighbor].Router
		var freeSlots = neighborRouter.FreeSlots(direction.GetReflexDirection(), ivc)

		if freeSlots > maxFreeSlots {
			maxFreeSlots = freeSlots
			bestDirections = []Direction{direction}
		} else if freeSlots == maxFreeSlots {
			bestDirections = append(bestDirections, direction)
		}
	}

	if len(bestDirections) > 0 {
		return bestDirections[selectionAlgorithm.Node.Network.Experiment.rand.Intn(len(bestDirections))]
	}

	return directions[0]
}