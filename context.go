package acogo

const (
	ContextState_IDLE = 0
	ContextState_BLOCKED = 1
	ContextState_RUNNING = 2
	ContextState_FINISHED = 3
)

type ContextState uint32

type Context struct {
	Id               uint32
	State            ContextState
	SignalMasks      *SignalMasks
	SignalFinish     uint32
	Regs             *ArchitecturalRegisterFile
	Kernel           *Kernel
	ThreadId         uint32
	UserId           uint32
	EffectiveUserId  uint32
	GroupId          uint32
	EffectiveGroupId uint32
	ProcessId        uint32
	Process          *Process
	Parent           *Context
}

func NewContext(kernel *Kernel, process *Process, parent *Context, regs *ArchitecturalRegisterFile, signalFinish uint32) *Context {
	var context = &Context{
		Kernel:kernel,
		Parent:parent,
		Regs:regs,
		SignalFinish:signalFinish,
		Id: kernel.CurrentContextId,
		ProcessId:kernel.CurrentPid,
		SignalMasks:NewSignalMasks(),
		State:ContextState_IDLE,
		Process:process,
	}

	kernel.CurrentContextId++
	kernel.CurrentPid++

	return context
}

func (context *Context) DecodeNextStaticInstruction() *StaticInst {
	context.Regs.Pc = context.Regs.Npc
	context.Regs.Npc = context.Regs.Nnpc
	context.Regs.Nnpc = context.Regs.Nnpc + 4
	context.Regs.Gpr[REGISTER_ZERO] = 0

	return context.Decode(context.Regs.Pc)
}

func (context *Context) Decode(mappedPc uint32) *StaticInst {
	//return context.Process.GetStaticInst(mappedPc)
	return nil // TODO
}

func (context *Context) Suspend() {
	if context.State == ContextState_BLOCKED {
		panic("Impossible")
	}

	context.State = ContextState_BLOCKED
}

func (context *Context) Resume() {
	if context.State != ContextState_BLOCKED {
		panic("Impossible")
	}

	context.State = ContextState_RUNNING
}

func (context *Context) Finish() {
	if context.State == ContextState_FINISHED {
		panic("Impossible")
	}

	context.State = ContextState_FINISHED

	for _, c := range context.Kernel.Contexts {
		if c.State != ContextState_FINISHED && c.Parent == context {
			c.Finish()
		}
	}

	if context.SignalFinish != 0 && context.Parent != nil {
		context.Parent.SignalMasks.Pending.Set(context.SignalFinish)
	}
}

type ContextMapping struct {
	ThreadId   uint32
	Executable string
	Arguments  string
}

func NewContextMapping(threadId uint32, executable string, arguments string) *ContextMapping {
	var contextMapping = &ContextMapping{
		ThreadId:threadId,
		Executable:executable,
		Arguments:arguments,
	}

	return contextMapping
}
