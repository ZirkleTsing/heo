package cpu

type PhysicalRegisterState string

const (
	PhysicalRegisterState_AVAILABLE = PhysicalRegisterState("AVAILABLE")
	PhysicalRegisterState_RENAME_BUFFER_INVALID = PhysicalRegisterState("RENAME_BUFFER_INVALID")
	PhysicalRegisterState_RENAME_BUFFER_VALID = PhysicalRegisterState("RENAME_BUFFER_VALID")
	PhysicalRegisterState_ARCHITECTURAL_REGISTER = PhysicalRegisterState("ARCHITECTURAL_REGISTER")
)

type PhysicalRegister struct {
	PhysicalRegisterFile                         *PhysicalRegisterFile
	State                                        PhysicalRegisterState
	Dependency                                   int32
	EffectiveAddressComputationOperandDependents []*ReorderBufferEntry
	StoreAddressDependents                       []*LoadStoreBufferEntry
	Dependents                                   []GeneralReorderBufferEntry
}

func NewPhysicalRegister(physicalRegisterFile *PhysicalRegisterFile) *PhysicalRegister {
	var physicalRegister = &PhysicalRegister{
		PhysicalRegisterFile:physicalRegisterFile,
		State:PhysicalRegisterState_AVAILABLE,
	}

	return physicalRegister
}

func (physicalRegister *PhysicalRegister) Reserve(dependency uint32) {
	if physicalRegister.State != PhysicalRegisterState_AVAILABLE {
		panic("Impossible")
	}

	physicalRegister.Dependency = int32(dependency)

	physicalRegister.State = PhysicalRegisterState_ARCHITECTURAL_REGISTER

	physicalRegister.PhysicalRegisterFile.NumFreePhysicalRegisters--
}

func (physicalRegister *PhysicalRegister) Allocate(dependency uint32) {
	if physicalRegister.State != PhysicalRegisterState_AVAILABLE {
		panic("Impossible")
	}

	physicalRegister.Dependency = int32(dependency)

	physicalRegister.State = PhysicalRegisterState_RENAME_BUFFER_INVALID

	physicalRegister.PhysicalRegisterFile.NumFreePhysicalRegisters--
}

func (physicalRegister *PhysicalRegister) Writeback() {
	if physicalRegister.State != PhysicalRegisterState_RENAME_BUFFER_INVALID {
		panic("Impossible")
	}

	physicalRegister.State = PhysicalRegisterState_RENAME_BUFFER_VALID

	for _, effectiveAddressComputationOperandDependent := range physicalRegister.EffectiveAddressComputationOperandDependents {
		effectiveAddressComputationOperandDependent.EffectiveAddressComputationOperandReady = true
	}

	for _, storeAddressDependent := range physicalRegister.StoreAddressDependents {
		storeAddressDependent.StoreAddressReady = true
	}

	for _, dependent := range physicalRegister.Dependents {
		dependent.SetNumNotReadyOperands(dependent.NumNotReadyOperands() + 1)
	}

	physicalRegister.EffectiveAddressComputationOperandDependents = []*ReorderBufferEntry{}
	physicalRegister.StoreAddressDependents = []*LoadStoreBufferEntry{}
	physicalRegister.Dependents = []GeneralReorderBufferEntry{}
}

func (physicalRegister *PhysicalRegister) Commit() {
	if physicalRegister.State != PhysicalRegisterState_RENAME_BUFFER_VALID {
		panic("Impossible")
	}

	physicalRegister.State = PhysicalRegisterState_ARCHITECTURAL_REGISTER
}

func (physicalRegister *PhysicalRegister) Recover() {
	if physicalRegister.State != PhysicalRegisterState_RENAME_BUFFER_INVALID &&
		physicalRegister.State != PhysicalRegisterState_RENAME_BUFFER_VALID {
		panic("Impossible")
	}

	physicalRegister.Dependency = -1

	physicalRegister.State = PhysicalRegisterState_AVAILABLE

	physicalRegister.PhysicalRegisterFile.NumFreePhysicalRegisters++
}

func (physicalRegister *PhysicalRegister) Reclaim() {
	if physicalRegister.State != PhysicalRegisterState_ARCHITECTURAL_REGISTER {
		panic("Impossible")
	}

	physicalRegister.Dependency = -1

	physicalRegister.State = PhysicalRegisterState_AVAILABLE

	physicalRegister.PhysicalRegisterFile.NumFreePhysicalRegisters++
}

func (physicalRegister *PhysicalRegister) Ready() bool {
	return physicalRegister.State == PhysicalRegisterState_RENAME_BUFFER_VALID ||
		physicalRegister.State == PhysicalRegisterState_ARCHITECTURAL_REGISTER
}

type PhysicalRegisterFile struct {
	PhysicalRegisters        []*PhysicalRegister
	NumFreePhysicalRegisters uint32
}

func NewPhysicalRegisterFile(size uint32) *PhysicalRegisterFile {
	var physicalRegisterFile = &PhysicalRegisterFile{
		NumFreePhysicalRegisters:size,
	}

	for i := uint32(0); i < size; i++ {
		physicalRegisterFile.PhysicalRegisters = append(
			physicalRegisterFile.PhysicalRegisters,
			NewPhysicalRegister(physicalRegisterFile),
		)
	}

	return physicalRegisterFile
}

func (physicalRegisterFile *PhysicalRegisterFile) Allocate(dependency uint32) *PhysicalRegister {
	for _, physicalRegister := range physicalRegisterFile.PhysicalRegisters {
		if physicalRegister.State == PhysicalRegisterState_AVAILABLE {
			physicalRegister.Allocate(dependency)
			return physicalRegister
		}
	}

	panic("Impossible")
}

func (physicalRegisterFile *PhysicalRegisterFile) Full() bool {
	return physicalRegisterFile.NumFreePhysicalRegisters == 0
}
