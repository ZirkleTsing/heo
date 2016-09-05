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
	var oooCore = &OoOCore{
		MemoryHierarchyCore: NewMemoryHierarchyCore(processor, num),
	}

	oooCore.fuPool = NewFUPool(oooCore)

	return oooCore
}

func (oooCore *OoOCore) FUPool() *FUPool {
	return oooCore.fuPool
}

func (oooCore *OoOCore) WaitingInstructionQueue() []GeneralReorderBufferEntry {
	return oooCore.waitingInstructionQueue
}

func (oooCore *OoOCore) ReadyInstructionQueue() []GeneralReorderBufferEntry {
	return oooCore.readyInstructionQueue
}

func (oooCore *OoOCore) ReadyLoadQueue() []GeneralReorderBufferEntry {
	return oooCore.readyLoadQueue
}

func (oooCore *OoOCore) WaitingStoreQueue() []GeneralReorderBufferEntry {
	return oooCore.waitingStoreQueue
}

func (oooCore *OoOCore) ReadyStoreQueue() []GeneralReorderBufferEntry {
	return oooCore.readyStoreQueue
}

func (oooCore *OoOCore) OoOEventQueue() []GeneralReorderBufferEntry {
	return oooCore.oooEventQueue
}

func (oooCore *OoOCore) SetOoOEventQueue(oooEventQueue []GeneralReorderBufferEntry) {
	oooCore.oooEventQueue = oooEventQueue
}
