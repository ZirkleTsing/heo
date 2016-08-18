package acogo

type Process struct {
	Id                   uint32
	ContextMapping       *ContextMapping
	Environments         []string
	StdInFileDescriptor  uint32
	StdOutFileDescriptor uint32
	StackBase            uint32
	StackSize            uint32
	TextSize             uint32
	EnvironmentBase      uint32
	HeapTop              uint32
	DataTop              uint32
	ProgramEntry         uint32
	LittleEndian         bool
	Memory               *PagedMemory
}

func NewProcess(kernel *Kernel, contextMapping *ContextMapping) *Process {
	var process = &Process{
	}

	return process
}
