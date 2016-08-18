package acogo

const (
	ContextState_IDLE = 0
	ContextState_BLOCKED = 1
	ContextState_RUNNING = 2
	ContextState_FINISHED = 3
)

type ContextState uint32

type Context struct {
	Id               int
	State            ContextState
	SignalMasks      SignalMasks
	SignalFinish     uint32
	Memory           *Memory
	Regs             *ArchitecturalRegisterFile
	Kernel           *Kernel
	ThreadId         int32
	UserId           int32
	EffectiveUserId  int32
	GroupId          int32
	EffectiveGroupId int32
	ProcessId        int32
	Process          *Process
	Parent           *Context
}

func NewContext(id int, littleEndian bool) *Context {
	var context = &Context{
		Id: id,
		Memory:NewMemory(littleEndian),
		Regs:NewArchitecturalRegisterFile(littleEndian),
	}

	return context
}

type ContextMapping struct {
}

func NewContextMapping() *ContextMapping {
	var contextMapping = &ContextMapping{

	}

	return contextMapping
}
