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
	var baseReorderBufferEntry = &BaseReorderBufferEntry{
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

	return baseReorderBufferEntry
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Id() int32 {
	return baseReorderBufferEntry.id
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Thread() Thread {
	return baseReorderBufferEntry.thread
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) DynamicInst() *DynamicInst {
	return baseReorderBufferEntry.dynamicInst
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Npc() uint32 {
	return baseReorderBufferEntry.npc
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Nnpc() uint32 {
	return baseReorderBufferEntry.nnpc
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) PredictedNnpc() uint32 {
	return baseReorderBufferEntry.predictedNnpc
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) ReturnAddressStackRecoverIndex() uint32 {
	return baseReorderBufferEntry.returnAddressStackRecoverIndex
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) BranchPredictorUpdate() *BranchPredictorUpdate {
	return baseReorderBufferEntry.branchPredictorUpdate
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Speculative() bool {
	return baseReorderBufferEntry.speculative
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) OldPhysicalRegisters() map[uint32]*PhysicalRegister {
	return baseReorderBufferEntry.oldPhysicalRegisters
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) TargetPhysicalRegisters() map[uint32]*PhysicalRegister {
	return baseReorderBufferEntry.targetPhysicalRegisters
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetTargetPhysicalRegisters(targetPhysicalRegisters map[uint32]*PhysicalRegister) {
	baseReorderBufferEntry.targetPhysicalRegisters = targetPhysicalRegisters
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SourcePhysicalRegisters() map[uint32]*PhysicalRegister {
	return baseReorderBufferEntry.sourcePhysicalRegisters
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetSourcePhysicalRegisters(sourcePhysicalRegisters map[uint32]*PhysicalRegister) {
	baseReorderBufferEntry.sourcePhysicalRegisters = sourcePhysicalRegisters
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Dispatched() bool {
	return baseReorderBufferEntry.dispatched
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetDispatched(dispatched bool) {
	baseReorderBufferEntry.dispatched = dispatched
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Issued() bool {
	return baseReorderBufferEntry.issued
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetIssued(issued bool) {
	baseReorderBufferEntry.issued = issued
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Completed() bool {
	return baseReorderBufferEntry.completed
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetCompleted(completed bool) {
	baseReorderBufferEntry.completed = completed
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) Squashed() bool {
	return baseReorderBufferEntry.squashed
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetSquashed(squashed bool) {
	baseReorderBufferEntry.squashed = squashed
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) NumNotReadyOperands() uint32 {
	return baseReorderBufferEntry.numNotReadyOperands
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) SetNumNotReadyOperands(numNotReadyOperands uint32) {
	baseReorderBufferEntry.numNotReadyOperands = numNotReadyOperands
}

func (baseReorderBufferEntry *BaseReorderBufferEntry) doWriteback() {
	for dependency, targetPhysicalRegister := range baseReorderBufferEntry.targetPhysicalRegisters {
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
