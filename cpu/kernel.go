package cpu

import "github.com/mcai/acogo/cpu/regs"

type Kernel struct {
	Pipes            []*Pipe
	SystemEvents     []SystemEvent
	SignalActions    []*SignalAction
	Contexts         []*Context
	Processes        []*Process
	SyscallEmulation *SyscallEmulation

	CurrentCycle     uint64
	CurrentPid       uint32
	CurrentMemoryId  uint32
	CurrentContextId uint32
	CurrentFd        int
}

func NewKernel() *Kernel {
	var kernel = &Kernel{
		SyscallEmulation:NewSyscallEmulation(),
		CurrentCycle:0,
		CurrentPid:1000,
		CurrentMemoryId:0,
		CurrentContextId:0,
		CurrentFd:100,
	}

	for i := 0; i < MAX_SIGNAL; i++ {
		kernel.SignalActions = append(kernel.SignalActions, NewSignalAction())
	}

	return kernel
}

func (kernel *Kernel) GetProcessFromId(id uint32) *Process {
	for _, process := range kernel.Processes {
		if process.Id == id {
			return process
		}
	}

	return nil
}

func (kernel *Kernel) GetContextFromId(id uint32) *Context {
	for _, context := range kernel.Contexts {
		if context.Id == id {
			return context
		}
	}

	return nil
}

func (kernel *Kernel) GetContextFromProcessId(processId uint32) *Context {
	for _, context := range kernel.Contexts {
		if context.ProcessId == processId {
			return context
		}
	}

	return nil
}

func (kernel *Kernel) ProcessSystemEvents() {
	//TODO
}

func (kernel *Kernel) ProcessSignals() {
	//TODO
}

func (kernel *Kernel) CreatePipe() []int {
	var fileDescriptors = make([]int, 2)

	fileDescriptors[0] = kernel.CurrentFd

	kernel.CurrentFd++

	fileDescriptors[1] = kernel.CurrentFd

	kernel.CurrentFd++

	kernel.Pipes = append(kernel.Pipes, NewPipe(fileDescriptors))

	return fileDescriptors
}

func (kernel *Kernel) getBuffer(fileDescriptor int, index uint32) *CircularByteBuffer {
	for _, pipe := range kernel.Pipes {
		if pipe.FileDescriptors[index] == fileDescriptor {
			return pipe.Buffer
		}
	}

	return nil
}

func (kernel *Kernel) GetReadBuffer(fileDescriptor int) *CircularByteBuffer {
	return kernel.getBuffer(fileDescriptor, 0)
}

func (kernel *Kernel) GetWriteBuffer(fileDescriptor int) *CircularByteBuffer {
	return kernel.getBuffer(fileDescriptor, 1)
}

func (kernel *Kernel) RunSignalHandler(context *Context, signal uint32) {
	if kernel.SignalActions[signal - 1].Handler == 0 {
		panic("Impossible")
	}

	context.SignalMasks.Pending.Clear(signal)

	var oldRegs = context.Regs.Clone()

	context.Regs.Gpr[regs.REGISTER_A0] = signal
	context.Regs.Gpr[regs.REGISTER_T9] = kernel.SignalActions[signal - 1].Handler
	context.Regs.Gpr[regs.REGISTER_RA] = 0xffffffff
	context.Regs.Npc = kernel.SignalActions[signal - 1].Handler
	context.Regs.Nnpc = context.Regs.Npc + 4

	for context.State == ContextState_RUNNING && context.Regs.Npc != 0xfffffff {
		//TODO
	}

	context.Regs = oldRegs
}

func (kernel *Kernel) MustProcessSignal(context *Context, signal uint32) bool {
	return context.SignalMasks.Pending.Contains(signal) && !context.SignalMasks.Blocked.Contains(signal)
}

func (kernel *Kernel) AdvanceOneCycle() {
	if kernel.CurrentCycle % 1000 == 0 {
		kernel.ProcessSystemEvents()
		kernel.ProcessSignals()
	}

	kernel.CurrentCycle++
}