package noc

type Transpose2TrafficGenerator struct {
	*GeneralTrafficGenerator
}

func NewTranspose2TrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int64, newPacket func(src int, dest int) Packet) *Transpose2TrafficGenerator {
	var generator = &Transpose2TrafficGenerator{
		GeneralTrafficGenerator: NewGeneralTrafficGenerator(network, packetInjectionRate, maxPackets, newPacket),
	}

	return generator
}

func (generator *Transpose2TrafficGenerator) AdvanceOneCycle() {
	generator.GeneralTrafficGenerator.AdvanceOneCycle(func(src int) int {
		var srcX, srcY = generator.Network.GetX(src), generator.Network.GetY(src)
		var destX, destY = srcY, srcX

		return destY * generator.Network.Width + destX
	})
}
