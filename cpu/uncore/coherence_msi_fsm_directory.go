package uncore

import (
	"github.com/mcai/acogo/simutil"
	"reflect"
)

type DirectoryEntry struct {
	Owner   CacheController
	Sharers []CacheController
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
	return directoryControllerFsm.DirectoryController.Cache().Sets[directoryControllerFsm.Set].Lines[directoryControllerFsm.Way]
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) fireTransition(event CacheControllerEvent) {
	event.Complete()
	directoryControllerFsm.DirectoryController.FsmFactory.FireTransition(directoryControllerFsm, event.EventType(), event)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) SendDataToRequester(producerFlow CacheCoherenceFlow, requester *CacheController, tag uint32, numInvAcks int32) {
	directoryControllerFsm.DirectoryController.TransferMessage(
		requester,
		directoryControllerFsm.DirectoryController.Cache().LineSize() + 8,
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
	directoryControllerFsm.DirectoryController.Cache().ReplacementPolicy.HandlePromotionOnHit(access, set, way)
	directoryControllerFsm.Line().Access = access
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) Stall(action func()) {
	directoryControllerFsm.StalledEvents = append(directoryControllerFsm.StalledEvents, action)
}

func (directoryControllerFsm *DirectoryControllerFiniteStateMachine) StallEvent(event CacheControllerEvent) {
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

