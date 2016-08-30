package uncore

import (
	"github.com/mcai/acogo/simutil"
	"reflect"
)

type CacheControllerFiniteStateMachine struct {
	*simutil.BaseFiniteStateMachine
	CacheController     CacheController
	PreviousState       CacheControllerState
	Set                 uint32
	Way                 uint32
	NumInvAcks          int32
	StalledEvents       []func()
	OnCompletedCallback func()
}

func NewCacheControllerFiniteStateMachine(set uint32, way uint32, cacheController CacheController) *CacheControllerFiniteStateMachine {
	var cacheControllerFsm = &CacheControllerFiniteStateMachine{
		BaseFiniteStateMachine: simutil.NewBaseFiniteStateMachine(CacheControllerState_I),
		Set:set,
		Way:way,
		CacheController:cacheController,
	}

	cacheControllerFsm.BlockingEventDispatcher.AddListener(reflect.TypeOf((*simutil.ExitStateEvent)(nil)), func(event interface{}){
		cacheControllerFsm.PreviousState = cacheControllerFsm.State().(CacheControllerState)
	})

	return cacheControllerFsm
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

//TODO...

type CacheControllerFiniteStateMachineFactory struct {
	*simutil.FiniteStateMachineFactory
}

func NewCacheControllerFiniteStateMachineFactory() *CacheControllerFiniteStateMachineFactory {
	var fsmFactory = &CacheControllerFiniteStateMachineFactory{

	}

	return fsmFactory
}