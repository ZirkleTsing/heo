package cpu

type Thread struct {
	Core            *Core
	Num             uint32
	Context         *Context
	NumInstructions uint64
}

func NewThread(core *Core, num uint32) *Thread {
	var thread = &Thread{
		Core:core,
		Num:num,
	}

	return thread
}

func (thread *Thread) AdvanceOneCycle() {
	if thread.Context != nil && thread.Context.State == ContextState_RUNNING {
		var staticInst *StaticInst

		for {
			staticInst = thread.Context.DecodeNextStaticInst()
			staticInst.Execute(thread.Context)

			if staticInst.Mnemonic.Name != Mnemonic_NOP {
				thread.NumInstructions++
			}

			if !(thread.Context != nil &&
				thread.Context.State == ContextState_RUNNING &&
				staticInst.Mnemonic.Name == Mnemonic_NOP) {
				break
			}
		}
	}
}
