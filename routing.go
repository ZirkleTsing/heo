package acogo

type RoutingAlgorithm interface {
	NextHop(packet Packet, parent int) []Direction
}