package uncore

type MemoryHierarchyAccessType uint32

const (
	MemoryHierarchyAccessType_IFETCH = MemoryHierarchyAccessType(0)
	MemoryHierarchyAccessType_LOAD = MemoryHierarchyAccessType(1)
	MemoryHierarchyAccessType_STORE = MemoryHierarchyAccessType(2)
	MemoryHierarchyAccessType_UNKNOWN = MemoryHierarchyAccessType(3)
)

func (memoryHierarchyAccessType MemoryHierarchyAccessType) IsRead() bool {
	return memoryHierarchyAccessType == MemoryHierarchyAccessType_IFETCH ||
		memoryHierarchyAccessType == MemoryHierarchyAccessType_LOAD
}

func (memoryHierarchyAccessType MemoryHierarchyAccessType) IsWrite() bool {
	return memoryHierarchyAccessType == MemoryHierarchyAccessType_STORE
}

type MemoryHierarchyAccess struct {
	MemoryHierarchy     MemoryHierarchy
	Id                  int32
	AccessType          MemoryHierarchyAccessType

	ThreadId            int32
	VirtualPc           int32
	PhysicalAddress     uint32
	PhysicalTag         uint32

	OnCompletedCallback func()

	Aliases             []*MemoryHierarchyAccess

	BeginCycle          int64
	EndCycle            int64
}

func NewMemoryHierarchyAccess(memoryHierarchy MemoryHierarchy, accessType MemoryHierarchyAccessType, threadId int32, virtualPc int32, physicalAddress uint32, physicalTag uint32, onCompletedCallback func()) *MemoryHierarchyAccess {
	var memoryHierarchyAccess = &MemoryHierarchyAccess{
		MemoryHierarchy:memoryHierarchy,
		Id:memoryHierarchy.CurrentMemoryHierarchyAccessId(),
		AccessType:accessType,
		ThreadId:threadId,
		VirtualPc:virtualPc,
		PhysicalAddress:physicalAddress,
		PhysicalTag:physicalTag,
		OnCompletedCallback:onCompletedCallback,
		BeginCycle:memoryHierarchy.Driver().CycleAccurateEventQueue().CurrentCycle,
	}

	memoryHierarchy.SetCurrentMemoryHierarchyAccessId(
		memoryHierarchy.CurrentMemoryHierarchyAccessId() + 1,
	)

	return memoryHierarchyAccess
}

func (memoryHierarchyAccess *MemoryHierarchyAccess) NumCycles() uint32 {
	return uint32(memoryHierarchyAccess.EndCycle - memoryHierarchyAccess.BeginCycle)
}

func (memoryHierarchyAccess *MemoryHierarchyAccess) Complete() {
	memoryHierarchyAccess.EndCycle = memoryHierarchyAccess.MemoryHierarchy.Driver().CycleAccurateEventQueue().CurrentCycle
	memoryHierarchyAccess.OnCompletedCallback()
	memoryHierarchyAccess.OnCompletedCallback = nil
}
