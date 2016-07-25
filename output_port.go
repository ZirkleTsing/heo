package acogo

type OutputPort struct {
	Router          *Router
	Direction       Direction
	VirtualChannels []*OutputVirtualChannel
}

func NewOutputPort(router *Router, direction Direction) *OutputPort {
	var outputPort = &OutputPort{
		Router:router,
		Direction:direction,
	}

	for i := 0; i < router.Node.Network.Experiment.Config.NumVirtualChannels; i++ {
		var outputVirtualChannel = NewOutputVirtualChannel(outputPort, i)
		outputPort.VirtualChannels = append(outputPort.VirtualChannels, outputVirtualChannel)
	}

	return outputPort
}
