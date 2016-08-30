package uncore

type ControllerEvent interface {
	CacheCoherenceFlow
}

type BaseControllerEvent struct {
	*BaseCacheCoherenceFlow
}

func NewBaseControllerEvent(generator Controller, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *BaseControllerEvent {
	var baseControllerEvent = &BaseControllerEvent{
		BaseCacheCoherenceFlow:NewBaseCacheCoherenceFlow(generator, producerFlow, access, tag),
	}

	return baseControllerEvent
}