package cpu

import (
	"reflect"
	"github.com/mcai/acogo/cpu/uncore"
)

type Thread interface {
	Core() Core
	Num() int32
	Id() int32
	Context() *Context
	SetContext(context *Context)
	FastForwardOneCycle()

	Itlb() *uncore.TranslationLookasideBuffer
	Dtlb() *uncore.TranslationLookasideBuffer
	WarmupOneCycle()

	NumDynamicInsts() int64
	ResetNumDynamicInsts()

	InstructionsPerCycle() float64
	CyclesPerInstructions() float64
}

type BaseThread struct {
	core                  Core
	num                   int32
	id                    int32
	context               *Context
	numDynamicInsts       int64

	ExecutedMnemonicNames map[MnemonicName]int32
	ExecutedSyscallNames  map[string]int32
}

func NewBaseThread(core Core, num int32) *BaseThread {
	var thread = &BaseThread{
		core:core,
		num:num,
		id:core.Num() * core.Processor().Experiment.CPUConfig.NumThreadsPerCore + num,
		ExecutedMnemonicNames:make(map[MnemonicName]int32),
		ExecutedSyscallNames:make(map[string]int32),
	}

	core.Processor().Experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)

		if staticInstExecutedEvent.Context == thread.Context() {
			var mnemonicName = staticInstExecutedEvent.StaticInst.Mnemonic.Name

			if _, ok := thread.ExecutedMnemonicNames[mnemonicName]; !ok {
				thread.ExecutedMnemonicNames[mnemonicName] = 0
			}

			thread.ExecutedMnemonicNames[mnemonicName]++
		}
	})

	core.Processor().Experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*SyscallExecutedEvent)(nil)), func(event interface{}) {
		var syscallExecutedEvent = event.(*SyscallExecutedEvent)

		if syscallExecutedEvent.Context == thread.Context() {
			var syscallName = syscallExecutedEvent.SyscallName

			if _, ok := thread.ExecutedSyscallNames[syscallName]; !ok {
				thread.ExecutedSyscallNames[syscallName] = 0
			}

			thread.ExecutedSyscallNames[syscallName]++
		}
	})

	return thread
}

func (thread *BaseThread) Core() Core {
	return thread.core
}

func (thread *BaseThread) Num() int32 {
	return thread.num
}

func (thread *BaseThread) Id() int32 {
	return thread.id
}

func (thread *BaseThread) Context() *Context {
	return thread.context
}

func (thread *BaseThread) SetContext(context *Context) {
	thread.context = context
}

func (thread *BaseThread) NumDynamicInsts() int64 {
	return thread.numDynamicInsts
}

func (thread *BaseThread) ResetNumDynamicInsts() {
	thread.numDynamicInsts = 0
}

func (thread *BaseThread) FastForwardOneCycle() {
	if thread.Context() != nil && thread.Context().State == ContextState_RUNNING {
		var staticInst *StaticInst

		for {
			staticInst = thread.Context().DecodeNextStaticInst()
			staticInst.Execute(thread.Context())

			if staticInst.Mnemonic.Name != Mnemonic_NOP {
				thread.numDynamicInsts++
			}

			if !(thread.Context() != nil &&
				thread.Context().State == ContextState_RUNNING &&
				staticInst.Mnemonic.Name == Mnemonic_NOP) {
				break
			}
		}
	}
}

func (thread *BaseThread) InstructionsPerCycle() float64 {
	if thread.Core().Processor().Experiment.CycleAccurateEventQueue().CurrentCycle == 0 {
		return float64(0)
	}
	return float64(thread.numDynamicInsts) / float64(thread.Core().Processor().Experiment.CycleAccurateEventQueue().CurrentCycle)
}

func (thread *BaseThread) CyclesPerInstructions() float64 {
	if thread.numDynamicInsts == 0 {
		return float64(0)
	}

	return float64(thread.Core().Processor().Experiment.CycleAccurateEventQueue().CurrentCycle) / float64(thread.numDynamicInsts)
}
