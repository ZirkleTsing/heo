package acogo

type AntPacket struct {
	Forward bool
	*DataPacket
}

func NewAntPacket(network *Network, src int, dest int, size int, onCompletedCallback func(), forward bool) *AntPacket {
	var packet = &AntPacket{
		Forward:forward,
		DataPacket:NewDataPacket(network, src, dest, size, forward, onCompletedCallback),
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
		return packet.DataPacket.DoRouteComputation(inputVirtualChannel)
	} else {
		if inputVirtualChannel.InputPort.Router.Node.Id != packet.Src {
			selectionAlgorithm.UpdatePheromoneTable(packet, inputVirtualChannel)
		}

		return selectionAlgorithm.BackwardAntPacket(packet)
	}
}
