package acogo

type TransposeTrafficGenerator struct {
	*GeneralTrafficGenerator
}

func NewTransposeTrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int64, newPacket func(src int, dest int) Packet) *TransposeTrafficGenerator {
	var generator = &TransposeTrafficGenerator{
		GeneralTrafficGenerator: NewGeneralTrafficGenerator(network, packetInjectionRate, maxPackets, newPacket),
	}

	return generator
}

func (generator *TransposeTrafficGenerator) AdvanceOneCycle() {
	generator.GeneralTrafficGenerator.AdvanceOneCycle(func(src int) int {
		var srcX, srcY = generator.Network.GetX(src), generator.Network.GetY(src)
		var destX, destY = generator.Network.Width - 1 - srcY, generator.Network.Width - 1 - srcX

		return destY * generator.Network.Width + destX
	})
}
