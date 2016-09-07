package cpu

import "github.com/mcai/acogo/cpu/regs"

type OoOThread struct {
	*MemoryHierarchyThread

	BranchPredictor                 *BranchPredictor

	IntPhysicalRegs                 *PhysicalRegisterFile
	FpPhysicalRegs                  *PhysicalRegisterFile
	MiscPhysicalRegs                *PhysicalRegisterFile

	RenameTable                     map[uint32]*PhysicalRegister

	DecodeBuffer                    *PipelineBuffer
	ReorderBuffer                   *PipelineBuffer
	LoadStoreQueue                  *PipelineBuffer

	FetchNpc                        uint32
	FetchNnpc                       uint32

	LastDecodedDynamicInst          *DynamicInst
	LastDecodedDynamicInstCommitted bool
}

func NewOoOThread(core Core, num int32) *OoOThread {
	var thread = &OoOThread{
		MemoryHierarchyThread:NewMemoryHierarchyThread(core, num),

		BranchPredictor:NewBranchPredictor(
			core.Processor().Experiment.CPUConfig.BranchPredictorSize,
			core.Processor().Experiment.CPUConfig.BranchTargetBufferNumSets,
			core.Processor().Experiment.CPUConfig.BranchTargetBufferAssoc,
			core.Processor().Experiment.CPUConfig.ReturnAddressStackSize,
		),

		IntPhysicalRegs:NewPhysicalRegisterFile(core.Processor().Experiment.CPUConfig.PhysicalRegisterFileSize),
		FpPhysicalRegs:NewPhysicalRegisterFile(core.Processor().Experiment.CPUConfig.PhysicalRegisterFileSize),
		MiscPhysicalRegs:NewPhysicalRegisterFile(core.Processor().Experiment.CPUConfig.PhysicalRegisterFileSize),

		RenameTable:make(map[uint32]*PhysicalRegister),

		DecodeBuffer:NewPipelineBuffer(core.Processor().Experiment.CPUConfig.DecodeBufferSize),
		ReorderBuffer:NewPipelineBuffer(core.Processor().Experiment.CPUConfig.ReorderBufferSize),
		LoadStoreQueue:NewPipelineBuffer(core.Processor().Experiment.CPUConfig.LoadStoreQueueSize),
	}

	for i := uint32(0); i < regs.NUM_INT_REGISTERS; i++ {
		var dep = NewRegisterDependency(RegisterDependencyType_INT, i).ToInt()
		var physicalReg = thread.IntPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		thread.RenameTable[dep] = physicalReg
	}

	for i := uint32(0); i < regs.NUM_FP_REGISTERS; i++ {
		var dep = NewRegisterDependency(RegisterDependencyType_FP, i).ToInt()
		var physicalReg = thread.FpPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		thread.RenameTable[dep] = physicalReg
	}

	for i := uint32(0); i < regs.NUM_MISC_REGISTERS; i++ {
		var dep = NewRegisterDependency(RegisterDependencyType_MISC, i).ToInt()
		var physicalReg = thread.MiscPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		thread.RenameTable[dep] = physicalReg
	}

	return thread
}

func (thread *OoOThread) UpdateFetchNpcAndNnpcFromRegs() {
	thread.FetchNpc = thread.Context().Regs().Npc
	thread.FetchNnpc = thread.Context().Regs().Nnpc
}

func (thread *OoOThread) CanFetch() bool {
	if thread.FetchStalled {
		return false
	}

	var cacheLineToFetch = thread.Core().L1IController().Cache.GetTag(thread.FetchNpc)
	if int32(cacheLineToFetch) != thread.LastFetchedCacheLine {
		if !thread.Core().CanIfetch(thread, thread.FetchNpc) {
			return false
		} else {
			thread.Core().Ifetch(thread, thread.FetchNpc, thread.FetchNpc, func() {
				thread.FetchStalled = false
			})

			thread.FetchStalled = true
			thread.LastFetchedCacheLine = int32(cacheLineToFetch)

			return false
		}
	}

	return true
}

func (thread *OoOThread) Fetch() {
	if !thread.CanFetch() {
		return
	}

	var hasDone = false

	for !hasDone {
		if thread.context.State != ContextState_RUNNING {
			break
		}

		if thread.DecodeBuffer.Full() {
			break
		}

		if thread.context.Regs().Npc != thread.FetchNpc {

		}
	}
}