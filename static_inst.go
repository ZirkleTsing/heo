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
	StaticInstFlag_NONE = 0x00000000
	StaticInstFlag_ICOMP = 0x00000001
	StaticInstFlag_FCOMP = 0x00000002
	StaticInstFlag_CTRL = 0x00000004
	StaticInstFlag_UNCOND = 0x00000008
	StaticInstFlag_COND = 0x00000010
	StaticInstFlag_MEM = 0x00000020
	StaticInstFlag_LOAD = 0x00000040
	StaticInstFlag_STORE = 0x00000080
	StaticInstFlag_DISP = 0x00000100
	StaticInstFlag_RR = 0x00000200
	StaticInstFlag_DIRECT = 0x00000400
	StaticInstFlag_TRAP = 0x00000800
	StaticInstFlag_LONGLAT = 0x00001000
	StaticInstFlag_DIRJMP = 0x00002000
	StaticInstFlag_INDIRJMP = 0x00004000
	StaticInstFlag_CALL = 0x00008000
	StaticInstFlag_FPCOND = 0x00010000
	StaticInstFlag_IMM = 0x00020000
	StaticInstFlag_RET = 0x00040000
)

type StaticInstFlag uint

type StaticInst struct {
	MachInst uint32
	Mnemonic string
	Flags    []StaticInstFlag
}