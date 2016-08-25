package cpu

const (
	StaticInstType_INTEGER_COMPUTATION = 0
	StaticInstType_FLOAT_COMPUTATION = 1
	StaticInstType_CONDITIONAL = 2
	StaticInstType_UNCONDITIONAL = 3
	StaticInstType_LOAD = 4
	StaticInstType_STORE = 5
	StaticInstType_FUNCTION_CALL = 6
	StaticInstType_FUNCTION_RETURN = 7
	StaticInstType_TRAP = 8
	StaticInstType_NOP = 9
	StaticInstType_UNIMPLEMENTED = 10
	StaticInstType_UNKNOWN = 11
)

type StaticInstType uint

const (
	StaticInstFlag_INTEGER_COMPUTATION = 0
	StaticInstFlag_FLOAT_COMPUTATION = 1
	StaticInstFlag_UNCONDITIONAL = 2
	StaticInstFlag_CONDITIONAL = 3
	StaticInstFlag_LOAD = 4
	StaticInstFlag_STORE = 5
	StaticInstFlag_DIRECT_JUMP = 6
	StaticInstFlag_INDIRECT_JUMP = 7
	StaticInstFlag_FUNCTION_CALL = 8
	StaticInstFlag_FUNCTION_RETURN = 9
	StaticInstFlag_IMMEDIATE = 10
	StaticInstFlag_DISPLACED_ADDRESSING = 11
	StaticInstFlag_TRAP = 12
	StaticInstFlag_NOP = 13
	StaticInstFlag_UNIMPLEMENTED = 14
	StaticInstFlag_UNKNOWN = 15
)

type StaticInstFlag uint

type StaticInst struct {
	Mnemonic *Mnemonic
	MachInst MachInst
}

func NewStaticInst(mnemonic *Mnemonic, machInst MachInst) *StaticInst {
	var staticInst = &StaticInst{
		Mnemonic:mnemonic,
		MachInst:machInst,
	}

	return staticInst
}

func (staticInst *StaticInst) Execute(context *Context) {
	var oldPc = context.Regs.Pc

	staticInst.Mnemonic.Execute(context, staticInst.MachInst)

	context.Kernel.Experiment.BlockingEventDispatcher.Dispatch(NewStaticInstExecutedEvent(context, oldPc, staticInst))
}

type StaticInstExecutedEvent struct {
	Context    *Context
	Pc         uint32
	StaticInst *StaticInst
}

func NewStaticInstExecutedEvent(context *Context, pc uint32, staticInst *StaticInst) *StaticInstExecutedEvent {
	var staticInstExecutedEvent = &StaticInstExecutedEvent{
		Context:context,
		Pc:pc,
		StaticInst:staticInst,
	}

	return staticInstExecutedEvent
}