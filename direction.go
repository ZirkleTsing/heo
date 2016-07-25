package acogo

import "fmt"

type Direction int

const (
	DirectionLocal = 0
	DirectionNorth = 1
	DirectionEast = 2
	DirectionSouth = 3
	DirectionWest = 4
)

func (direction Direction) GetReflexDirection() Direction {
	switch direction {
	case DirectionLocal:
		return DirectionLocal
	case DirectionNorth:
		return DirectionSouth
	case DirectionEast:
		return DirectionWest
	case DirectionSouth:
		return DirectionNorth
	case DirectionWest:
		return DirectionEast
	default:
		panic(fmt.Sprintf("%d", direction))
	}
}
