package acogo

type RoutingAlgorithm interface {
	NextHop(src int, dest int, parent int) []Direction
}

type XYRoutingAlgorithm struct {
	Node *Node
}

func NewXYRoutingAlgorithm(node *Node) *XYRoutingAlgorithm {
	var xyRoutingAlgorithm = &XYRoutingAlgorithm{
		Node:node,
	}

	return xyRoutingAlgorithm
}

func (routingAlgorithm *XYRoutingAlgorithm) NextHop(src int, dest int, parent int) []Direction {
	var directions []Direction

	var destX = routingAlgorithm.Node.Network.GetX(dest)
	var destY = routingAlgorithm.Node.Network.GetY(dest)

	if destX != routingAlgorithm.Node.X {
		if destX > routingAlgorithm.Node.X {
			directions = append(directions, DirectionEast)
		} else {
			directions = append(directions, DirectionWest)
		}
	} else {
		if destY > routingAlgorithm.Node.Y {
			directions = append(directions, DirectionSouth)
		} else {
			directions = append(directions, DirectionNorth)
		}
	}

	return directions
}

type OddEvenRoutingAlgorithm struct {
	Node *Node
}

func NewOddEvenRoutingAlgorithm(node *Node) *OddEvenRoutingAlgorithm {
	var xyRoutingAlgorithm = &OddEvenRoutingAlgorithm{
		Node:node,
	}

	return xyRoutingAlgorithm
}

func (routingAlgorithm *OddEvenRoutingAlgorithm) NextHop(src int, dest int, parent int) []Direction {
	var directions []Direction

	var c0 = routingAlgorithm.Node.X
	var c1 = routingAlgorithm.Node.Y

	var s0 = routingAlgorithm.Node.Network.GetX(src)
	//var s1 = routingAlgorithm.Node.Network.GetY(src)

	var d0 = routingAlgorithm.Node.Network.GetX(dest)
	var d1 = routingAlgorithm.Node.Network.GetY(dest)

	var e0 = d0 - c0;
	var e1 = -(d1 - c1);

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
