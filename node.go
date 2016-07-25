package acogo

import (
	"fmt"
)

type Node struct {
	Network            *Network
	Id                 int
	X, Y               int
	Neighbors          map[Direction]int
	Router             *Router
	RoutingAlgorithm   RoutingAlgorithm
	SelectionAlgorithm SelectionAlgorithm
}

func NewNode(network *Network, id int) *Node {
	var node = &Node{
		Network:network,
		Id:id,
		X:network.GetX(id),
		Y:network.GetY(id),
		Neighbors:make(map[Direction]int),
	}

	if (id / network.Width > 0) {
		node.Neighbors[DIRECTION_NORTH] = id - network.Width
	}

	if ( (id % network.Width) != network.Width - 1) {
		node.Neighbors[DIRECTION_EAST] = id + 1
	}

	if (id / network.Width < network.Width - 1) {
		node.Neighbors[DIRECTION_SOUTH] = id + network.Width
	}

	if (id % network.Width != 0) {
		node.Neighbors[DIRECTION_WEST] = id - 1
	}

	node.Router = NewRouter(node)

	//TODO
	node.RoutingAlgorithm = NewOddEvenRoutingAlgorithm(node)
	//node.RoutingAlgorithm = NewXYRoutingAlgorithm(node)

	node.SelectionAlgorithm = NewACOSelectionAlgorithm(node)
	//node.SelectionAlgorithm = NewBufferLevelSelectionAlgorithm(node)

	return node
}

func (node *Node) DumpNeighbors() {
	for direction, neighbor := range node.Neighbors {
		fmt.Printf("node#%d.neighbors[%s]=%d\n", node.Id, direction, neighbor)
	}

	fmt.Println()
}
