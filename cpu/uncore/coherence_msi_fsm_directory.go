package uncore

import (
	"github.com/mcai/acogo/simutil"
	"reflect"
)

type DirectoryEntry struct {
	Owner   *CacheController
	Sharers []*CacheController
}

func NewDirectoryEntry() *DirectoryEntry {
	var directoryEntry = &DirectoryEntry{
	}

	return directoryEntry
}

type DirectoryControllerFiniteStateMachine struct {
	*simutil.BaseFiniteStateMachine
	DirectoryController *DirectoryController
	DirectoryEntry      *DirectoryEntry
	PreviousState       DirectoryControllerState
	Set                 uint32
	Way                 uint32
	NumRecallAcks       int32
	StalledEvents       []func()
	OnCompletedCallback func()
	EvicterTag          int32
	VictimTag           int32
}

func NewDirectoryControllerFiniteStateMachine(set uint32, way uint32, directoryController *DirectoryController) *DirectoryControllerFiniteStateMachine {
	var directoryControllerFsm = &DirectoryControllerFiniteStateMachine{
		BaseFiniteStateMachine:simutil.NewBaseFiniteStateMachine(DirectoryControllerState_I),
		DirectoryController:directoryController,
		DirectoryEntry:NewDirectoryEntry(),
		Set:set,
		Way:way,
		EvicterTag:INVALID_TAG,
		VictimTag:INVALID_TAG,
	}

	directoryControllerFsm.BlockingEventDispatcher.AddListener(
		reflect.TypeOf((*simutil.ExitStateEvent)(nil)),
		func(event interface{}) {
			directoryControllerFsm.PreviousState = directoryControllerFsm.State().(DirectoryControllerState)
		},
	)

	return directoryControllerFsm
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) Line() *CacheLine {
	return directoryControllerFsm.DirectoryController.Cache.Sets[directoryControllerFsm.Set].Lines[directoryControllerFsm.Way]
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) fireTransition(event DirectoryControllerEvent) {
	event.Complete()
	directoryControllerFsm.DirectoryController.FsmFactory.FireTransition(directoryControllerFsm, event.EventType(), event)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventGetS(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController, onStalledCallback func()) {
	var getSEvent = NewGetSEvent(
		directoryControllerFsm.DirectoryController,
		producerFlow,
		producerFlow.Access(),
		tag,
		requester,
		directoryControllerFsm.Set,
		directoryControllerFsm.Way,
		onStalledCallback,
	)

	directoryControllerFsm.fireTransition(getSEvent)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventGetM(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController, onStalledCallback func()) {
	var getMEvent = NewGetMEvent(
		directoryControllerFsm.DirectoryController,
		producerFlow,
		producerFlow.Access(),
		tag,
		requester,
		directoryControllerFsm.Set,
		directoryControllerFsm.Way,
		onStalledCallback,
	)

	directoryControllerFsm.fireTransition(getMEvent)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventReplacement(producerFlow CacheCoherenceFlow, tag uint32, cacheAccess *CacheAccess, requester *CacheController, onCompletedCallback func(), onStalledCallback func()) {
	var replacementEvent = NewDirReplacementEvent(
		directoryControllerFsm.DirectoryController,
		producerFlow,
		producerFlow.Access(),
		tag,
		cacheAccess,
		directoryControllerFsm.Set,
		directoryControllerFsm.Way,
		onCompletedCallback,
		onStalledCallback,
	)

	directoryControllerFsm.fireTransition(replacementEvent)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventRecallAck(producerFlow CacheCoherenceFlow, sender *CacheController, tag uint32) {
	var recallAckEvent = NewRecallAckEvent(
		directoryControllerFsm.DirectoryController,
		producerFlow,
		producerFlow.Access(),
		tag,
		sender,
	)

	directoryControllerFsm.fireTransition(recallAckEvent)

	if directoryControllerFsm.NumRecallAcks == 0 {
		var lastRecallAckEvent = NewLastRecallAckEvent(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
		)

		directoryControllerFsm.fireTransition(lastRecallAckEvent)
	}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventPutS(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32) {
	if len(directoryControllerFsm.DirectoryEntry.Sharers) > 1 {
		var putSNotLastEvent = NewPutSNotLastEvent(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		)

		directoryControllerFsm.fireTransition(putSNotLastEvent)
	} else {
		var putSLastEvent = NewPutSLastEvent(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		)

		directoryControllerFsm.fireTransition(putSLastEvent)
	}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventPutMAndData(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32) {
	if requester == directoryControllerFsm.DirectoryEntry.Owner {
		var putMAndDataFromOwnerEvent = NewPutMAndDataFromOwnerEvent(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		)

		directoryControllerFsm.fireTransition(putMAndDataFromOwnerEvent)
	} else {
		var putMAndDataFromNonOwnerEvent = NewPutMAndDataFromNonOwnerEvent(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		)

		directoryControllerFsm.fireTransition(putMAndDataFromNonOwnerEvent)
	}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventData(producerFlow CacheCoherenceFlow, tag uint32, sender *CacheController) {
	var dataEvent = NewDataEvent(
		directoryControllerFsm.DirectoryController,
		producerFlow,
		producerFlow.Access(),
		tag,
		sender,
	)

	directoryControllerFsm.fireTransition(dataEvent)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendDataToRequester(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, numInvAcks int32) {
	directoryControllerFsm.DirectoryController.TransferMessage(
		requester,
		directoryControllerFsm.DirectoryController.Cache.LineSize() + 8,
		NewDataMessage(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			directoryControllerFsm.DirectoryController,
			numInvAcks,
		),
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendPutAckToRequester(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32) {
	directoryControllerFsm.DirectoryController.SendPutAckToRequester(
		producerFlow,
		tag,
		requester,
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) CopyDataToMem(tag uint32) {
	directoryControllerFsm.DirectoryController.Next().(*MemoryController).ReceiveMemWriteRequest(
		directoryControllerFsm.DirectoryController,
		tag,
		func() {},
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendFwdGetSToOwner(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32) {
	directoryControllerFsm.DirectoryController.TransferMessage(
		directoryControllerFsm.DirectoryEntry.Owner,
		8,
		NewFwdGetSMessage(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		),
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendFwdGetMToOwner(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	directoryControllerFsm.DirectoryController.TransferMessage(
		directoryControllerFsm.DirectoryEntry.Owner,
		8,
		NewFwdGetMMessage(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
			requester,
		),
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendInvToSharers(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	for _, sharer := range directoryControllerFsm.DirectoryEntry.Sharers {
		if requester != sharer {
			directoryControllerFsm.DirectoryController.TransferMessage(
				sharer,
				8,
				NewInvMessage(
					directoryControllerFsm.DirectoryController,
					producerFlow,
					producerFlow.Access(),
					tag,
					requester,
				),
			)
		}
	}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendRecallToOwner(producerFlow CacheCoherenceFlow, tag uint32) {
	var owner = directoryControllerFsm.DirectoryEntry.Owner

	if owner.Cache.FindWay(tag) == INVALID_WAY {
		panic("Impossible")
	}

	directoryControllerFsm.DirectoryController.TransferMessage(
		owner,
		8,
		NewRecallMessage(
			directoryControllerFsm.DirectoryController,
			producerFlow,
			producerFlow.Access(),
			tag,
		),
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendRecallToSharers(producerFlow CacheCoherenceFlow, tag uint32) {
	for _, sharer := range directoryControllerFsm.DirectoryEntry.Sharers {
		if sharer.Cache.FindWay(tag) == INVALID_WAY {
			panic("Impossible")
		}

		directoryControllerFsm.DirectoryController.TransferMessage(
			sharer,
			8,
			NewRecallMessage(
				directoryControllerFsm.DirectoryController,
				producerFlow,
				producerFlow.Access(),
				tag,
			),
		)
	}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) AddRequesterAndOwnerToSharers(requester *CacheController) {
	directoryControllerFsm.DirectoryEntry.Sharers = append(
		directoryControllerFsm.DirectoryEntry.Sharers,
		requester,
		directoryControllerFsm.DirectoryEntry.Owner,
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) AddRequesterToSharers(requester *CacheController) {
	directoryControllerFsm.DirectoryEntry.Sharers = append(
		directoryControllerFsm.DirectoryEntry.Sharers,
		requester,
	)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) RemoveRequesterFromSharers(requester *CacheController) {
	var sharersToPreserve []*CacheController

	for _, sharer := range directoryControllerFsm.DirectoryEntry.Sharers {
		if requester != sharer {
			sharersToPreserve = append(sharersToPreserve, sharer)
		}
	}

	directoryControllerFsm.DirectoryEntry.Sharers = sharersToPreserve
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SetOwnerToRequester(requester *CacheController) {
	directoryControllerFsm.DirectoryEntry.Owner = requester
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) ClearSharers() {
	directoryControllerFsm.DirectoryEntry.Sharers = []*CacheController{}
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) ClearOwner() {
	directoryControllerFsm.DirectoryEntry.Owner = nil
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) FireServiceNonblockingRequestEvent(access *MemoryHierarchyAccess, tag uint32, hitInCache bool) {
	//TODO
	directoryControllerFsm.DirectoryController.UpdateStats(access.AccessType.IsRead(), hitInCache)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) FireCacheLineInsertEvent(access *MemoryHierarchyAccess, tag uint32, victimTag uint32) {
	//TODO
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) FireReplacementEvent(access *MemoryHierarchyAccess, tag uint32) {
	//TODO
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) FirePutSOrPutMAndDataFromOwnerEvent(access *MemoryHierarchyAccess, tag uint32) {
	//TODO
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) FireNonblockingRequestHitToTransientTagEvent(access *MemoryHierarchyAccess, tag uint32) {
	//TODO
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) Hit(access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32) {
	directoryControllerFsm.FireServiceNonblockingRequestEvent(access, tag, true)
	directoryControllerFsm.DirectoryController.Cache.ReplacementPolicy.HandlePromotionOnHit(access, set, way)
	directoryControllerFsm.Line().Access = access
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) Stall(action func()) {
	directoryControllerFsm.StalledEvents = append(directoryControllerFsm.StalledEvents, action)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) StallEvent(event DirectoryControllerEvent) {
	directoryControllerFsm.Stall(func() {
		directoryControllerFsm.fireTransition(event)
	})
}

type DirectoryControllerFiniteStateMachineFactory struct {
	*simutil.FiniteStateMachineFactory
}

func NewDirectoryControllerFiniteStateMachineFactory() *DirectoryControllerFiniteStateMachineFactory {
	var directoryControllerFsmFactory = &DirectoryControllerFiniteStateMachineFactory{
		FiniteStateMachineFactory:simutil.NewFiniteStateMachineFactory(),
	}

	//TODO...

	return directoryControllerFsmFactory
}

