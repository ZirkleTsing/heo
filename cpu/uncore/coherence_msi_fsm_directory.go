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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventRecallAck(producerFlow CacheCoherenceFlow, tag uint32, sender *CacheController) {
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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventPutS(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) OnEventPutMAndData(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendDataToRequester(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController, numInvAcks int32) {
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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendPutAckToRequester(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
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

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendFwdGetSToOwner(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
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

	var actionWhenStateChanged = func(fsm simutil.FiniteStateMachine) {
		var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)

		if directoryControllerFsm.PreviousState != directoryControllerFsm.State() {
			if directoryControllerFsm.State().(DirectoryControllerState).Stable() {
				var onCompletedCallback = directoryControllerFsm.OnCompletedCallback
				if onCompletedCallback != nil {
					directoryControllerFsm.OnCompletedCallback = nil
					onCompletedCallback()
				}
			}

			var stalledEventsToProcess []func()

			for _, stalledEvent := range directoryControllerFsm.StalledEvents {
				stalledEventsToProcess = append(stalledEventsToProcess, stalledEvent)
			}

			directoryControllerFsm.StalledEvents = []func(){}

			for _, stalledEventToProcess := range stalledEventsToProcess {
				stalledEventToProcess()
			}
		}
	}

	directoryControllerFsmFactory.InState(DirectoryControllerState_I).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.DirectoryController.NumPendingMemoryAccesses++

			directoryControllerFsm.DirectoryController.Transfer(
				directoryControllerFsm.DirectoryController.Next(),
				8,
				func() {
					directoryControllerFsm.DirectoryController.Next().(*MemoryController).ReceiveMemReadRequest(
						directoryControllerFsm.DirectoryController,
						event.Tag(),
						func() {
							directoryControllerFsm.DirectoryController.MemoryHierarchy().Driver().CycleAccurateEventQueue().Schedule(
								func() {
									directoryControllerFsm.DirectoryController.NumPendingMemoryAccesses--

									var dataFromMemEvent = NewDataFromMemEvent(
										directoryControllerFsm.DirectoryController,
										event,
										event.Access(),
										event.Tag(),
										event.Requester,
									)

									directoryControllerFsm.fireTransition(dataFromMemEvent)
								},
								int(directoryControllerFsm.DirectoryController.HitLatency()),
							)
						},
					)
				},
			)

			directoryControllerFsm.FireServiceNonblockingRequestEvent(event.Access(), event.Tag(), false)
			directoryControllerFsm.Line().Access = event.Access()
			directoryControllerFsm.Line().SetTag(int32(event.Tag()))
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.DirectoryController.NumPendingMemoryAccesses++

			directoryControllerFsm.DirectoryController.Transfer(
				directoryControllerFsm.DirectoryController.Next(),
				8,
				func() {
					directoryControllerFsm.DirectoryController.Next().(*MemoryController).ReceiveMemReadRequest(
						directoryControllerFsm.DirectoryController,
						event.Tag(),
						func() {
							directoryControllerFsm.DirectoryController.MemoryHierarchy().Driver().CycleAccurateEventQueue().Schedule(
								func() {
									directoryControllerFsm.DirectoryController.NumPendingMemoryAccesses--

									var dataFromMemEvent = NewDataFromMemEvent(
										directoryControllerFsm.DirectoryController,
										event,
										event.Access(),
										event.Tag(),
										event.Requester,
									)

									directoryControllerFsm.fireTransition(dataFromMemEvent)
								},
								int(directoryControllerFsm.DirectoryController.HitLatency()),
							)
						},
					)
				},
			)

			directoryControllerFsm.FireServiceNonblockingRequestEvent(event.Access(), event.Tag(), false)
			directoryControllerFsm.Line().Access = event.Access()
			directoryControllerFsm.Line().SetTag(int32(event.Tag()))
		},
		DirectoryControllerState_IM_D,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_IS_D).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
			directoryControllerFsm.FireNonblockingRequestHitToTransientTagEvent(
				event.Access(),
				event.Tag(),
			)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
			directoryControllerFsm.FireNonblockingRequestHitToTransientTagEvent(
				event.Access(),
				event.Tag(),
			)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IS_D,
	).OnCondition(
		DirectoryControllerEventType_DATA_FROM_MEM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DataFromMemEvent)

			directoryControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester, 0)
			directoryControllerFsm.AddRequesterToSharers(event.Requester)
			directoryControllerFsm.FireCacheLineInsertEvent(event.Access(), event.Tag(), uint32(directoryControllerFsm.VictimTag))
			directoryControllerFsm.EvicterTag = INVALID_TAG
			directoryControllerFsm.VictimTag = INVALID_TAG
			directoryControllerFsm.DirectoryController.Cache.ReplacementPolicy.HandleInsertionOnMiss(
				event.Access(),
				directoryControllerFsm.Set,
				directoryControllerFsm.Way,
			)
		},
		DirectoryControllerState_S,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_IM_D).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
			directoryControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
			directoryControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.StallEvent(event)
		},
		DirectoryControllerState_IM_D,
	).OnCondition(
		DirectoryControllerEventType_DATA_FROM_MEM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DataFromMemEvent)

			directoryControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester, 0)
			directoryControllerFsm.SetOwnerToRequester(event.Requester)
			directoryControllerFsm.FireCacheLineInsertEvent(event.Access(), event.Tag(), uint32(directoryControllerFsm.VictimTag))
			directoryControllerFsm.EvicterTag = INVALID_TAG
			directoryControllerFsm.VictimTag = INVALID_TAG
			directoryControllerFsm.DirectoryController.Cache.ReplacementPolicy.HandleInsertionOnMiss(
				event.Access(),
				directoryControllerFsm.Set,
				directoryControllerFsm.Way,
			)
		},
		DirectoryControllerState_IM_D,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_S).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester, 0)
			directoryControllerFsm.AddRequesterToSharers(event.Requester)
			directoryControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
		},
		DirectoryControllerState_S,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			var numInvAcks = int32(0)

			for _, sharer := range directoryControllerFsm.DirectoryEntry.Sharers {
				if sharer != event.Requester {
					numInvAcks++
				}
			}

			directoryControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester, numInvAcks)

			directoryControllerFsm.SendInvToSharers(event, event.Tag(), event.Requester)
			directoryControllerFsm.ClearSharers()
			directoryControllerFsm.SetOwnerToRequester(event.Requester)
			directoryControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
		},
		DirectoryControllerState_M,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.NumRecallAcks = int32(len(directoryControllerFsm.DirectoryEntry.Sharers))
			directoryControllerFsm.SendRecallToSharers(event, uint32(directoryControllerFsm.Line().Tag()))
			directoryControllerFsm.ClearSharers()
			directoryControllerFsm.OnCompletedCallback = event.OnCompletedCallback
			directoryControllerFsm.FireReplacementEvent(event.Access(), event.Tag())
			directoryControllerFsm.EvicterTag = int32(event.Tag())
			directoryControllerFsm.VictimTag = directoryControllerFsm.Line().Tag()
			directoryControllerFsm.DirectoryController.NumEvictions++
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_S,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
			directoryControllerFsm.FirePutSOrPutMAndDataFromOwnerEvent(event.Access(), event.Tag())
			directoryControllerFsm.Line().Access = nil
			directoryControllerFsm.Line().SetTag(INVALID_TAG)
		},
		DirectoryControllerState_I,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_S,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_M).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.SendFwdGetSToOwner(event, event.Tag(), event.Requester)
			directoryControllerFsm.AddRequesterAndOwnerToSharers(event.Requester)
			directoryControllerFsm.ClearOwner()
			directoryControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.SendFwdGetMToOwner(event, event.Tag(), event.Requester)
			directoryControllerFsm.SetOwnerToRequester(event.Requester)
			directoryControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
		},
		DirectoryControllerState_M,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.NumRecallAcks = 1
			directoryControllerFsm.SendRecallToOwner(event, uint32(directoryControllerFsm.Line().Tag()))
			directoryControllerFsm.ClearOwner()
			directoryControllerFsm.OnCompletedCallback = event.OnCompletedCallback
			directoryControllerFsm.FireReplacementEvent(event.Access(), event.Tag())
			directoryControllerFsm.EvicterTag = int32(event.Tag())
			directoryControllerFsm.VictimTag = directoryControllerFsm.Line().Tag()
			directoryControllerFsm.DirectoryController.NumEvictions++
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_M,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_M,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_OWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromOwnerEvent)

			directoryControllerFsm.CopyDataToMem(event.Tag())
			directoryControllerFsm.ClearOwner()
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
			directoryControllerFsm.FirePutSOrPutMAndDataFromOwnerEvent(event.Access(), event.Tag())
			directoryControllerFsm.Line().Access = nil
			directoryControllerFsm.Line().SetTag(INVALID_TAG)
		},
		DirectoryControllerState_I,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_M,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_S_D).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.RemoveRequesterFromSharers(event.Requester)
			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_S_D,
	).OnCondition(
		DirectoryControllerEventType_DATA,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DataEvent)

			directoryControllerFsm.CopyDataToMem(event.Tag())
		},
		DirectoryControllerState_S,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_MI_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_RECALL_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)

			directoryControllerFsm.NumRecallAcks--
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_LAST_RECALL_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*LastRecallAckEvent)

			directoryControllerFsm.CopyDataToMem(event.Tag())
			directoryControllerFsm.Line().Access = nil
			directoryControllerFsm.Line().SetTag(INVALID_TAG)
		},
		DirectoryControllerState_I,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_MI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_MI_A,
	)

	directoryControllerFsmFactory.InState(DirectoryControllerState_SI_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		DirectoryControllerEventType_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetSEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*GetMEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_DIR_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*DirReplacementEvent)

			directoryControllerFsm.Stall(event.OnStalledCallback)
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_RECALL_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)

			directoryControllerFsm.NumRecallAcks--
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_LAST_RECALL_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)

			directoryControllerFsm.Line().Access = nil
			directoryControllerFsm.Line().SetTag(INVALID_TAG)
		},
		DirectoryControllerState_I,
	).OnCondition(
		DirectoryControllerEventType_PUTS_NOT_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSNotLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTS_LAST,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutSLastEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_SI_A,
	).OnCondition(
		DirectoryControllerEventType_PUTM_AND_DATA_FROM_NONOWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var directoryControllerFsm = fsm.(*DirectoryControllerFiniteStateMachine)
			var event = params.(*PutMAndDataFromNonOwnerEvent)

			directoryControllerFsm.SendPutAckToRequester(event, event.Tag(), event.Requester)
		},
		DirectoryControllerState_SI_A,
	)

	return directoryControllerFsmFactory
}

