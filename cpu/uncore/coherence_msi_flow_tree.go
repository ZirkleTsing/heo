package uncore

func SetupCacheCoherenceFlowTree(cacheCoherenceFlow CacheCoherenceFlow) {
	if cacheCoherenceFlow.ProducerFlow() == nil {
		cacheCoherenceFlow.SetAncestorFlow(cacheCoherenceFlow)
		cacheCoherenceFlow.Generator().MemoryHierarchy().SetPendingFlows(
			append(cacheCoherenceFlow.Generator().MemoryHierarchy().PendingFlows(), cacheCoherenceFlow),
		)
	} else {
		cacheCoherenceFlow.SetAncestorFlow(cacheCoherenceFlow.ProducerFlow().AncestorFlow())
		cacheCoherenceFlow.ProducerFlow().SetChildFlows(append(cacheCoherenceFlow.ProducerFlow().ChildFlows(), cacheCoherenceFlow))
	}

	cacheCoherenceFlow.AncestorFlow().SetNumPendingDescendantFlows(
		cacheCoherenceFlow.AncestorFlow().NumPendingDescendantFlows() + 1)
}
