package acogo

type Pheromone struct {
	PheromoneTable *PheromoneTable
	Dest           int
	Direction      Direction
	Value          float64
}

func NewPheromone(pheromoneTable *PheromoneTable, dest int, direction Direction, value float64) *Pheromone {
	var pheromone = &Pheromone{
		PheromoneTable:pheromoneTable,
		Dest:dest,
		Direction:direction,
		Value:value,
	}

	return pheromone
}

type PheromoneTable struct {
	Node       *Node
	Pheromones map[int](map[Direction]*Pheromone)
}

func NewPheromoneTable(node *Node) *PheromoneTable {
	var pheromoneTable = &PheromoneTable{
		Node:node,
		Pheromones:make(map[int](map[Direction]*Pheromone)),
	}

	return pheromoneTable
}

func (pheromoneTable *PheromoneTable) Append(dest int, direction Direction, pheromoneValue float64) {
	var pheromone = NewPheromone(pheromoneTable, dest, direction, pheromoneValue)

	if _, exists := pheromoneTable.Pheromones[dest]; !exists {
		pheromoneTable.Pheromones[dest] = make(map[Direction]*Pheromone)
	}

	pheromoneTable.Pheromones[dest][direction] = pheromone
}

func (pheromoneTable *PheromoneTable) Update(dest int, direction Direction) {
	for _, pheromone := range pheromoneTable.Pheromones[dest] {
		if pheromone.Direction == direction {
			pheromone.Value += pheromoneTable.Node.Network.Experiment.Config.ReinforcementFactor * (1 - pheromone.Value)
		} else {
			pheromone.Value -= pheromoneTable.Node.Network.Experiment.Config.ReinforcementFactor * pheromone.Value
		}
	}
}

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
	var bestDirection = Direction(-1)

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
