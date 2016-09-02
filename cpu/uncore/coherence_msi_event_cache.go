package uncore

type CacheControllerEventType string

const (
	CacheControllerEventType_LOAD = CacheControllerEventType("LOAD")
	CacheControllerEventType_STORE = CacheControllerEventType("STORE")
	CacheControllerEventType_REPLACEMENT = CacheControllerEventType("REPLACEMENT")
	CacheControllerEventType_FWD_GETS = CacheControllerEventType("FWD_GETS")
	CacheControllerEventType_FWD_GETM = CacheControllerEventType("FWD_GETM")
	CacheControllerEventType_INV = CacheControllerEventType("INV")
	CacheControllerEventType_RECALL = CacheControllerEventType("RECALL")
	CacheControllerEventType_PUT_ACK = CacheControllerEventType("PUT_ACK")
	CacheControllerEventType_DATA_FROM_DIR_ACKS_EQ_0 = CacheControllerEventType("DATA_FROM_DIR_ACKS_EQ_0")
	CacheControllerEventType_DATA_FROM_DIR_ACKS_GT_0 = CacheControllerEventType("DATA_FROM_DIR_ACKS_GT_0")
	CacheControllerEventType_DATA_FROM_OWNER = CacheControllerEventType("DATA_FROM_OWNER")
	CacheControllerEventType_INV_ACK = CacheControllerEventType("INV_ACK")
	CacheControllerEventType_LAST_INV_ACK = CacheControllerEventType("LAST_INV_ACK")
)

type CacheControllerEvent interface {
	ControllerEvent
	EventType() CacheControllerEventType
}

type BaseCacheControllerEvent struct {
	*BaseControllerEvent
	eventType CacheControllerEventType
}

func NewBaseCacheControllerEvent(generator *CacheController, producerFlow CacheCoherenceFlow, eventType CacheControllerEventType, access *MemoryHierarchyAccess, tag uint32) *BaseCacheControllerEvent {
	var baseCacheControllerEvent = &BaseCacheControllerEvent{
		BaseControllerEvent:NewBaseControllerEvent(generator, producerFlow, access, tag),
		eventType:eventType,
	}

	return baseCacheControllerEvent
}

func (baseCacheControllerEvent *BaseCacheControllerEvent) EventType() CacheControllerEventType {
	return baseCacheControllerEvent.eventType
}

type DataFromDirAcksEq0Event struct {
	*BaseCacheControllerEvent
	Sender Controller
}

func NewDataFromDirAcksEq0Event(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender Controller) *DataFromDirAcksEq0Event {
	var dataFromDirAcksEq0Event = &DataFromDirAcksEq0Event{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_DATA_FROM_DIR_ACKS_EQ_0, access, tag),
		Sender:sender,
	}

	SetupCacheCoherenceFlowTree(dataFromDirAcksEq0Event)

	return dataFromDirAcksEq0Event
}

type DataFromDirAcksGt0Event struct {
	*BaseCacheControllerEvent
	Sender Controller
}

func NewDataFromDirAcksGt0Event(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender Controller) *DataFromDirAcksGt0Event {
	var dataFromDirAcksGt0Event = &DataFromDirAcksGt0Event{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_DATA_FROM_DIR_ACKS_GT_0, access, tag),
		Sender:sender,
	}

	SetupCacheCoherenceFlowTree(dataFromDirAcksGt0Event)

	return dataFromDirAcksGt0Event
}

type DataFromOwnerEvent struct {
	*BaseCacheControllerEvent
	Sender Controller
}

func NewDataFromOwnerEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender Controller) *DataFromOwnerEvent {
	var dataFromOwnerEvent = &DataFromOwnerEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_DATA_FROM_OWNER, access, tag),
		Sender:sender,
	}

	SetupCacheCoherenceFlowTree(dataFromOwnerEvent)

	return dataFromOwnerEvent
}

type FwdGetMEvent struct {
	*BaseCacheControllerEvent
	Requester *CacheController
}

func NewFwdGetMEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *FwdGetMEvent {
	var fwdGetMEvent = &FwdGetMEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_FWD_GETM, access, tag),
		Requester:requester,
	}

	SetupCacheCoherenceFlowTree(fwdGetMEvent)

	return fwdGetMEvent
}

type FwdGetSEvent struct {
	*BaseCacheControllerEvent
	Requester *CacheController
}

func NewFwdGetSEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *FwdGetSEvent {
	var fwdGetSEvent = &FwdGetSEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_FWD_GETS, access, tag),
		Requester:requester,
	}

	SetupCacheCoherenceFlowTree(fwdGetSEvent)

	return fwdGetSEvent
}

type InvAckEvent struct {
	*BaseCacheControllerEvent
	Sender *CacheController
}

func NewInvAckEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender *CacheController) *InvAckEvent {
	var invAckEvent = &InvAckEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_INV_ACK, access, tag),
		Sender:sender,
	}

	SetupCacheCoherenceFlowTree(invAckEvent)

	return invAckEvent
}

type InvEvent struct {
	*BaseCacheControllerEvent
	Requester *CacheController
}

func NewInvEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *InvEvent {
	var invEvent = &InvEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_INV, access, tag),
		Requester:requester,
	}

	SetupCacheCoherenceFlowTree(invEvent)

	return invEvent
}

type LastInvAckEvent struct {
	*BaseCacheControllerEvent
}

func NewLastInvAckEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *LastInvAckEvent {
	var lastInvAckEvent = &LastInvAckEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_LAST_INV_ACK, access, tag),
	}

	SetupCacheCoherenceFlowTree(lastInvAckEvent)

	return lastInvAckEvent
}

type LoadEvent struct {
	*BaseCacheControllerEvent
	Set                 uint32
	Way                 uint32
	OnCompletedCallback func()
	OnStalledCallback   func()
}

func NewLoadEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32, onCompletedCallback func(), onStalledCallback func()) *LoadEvent {
	var loadEvent = &LoadEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_LOAD, access, tag),
		Set:set,
		Way:way,
		OnCompletedCallback:onCompletedCallback,
		OnStalledCallback:onStalledCallback,
	}

	SetupCacheCoherenceFlowTree(loadEvent)

	return loadEvent
}

type PutAckEvent struct {
	*BaseCacheControllerEvent
}

func NewPutAckEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *PutAckEvent {
	var putAckEvent = &PutAckEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_PUT_ACK, access, tag),
	}

	SetupCacheCoherenceFlowTree(putAckEvent)

	return putAckEvent
}

type RecallEvent struct {
	*BaseCacheControllerEvent
}

func NewRecallEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *RecallEvent {
	var recallEvent = &RecallEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_RECALL, access, tag),
	}

	SetupCacheCoherenceFlowTree(recallEvent)

	return recallEvent
}

type ReplacementEvent struct {
	*BaseCacheControllerEvent
	CacheAccess         *CacheAccess
	Set                 uint32
	Way                 uint32
	OnCompletedCallback func()
	OnStalledCallback   func()
}

func NewReplacementEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, cacheAccess *CacheAccess, set uint32, way uint32, onCompletedCallback func(), onStalledCallback func()) *ReplacementEvent {
	var replacementEvent = &ReplacementEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_REPLACEMENT, access, tag),
		CacheAccess:cacheAccess,
		Set:set,
		Way:way,
		OnCompletedCallback:onCompletedCallback,
		OnStalledCallback:onStalledCallback,
	}

	SetupCacheCoherenceFlowTree(replacementEvent)

	return replacementEvent
}

type StoreEvent struct {
	*BaseCacheControllerEvent
	Set                 uint32
	Way                 uint32
	OnCompletedCallback func()
	OnStalledCallback   func()
}

func NewStoreEvent(generator *CacheController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32, onCompletedCallback func(), onStalledCallback func()) *StoreEvent {
	var storeEvent = &StoreEvent{
		BaseCacheControllerEvent:NewBaseCacheControllerEvent(generator, producerFlow, CacheControllerEventType_STORE, access, tag),
		Set:set,
		Way:way,
		OnCompletedCallback:onCompletedCallback,
		OnStalledCallback:onStalledCallback,
	}

	SetupCacheCoherenceFlowTree(storeEvent)

	return storeEvent
}