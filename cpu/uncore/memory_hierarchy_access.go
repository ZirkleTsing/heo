package uncore

type MemoryHierarchyAccessType uint32

const (
	MemoryHierarchyAccess_IFETCH = 0
	MemoryHierarchyAccess_LOAD = 1
	MemoryHierarchyAccess_STORE = 2
	MemoryHierarchyAccess_UNKNOWN = 3
)

func (memoryHierarchyAccessType MemoryHierarchyAccessType) IsRead() bool {
	return memoryHierarchyAccessType == MemoryHierarchyAccess_IFETCH ||
		memoryHierarchyAccessType == MemoryHierarchyAccess_LOAD
}

func (memoryHierarchyAccessType MemoryHierarchyAccessType) IsWrite() bool {
	return memoryHierarchyAccessType == MemoryHierarchyAccess_STORE
}

type MemoryHierarchyAccess struct {
	MemoryHierarchy *MemoryHierarchy
	Id int32
	AccessType MemoryHierarchyAccessType

	ThreadId int32
	VirtualPc int32
	PhysicalAddress uint32
	PhysicalTag uint32

	OnCompletedCallback func()

	Aliases []*MemoryHierarchyAccess

	BeginCycle int64
	EndCycle int64
}

func NewMemoryHierarchyAccess(memoryHierarchy *MemoryHierarchy, accessType MemoryHierarchyAccessType, threadId int32, virtualPc int32, physicalAddress uint32, physicalTag uint32, onCompletedCallback func()) *MemoryHierarchyAccess {
	var memoryHierarchyAccess = &MemoryHierarchyAccess{
		MemoryHierarchy:memoryHierarchy,
		Id:memoryHierarchy.CurrentMemoryHierarchyAccessId,
		AccessType:accessType,
		ThreadId:threadId,
		VirtualPc:virtualPc,
		PhysicalAddress:physicalAddress,
		PhysicalTag:physicalTag,
		OnCompletedCallback:onCompletedCallback,
		BeginCycle:memoryHierarchy.Driver.CycleAccurateEventQueue().CurrentCycle,
	}

	memoryHierarchy.CurrentMemoryHierarchyAccessId++

	return memoryHierarchyAccess
}

func (memoryHierarchyAccess *MemoryHierarchyAccess) NumCycles() uint32 {
	return uint32(memoryHierarchyAccess.EndCycle - memoryHierarchyAccess.BeginCycle)
}

func (memoryHierarchyAccess *MemoryHierarchyAccess) Complete() {
	memoryHierarchyAccess.EndCycle = memoryHierarchyAccess.MemoryHierarchy.Driver.CycleAccurateEventQueue().CurrentCycle
	memoryHierarchyAccess.OnCompletedCallback()
	memoryHierarchyAccess.OnCompletedCallback = nil
}
