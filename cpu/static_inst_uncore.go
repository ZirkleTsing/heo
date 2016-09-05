package cpu

type StaticInstFlag string

const (
	StaticInstFlag_INT_COMP = StaticInstFlag("INT_COMP")
	StaticInstFlag_FP_COMP = StaticInstFlag("FP_COMP")
	StaticInstFlag_UNCOND = StaticInstFlag("UNCOND")
	StaticInstFlag_COND = StaticInstFlag("COND")
	StaticInstFlag_LD = StaticInstFlag("LD")
	StaticInstFlag_ST = StaticInstFlag("ST")
	StaticInstFlag_DIRECT_JMP = StaticInstFlag("DIRECT_JUMP")
	StaticInstFlag_INDIRECT_JMP = StaticInstFlag("INDIRECT_JUMP")
	StaticInstFlag_FUNC_CALL = StaticInstFlag("FUNC_CALL")
	StaticInstFlag_FUNC_RET = StaticInstFlag("FUNC_RET")
	StaticInstFlag_IMM = StaticInstFlag("IMM")
	StaticInstFlag_DISPLACED_ADDRESSING = StaticInstFlag("DISPLACED_ADDRESSING")
	StaticInstFlag_TRAP = StaticInstFlag("TRAP")
	StaticInstFlag_NOP = StaticInstFlag("NOP")
)

type StaticInstType string

const (
	StaticInstType_INT_COMP = StaticInstType("INT_COMP")
	StaticInstType_FP_COMP = StaticInstType("FP_COMP")
	StaticInstType_COND = StaticInstType("COND")
	StaticInstType_UNCOND = StaticInstType("UNCOND")
	StaticInstType_LD = StaticInstType("LD")
	StaticInstType_ST = StaticInstType("ST")
	StaticInstType_FUNC_CALL = StaticInstType("FUNC_CALL")
	StaticInstType_FUNC_RET = StaticInstType("FUNC_RET")
	StaticInstType_TRAP = StaticInstType("TRAP")
	StaticInstType_NOP = StaticInstType("NOP")
)
