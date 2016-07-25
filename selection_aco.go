package acogo

type ACOSelectionAlgorithm struct {
	Node *Node
	PheromoneTable *PheromoneTable
}

func NewACOSelectionAlgorithm(node *Node) *ACOSelectionAlgorithm {
	var selectionAlgorithm = &ACOSelectionAlgorithm{
		Node:node,
		PheromoneTable:NewPheromoneTable(node),
	}

	var pheromoneValue = 1.0 / float64(len(node.Neighbors))

	for dest := 0; dest < node.Network.NumNodes; dest++ {
		if node.Id != dest {
			for i := range node.Neighbors {
				var direction = Direction(i)
				selectionAlgorithm.PheromoneTable.Append(dest, direction, pheromoneValue)
			}
		}
	}

	return selectionAlgorithm
}

func (selectionAlgorithm *ACOSelectionAlgorithm) CreateAndSendBackwardAntPacket(packet *AntPacket) {
	var newPacket = NewAntPacket(packet.Network, packet.Dest, packet.Src, selectionAlgorithm.Node.Network.Experiment.Config.AntPacketSize, func() {}, false)

	newPacket.Memory = packet.Memory

	selectionAlgorithm.Node.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
		selectionAlgorithm.Node.Network.Receive(newPacket)
	}, 1)
}

func (selectionAlgorithm *ACOSelectionAlgorithm) BackwardAntPacket(packet *AntPacket) Direction {
	var i int

	for i = len(packet.Memory) - 1; i > 0; i-- {
		var entry = packet.Memory[i]
		if entry.NodeId == selectionAlgorithm.Node.Id {
			break
		}
	}

	var prev = packet.Memory[i - 1].NodeId

	for direction, neighbor := range selectionAlgorithm.Node.Neighbors {
		if neighbor == prev {
			return direction
		}
	}

	panic("Impossible")
}

func (selectionAlgorithm *ACOSelectionAlgorithm) UpdatePheromoneTable(packet *AntPacket, inputVirtualChannel *InputVirtualChannel) {
	var i int
	for i = 0; i < len(packet.Memory); i++ {
		var entry = packet.Memory[i]
		if entry.NodeId == selectionAlgorithm.Node.Id {
			break
		}
	}

	for j := i + 1; j < len(packet.Memory); j++ {
		var entry = packet.Memory[j]
		var dest = entry.NodeId
		selectionAlgorithm.PheromoneTable.Update(dest, inputVirtualChannel.InputPort.Direction)
	}
}

func (selectionAlgorithm *ACOSelectionAlgorithm) Select(packet Packet, ivc int, directions []Direction) Direction {
	var maxProbability = -1.0
	var bestDirection = DIRECTION_UNKNOWN

	for direction, neighbor := range selectionAlgorithm.Node.Neighbors {
		var neighborRouter = selectionAlgorithm.Node.Network.Nodes[neighbor].Router
		var pheromone = selectionAlgorithm.PheromoneTable.Pheromones[packet.GetDest()][direction]
		var freeSlots = neighborRouter.FreeSlots(direction, ivc)

		var alpha = selectionAlgorithm.Node.Network.Experiment.Config.AcoSelectionAlpha
		var qTotal = selectionAlgorithm.Node.Network.Experiment.Config.MaxInputBufferSize
		var n = len(selectionAlgorithm.Node.Neighbors)

		var probability = (pheromone.Value + alpha * (float64(freeSlots) / float64(qTotal))) / (1 + alpha * float64(n - 1))
		if probability > maxProbability {
			maxProbability = probability
			bestDirection = direction
		}
	}

	return bestDirection
}
