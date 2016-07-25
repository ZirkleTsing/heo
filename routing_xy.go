package acogo

import (
	"fmt"
)

type XYRoutingAlgorithm struct {
	Node *Node
}

func NewXYRoutingAlgorithm(node *Node) *XYRoutingAlgorithm {
	var routingAlgorithm = &XYRoutingAlgorithm{
		Node:node,
	}

	return routingAlgorithm
}

func (routingAlgorithm *XYRoutingAlgorithm) NextHop(packet Packet, parent int) []Direction {
	var directions []Direction

	var destX = routingAlgorithm.Node.Network.GetX(packet.GetDest())
	var destY = routingAlgorithm.Node.Network.GetY(packet.GetDest())

	var x = routingAlgorithm.Node.X
	var y = routingAlgorithm.Node.Y

	if destX != x {
		if destX > x {
			directions = append(directions, DIRECTION_EAST)
		} else {
			directions = append(directions, DIRECTION_WEST)
		}
	} else {
		if destY > y {
			directions = append(directions, DIRECTION_SOUTH)
		} else {
			directions = append(directions, DIRECTION_NORTH)
		}
	}

	fmt.Printf("NextHop(packet#%d(src=%d, dest=%d), current=%d) = %s\n", packet.GetId(), packet.GetSrc(), packet.GetDest(), routingAlgorithm.Node.Id, directions[0])

	return directions
}
