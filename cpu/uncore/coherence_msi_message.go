package uncore

type CoherenceMessageType string

const (
	CoherenceMessageType_GETS = CoherenceMessageType("GETS")
	CoherenceMessageType_GETM = CoherenceMessageType("GETM")
	CoherenceMessageType_PUTS = CoherenceMessageType("PUTS")
	CoherenceMessageType_PUTM_AND_DATA = CoherenceMessageType("PUTM_AND_DATA")
	CoherenceMessageType_FWD_GETS = CoherenceMessageType("FWD_GETS")
	CoherenceMessageType_FWD_GETM = CoherenceMessageType("FWD_GETM")
	CoherenceMessageType_INV = CoherenceMessageType("INV")
	CoherenceMessageType_RECALL = CoherenceMessageType("RECALL")
	CoherenceMessageType_PUT_ACK = CoherenceMessageType("PUT_ACK")
	CoherenceMessageType_DATA = CoherenceMessageType("DATA")
	CoherenceMessageType_INV_ACK = CoherenceMessageType("INV_ACK")
	CoherenceMessageType_RECALL_ACK = CoherenceMessageType("RECALL_ACK")
)

type CoherenceMessage interface {
	CacheCoherenceFlow
	MessageType() CoherenceMessageType
	DestArrived() bool
	SetDestArrived(destArrived bool)
}

type BaseCoherenceMessage struct {
	*BaseCacheCoherenceFlow
	messageType CoherenceMessageType
	destArrived bool
}

func NewBaseCoherenceMessage(generator Controller, producerFlow CacheCoherenceFlow, messageType CoherenceMessageType, access *MemoryHierarchyAccess, tag uint32) *BaseCoherenceMessage {
	var coherenceMessage = &BaseCoherenceMessage{
		BaseCacheCoherenceFlow:NewBaseCacheCoherenceFlow(generator, producerFlow, access, tag),
		messageType:messageType,
	}

	return coherenceMessage
}

func (baseCoherenceMessage *BaseCoherenceMessage) MessageType() CoherenceMessageType {
	return baseCoherenceMessage.messageType
}

func (baseCoherenceMessage *BaseCoherenceMessage) DestArrived() bool {
	return baseCoherenceMessage.destArrived
}

func (baseCoherenceMessage *BaseCoherenceMessage) SetDestArrived(destArrived bool) {
	baseCoherenceMessage.destArrived = destArrived
}

type DataMessage struct {
	*BaseCoherenceMessage
	Sender     Controller
	NumInvAcks int32
}

func NewDataMessage(generator Controller, producerFlow CacheCoherenceFlow, sender Controller, tag uint32, numInvAcks int32, access *MemoryHierarchyAccess) *DataMessage {
	var dataMessage = &DataMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_DATA, access, tag),
		Sender:sender,
		NumInvAcks:numInvAcks,
	}

	return dataMessage
}

type FwdGetMMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewFwdGetMMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *FwdGetMMessage {
	var fwdGetMMessage = &FwdGetMMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_FWD_GETM, access, tag),
		Requester:requester,
	}

	return fwdGetMMessage
}

type FwdGetSMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewFwdGetSMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *FwdGetSMessage {
	var fwdGetSMessage = &FwdGetSMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_FWD_GETS, access, tag),
		Requester:requester,
	}

	return fwdGetSMessage
}

type GetMMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewGetMMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *GetMMessage {
	var getMMessage = &GetMMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_GETM, access, tag),
		Requester:requester,
	}

	return getMMessage
}

type GetSMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewGetSMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *GetSMessage {
	var getSMessage = &GetSMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_GETS, access, tag),
		Requester:requester,
	}

	return getSMessage
}

type InvAckMessage struct {
	*BaseCoherenceMessage
	Sender *CacheController
}

func NewInvAckMessage(generator Controller, producerFlow CacheCoherenceFlow, sender *CacheController, tag uint32, access *MemoryHierarchyAccess) *InvAckMessage {
	var invAckMessage = &InvAckMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_INV_ACK, access, tag),
		Sender:sender,
	}

	return invAckMessage
}

type InvMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewInvMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *InvMessage {
	var invMessage = &InvMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_INV, access, tag),
		Requester:requester,
	}

	return invMessage
}

type PutAckMessage struct {
	*BaseCoherenceMessage
}

func NewPutAckMessage(generator Controller, producerFlow CacheCoherenceFlow, tag uint32, access *MemoryHierarchyAccess) *PutAckMessage {
	var putAckMessage = &PutAckMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_PUT_ACK, access, tag),
	}

	return putAckMessage
}

type PutMAndDataMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewPutMAndDataMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *PutMAndDataMessage {
	var putMAndDataMessage = &PutMAndDataMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_PUTM_AND_DATA, access, tag),
		Requester:requester,
	}

	return putMAndDataMessage
}

type PutSMessage struct {
	*BaseCoherenceMessage
	Requester *CacheController
}

func NewPutSMessage(generator Controller, producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, access *MemoryHierarchyAccess) *PutSMessage {
	var putSMessage = &PutSMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_PUTS, access, tag),
		Requester:requester,
	}

	return putSMessage
}

type RecallAckMessage struct {
	*BaseCoherenceMessage
	Sender *CacheController
}

func NewRecallAckMessage(generator Controller, producerFlow CacheCoherenceFlow, sender *CacheController, tag uint32, access *MemoryHierarchyAccess) *RecallAckMessage {
	var recallAckMessage = &RecallAckMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_RECALL_ACK, access, tag),
		Sender:sender,
	}

	return recallAckMessage
}

type RecallMessage struct {
	*BaseCoherenceMessage
}

func NewRecallMessage(generator Controller, producerFlow CacheCoherenceFlow, tag uint32, access *MemoryHierarchyAccess) *RecallMessage {
	var recallMessage = &RecallMessage{
		BaseCoherenceMessage:NewBaseCoherenceMessage(generator, producerFlow, CoherenceMessageType_RECALL, access, tag),
	}

	return recallMessage
}