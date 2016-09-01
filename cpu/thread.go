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
	NumDynamicInsts() int32
	FastForwardOneCycle()

	Itlb() *uncore.TranslationLookasideBuffer
	Dtlb() *uncore.TranslationLookasideBuffer
	WarmupOneCycle()
}

type BaseThread struct {
	core                  Core
	num                   int32
	id                    int32
	context               *Context
	numDynamicInsts       int32

	ExecutedMnemonicNames map[MnemonicName]int32
	ExecutedSyscallNames  map[string]int32
}

func NewBaseThread(core Core, num int32) *BaseThread {
	var baseThread = &BaseThread{
		core:core,
		num:num,
		id:core.Num() * core.Processor().Experiment.CPUConfig.NumThreadsPerCore + num,
		ExecutedMnemonicNames:make(map[MnemonicName]int32),
		ExecutedSyscallNames:make(map[string]int32),
	}

	core.Processor().Experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)

		if staticInstExecutedEvent.Context == baseThread.Context() {
			var mnemonicName = staticInstExecutedEvent.StaticInst.Mnemonic.Name

			if _, ok := baseThread.ExecutedMnemonicNames[mnemonicName]; !ok {
				baseThread.ExecutedMnemonicNames[mnemonicName] = 0
			}

			baseThread.ExecutedMnemonicNames[mnemonicName]++
		}
	})

	core.Processor().Experiment.BlockingEventDispatcher().AddListener(reflect.TypeOf((*SyscallExecutedEvent)(nil)), func(event interface{}) {
		var syscallExecutedEvent = event.(*SyscallExecutedEvent)

		if syscallExecutedEvent.Context == baseThread.Context() {
			var syscallName = syscallExecutedEvent.SyscallName

			if _, ok := baseThread.ExecutedSyscallNames[syscallName]; !ok {
				baseThread.ExecutedSyscallNames[syscallName] = 0
			}

			baseThread.ExecutedSyscallNames[syscallName]++
		}
	})

	return baseThread
}

func (baseThread *BaseThread) Core() Core {
	return baseThread.core
}

func (baseThread *BaseThread) Num() int32 {
	return baseThread.num
}

func (baseThread *BaseThread) Id() int32 {
	return baseThread.id
}

func (baseThread *BaseThread) Context() *Context {
	return baseThread.context
}

func (baseThread *BaseThread) SetContext(context *Context) {
	baseThread.context = context
}

func (baseThread *BaseThread) NumDynamicInsts() int32 {
	return baseThread.numDynamicInsts
}

func (baseThread *BaseThread) FastForwardOneCycle() {
	if baseThread.Context() != nil && baseThread.Context().State == ContextState_RUNNING {
		var staticInst *StaticInst

		for {
			staticInst = baseThread.Context().DecodeNextStaticInst()
			staticInst.Execute(baseThread.Context())

			if staticInst.Mnemonic.Name != Mnemonic_NOP {
				baseThread.numDynamicInsts++
			}

			if !(baseThread.Context() != nil &&
				baseThread.Context().State == ContextState_RUNNING &&
				staticInst.Mnemonic.Name == Mnemonic_NOP) {
				break
			}
		}
	}
}
