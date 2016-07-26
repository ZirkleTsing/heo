package acogo

import "math/rand"

type GeneralTrafficGenerator struct {
	Network             *Network
	PacketInjectionRate float64
	MaxPackets          int64
	NewPacket           func(src int, dest int) Packet
}

func NewGeneralTrafficGenerator(network *Network, packetInjectionRate float64, maxPackets int64, newPacket func(src int, dest int) Packet) *GeneralTrafficGenerator {
	var generator = &GeneralTrafficGenerator{
		Network:network,
		PacketInjectionRate:packetInjectionRate,
		MaxPackets:maxPackets,
		NewPacket:newPacket,
	}

	return generator
}

func (generator *GeneralTrafficGenerator) AdvanceOneCycle(dest func(src int) int) {
	for _, node := range generator.Network.Nodes {
		if !generator.Network.AcceptPacket || generator.MaxPackets != -1 && generator.Network.NumPacketsReceived > generator.MaxPackets {
			break
		}

		var valid = rand.Float64() <= generator.PacketInjectionRate
		if valid {
			var src = node.Id
			var dest = dest(src)

			if src != dest {
				generator.Network.Experiment.CycleAccurateEventQueue.Schedule(func() {
					generator.Network.Receive(generator.NewPacket(src, dest))
				}, 1)
			}
		}
	}
}
