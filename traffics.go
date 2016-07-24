package acogo

type TransposeTrafficGenerator struct {
	Network             *Network
	PacketInjectionRate float32
	MaxPackets          int
	Hotspots            []int //TODO
	NewPacket           func(src int, dest int) Packet
}

func NewTransposeTrafficGenerator(network *Network, packetInjectionRate float32, maxPackets int, newPacket func(src int, dest int) Packet) *TransposeTrafficGenerator {
	var gen = &TransposeTrafficGenerator{
		Network:network,
		PacketInjectionRate:packetInjectionRate,
		MaxPackets:maxPackets,
		NewPacket:newPacket,
	}

	network.Experiment.CycleAccurateEventQueue.AddPerCycleEvent(gen.generateTraffic)

	return gen
}

func (gen *TransposeTrafficGenerator) generateTraffic() {
	for i := 0; i < len(gen.Network.Nodes); i++ {
		if !gen.Network.AcceptPacket || gen.MaxPackets != -1 && gen.Network.NumPacketsReceived > gen.MaxPackets {
			break
		}

		var valid = gen.Network.Experiment.rand.Float32() <= gen.PacketInjectionRate
		if valid {
			var node = gen.Network.Nodes[i]
			var src = node.Id
			var dest = gen.dest(src)

			if src != dest {
				gen.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
					gen.Network.Receive(gen.NewPacket(src, dest))
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
