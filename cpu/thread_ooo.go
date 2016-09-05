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
	var oooThread = &OoOThread{
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
		var physicalReg = oooThread.IntPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		oooThread.RenameTable[dep] = physicalReg
	}

	for i := uint32(0); i < regs.NUM_FP_REGISTERS; i++ {
		var dep = NewRegisterDependency(RegisterDependencyType_FP, i).ToInt()
		var physicalReg = oooThread.FpPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		oooThread.RenameTable[dep] = physicalReg
	}

	for i := uint32(0); i < regs.NUM_MISC_REGISTERS; i++ {
		var dep = NewRegisterDependency(RegisterDependencyType_MISC, i).ToInt()
		var physicalReg = oooThread.MiscPhysicalRegs.PhysicalRegisters[i]
		physicalReg.Reserve(dep)
		oooThread.RenameTable[dep] = physicalReg
	}

	return oooThread
}
