package cpu

import (
	"github.com/mcai/acogo/cpu/regs"
	"github.com/mcai/acogo/cpu/mem"
)

type Kernel struct {
	Experiment          *CPUExperiment

	Pipes               []*Pipe
	SystemEvents        []SystemEvent
	SignalActions       []*SignalAction
	Contexts            []*Context
	Processes           []*Process
	SyscallEmulation    *SyscallEmulation

	CurrentCycle        int
	CurrentPid          int
	CurrentProcessId    int
	CurrentMemoryId     int
	CurrentMemoryPageId int
	CurrentContextId    int
	CurrentFd           int
}

func NewKernel(experiment *CPUExperiment) *Kernel {
	var kernel = &Kernel{
		Experiment:experiment,
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

	for _, contextMapping := range experiment.Config.ContextMappings {
		var context = LoadContext(kernel, contextMapping)

		if !kernel.Map(context, func(candidateThreadId int) bool {
			return candidateThreadId == contextMapping.ThreadId
		}) {
			panic("Impossible")
		}

		kernel.Contexts = append(kernel.Contexts, context)
	}

	return kernel
}

func (kernel *Kernel) GetProcessFromId(id int) *Process {
	for _, process := range kernel.Processes {
		if process.Id == id {
			return process
		}
	}

	return nil
}

func (kernel *Kernel) GetContextFromId(id int) *Context {
	for _, context := range kernel.Contexts {
		if context.Id == id {
			return context
		}
	}

	return nil
}

func (kernel *Kernel) GetContextFromProcessId(processId int) *Context {
	for _, context := range kernel.Contexts {
		if context.ProcessId == processId {
			return context
		}
	}

	return nil
}

func (kernel *Kernel) Map(contextToMap *Context, predicate func(candidateThreadId int) bool) bool {
	if contextToMap.ThreadId != -1 {
		panic("Impossible")
	}

	for coreNum := 0; coreNum < kernel.Experiment.Config.NumCores; coreNum++ {
		for threadNum := 0; threadNum < kernel.Experiment.Config.NumThreadsPerCore; threadNum++ {
			var threadId = coreNum * kernel.Experiment.Config.NumThreadsPerCore + threadNum

			var hasMapped bool

			for _, context := range kernel.Contexts {
				if context.ThreadId == threadId {
					hasMapped = true
					break
				}
			}

			if !hasMapped && predicate(threadId) {
				contextToMap.ThreadId = threadId
				return true
			}
		}
	}

	return false
}

func (kernel *Kernel) ProcessSystemEvents() {
	var systemEventsToPreserve []SystemEvent

	for _, e := range kernel.SystemEvents {
		if (e.Context().State == ContextState_RUNNING || e.Context().State == ContextState_BLOCKED) && !e.NeedProcess() {
			e.Process()
		} else {
			systemEventsToPreserve = append(systemEventsToPreserve, e)
		}
	}

	kernel.SystemEvents = systemEventsToPreserve
}

func (kernel *Kernel) ProcessSignals() {
	for _, context := range kernel.Contexts {
		if context.State == ContextState_RUNNING || context.State == ContextState_BLOCKED {
			for signal := uint32(1); signal <= MAX_SIGNAL; signal++ {
				if kernel.MustProcessSignal(context, signal) {
					kernel.RunSignalHandler(context, signal)
				}
			}
		}
	}
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

func (kernel *Kernel) getBuffer(fileDescriptor int, index uint32) *mem.CircularByteBuffer {
	for _, pipe := range kernel.Pipes {
		if pipe.FileDescriptors[index] == fileDescriptor {
			return pipe.Buffer
		}
	}

	return nil
}

func (kernel *Kernel) GetReadBuffer(fileDescriptor int) *mem.CircularByteBuffer {
	return kernel.getBuffer(fileDescriptor, 0)
}

func (kernel *Kernel) GetWriteBuffer(fileDescriptor int) *mem.CircularByteBuffer {
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
		context.DecodeNextStaticInst().Execute(context)
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