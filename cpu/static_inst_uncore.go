package cpu

type StaticInstType string

const (
	StaticInstType_INTEGER_COMPUTATION = StaticInstType("INTEGER_COMPUTATION")
	StaticInstType_FLOAT_COMPUTATION = StaticInstType("FLOAT_COMPUTATION")
	StaticInstType_CONDITIONAL = StaticInstType("CONDITIONAL")
	StaticInstType_UNCONDITIONAL = StaticInstType("UNCONDITIONAL")
	StaticInstType_LOAD = StaticInstType("LOAD")
	StaticInstType_STORE = StaticInstType("STORE")
	StaticInstType_FUNCTION_CALL = StaticInstType("FUNCTION_CALL")
	StaticInstType_FUNCTION_RETURN = StaticInstType("FUNCTION_RETURN")
	StaticInstType_TRAP = StaticInstType("TRAP")
	StaticInstType_NOP = StaticInstType("NOP")
	StaticInstType_UNIMPLEMENTED = StaticInstType("UNIMPLEMENTED")
	StaticInstType_UNKNOWN = StaticInstType("UNKNOWN")
)

type StaticInstFlag string

const (
	StaticInstFlag_INTEGER_COMPUTATION = StaticInstFlag("INTEGER_COMPUTATION")
	StaticInstFlag_FLOAT_COMPUTATION = StaticInstFlag("FLOAT_COMPUTATION")
	StaticInstFlag_UNCONDITIONAL = StaticInstFlag("UNCONDITIONAL")
	StaticInstFlag_CONDITIONAL = StaticInstFlag("CONDITIONAL")
	StaticInstFlag_LOAD = StaticInstFlag("LOAD")
	StaticInstFlag_STORE = StaticInstFlag("STORE")
	StaticInstFlag_DIRECT_JUMP = StaticInstFlag("DIRECT_JUMP")
	StaticInstFlag_INDIRECT_JUMP = StaticInstFlag("INDIRECT_JUMP")
	StaticInstFlag_FUNCTION_CALL = StaticInstFlag("FUNCTION_CALL")
	StaticInstFlag_FUNCTION_RETURN = StaticInstFlag("FUNCTION_RETURN")
	StaticInstFlag_IMMEDIATE = StaticInstFlag("IMMEDIATE")
	StaticInstFlag_DISPLACED_ADDRESSING = StaticInstFlag("DISPLACED_ADDRESSING")
	StaticInstFlag_TRAP = StaticInstFlag("TRAP")
	StaticInstFlag_NOP = StaticInstFlag("NOP")
	StaticInstFlag_UNIMPLEMENTED = StaticInstFlag("UNIMPLEMENTED")
	StaticInstFlag_UNKNOWN = StaticInstFlag("UNKNOWN")
)