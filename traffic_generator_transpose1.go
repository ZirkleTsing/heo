package acogo

type Transpose1TrafficGenerator struct {
	*GeneralTrafficGenerator
}

func NewTranspose1TrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int64, newPacket func(src int, dest int) Packet) *Transpose1TrafficGenerator {
	var generator = &Transpose1TrafficGenerator{
		GeneralTrafficGenerator: NewGeneralTrafficGenerator(network, packetInjectionRate, maxPackets, newPacket),
	}

	return generator
}

func (generator *Transpose1TrafficGenerator) AdvanceOneCycle() {
	generator.GeneralTrafficGenerator.AdvanceOneCycle(func(src int) int {
		var srcX, srcY = generator.Network.GetX(src), generator.Network.GetY(src)
		var destX, destY = generator.Network.Width - 1 - srcY, generator.Network.Width - 1 - srcX

		return destY * generator.Network.Width + destX
	})
}
