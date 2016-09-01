package uncore

import "github.com/mcai/acogo/noc"

type NoCMemoryHierarchy struct {
	*MemoryHierarchy

	Network *noc.Network
	DevicesToNodeIds map[interface{}]uint32
}

func NewNocMemoryHierarchy(driver MemoryHierarchyDriver, config *UncoreConfig) *NoCMemoryHierarchy {
	var nocMemoryHierarchy = &NoCMemoryHierarchy{
		MemoryHierarchy:NewMemoryHierarchy(driver, config),
	}

	return nocMemoryHierarchy
}