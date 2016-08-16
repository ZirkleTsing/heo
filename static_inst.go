package acogo

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
	StaticInstFlag_FUNCTIONAL_CALL = 8
	StaticInstFlag_IMMEDIATE = 9
	StaticInstFlag_DISPLACED_ADDRESSING = 10
	StaticInstFlag_TRAP = 11
	StaticInstFlag_NOP = 12
	StaticInstFlag_UNIMPLEMENTED = 13
	StaticInstFlag_UNKNOWN = 14
)

type StaticInstFlag uint

type StaticInst struct {
	MachInst MachInst
	Mnemonic *Mnemonic
}

func NewStaticInst(machInst MachInst, mnemonic *Mnemonic) *StaticInst {
	var staticInst = &StaticInst{
		MachInst:machInst,
		Mnemonic:mnemonic,
	}

	return staticInst
}