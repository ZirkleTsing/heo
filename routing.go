package acogo

import (
	"fmt"
)

type RoutingAlgorithm interface {
	NextHop(packet Packet, parent int) []Direction
}

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
			directions = append(directions, DirectionEast)
		} else {
			directions = append(directions, DirectionWest)
		}
	} else {
		if destY > y {
			directions = append(directions, DirectionSouth)
		} else {
			directions = append(directions, DirectionNorth)
		}
	}

	fmt.Printf("NextHop(packet#%d(src=%d, dest=%d), current=%d): %d\n", packet.GetId(), packet.GetSrc(), packet.GetDest(), routingAlgorithm.Node.Id, directions[0])

	return directions
}

type OddEvenRoutingAlgorithm struct {
	Node *Node
}

func NewOddEvenRoutingAlgorithm(node *Node) *OddEvenRoutingAlgorithm {
	var routingAlgorithm = &OddEvenRoutingAlgorithm{
		Node:node,
	}

	return routingAlgorithm
}

func (routingAlgorithm *OddEvenRoutingAlgorithm) NextHop(packet Packet, parent int) []Direction {
	var directions []Direction

	var c0 = routingAlgorithm.Node.X
	var c1 = routingAlgorithm.Node.Y

	var s0 = routingAlgorithm.Node.Network.GetX(packet.GetSrc())
	//var s1 = routingAlgorithm.Node.Network.GetY(packet.GetSrc())

	var d0 = routingAlgorithm.Node.Network.GetX(packet.GetDest())
	var d1 = routingAlgorithm.Node.Network.GetY(packet.GetDest())

	var e0 = d0 - c0
	var e1 = -(d1 - c1)

	if e0 == 0 {
		if e1 > 0 {
			directions = append(directions, DirectionNorth)
		} else {
			directions = append(directions, DirectionSouth)
		}
	} else {
		if e0 > 0 {
			if e1 == 0 {
				directions = append(directions, DirectionEast)
			} else {
				if c0 % 2 == 1 || c0 == s0 {
					if e1 > 0 {
						directions = append(directions, DirectionNorth)
					} else {
						directions = append(directions, DirectionSouth)
					}
				}

				if d0 % 2 == 1 || e0 != 1 {
					directions = append(directions, DirectionEast)
				}
			}
		} else {
			directions = append(directions, DirectionWest)
			if c0 % 2 == 0 {
				if e1 > 0 {
					directions = append(directions, DirectionNorth)
				}
				if e1 < 0 {
					directions = append(directions, DirectionSouth)
				}
			}
		}
	}

	return directions
}
