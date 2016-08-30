package uncore

type CacheCoherenceFlow interface {
	Id() int32
	Generator() Controller
	ProducerFlow() CacheCoherenceFlow
	AncestorFlow() CacheCoherenceFlow
	ChildFlows() []CacheCoherenceFlow
	SetChildFlows(childFlows []CacheCoherenceFlow)
	NumPendingDescendantFlows() int32
	SetNumPendingDescendantFlows(numPendingDescendantFlows int32)
	BeginCycle() int64
	EndCycle() int64
	Completed() bool
	Access() *MemoryHierarchyAccess
	Tag() uint32
	Complete()
}

type BaseCacheCoherenceFlow struct {
	id                        int32
	generator                 Controller
	producerFlow              CacheCoherenceFlow
	ancestorFlow              CacheCoherenceFlow
	childFlows                []CacheCoherenceFlow
	numPendingDescendantFlows int32
	beginCycle                int64
	endCycle                  int64
	completed                 bool
	access                    *MemoryHierarchyAccess
	tag                       uint32
}

func NewBaseCacheCoherenceFlow(generator Controller, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *BaseCacheCoherenceFlow {
	var baseCacheCoherenceFlow = &BaseCacheCoherenceFlow{
		id:generator.GetMemoryDevice().MemoryHierarchy.CurrentCacheCoherenceFlowId,
		generator:generator,
		producerFlow:producerFlow,
		access:access,
		tag:tag,
	}

	generator.GetMemoryDevice().MemoryHierarchy.CurrentCacheCoherenceFlowId++

	baseCacheCoherenceFlow.beginCycle = generator.GetMemoryDevice().MemoryHierarchy.Driver.CycleAccurateEventQueue().CurrentCycle

	if producerFlow == nil {
		baseCacheCoherenceFlow.ancestorFlow = baseCacheCoherenceFlow
		generator.GetMemoryDevice().MemoryHierarchy.PendingFlows = append(generator.GetMemoryDevice().MemoryHierarchy.PendingFlows, baseCacheCoherenceFlow)
	} else {
		baseCacheCoherenceFlow.ancestorFlow = producerFlow.AncestorFlow()
		producerFlow.SetChildFlows(append(producerFlow.ChildFlows(), baseCacheCoherenceFlow))
	}

	baseCacheCoherenceFlow.ancestorFlow.SetNumPendingDescendantFlows(
		baseCacheCoherenceFlow.ancestorFlow.NumPendingDescendantFlows() + 1)

	return baseCacheCoherenceFlow
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Id() int32 {
	return cacheCoherenceFlow.id
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Generator() Controller {
	return cacheCoherenceFlow.generator
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) ProducerFlow() CacheCoherenceFlow {
	return cacheCoherenceFlow.producerFlow
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) AncestorFlow() CacheCoherenceFlow {
	return cacheCoherenceFlow.ancestorFlow
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) ChildFlows() []CacheCoherenceFlow {
	return cacheCoherenceFlow.childFlows
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) SetChildFlows(childFlows []CacheCoherenceFlow) {
	cacheCoherenceFlow.childFlows = childFlows
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) NumPendingDescendantFlows() int32 {
	return cacheCoherenceFlow.numPendingDescendantFlows
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) SetNumPendingDescendantFlows(numPendingDescendantFlows int32) {
	cacheCoherenceFlow.numPendingDescendantFlows = numPendingDescendantFlows
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) BeginCycle() int64 {
	return cacheCoherenceFlow.beginCycle
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) EndCycle() int64 {
	return cacheCoherenceFlow.endCycle
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Completed() bool {
	return cacheCoherenceFlow.completed
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Access() *MemoryHierarchyAccess {
	return cacheCoherenceFlow.access
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Tag() uint32 {
	return cacheCoherenceFlow.tag
}

func (cacheCoherenceFlow *BaseCacheCoherenceFlow) Complete() {
	cacheCoherenceFlow.completed = true
	cacheCoherenceFlow.endCycle = cacheCoherenceFlow.generator.GetMemoryDevice().MemoryHierarchy.Driver.CycleAccurateEventQueue().CurrentCycle
	cacheCoherenceFlow.ancestorFlow.SetNumPendingDescendantFlows(
		cacheCoherenceFlow.ancestorFlow.NumPendingDescendantFlows() - 1)

	if cacheCoherenceFlow.ancestorFlow.NumPendingDescendantFlows() == 0 {
		var pendingFlowsToReserve []CacheCoherenceFlow

		for _, pendingFlow := range cacheCoherenceFlow.generator.GetMemoryDevice().MemoryHierarchy.PendingFlows {
			if pendingFlow != cacheCoherenceFlow.ancestorFlow {
				pendingFlowsToReserve = append(pendingFlowsToReserve, pendingFlow)
			}
		}

		cacheCoherenceFlow.generator.GetMemoryDevice().MemoryHierarchy.PendingFlows = pendingFlowsToReserve
	}
}

type LoadFlow struct {
	*BaseCacheCoherenceFlow
	OnCompletedCallback func()
}

func NewLoadFlow(generator *CacheController, access *MemoryHierarchyAccess, tag uint32, onCompletedCallback func()) *LoadFlow {
	var loadFlow = &LoadFlow{
		BaseCacheCoherenceFlow:NewBaseCacheCoherenceFlow(generator, nil, access, tag),
	}

	loadFlow.OnCompletedCallback = func() {
		onCompletedCallback()
		loadFlow.Complete()
	}

	return loadFlow
}

type StoreFlow struct {
	*BaseCacheCoherenceFlow
	OnCompletedCallback func()
}

func NewStoreFlow(generator *CacheController, access *MemoryHierarchyAccess, tag uint32, onCompletedCallback func()) *StoreFlow {
	var storeFlow = &StoreFlow{
		BaseCacheCoherenceFlow:NewBaseCacheCoherenceFlow(generator, nil, access, tag),
	}

	storeFlow.OnCompletedCallback = func() {
		onCompletedCallback()
		storeFlow.Complete()
	}

	return storeFlow
}