package acogo

type TransposeTrafficGenerator struct {
	network             *Network
	packetInjectionRate float32
	maxPackets          int
	hotspots            []int
	newPacket           func(src int, dest int) *Packet
}

func NewTransposeTrafficGenerator(network *Network, packetInjectionRate float32, maxPackets int, newPacket func(src int, dest int) *Packet) *TransposeTrafficGenerator {
	var gen = &TransposeTrafficGenerator{
		network:network,
		packetInjectionRate:packetInjectionRate,
		maxPackets:maxPackets,
		newPacket:newPacket,
	}

	network.experiment.cycleAccurateEventQueue.AddPerCycleEvent(gen.generateTraffic)

	return gen
}

func (gen *TransposeTrafficGenerator) generateTraffic() {
	for i := 0; i < len(gen.network.nodes); i++ {
		if !gen.network.acceptPacket || gen.maxPackets != -1 && gen.network.numPacketsReceived > gen.maxPackets {
			break
		}

		var valid = gen.network.experiment.rand.Float32() <= gen.packetInjectionRate
		if valid {
			var node = gen.network.nodes[i]
			var src = node.id
			var dest = gen.dest(src)

			if src != dest {
				gen.network.experiment.cycleAccurateEventQueue.Schedule(func() {
					gen.network.Receive(gen.newPacket(src, dest))
				}, 1)
			}
		}
	}
}

func (gen *TransposeTrafficGenerator) dest(src int) int {
	var srcX, srcY = gen.network.GetX(src), gen.network.GetY(src)
	var destX, destY = gen.network.width - 1 - srcY, gen.network.width - 1 - srcX

	return destY * gen.network.width + destX
}
