package uncore

import (
	"github.com/mcai/acogo/noc"
	"math"
)

type NoCMemoryHierarchy struct {
	*MemoryHierarchy

	Network          *noc.Network
	DevicesToNodeIds map[interface{}]uint32
}

func NewNocMemoryHierarchy(driver interface{}, uncoreConfig *UncoreConfig, nocConfig *noc.NoCConfig) *NoCMemoryHierarchy {
	var nocMemoryHierarchy = &NoCMemoryHierarchy{
		MemoryHierarchy:NewMemoryHierarchy(driver.(UncoreDriver), uncoreConfig),
		DevicesToNodeIds:make(map[interface{}]uint32),
	}

	var numNodes = uint32(0)

	for i, l1IController := range nocMemoryHierarchy.L1IControllers {
		nocMemoryHierarchy.DevicesToNodeIds[l1IController] = numNodes

		var l1DController = nocMemoryHierarchy.L1DControllers[i]

		nocMemoryHierarchy.DevicesToNodeIds[l1DController] = numNodes

		numNodes++
	}

	nocMemoryHierarchy.DevicesToNodeIds[nocMemoryHierarchy.L2Controller] = numNodes

	numNodes++

	nocMemoryHierarchy.DevicesToNodeIds[nocMemoryHierarchy.MemoryController] = numNodes

	numNodes++

	var width = uint32(math.Sqrt(float64(numNodes)))

	if width * width != numNodes {
		numNodes = (width + 1) * (width + 1)
	}

	nocConfig.NumNodes = int(numNodes)

	nocMemoryHierarchy.Network = noc.NewNetwork(driver.(noc.NetworkDriver), nocConfig)

	return nocMemoryHierarchy
}

func (nocMemoryHierarchy *NoCMemoryHierarchy) Transfer(from MemoryDevice, to MemoryDevice, size uint32, onCompletedCallback func()) {
	var src = nocMemoryHierarchy.DevicesToNodeIds[from]
	var dest = nocMemoryHierarchy.DevicesToNodeIds[to]

	var packet = noc.NewDataPacket(nocMemoryHierarchy.Network, int(src), int(dest), int(size), true, onCompletedCallback)

	nocMemoryHierarchy.Driver.CycleAccurateEventQueue().Schedule(
		func() {
			nocMemoryHierarchy.Network.Receive(packet)
		},
		1,
	)
}