package cpu

type GeneralReorderBufferEntry interface {
	Id() int32
	Thread() Thread
	DynamicInst() *DynamicInst

	Npc() uint32
	Nnpc() uint32
	PredictedNnpc() uint32

	ReturnAddressStackRecoverIndex() uint32
	BranchPredictorUpdate() *BranchPredictorUpdate
	Speculative() bool

	OldPhysicalRegisters() map[uint32]*PhysicalRegister
	TargetPhysicalRegisters() map[uint32]*PhysicalRegister
	SetTargetPhysicalRegisters(targetPhysicalRegisters map[uint32]*PhysicalRegister)
	SourcePhysicalRegisters() map[uint32]*PhysicalRegister
	SetSourcePhysicalRegisters(sourcePhysicalRegisters map[uint32]*PhysicalRegister)

	Dispatched() bool
	SetDispatched(dispatched bool)
	Issued() bool
	SetIssued(issued bool)
	Completed() bool
	SetCompleted(completed bool)
	Squashed() bool
	SetSquashed(squashed bool)

	NumNotReadyOperands() uint32
	SetNumNotReadyOperands(numNotReadyOperands uint32)

	Writeback()
}

type BaseReorderBufferEntry struct {
	id                             int32
	thread                         Thread
	dynamicInst                    *DynamicInst

	npc                            uint32
	nnpc                           uint32
	predictedNnpc                  uint32

	returnAddressStackRecoverIndex uint32
	branchPredictorUpdate          *BranchPredictorUpdate
	speculative                    bool

	oldPhysicalRegisters           map[uint32]*PhysicalRegister
	targetPhysicalRegisters        map[uint32]*PhysicalRegister
	sourcePhysicalRegisters        map[uint32]*PhysicalRegister

	dispatched                     bool
	issued                         bool
	completed                      bool
	squashed                       bool

	numNotReadyOperands            uint32
}

func NewBaseReorderBufferEntry(thread Thread, dynamicInst *DynamicInst, npc uint32, nnpc uint32, predictedNnpc uint32, returnAddressStackRecoverIndex uint32, branchPredictorUpdate *BranchPredictorUpdate, speculative bool) *BaseReorderBufferEntry {
	var reorderBufferEntry = &BaseReorderBufferEntry{
		id:thread.Core().Processor().Experiment.OoO.CurrentReorderBufferEntryId,
		thread:thread,
		dynamicInst:dynamicInst,

		npc:npc,
		nnpc:nnpc,
		predictedNnpc:predictedNnpc,

		returnAddressStackRecoverIndex:returnAddressStackRecoverIndex,
		branchPredictorUpdate:branchPredictorUpdate,
		speculative:speculative,
	}

	thread.Core().Processor().Experiment.OoO.CurrentReorderBufferEntryId++

	return reorderBufferEntry
}

func (reorderBufferEntry *BaseReorderBufferEntry) Id() int32 {
	return reorderBufferEntry.id
}

func (reorderBufferEntry *BaseReorderBufferEntry) Thread() Thread {
	return reorderBufferEntry.thread
}

func (reorderBufferEntry *BaseReorderBufferEntry) DynamicInst() *DynamicInst {
	return reorderBufferEntry.dynamicInst
}

func (reorderBufferEntry *BaseReorderBufferEntry) Npc() uint32 {
	return reorderBufferEntry.npc
}

func (reorderBufferEntry *BaseReorderBufferEntry) Nnpc() uint32 {
	return reorderBufferEntry.nnpc
}

func (reorderBufferEntry *BaseReorderBufferEntry) PredictedNnpc() uint32 {
	return reorderBufferEntry.predictedNnpc
}

func (reorderBufferEntry *BaseReorderBufferEntry) ReturnAddressStackRecoverIndex() uint32 {
	return reorderBufferEntry.returnAddressStackRecoverIndex
}

func (reorderBufferEntry *BaseReorderBufferEntry) BranchPredictorUpdate() *BranchPredictorUpdate {
	return reorderBufferEntry.branchPredictorUpdate
}

func (reorderBufferEntry *BaseReorderBufferEntry) Speculative() bool {
	return reorderBufferEntry.speculative
}

func (reorderBufferEntry *BaseReorderBufferEntry) OldPhysicalRegisters() map[uint32]*PhysicalRegister {
	return reorderBufferEntry.oldPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) TargetPhysicalRegisters() map[uint32]*PhysicalRegister {
	return reorderBufferEntry.targetPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetTargetPhysicalRegisters(targetPhysicalRegisters map[uint32]*PhysicalRegister) {
	reorderBufferEntry.targetPhysicalRegisters = targetPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SourcePhysicalRegisters() map[uint32]*PhysicalRegister {
	return reorderBufferEntry.sourcePhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetSourcePhysicalRegisters(sourcePhysicalRegisters map[uint32]*PhysicalRegister) {
	reorderBufferEntry.sourcePhysicalRegisters = sourcePhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) Dispatched() bool {
	return reorderBufferEntry.dispatched
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetDispatched(dispatched bool) {
	reorderBufferEntry.dispatched = dispatched
}

func (reorderBufferEntry *BaseReorderBufferEntry) Issued() bool {
	return reorderBufferEntry.issued
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetIssued(issued bool) {
	reorderBufferEntry.issued = issued
}

func (reorderBufferEntry *BaseReorderBufferEntry) Completed() bool {
	return reorderBufferEntry.completed
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetCompleted(completed bool) {
	reorderBufferEntry.completed = completed
}

func (reorderBufferEntry *BaseReorderBufferEntry) Squashed() bool {
	return reorderBufferEntry.squashed
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetSquashed(squashed bool) {
	reorderBufferEntry.squashed = squashed
}

func (reorderBufferEntry *BaseReorderBufferEntry) NumNotReadyOperands() uint32 {
	return reorderBufferEntry.numNotReadyOperands
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetNumNotReadyOperands(numNotReadyOperands uint32) {
	reorderBufferEntry.numNotReadyOperands = numNotReadyOperands
}

func (reorderBufferEntry *BaseReorderBufferEntry) doWriteback() {
	for dependency, targetPhysicalRegister := range reorderBufferEntry.targetPhysicalRegisters {
		if dependency != 0 {
			targetPhysicalRegister.Writeback()
		}
	}
}

type ReorderBufferEntry struct {
	*BaseReorderBufferEntry

	EffectiveAddressComputation             bool
	LoadStoreBufferEntry                    *LoadStoreBufferEntry
	EffectiveAddressComputationOperandReady bool
}

func NewReorderBufferEntry(thread Thread, dynamicInst *DynamicInst, npc uint32, nnpc uint32, predictedNnpc uint32, returnAddressStackRecoverIndex uint32, branchPredictorUpdate *BranchPredictorUpdate, speculative bool) *ReorderBufferEntry {
	var reorderBufferEntry = &ReorderBufferEntry{
		BaseReorderBufferEntry:NewBaseReorderBufferEntry(
			thread,
			dynamicInst,
			npc,
			nnpc,
			predictedNnpc,
			returnAddressStackRecoverIndex,
			branchPredictorUpdate,
			speculative,
		),
	}

	return reorderBufferEntry
}

func (reorderBufferEntry *ReorderBufferEntry) Writeback() {
	if !reorderBufferEntry.EffectiveAddressComputation {
		reorderBufferEntry.doWriteback()
	}
}

func (reorderBufferEntry *ReorderBufferEntry) AllOperandReady() bool {
	if reorderBufferEntry.EffectiveAddressComputation {
		return reorderBufferEntry.EffectiveAddressComputationOperandReady
	}

	return reorderBufferEntry.numNotReadyOperands == 0
}

type LoadStoreBufferEntry struct {
	*BaseReorderBufferEntry

	EffectiveAddress  uint32
	StoreAddressReady bool
}

func NewLoadStoreBufferEntry(thread Thread, dynamicInst *DynamicInst, npc uint32, nnpc uint32, predictedNnpc uint32, returnAddressStackRecoverIndex uint32, branchPredictorUpdate *BranchPredictorUpdate, speculative bool) *LoadStoreBufferEntry {
	var loadStoreBufferEntry = &LoadStoreBufferEntry{
		BaseReorderBufferEntry:NewBaseReorderBufferEntry(
			thread,
			dynamicInst,
			npc,
			nnpc,
			predictedNnpc,
			returnAddressStackRecoverIndex,
			branchPredictorUpdate,
			speculative,
		),
	}

	return loadStoreBufferEntry
}

func (loadStoreBufferEntry *LoadStoreBufferEntry) Writeback() {
	loadStoreBufferEntry.doWriteback()
}

func (loadStoreBufferEntry *LoadStoreBufferEntry) AllOperandReady() bool {
	return loadStoreBufferEntry.numNotReadyOperands == 0
}

func SignalCompleted(generalReorderBufferEntry GeneralReorderBufferEntry) {
	if !generalReorderBufferEntry.Squashed() {
		generalReorderBufferEntry.Thread().Core().SetOoOEventQueue(
			append(
				generalReorderBufferEntry.Thread().Core().OoOEventQueue(),
				generalReorderBufferEntry,
			),
		)
	}
}
