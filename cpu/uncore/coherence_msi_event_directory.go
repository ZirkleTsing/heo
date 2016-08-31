package uncore

type DirectoryControllerEventType string

const (
	DirectoryControllerEventType_GETS = DirectoryControllerEventType("GETS")
	DirectoryControllerEventType_GETM = DirectoryControllerEventType("GETM")
	DirectoryControllerEventType_DIR_REPLACEMENT = DirectoryControllerEventType("DIR_REPLACEMENT")
	DirectoryControllerEventType_RECALL_ACK = DirectoryControllerEventType("RECALL_ACK")
	DirectoryControllerEventType_LAST_RECALL_ACK = DirectoryControllerEventType("LAST_RECALL_ACK")
	DirectoryControllerEventType_PUTS_NOT_LAST = DirectoryControllerEventType("PUTS_NOT_LAST")
	DirectoryControllerEventType_PUTS_LAST = DirectoryControllerEventType("PUTS_LAST")
	DirectoryControllerEventType_PUTM_AND_DATA_FROM_OWNER = DirectoryControllerEventType("PUTM_AND_DATA_FROM_OWNER")
	DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER = DirectoryControllerEventType("PUTM_AND_DATA_FROM_NONOWNER")
	DirectoryControllerEventType_DATA = DirectoryControllerEventType("DATA")
	DirectoryControllerEventType_DATA_FROM_MEM = DirectoryControllerEventType("DATA_FROM_MEM")
)

type DirectoryControllerEvent interface {
	ControllerEvent
	EventType() DirectoryControllerEventType
}

type BaseDirectoryControllerEvent struct {
	*BaseControllerEvent
	eventType DirectoryControllerEventType
}

func NewBaseDirectoryControllerEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, eventType DirectoryControllerEventType, access *MemoryHierarchyAccess, tag uint32) *BaseDirectoryControllerEvent {
	var baseDirectoryControllerEvent = &BaseDirectoryControllerEvent{
		BaseControllerEvent:NewBaseControllerEvent(generator, producerFlow, access, tag),
		eventType:eventType,
	}

	return baseDirectoryControllerEvent
}

func (baseDirectoryControllerEvent *BaseDirectoryControllerEvent) EventType() DirectoryControllerEventType {
	return baseDirectoryControllerEvent.eventType
}

type DataEvent struct {
	*BaseDirectoryControllerEvent
	Sender *CacheController
}

func NewDataEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender *CacheController) *DataEvent {
	var dataEvent = &DataEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_DATA, access, tag),
		Sender:sender,
	}

	return dataEvent
}

type DataFromMemEvent struct {
	*BaseDirectoryControllerEvent
	Requester *CacheController
}

func NewDataFromMemEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *DataFromMemEvent {
	var dataFromMemEvent = &DataFromMemEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_DATA_FROM_MEM, access, tag),
		Requester:requester,
	}

	return dataFromMemEvent
}

type GetMEvent struct {
	*BaseDirectoryControllerEvent
	Requester         *CacheController
	Set               uint32
	Way               uint32
	OnStalledCallback func()
}

func NewGetMEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController, set uint32, way uint32, onStalledCallback func()) *GetMEvent {
	var getMEvent = &GetMEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_GETM, access, tag),
		Requester:requester,
		Set:set,
		Way:way,
		OnStalledCallback:onStalledCallback,
	}

	return getMEvent
}

type GetSEvent struct {
	*BaseDirectoryControllerEvent
	Requester         *CacheController
	Set               uint32
	Way               uint32
	OnStalledCallback func()
}

func NewGetSEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController, set uint32, way uint32, onStalledCallback func()) *GetSEvent {
	var getSEvent = &GetSEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_GETS, access, tag),
		Requester:requester,
		Set:set,
		Way:way,
		OnStalledCallback:onStalledCallback,
	}

	return getSEvent
}

type LastRecallAckEvent struct {
	*BaseDirectoryControllerEvent
}

func NewLastRecallAckEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32) *LastRecallAckEvent {
	var lastRecallAckEvent = &LastRecallAckEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_LAST_RECALL_ACK, access, tag),
	}

	return lastRecallAckEvent
}

type PutMAndDataFromNonOwnerEvent struct {
	*BaseDirectoryControllerEvent
	Requester *CacheController
}

func NewPutMAndDataFromNonOwnerEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *PutMAndDataFromNonOwnerEvent {
	var putMAndDataFromNonOwnerEvent = &PutMAndDataFromNonOwnerEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER, access, tag),
		Requester:requester,
	}

	return putMAndDataFromNonOwnerEvent
}

type PutMAndDataFromOwnerEvent struct {
	*BaseDirectoryControllerEvent
	Requester *CacheController
}

func NewPutMAndDataFromOwnerEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *PutMAndDataFromOwnerEvent {
	var putMAndDataFromOwnerEvent = &PutMAndDataFromOwnerEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_PUTM_AND_DATA_FROM_OWNER, access, tag),
		Requester:requester,
	}

	return putMAndDataFromOwnerEvent
}

type PutSLastEvent struct {
	*BaseDirectoryControllerEvent
	Requester *CacheController
}

func NewPutSLastEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *PutSLastEvent {
	var putSLastEvent = &PutSLastEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_PUTS_LAST, access, tag),
		Requester:requester,
	}

	return putSLastEvent
}

type PutSNotLastEvent struct {
	*BaseDirectoryControllerEvent
	Requester *CacheController
}

func NewPutSNotLastEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, requester *CacheController) *PutSNotLastEvent {
	var putSNotLastEvent = &PutSNotLastEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_PUTS_NOT_LAST, access, tag),
		Requester:requester,
	}

	return putSNotLastEvent
}

type RecallAckEvent struct {
	*BaseDirectoryControllerEvent
	Sender *CacheController
}

func NewRecallAckEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, sender *CacheController) *RecallAckEvent {
	var recallAckEvent = &RecallAckEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_RECALL_ACK, access, tag),
		Sender:sender,
	}

	return recallAckEvent
}

type DirReplacementEvent struct {
	*BaseDirectoryControllerEvent
	CacheAccess *CacheAccess
	Set               uint32
	Way               uint32
	OnCompletedCallback func()
	OnStalledCallback func()
}

func NewDirReplacementEvent(generator *DirectoryController, producerFlow CacheCoherenceFlow, access *MemoryHierarchyAccess, tag uint32, cacheAccess *CacheAccess, set uint32, way uint32, onCompletedCallback func(), onStalledCallback func()) *DirReplacementEvent {
	var dirReplacementEvent = &DirReplacementEvent{
		BaseDirectoryControllerEvent:NewBaseDirectoryControllerEvent(generator, producerFlow, DirectoryControllerEventType_DIR_REPLACEMENT, access, tag),
		CacheAccess:cacheAccess,
		Set:set,
		Way:way,
		OnCompletedCallback:onCompletedCallback,
		OnStalledCallback:onStalledCallback,
	}

	return dirReplacementEvent
}