package acogo

const (
	NOP = "nop"
	BREAK = "break"
	SYSTEM_CALL = "system_call"
	ADD = "add"
	ADDI = "addi"
	ADDIU = "addiu"
	ADDU = "addu"
	SUB = "sub"
	SUBU = "subu"
	AND = "and"
	ANDI = "andi"
	NOR = "nor"
	OR = "or"
	ORI = "ori"
	XOR = "xor"
	XORI = "xori"
	MULT = "mult"
	MULTU = "multu"
	DIV = "div"
	DIVU = "divu"
	SLL = "sll"
	SLLV = "sllv"
	SLT = "slt"
	SLTI = "slti"
	SLTIU = "sltiu"
	SLTU = "sltu"
	SRA = "sra"
	SRAV = "srav"
	SRL = "srl"
	SRLV = "srlv"
	MADD = "madd"
	MSUB = "msub"
	B = "b"
	BAL = "bal"
	BEQ = "beq"
	BEQL = "beql"
	BGEZ = "bgez"
	BGEZL = "bgezl"
	BGEZAL = "bgezal"
	BGEZALL = "bgezall"
	BGTZ = "bgtz"
	BGTZL = "bgtzl"
	BLEZ = "blez"
	BLEZL = "blezl"
	BLTZ = "bltz"
	BLTZL = "bltzl"
	BLTZAL = "bltzal"
	BLTZALL = "bltzall"
	BNE = "bne"
	BNEL = "bnel"
	J = "j"
	JAL = "jal"
	JALR = "jalr"
	JR = "jr"
	LB = "lb"
	LBU = "lbu"
	LH = "lh"
	LHU = "lhu"
	LUI = "lui"
	LW = "lw"
	LWL = "lwl"
	LWR = "lwr"
	SB = "sb"
	SH = "sh"
	SW = "sw"
	SWL = "swl"
	SWR = "swr"
	LDC1 = "ldc1"
	LWC1 = "lwc1"
	SDC1 = "sdc1"
	SWC1 = "swc1"
	MFHI = "mfhi"
	MFLO = "mflo"
	MTHI = "mthi"
	MTLO = "mtlo"
	CFC1 = "cfc1"
	CTC1 = "ctc1"
	MFC1 = "mfc1"
	LL = "ll"
	SC = "sc"
	NEG_D = "neg_d"
	MOV_D = "mov_d"
	SQRT_D = "sqrt_d"
	ABS_D = "abs_d"
	MUL_D = "mul_d"
	DIV_D = "div_d"
	ADD_D = "add_d"
	SUB_D = "sub_d"
	MUL_S = "mul_s"
	DIV_S = "div_s"
	ADD_S = "add_s"
	SUB_S = "sub_s"
	MOV_S = "mov_s"
	NEG_S = "neg_s"
	ABS_S = "abs_s"
	SQRT_S = "sqrt_s"
	C_COND_D = "c_cond_d"
	C_COND_S = "c_cond_s"
	CVT_D_L = "cvt_d_l"
	CVT_S_L = "cvt_s_l"
	CVT_D_W = "cvt_d_w"
	CVT_S_W = "cvt_s_w"
	CVT_L_D = "cvt_l_d"
	CVT_W_D = "cvt_w_d"
	CVT_S_D = "cvt_s_d"
	CVT_L_S = "cvt_l_s"
	CVT_W_S = "cvt_w_s"
	CVT_D_S = "cvt_d_s"
	BC1FL = "bc1fl"
	BC1TL = "bc1tl"
	BC1F = "bc1f"
	BC1T = "bc1t"
	MOVF = "movf"
	_MOVF = "_movf"
	MOVN = "movn"
	_MOVN = "movn"
	_MOVT = "movt"
	MOVZ = "movz"
	_MOVZ = "_movz"
	MUL = "mul"
	TRUNC_W = "trunc_w"
	UNKNOWN = "unknown"
)

type MnemonicName string

type DecodeMethod struct {
	Mask uint32
	Bits uint32
}

func NewDecodeMethod(mask uint32, bits uint32) DecodeMethod {
	var decodeMethod = DecodeMethod{
		Mask:mask,
		Bits:bits,
	}

	return decodeMethod
}

type DecodeCondition struct {
	BitField *BitField
	Value    uint32
}

func NewDecodeCondition(bitField *BitField, value uint32) DecodeCondition {
	var decodeCondition = DecodeCondition{
		BitField:bitField,
		Value:value,
	}

	return decodeCondition
}

type Mnemonic struct {
	Name               MnemonicName
	StaticInstType     StaticInstType
	StaticInstFlags    []StaticInstFlag
	DecodeMethod       *DecodeMethod
	DecodeCondition    *DecodeCondition
	Mask               uint32
	Bits               uint32
	ExtraBitField      *BitField
	ExtraBitFieldValue uint32
	Execute            func(context *Context, machInst MachInst)
}

func NewMnemonic(name MnemonicName, staticInstType StaticInstType, staticInstFlags []StaticInstFlag, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, execute func(context *Context, machInst MachInst)) *Mnemonic {
	var mnemonic = &Mnemonic{
		Name:name,
		StaticInstType:staticInstType,
		StaticInstFlags:staticInstFlags,
		DecodeMethod:decodeMethod,
		DecodeCondition:decodeCondition,
		Mask:decodeMethod.Mask,
		Bits:decodeMethod.Bits,
		Execute:execute,
	}

	if decodeCondition != nil {
		mnemonic.ExtraBitField = decodeCondition.BitField
		mnemonic.ExtraBitFieldValue = decodeCondition.Value
	}

	return mnemonic
}

const (
	FMT_SINGLE = 16
	FMT_DOUBLE = 17
	FMT_WORD = 20
	FMT_LONG = 21
	FMT_PS = 22
)

const (
	FMT3_SINGLE = 0
	FMT3_DOUBLE = 1
	FMT3_WORD = 4
	FMT3_LONG = 5
	FMT3_PS = 6
)

var (
	//TODO...
	Add = NewMnemonic(ADD, StaticInstType_INTEGER_COMPUTATION, nil, nil, nil, func(context *Context, machInst MachInst) {

	})
)
