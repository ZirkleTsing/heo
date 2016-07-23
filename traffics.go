package acogo

type SyntheticTrafficGenerator struct {
	network             *Network
	packetInjectionRate float32
	packetSize          int
	maxPackets          int
	hotspots            []int
}

func newSyntheticTrafficGenerator(network *Network, packetInjectionRate float32, packetSize int, maxPackets int) *SyntheticTrafficGenerator {
	var gen = &SyntheticTrafficGenerator{
		network:network,
		packetInjectionRate:packetInjectionRate,
		packetSize:packetSize,
		maxPackets:maxPackets,
	}

	network.cycleAccurateEventQueue.AddPerCycleEvent(gen.generateTraffic)

	return gen
}

func (gen *SyntheticTrafficGenerator) generateTraffic() {
	for i:=0; i < len(gen.network.nodes); i++ {
		if !gen.network.acceptPacket || gen.maxPackets != -1 && gen.network.numPacketsReceived > gen.maxPackets {
			break
		}

		var valid = gen.network.random.nextDouble() <= gen.packetInjectionRate
		if valid {
			var node  = gen.network.nodes[i]
			var src = node.id
			var dest = gen.dest(src)

			if src != dest {
				gen.network.cycleAccurateEventQueue.Schedule(func() {
					gen.network.receive(gen.newPacket(src, dest, gen.packetSize))
				}, 1)
			}
		}
	}
}
