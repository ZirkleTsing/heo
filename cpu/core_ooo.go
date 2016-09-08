package cpu

type OoOCore struct {
	*MemoryHierarchyCore

	fuPool                  *FUPool
	waitingInstructionQueue []GeneralReorderBufferEntry
	readyInstructionQueue   []GeneralReorderBufferEntry
	readyLoadQueue          []GeneralReorderBufferEntry
	waitingStoreQueue       []GeneralReorderBufferEntry
	readyStoreQueue         []GeneralReorderBufferEntry
	oooEventQueue           []GeneralReorderBufferEntry
}

func NewOoOCore(processor *Processor, num int32) *OoOCore {
	var core = &OoOCore{
		MemoryHierarchyCore: NewMemoryHierarchyCore(processor, num),
	}

	core.fuPool = NewFUPool(core)

	return core
}

func (core *OoOCore) FUPool() *FUPool {
	return core.fuPool
}

func (core *OoOCore) WaitingInstructionQueue() []GeneralReorderBufferEntry {
	return core.waitingInstructionQueue
}

func (core *OoOCore) SetWaitingInstructionQueue(waitingInstructionQueue []GeneralReorderBufferEntry) {
	core.waitingInstructionQueue = waitingInstructionQueue
}

func (core *OoOCore) ReadyInstructionQueue() []GeneralReorderBufferEntry {
	return core.readyInstructionQueue
}

func (core *OoOCore) SetReadyInstructionQueue(readyInstructionQueue []GeneralReorderBufferEntry) {
	core.readyInstructionQueue = readyInstructionQueue
}

func (core *OoOCore) ReadyLoadQueue() []GeneralReorderBufferEntry {
	return core.readyLoadQueue
}

func (core *OoOCore) SetReadyLoadQueue(readyLoadQueue []GeneralReorderBufferEntry) {
	core.readyLoadQueue = readyLoadQueue
}

func (core *OoOCore) WaitingStoreQueue() []GeneralReorderBufferEntry {
	return core.waitingStoreQueue
}

func (core *OoOCore) SetWaitingStoreQueue(waitingStoreQueue []GeneralReorderBufferEntry) {
	core.waitingStoreQueue = waitingStoreQueue
}

func (core *OoOCore) ReadyStoreQueue() []GeneralReorderBufferEntry {
	return core.readyStoreQueue
}

func (core *OoOCore) SetReadyStoreQueue(readyStoreQueue []GeneralReorderBufferEntry) {
	core.readyStoreQueue = readyStoreQueue
}

func (core *OoOCore) OoOEventQueue() []GeneralReorderBufferEntry {
	return core.oooEventQueue
}

func (core *OoOCore) SetOoOEventQueue(oooEventQueue []GeneralReorderBufferEntry) {
	core.oooEventQueue = oooEventQueue
}

func (core *OoOCore) DoMeasurementOneCycle() {

}

func (core *OoOCore) Fetch() {
	//TODO
}

func (core *OoOCore) RegisterRename() {
	//TODO
}

func (core *OoOCore) Dispatch() {
	//TODO
}

func (core *OoOCore) Wakeup() {
	//TODO
}

func (core *OoOCore) Issue() {
	//TODO
}

func (core *OoOCore) Writeback() {
	//TODO
}

func (core *OoOCore) RefreshLoadStoreQueue() {
	//TODO
}

func (core *OoOCore) Commit() {
	//TODO
}

func (core *OoOCore) RemoveFromQueues(entryToRemove GeneralReorderBufferEntry) {
	var waitingInstructionQueueToReserve []GeneralReorderBufferEntry
	var readyInstructionQueueToReserve   []GeneralReorderBufferEntry
	var readyLoadQueueToReserve          []GeneralReorderBufferEntry
	var waitingStoreQueueToReserve       []GeneralReorderBufferEntry
	var readyStoreQueueToReserve         []GeneralReorderBufferEntry
	var oooEventQueueToReserve           []GeneralReorderBufferEntry

	for _, entry := range core.waitingInstructionQueue {
		if entry != entryToRemove {
			waitingInstructionQueueToReserve = append(waitingInstructionQueueToReserve, entry)
		}
	}

	for _, entry := range core.readyInstructionQueue {
		if entry != entryToRemove {
			readyInstructionQueueToReserve = append(readyInstructionQueueToReserve, entry)
		}
	}

	for _, entry := range core.readyLoadQueue {
		if entry != entryToRemove {
			readyLoadQueueToReserve = append(readyLoadQueueToReserve, entry)
		}
	}

	for _, entry := range core.waitingStoreQueue {
		if entry != entryToRemove {
			waitingStoreQueueToReserve = append(waitingStoreQueueToReserve, entry)
		}
	}

	for _, entry := range core.readyStoreQueue {
		if entry != entryToRemove {
			readyStoreQueueToReserve = append(readyStoreQueueToReserve, entry)
		}
	}

	for _, entry := range core.oooEventQueue {
		if entry != entryToRemove {
			oooEventQueueToReserve = append(oooEventQueueToReserve, entry)
		}
	}

	core.waitingInstructionQueue = waitingInstructionQueueToReserve
	core.readyInstructionQueue = readyInstructionQueueToReserve
	core.readyLoadQueue = readyLoadQueueToReserve
	core.waitingStoreQueue = waitingStoreQueueToReserve
	core.readyStoreQueue = readyStoreQueueToReserve
	core.oooEventQueue = oooEventQueueToReserve

	entryToRemove.SetSquashed(true)
}