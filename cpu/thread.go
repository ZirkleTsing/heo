package cpu

import "reflect"

type Thread struct {
	Core                  *Core
	Num                   uint32
	Context               *Context
	NumDynamicInsts       uint64
	ExecutedMnemonicNames map[MnemonicName]uint64
	ExecutedSyscallNames  map[string]uint64
}

func NewThread(core *Core, num uint32) *Thread {
	var thread = &Thread{
		Core:core,
		Num:num,
		ExecutedMnemonicNames:make(map[MnemonicName]uint64),
		ExecutedSyscallNames:make(map[string]uint64),
	}

	core.Processor.Experiment.BlockingEventDispatcher.AddListener(reflect.TypeOf((*StaticInstExecutedEvent)(nil)), func(event interface{}) {
		var staticInstExecutedEvent = event.(*StaticInstExecutedEvent)

		if staticInstExecutedEvent.Context == thread.Context {
			var mnemonicName = staticInstExecutedEvent.StaticInst.Mnemonic.Name

			if _, ok := thread.ExecutedMnemonicNames[mnemonicName]; !ok {
				thread.ExecutedMnemonicNames[mnemonicName] = 0
			}

			thread.ExecutedMnemonicNames[mnemonicName]++
		}
	})

	core.Processor.Experiment.BlockingEventDispatcher.AddListener(reflect.TypeOf((*SyscallExecutedEvent)(nil)), func(event interface{}) {
		var syscallExecutedEvent = event.(*SyscallExecutedEvent)

		if syscallExecutedEvent.Context == thread.Context {
			var syscallName = syscallExecutedEvent.SyscallName

			if _, ok := thread.ExecutedSyscallNames[syscallName]; !ok {
				thread.ExecutedSyscallNames[syscallName] = 0
			}

			thread.ExecutedSyscallNames[syscallName]++
		}
	})

	return thread
}

func (thread *Thread) AdvanceOneCycle() {
	if thread.Context != nil && thread.Context.State == ContextState_RUNNING {
		var staticInst *StaticInst

		for {
			staticInst = thread.Context.DecodeNextStaticInst()
			staticInst.Execute(thread.Context)

			if staticInst.Mnemonic.Name != Mnemonic_NOP {
				thread.NumDynamicInsts++
			}

			if !(thread.Context != nil &&
				thread.Context.State == ContextState_RUNNING &&
				staticInst.Mnemonic.Name == Mnemonic_NOP) {
				break
			}
		}
	}
}
