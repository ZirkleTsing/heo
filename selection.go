package acogo

type SelectionAlgorithm interface {
	Select(src int, dest int, ivc int, directions []Direction) Direction
}

type BufferLevelSelectionAlgorithm struct {
	Node *Node
}

func NewBufferLevelSelectionAlgorithm(node *Node) *BufferLevelSelectionAlgorithm {
	var bufferLevelSelectionAlgorithm = &BufferLevelSelectionAlgorithm{
		Node:node,
	}

	return bufferLevelSelectionAlgorithm
}

func (selectionAlgorithm *BufferLevelSelectionAlgorithm) Select(src int, dest int, ivc int, directions []Direction) Direction {
	var bestDirections []Direction

	var maxFreeSlots = -1

	for i := 0; i < NumDirections; i++ {
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

	return directions[0]
}