package cpu

type GeneralReorderBufferEntry interface {
	Id() int32
	Thread() Thread
	DynamicInst() *DynamicInst

	Npc() uint32
	Nnpc() uint32
	PredictedNnpc() uint32

	ReturnAddressStackRecoverTop() uint32
	BranchPredictorUpdate() *BranchPredictorUpdate
	Speculative() bool

	OldPhysicalRegisters() map[*RegisterDependency]*PhysicalRegister

	TargetPhysicalRegisters() map[*RegisterDependency]*PhysicalRegister
	SetTargetPhysicalRegisters(targetPhysicalRegisters map[*RegisterDependency]*PhysicalRegister)

	SourcePhysicalRegisters() map[*RegisterDependency]*PhysicalRegister
	SetSourcePhysicalRegisters(sourcePhysicalRegisters map[*RegisterDependency]*PhysicalRegister)

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
	id                           int32
	thread                       Thread
	dynamicInst                  *DynamicInst

	npc                          uint32
	nnpc                         uint32
	predictedNnpc                uint32

	returnAddressStackRecoverTop uint32
	branchPredictorUpdate        *BranchPredictorUpdate
	speculative                  bool

	oldPhysicalRegisters         map[*RegisterDependency]*PhysicalRegister
	targetPhysicalRegisters      map[*RegisterDependency]*PhysicalRegister
	sourcePhysicalRegisters      map[*RegisterDependency]*PhysicalRegister

	dispatched                   bool
	issued                       bool
	completed                    bool
	squashed                     bool

	numNotReadyOperands          uint32
}

func NewBaseReorderBufferEntry(thread Thread, dynamicInst *DynamicInst, npc uint32, nnpc uint32, predictedNnpc uint32, returnAddressStackRecoverTop uint32, branchPredictorUpdate *BranchPredictorUpdate, speculative bool) *BaseReorderBufferEntry {
	var reorderBufferEntry = &BaseReorderBufferEntry{
		id:thread.Core().Processor().Experiment.OoO.CurrentReorderBufferEntryId,
		thread:thread,
		dynamicInst:dynamicInst,

		npc:npc,
		nnpc:nnpc,
		predictedNnpc:predictedNnpc,

		returnAddressStackRecoverTop:returnAddressStackRecoverTop,
		branchPredictorUpdate:branchPredictorUpdate,
		speculative:speculative,

		oldPhysicalRegisters:make(map[*RegisterDependency]*PhysicalRegister),
		targetPhysicalRegisters:make(map[*RegisterDependency]*PhysicalRegister),
		sourcePhysicalRegisters:make(map[*RegisterDependency]*PhysicalRegister),
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

func (reorderBufferEntry *BaseReorderBufferEntry) ReturnAddressStackRecoverTop() uint32 {
	return reorderBufferEntry.returnAddressStackRecoverTop
}

func (reorderBufferEntry *BaseReorderBufferEntry) BranchPredictorUpdate() *BranchPredictorUpdate {
	return reorderBufferEntry.branchPredictorUpdate
}

func (reorderBufferEntry *BaseReorderBufferEntry) Speculative() bool {
	return reorderBufferEntry.speculative
}

func (reorderBufferEntry *BaseReorderBufferEntry) OldPhysicalRegisters() map[*RegisterDependency]*PhysicalRegister {
	return reorderBufferEntry.oldPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) TargetPhysicalRegisters() map[*RegisterDependency]*PhysicalRegister {
	return reorderBufferEntry.targetPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetTargetPhysicalRegisters(targetPhysicalRegisters map[*RegisterDependency]*PhysicalRegister) {
	reorderBufferEntry.targetPhysicalRegisters = targetPhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SourcePhysicalRegisters() map[*RegisterDependency]*PhysicalRegister {
	return reorderBufferEntry.sourcePhysicalRegisters
}

func (reorderBufferEntry *BaseReorderBufferEntry) SetSourcePhysicalRegisters(sourcePhysicalRegisters map[*RegisterDependency]*PhysicalRegister) {
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
		if dependency != nil {
			targetPhysicalRegister.Writeback()
		}
	}
}

type ReorderBufferEntry struct {
	*BaseReorderBufferEntry

	EffectiveAddressComputation             bool
	LoadStoreBufferEntry                    *LoadStoreQueueEntry
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

type LoadStoreQueueEntry struct {
	*BaseReorderBufferEntry

	EffectiveAddress  int32
	StoreAddressReady bool
}

func NewLoadStoreQueueEntry(thread Thread, dynamicInst *DynamicInst, npc uint32, nnpc uint32, predictedNnpc uint32, returnAddressStackRecoverIndex uint32, branchPredictorUpdate *BranchPredictorUpdate, speculative bool) *LoadStoreQueueEntry {
	var loadStoreQueueEntry = &LoadStoreQueueEntry{
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

		EffectiveAddress:-1,
	}

	return loadStoreQueueEntry
}

func (loadStoreQueueEntry *LoadStoreQueueEntry) Writeback() {
	loadStoreQueueEntry.doWriteback()
}

func (loadStoreQueueEntry *LoadStoreQueueEntry) AllOperandReady() bool {
	return loadStoreQueueEntry.numNotReadyOperands == 0
}

func SignalCompleted(loadStoreQueueEntry GeneralReorderBufferEntry) {
	if !loadStoreQueueEntry.Squashed() {
		loadStoreQueueEntry.Thread().Core().SetOoOEventQueue(
			append(
				loadStoreQueueEntry.Thread().Core().OoOEventQueue(),
				loadStoreQueueEntry,
			),
		)
	}
}
