package acogo

type TransposeTrafficGenerator struct {
	Network             *Network
	PacketInjectionRate float64
	MaxPackets          int
	NewPacket           func(src int, dest int) Packet
}

func NewTransposeTrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int, newPacket func(src int, dest int) Packet) *TransposeTrafficGenerator {
	var generator = &TransposeTrafficGenerator{
		Network:network,
		PacketInjectionRate:packetInjectionRate,
		MaxPackets:maxPackets,
		NewPacket:newPacket,
	}

	network.Experiment.CycleAccurateEventQueue.AddPerCycleEvent(generator.GenerateTraffic)

	return generator
}

func (generator *TransposeTrafficGenerator) GenerateTraffic() {
	for _, node := range generator.Network.Nodes {
		if !generator.Network.AcceptPacket || generator.MaxPackets != -1 && generator.Network.NumPacketsReceived > generator.MaxPackets {
			break
		}

		var valid = generator.Network.Experiment.rand.Float64() <= generator.PacketInjectionRate
		if valid {
			var src = node.Id
			var dest = generator.dest(src)

			if src != dest {
				generator.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
					generator.Network.Receive(generator.NewPacket(src, dest))
				}, 1)
			}
		}
	}
}

func (gen *TransposeTrafficGenerator) dest(src int) int {
	var srcX, srcY = gen.Network.GetX(src), gen.Network.GetY(src)
	var destX, destY = gen.Network.Width - 1 - srcY, gen.Network.Width - 1 - srcX

	return destY * gen.Network.Width + destX
}
