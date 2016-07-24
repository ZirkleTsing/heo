package acogo

type AntPacket struct {
	Forward bool
	*Packet
}

func NewAntPacket(network *Network, src int, dest int, size int, onCompletedCallback func(), forward bool) *AntPacket {
	var packet = &AntPacket{
		Forward:forward,
		Packet:NewPacket(network, src, dest, size, forward, onCompletedCallback),
	}

	return packet
}

func (packet *AntPacket) HandleDestArrived(inputVirtualChannel *InputVirtualChannel) {
	var selectionAlgorithm = inputVirtualChannel.InputPort.Router.Node.SelectionAlgorithm.(*ACOSelectionAlgorithm)

	if packet.Forward {
		packet.Memorize(inputVirtualChannel.InputPort.Router.Node.Id)
		selectionAlgorithm.CreateAndSendBackwardAntPacket(packet)
	} else {
		selectionAlgorithm.UpdatePheromoneTable(packet, inputVirtualChannel)
	}

	packet.EndCycle = inputVirtualChannel.InputPort.Router.Node.Network.Experiment.CycleAccurateEventQueue.CurrentCycle

	if packet.OnCompletedCallback != nil {
		packet.OnCompletedCallback()
	}
}

func (packet *AntPacket) DoRouteComputation(inputVirtualChannel *InputVirtualChannel) Direction {
	var selectionAlgorithm = inputVirtualChannel.InputPort.Router.Node.SelectionAlgorithm.(*ACOSelectionAlgorithm)

	if packet.Forward {
		return packet.Packet.DoRouteComputation(inputVirtualChannel)
	} else {
		if inputVirtualChannel.InputPort.Router.Node.Id != packet.Src {
			selectionAlgorithm.UpdatePheromoneTable(packet, inputVirtualChannel)
		}

		return selectionAlgorithm.BackwardAntPacket(packet)
	}
}

type Pheromone struct {
	PheromoneTable *PheromoneTable
	Dest           int
	Direction      Direction
	Value          float32
}

func NewPheromone(pheromoneTable *PheromoneTable, dest int, direction Direction, value float32) *Pheromone {
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

func (pheromoneTable *PheromoneTable) Append(dest int, direction Direction, pheromoneValue float32) {
	var pheromone = NewPheromone(pheromoneTable, dest, direction, pheromoneValue)
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
	PheromoneTable *PheromoneTable
	*GeneralSelectionAlgorithm
}

func NewACOSelectionAlgorithm(node *Node) *ACOSelectionAlgorithm {
	var acoSelectionAlgorithm = &ACOSelectionAlgorithm{
		PheromoneTable:NewPheromoneTable(node),
		GeneralSelectionAlgorithm:NewGeneralSelectionAlgorithm(node),
	}

	var pheromoneValue = 1.0 / float32(len(node.Neighbors))

	for dest := 0; dest < node.Network.NumNodes; dest++ {
		if (node.Id != dest) {
			for i := range node.Neighbors {
				var direction = Direction(i)
				acoSelectionAlgorithm.PheromoneTable.Append(dest, direction, pheromoneValue)
			}
		}
	}

	return acoSelectionAlgorithm
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

	return Direction(-1)
}

func (selectionAlgorithm *ACOSelectionAlgorithm) UpdatePheromoneTable(packet *AntPacket, inputVirtualChannel *InputVirtualChannel) {
	var i int
	for i = 0; i < len(packet.Memory); i++ {
		var entry = packet.Memory[i]
		if entry.NodeId == selectionAlgorithm.Node.Id {
			break
		}
	}

	for j := i + 1; j < len(packet.Memory); i++ {
		var entry = packet.Memory[j]
		var dest = entry.NodeId
		selectionAlgorithm.PheromoneTable.Update(dest, inputVirtualChannel.InputPort.Direction)
	}
}

func (selectionAlgorithm *ACOSelectionAlgorithm) Select(src int, dest int, ivc int, directions []Direction) Direction {
	var maxProbability = float32(-1.0)
	var bestDirection = Direction(-1)

	for i := 0; i < DirectionWest; i++ {
		var direction = Direction(i)
		var pheromone = selectionAlgorithm.PheromoneTable.Pheromones[dest][direction]
		var neighborRouter = selectionAlgorithm.Node.Network.Nodes[selectionAlgorithm.Node.Neighbors[direction]].Router
		var freeSlots = neighborRouter.FreeSlots(direction, ivc)

		var alpha = selectionAlgorithm.Node.Network.Experiment.Config.AcoSelectionAlpha
		var qTotal = selectionAlgorithm.Node.Network.Experiment.Config.MaxInputBufferSize
		var n = len(selectionAlgorithm.Node.Neighbors)

		var probability = (pheromone.Value + alpha * (float32(freeSlots) / float32(qTotal))) / (1 + alpha * float32(n - 1));
		if probability > maxProbability {
			maxProbability = probability
			bestDirection = direction
		}
	}

	return bestDirection
}
