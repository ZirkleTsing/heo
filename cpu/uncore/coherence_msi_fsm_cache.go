package uncore

import (
	"github.com/mcai/acogo/simutil"
	"reflect"
)

type CacheControllerFiniteStateMachine struct {
	*simutil.BaseFiniteStateMachine
	CacheController     *CacheController
	PreviousState       CacheControllerState
	Set                 uint32
	Way                 uint32
	NumInvAcks          int32
	StalledEvents       []func()
	OnCompletedCallback func()
}

func NewCacheControllerFiniteStateMachine(set uint32, way uint32, cacheController *CacheController) *CacheControllerFiniteStateMachine {
	var cacheControllerFsm = &CacheControllerFiniteStateMachine{
		BaseFiniteStateMachine: simutil.NewBaseFiniteStateMachine(CacheControllerState_I),
		Set:set,
		Way:way,
		CacheController:cacheController,
	}

	cacheControllerFsm.BlockingEventDispatcher.AddListener(reflect.TypeOf((*simutil.ExitStateEvent)(nil)), func(event interface{}) {
		cacheControllerFsm.PreviousState = cacheControllerFsm.State().(CacheControllerState)
	})

	return cacheControllerFsm
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) Line() *CacheLine {
	return cacheControllerFsm.CacheController.Cache.Sets[cacheControllerFsm.Set].Lines[cacheControllerFsm.Way]
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) fireTransition(event CacheControllerEvent) {
	event.Complete()
	cacheControllerFsm.CacheController.FsmFactory.FireTransition(cacheControllerFsm, event.EventType(), event)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventLoad(producerFlow *LoadFlow, tag uint32, onCompletedCallback func(), onStalledCallback func()) {
	var loadEvent = NewLoadEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, cacheControllerFsm.Set, cacheControllerFsm.Way, onCompletedCallback, onStalledCallback)
	cacheControllerFsm.fireTransition(loadEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventStore(producerFlow *StoreFlow, tag uint32, onCompletedCallback func(), onStalledCallback func()) {
	var storeEvent = NewStoreEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, cacheControllerFsm.Set, cacheControllerFsm.Way, onCompletedCallback, onStalledCallback)
	cacheControllerFsm.fireTransition(storeEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventReplacement(producerFlow CacheCoherenceFlow, tag uint32, cacheAccess *CacheAccess, onCompletedCallback func(), onStalledCallback func()) {
	var replacementEvent = NewReplacementEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, cacheAccess, cacheControllerFsm.Set, cacheControllerFsm.Way, onCompletedCallback, onStalledCallback)
	cacheControllerFsm.fireTransition(replacementEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventFwdGetS(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	var fwdGetSEvent = NewFwdGetSEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, requester)
	cacheControllerFsm.fireTransition(fwdGetSEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventFwdGetM(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	var fwdGetMEvent = NewFwdGetMEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, requester)
	cacheControllerFsm.fireTransition(fwdGetMEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventInv(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	var invEvent = NewInvEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag, requester)
	cacheControllerFsm.fireTransition(invEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventRecall(producerFlow CacheCoherenceFlow, tag uint32) {
	var recallEvent = NewRecallEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag)
	cacheControllerFsm.fireTransition(recallEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventPutAck(producerFlow CacheCoherenceFlow, tag uint32) {
	var putAckEvent = NewPutAckEvent(cacheControllerFsm.CacheController, producerFlow, producerFlow.Access(), tag)
	cacheControllerFsm.fireTransition(putAckEvent)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventData(producerFlow CacheCoherenceFlow, tag uint32, sender Controller, numInvAcks int32) {
	cacheControllerFsm.NumInvAcks += numInvAcks

	switch sender.(type) {
	case *DirectoryController:
		if numInvAcks == 0 {
			var dataFromDirAcksEq0Event = NewDataFromDirAcksEq0Event(
				cacheControllerFsm.CacheController,
				producerFlow,
				producerFlow.Access(),
				tag,
				sender,
			)

			cacheControllerFsm.fireTransition(dataFromDirAcksEq0Event)

		} else {
			var dataFromDirAcksGt0Event = NewDataFromDirAcksGt0Event(
				cacheControllerFsm.CacheController,
				producerFlow,
				producerFlow.Access(),
				tag,
				sender,
			)

			cacheControllerFsm.fireTransition(dataFromDirAcksGt0Event)

			if cacheControllerFsm.NumInvAcks == 0 {
				cacheControllerFsm.OnEventLastInvAck(producerFlow, tag)
			}
		}
	default:
		var dataFromOwnerEvent = NewDataFromOwnerEvent(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			sender,
		)

		cacheControllerFsm.fireTransition(dataFromOwnerEvent)
	}
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventInvAck(producerFlow CacheCoherenceFlow, tag uint32, sender *CacheController) {
	var invAckEvent = NewInvAckEvent(
		cacheControllerFsm.CacheController,
		producerFlow,
		producerFlow.Access(),
		tag,
		sender,
	)

	cacheControllerFsm.fireTransition(invAckEvent)

	if cacheControllerFsm.NumInvAcks == 0 {
		cacheControllerFsm.OnEventLastInvAck(producerFlow, tag)
	}
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) OnEventLastInvAck(producerFlow CacheCoherenceFlow, tag uint32) {
	var lastInvAckEvent = NewLastInvAckEvent(
		cacheControllerFsm.CacheController,
		producerFlow,
		producerFlow.Access(),
		tag,
	)

	cacheControllerFsm.fireTransition(lastInvAckEvent)

	cacheControllerFsm.NumInvAcks = 0
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendGetSToDir(producerFlow CacheCoherenceFlow, tag uint32) {
	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		8,
		NewGetSMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendGetMToDir(producerFlow CacheCoherenceFlow, tag uint32) {
	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		8,
		NewGetMMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendPutSToDir(producerFlow CacheCoherenceFlow, tag uint32) {
	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		8,
		NewPutSMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendPutMAndDataToDir(producerFlow CacheCoherenceFlow, tag uint32) {
	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		cacheControllerFsm.CacheController.Cache.LineSize() + 8,
		NewPutMAndDataMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendDataToRequesterAndDir(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	cacheControllerFsm.CacheController.TransferMessage(
		requester,
		10,
		NewDataMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
			0,
		),
	)

	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		cacheControllerFsm.CacheController.Cache.LineSize() + 8,
		NewDataMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
			0,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendDataToRequester(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	cacheControllerFsm.CacheController.TransferMessage(
		requester,
		cacheControllerFsm.CacheController.Cache.LineSize() + 8,
		NewDataMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
			0,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendInvAckToRequester(producerFlow CacheCoherenceFlow, tag uint32, requester *CacheController) {
	cacheControllerFsm.CacheController.TransferMessage(
		requester,
		8,
		NewInvAckMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) SendRecallAckToDir(producerFlow CacheCoherenceFlow, tag uint32, size uint32) {
	cacheControllerFsm.CacheController.TransferMessage(
		cacheControllerFsm.CacheController.Next().(*DirectoryController),
		size,
		NewRecallAckMessage(
			cacheControllerFsm.CacheController,
			producerFlow,
			producerFlow.Access(),
			tag,
			cacheControllerFsm.CacheController,
		),
	)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) FireServiceNonblockingRequestEvent(access *MemoryHierarchyAccess, tag uint32, hitInCache bool) {
	//TODO
	cacheControllerFsm.CacheController.UpdateStats(access.AccessType.IsWrite(), hitInCache)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) FireReplacementEvent(access *MemoryHierarchyAccess, tag uint32) {
	//TODO
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) FireNonblockingRequestHitToTransientTagEvent(access *MemoryHierarchyAccess, tag uint32) {
	//TODO
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) Hit(access *MemoryHierarchyAccess, tag uint32, set uint32, way uint32) {
	cacheControllerFsm.FireServiceNonblockingRequestEvent(access, tag, true)
	cacheControllerFsm.CacheController.Cache.ReplacementPolicy.HandlePromotionOnHit(access, set, way)
	cacheControllerFsm.Line().Access = access
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) Stall(action func()) {
	cacheControllerFsm.StalledEvents = append(cacheControllerFsm.StalledEvents, action)
}

func (cacheControllerFsm *CacheControllerFiniteStateMachine) StallEvent(event CacheControllerEvent) {
	cacheControllerFsm.Stall(func() {
		cacheControllerFsm.fireTransition(event)
	})
}

type CacheControllerFiniteStateMachineFactory struct {
	*simutil.FiniteStateMachineFactory
}

func NewCacheControllerFiniteStateMachineFactory() *CacheControllerFiniteStateMachineFactory {
	var fsmFactory = &CacheControllerFiniteStateMachineFactory{
		FiniteStateMachineFactory:simutil.NewFiniteStateMachineFactory(),
	}

	var actionWhenStateChanged = func(fsm simutil.FiniteStateMachine) {
		var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

		if cacheControllerFsm.PreviousState != cacheControllerFsm.State() {
			if cacheControllerFsm.State().(CacheControllerState).Stable() {
				var onCompletedCallback = cacheControllerFsm.OnCompletedCallback
				if onCompletedCallback != nil {
					cacheControllerFsm.OnCompletedCallback = nil
					onCompletedCallback()
				}
			}

			var stalledEventsToProcess []func()

			for _, stalledEvent := range cacheControllerFsm.StalledEvents {
				stalledEventsToProcess = append(stalledEventsToProcess, stalledEvent)
			}

			cacheControllerFsm.StalledEvents = []func(){}

			for _, stalledEventToProcess := range stalledEventsToProcess {
				stalledEventToProcess()
			}
		}
	}

	fsmFactory.InState(CacheControllerState_I).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.SendGetSToDir(event, event.Tag())
			cacheControllerFsm.FireServiceNonblockingRequestEvent(event.Access(), event.Tag(), false)
			cacheControllerFsm.Line().Access = event.Access()
			cacheControllerFsm.Line().SetTag(int32(event.Tag()))
			cacheControllerFsm.OnCompletedCallback = func() {
				cacheControllerFsm.CacheController.Cache.ReplacementPolicy.HandleInsertionOnMiss(
					event.Access(),
					cacheControllerFsm.Set,
					cacheControllerFsm.Way,
				)
				event.OnCompletedCallback()
			}
		},
		CacheControllerState_IS_D,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.SendGetMToDir(event, event.Tag())
			cacheControllerFsm.FireServiceNonblockingRequestEvent(event.Access(), event.Tag(), false)
			cacheControllerFsm.Line().Access = event.Access()
			cacheControllerFsm.Line().SetTag(int32(event.Tag()))
			cacheControllerFsm.OnCompletedCallback = func() {
				cacheControllerFsm.CacheController.Cache.ReplacementPolicy.HandleInsertionOnMiss(
					event.Access(),
					cacheControllerFsm.Set,
					cacheControllerFsm.Way,
				)
				event.OnCompletedCallback()
			}
		},
		CacheControllerState_IM_AD,
	)

	fsmFactory.InState(CacheControllerState_IS_D).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IS_D,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IS_D,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_IS_D,
	).OnCondition(
		CacheControllerEventType_INV,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*InvEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IS_D,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_DIR_ACKS_EQ_0,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_S,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_OWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_S,
	)

	fsmFactory.InState(CacheControllerState_IM_AD).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IM_AD,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IM_AD,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)
			event.OnStalledCallback()
		},
		CacheControllerState_IM_AD,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IM_AD,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IM_AD,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_DIR_ACKS_EQ_0,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_M,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_DIR_ACKS_GT_0,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_OWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_M,
	).OnCondition(
		CacheControllerEventType_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.NumInvAcks--
		},
		CacheControllerState_IM_AD,
	)

	fsmFactory.InState(CacheControllerState_IM_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.FireNonblockingRequestHitToTransientTagEvent(event.Access(), event.Tag())
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			cacheControllerFsm.NumInvAcks--
		},
		CacheControllerState_IM_A,
	).OnCondition(
		CacheControllerEventType_LAST_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_M,
	)

	fsmFactory.InState(CacheControllerState_S).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
			cacheControllerFsm.CacheController.memoryHierarchy.Driver.CycleAccurateEventQueue().Schedule(
				event.OnCompletedCallback,
				0,
			)
		},
		CacheControllerState_S,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.SendGetMToDir(event, event.Tag())
			cacheControllerFsm.OnCompletedCallback = event.OnCompletedCallback
			cacheControllerFsm.FireServiceNonblockingRequestEvent(event.Access(), event.Tag(), true)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*ReplacementEvent)

			cacheControllerFsm.SendPutSToDir(event, uint32(cacheControllerFsm.Line().Tag()))
			cacheControllerFsm.OnCompletedCallback = event.OnCompletedCallback
			cacheControllerFsm.FireReplacementEvent(event.Access(), event.Tag())
			cacheControllerFsm.CacheController.NumEvictions++
		},
		CacheControllerState_SI_A,
	).OnCondition(
		CacheControllerEventType_INV,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*InvEvent)

			cacheControllerFsm.SendInvAckToRequester(event, event.Tag(), event.Requester)
			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.SendRecallAckToDir(event, event.Tag(), 8)
			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	)

	fsmFactory.InState(CacheControllerState_SM_AD).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
			cacheControllerFsm.CacheController.MemoryHierarchy().Driver.CycleAccurateEventQueue().Schedule(
				event.OnCompletedCallback,
				0,
			)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_INV,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*InvEvent)

			cacheControllerFsm.SendInvAckToRequester(event, event.Tag(), event.Requester)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.SendRecallAckToDir(event, event.Tag(), 8)
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_DIR_ACKS_EQ_0,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_DIR_ACKS_GT_0,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_DATA_FROM_OWNER,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_SM_AD,
	).OnCondition(
		CacheControllerEventType_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.NumInvAcks--
		},
		CacheControllerState_SM_AD,
	)

	fsmFactory.InState(CacheControllerState_SM_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
			cacheControllerFsm.CacheController.MemoryHierarchy().Driver.CycleAccurateEventQueue().Schedule(
				event.OnCompletedCallback,
				0,
			)
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.NumInvAcks--
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_LAST_INV_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
		},
		CacheControllerState_SM_A,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.StallEvent(event)
		},
		CacheControllerState_SM_A,
	)

	fsmFactory.InState(CacheControllerState_M).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
			cacheControllerFsm.CacheController.MemoryHierarchy().Driver.CycleAccurateEventQueue().Schedule(
				event.OnCompletedCallback,
				0,
			)
		},
		CacheControllerState_M,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Hit(event.Access(), event.Tag(), event.Set, event.Way)
			cacheControllerFsm.CacheController.MemoryHierarchy().Driver.CycleAccurateEventQueue().Schedule(
				event.OnCompletedCallback,
				0,
			)
		},
		CacheControllerState_M,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*ReplacementEvent)

			cacheControllerFsm.SendPutMAndDataToDir(event, uint32(cacheControllerFsm.Line().Tag()))
			cacheControllerFsm.OnCompletedCallback = event.OnCompletedCallback
			cacheControllerFsm.FireReplacementEvent(event.Access(), event.Tag())
			cacheControllerFsm.CacheController.NumEvictions++
		},
		CacheControllerState_MI_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.SendDataToRequesterAndDir(event, event.Tag(), event.Requester)
		},
		CacheControllerState_S,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester)
			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.SendRecallAckToDir(
				event,
				event.Tag(),
				cacheControllerFsm.CacheController.Cache.LineSize() + 8,
			)
			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	)

	fsmFactory.InState(CacheControllerState_MI_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_MI_A,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_MI_A,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_MI_A,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.SendRecallAckToDir(event, event.Tag(), 8)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETS,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetSEvent)

			cacheControllerFsm.SendDataToRequesterAndDir(event, event.Tag(), event.Requester)
		},
		CacheControllerState_SI_A,
	).OnCondition(
		CacheControllerEventType_FWD_GETM,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*FwdGetMEvent)

			cacheControllerFsm.SendDataToRequester(event, event.Tag(), event.Requester)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_PUT_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	)

	fsmFactory.InState(CacheControllerState_SI_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_SI_A,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_SI_A,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_SI_A,
	).OnCondition(
		CacheControllerEventType_INV,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*InvEvent)

			cacheControllerFsm.SendInvAckToRequester(event, event.Tag(), event.Requester)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_RECALL,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*RecallEvent)

			cacheControllerFsm.SendRecallAckToDir(event, event.Tag(), 8)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_PUT_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	)

	fsmFactory.InState(CacheControllerState_II_A).SetOnCompletedCallback(actionWhenStateChanged).OnCondition(
		CacheControllerEventType_LOAD,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*LoadEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_STORE,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)
			var event = params.(*StoreEvent)

			cacheControllerFsm.Stall(event.OnStalledCallback)
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_REPLACEMENT,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var event = params.(*ReplacementEvent)

			event.OnStalledCallback()
		},
		CacheControllerState_II_A,
	).OnCondition(
		CacheControllerEventType_PUT_ACK,
		func(fsm simutil.FiniteStateMachine, condition interface{}, params interface{}) {
			var cacheControllerFsm = fsm.(*CacheControllerFiniteStateMachine)

			cacheControllerFsm.Line().Access = nil
			cacheControllerFsm.Line().SetTag(INVALID_TAG)
		},
		CacheControllerState_I,
	)

	return fsmFactory
}