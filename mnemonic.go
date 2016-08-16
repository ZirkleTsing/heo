package acogo

const (
	Mnemonic_NOP = "nop"
	Mnemonic_BREAK = "break"
	Mnemonic_SYSTEM_CALL = "system_call"
	Mnemonic_ADD = "add"
	Mnemonic_ADDI = "addi"
	Mnemonic_ADDIU = "addiu"
	Mnemonic_ADDU = "addu"
	Mnemonic_SUB = "sub"
	Mnemonic_SUBU = "subu"
	Mnemonic_AND = "and"
	Mnemonic_ANDI = "andi"
	Mnemonic_NOR = "nor"
	Mnemonic_OR = "or"
	Mnemonic_ORI = "ori"
	Mnemonic_XOR = "xor"
	Mnemonic_XORI = "xori"
	Mnemonic_MULT = "mult"
	Mnemonic_MULTU = "multu"
	Mnemonic_DIV = "div"
	Mnemonic_DIVU = "divu"
	Mnemonic_SLL = "sll"
	Mnemonic_SLLV = "sllv"
	Mnemonic_SLT = "slt"
	Mnemonic_SLTI = "slti"
	Mnemonic_SLTIU = "sltiu"
	Mnemonic_SLTU = "sltu"
	Mnemonic_SRA = "sra"
	Mnemonic_SRAV = "srav"
	Mnemonic_SRL = "srl"
	Mnemonic_SRLV = "srlv"
	Mnemonic_MADD = "madd"
	Mnemonic_MSUB = "msub"
	Mnemonic_B = "b"
	Mnemonic_BAL = "bal"
	Mnemonic_BEQ = "beq"
	Mnemonic_BEQL = "beql"
	Mnemonic_BGEZ = "bgez"
	Mnemonic_BGEZL = "bgezl"
	Mnemonic_BGEZAL = "bgezal"
	Mnemonic_BGEZALL = "bgezall"
	Mnemonic_BGTZ = "bgtz"
	Mnemonic_BGTZL = "bgtzl"
	Mnemonic_BLEZ = "blez"
	Mnemonic_BLEZL = "blezl"
	Mnemonic_BLTZ = "bltz"
	Mnemonic_BLTZL = "bltzl"
	Mnemonic_BLTZAL = "bltzal"
	Mnemonic_BLTZALL = "bltzall"
	Mnemonic_BNE = "bne"
	Mnemonic_BNEL = "bnel"
	Mnemonic_J = "j"
	Mnemonic_JAL = "jal"
	Mnemonic_JALR = "jalr"
	Mnemonic_JR = "jr"
	Mnemonic_LB = "lb"
	Mnemonic_LBU = "lbu"
	Mnemonic_LH = "lh"
	Mnemonic_LHU = "lhu"
	Mnemonic_LUI = "lui"
	Mnemonic_LW = "lw"
	Mnemonic_LWL = "lwl"
	Mnemonic_LWR = "lwr"
	Mnemonic_SB = "sb"
	Mnemonic_SH = "sh"
	Mnemonic_SW = "sw"
	Mnemonic_SWL = "swl"
	Mnemonic_SWR = "swr"
	Mnemonic_LDC1 = "ldc1"
	Mnemonic_LWC1 = "lwc1"
	Mnemonic_SDC1 = "sdc1"
	Mnemonic_SWC1 = "swc1"
	Mnemonic_MFHI = "mfhi"
	Mnemonic_MFLO = "mflo"
	Mnemonic_MTHI = "mthi"
	Mnemonic_MTLO = "mtlo"
	Mnemonic_CFC1 = "cfc1"
	Mnemonic_CTC1 = "ctc1"
	Mnemonic_MFC1 = "mfc1"
	Mnemonic_MTC1 = "mtc1"
	Mnemonic_LL = "ll"
	Mnemonic_SC = "sc"
	Mnemonic_NEG_D = "neg_d"
	Mnemonic_MOV_D = "mov_d"
	Mnemonic_SQRT_D = "sqrt_d"
	Mnemonic_ABS_D = "abs_d"
	Mnemonic_MUL_D = "mul_d"
	Mnemonic_DIV_D = "div_d"
	Mnemonic_ADD_D = "add_d"
	Mnemonic_SUB_D = "sub_d"
	Mnemonic_MUL_S = "mul_s"
	Mnemonic_DIV_S = "div_s"
	Mnemonic_ADD_S = "add_s"
	Mnemonic_SUB_S = "sub_s"
	Mnemonic_MOV_S = "mov_s"
	Mnemonic_NEG_S = "neg_s"
	Mnemonic_ABS_S = "abs_s"
	Mnemonic_SQRT_S = "sqrt_s"
	Mnemonic_C_COND_D = "c_cond_d"
	Mnemonic_C_COND_S = "c_cond_s"
	Mnemonic_CVT_D_L = "cvt_d_l"
	Mnemonic_CVT_S_L = "cvt_s_l"
	Mnemonic_CVT_D_W = "cvt_d_w"
	Mnemonic_CVT_S_W = "cvt_s_w"
	Mnemonic_CVT_L_D = "cvt_l_d"
	Mnemonic_CVT_W_D = "cvt_w_d"
	Mnemonic_CVT_S_D = "cvt_s_d"
	Mnemonic_CVT_L_S = "cvt_l_s"
	Mnemonic_CVT_W_S = "cvt_w_s"
	Mnemonic_CVT_D_S = "cvt_d_s"
	Mnemonic_BC1FL = "bc1fl"
	Mnemonic_BC1TL = "bc1tl"
	Mnemonic_BC1F = "bc1f"
	Mnemonic_BC1T = "bc1t"
	Mnemonic_MOVF = "movf"
	Mnemonic__MOVF = "_movf"
	Mnemonic_MOVN = "movn"
	Mnemonic__MOVN = "_movn"
	Mnemonic__MOVT = "_movt"
	Mnemonic_MOVZ = "movz"
	Mnemonic__MOVZ = "_movz"
	Mnemonic_MUL = "mul"
	Mnemonic_TRUNC_W = "trunc_w"
	Mnemonic_UNKNOWN = "unknown"
)

type MnemonicName string

type DecodeMethod struct {
	Bits uint32
	Mask uint32
}

func NewDecodeMethod(bits uint32, mask uint32) *DecodeMethod {
	var decodeMethod = &DecodeMethod{
		Bits:bits,
		Mask:mask,
	}

	return decodeMethod
}

type DecodeCondition struct {
	BitField *BitField
	Value    uint32
}

func NewDecodeCondition(bitField *BitField, value uint32) *DecodeCondition {
	var decodeCondition = &DecodeCondition{
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
	Bits               uint32
	Mask               uint32
	ExtraBitField      *BitField
	ExtraBitFieldValue uint32
	Execute            func(context *Context, machInst MachInst)
}

func NewMnemonic(name MnemonicName, staticInstFlags []StaticInstFlag, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, execute func(context *Context, machInst MachInst)) *Mnemonic {
	var mnemonic = &Mnemonic{
		Name:name,
		StaticInstFlags:staticInstFlags,
		DecodeMethod:decodeMethod,
		DecodeCondition:decodeCondition,
		Bits:decodeMethod.Bits,
		Mask:decodeMethod.Mask,
		Execute:execute,
	}

	mnemonic.StaticInstType = StaticInstType_UNKNOWN // TODO

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
	Add = NewMnemonic(
		Mnemonic_ADD,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		add)

	Addi = NewMnemonic(
		Mnemonic_ADDI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		addi)

	Addiu = NewMnemonic(
		Mnemonic_ADDIU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		addiu)

	Addu = NewMnemonic(
		Mnemonic_ADDU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		addu)

	And = NewMnemonic(
		Mnemonic_AND,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		and)

	Andi = NewMnemonic(
		Mnemonic_ANDI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		andi)

	Div = NewMnemonic(
		Mnemonic_DIV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		div)

	Divu = NewMnemonic(
		Mnemonic_DIVU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		divu)

	Lui = NewMnemonic(
		Mnemonic_LUI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		nil)

	Madd = NewMnemonic(
		Mnemonic_MADD,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		nil)

	Mfhi = NewMnemonic(
		Mnemonic_MFHI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		nil)

	Mflo = NewMnemonic(
		Mnemonic_MFLO,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		nil)

	Msub = NewMnemonic(
		Mnemonic_MSUB,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		nil)

	Mthi = NewMnemonic(
		Mnemonic_MTHI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Mtlo = NewMnemonic(
		Mnemonic_MTLO,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		nil)

	Mult = NewMnemonic(
		Mnemonic_MULT,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		nil)

	Multu = NewMnemonic(
		Mnemonic_MULTU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		nil)

	Nor = NewMnemonic(
		Mnemonic_NOR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		nil)

	Or = NewMnemonic(
		Mnemonic_OR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		nil)

	Ori = NewMnemonic(
		Mnemonic_ORI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		nil)

	Sll = NewMnemonic(
		Mnemonic_SLL,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		nil)

	Sllv = NewMnemonic(
		Mnemonic_SLLV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		nil)

	Slt = NewMnemonic(
		Mnemonic_SLT,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		nil)

	Slti = NewMnemonic(
		Mnemonic_SLTI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		nil)

	Sltiu = NewMnemonic(
		Mnemonic_SLTIU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		nil)

	Sltu = NewMnemonic(
		Mnemonic_SLTU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		nil)

	Sra = NewMnemonic(
		Mnemonic_SRA,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		nil)

	Srav = NewMnemonic(
		Mnemonic_SRAV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		nil)

	Srl = NewMnemonic(
		Mnemonic_SRL,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		nil)

	Srlv = NewMnemonic(
		Mnemonic_SRLV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		nil)

	Sub = NewMnemonic(
		Mnemonic_SUB,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), // TODO: missing decoding information
		nil,
		nil)

	Subu = NewMnemonic(
		Mnemonic_SUBU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		nil)

	Xor = NewMnemonic(
		Mnemonic_XOR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		nil)

	Xori = NewMnemonic(
		Mnemonic_XORI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		nil)

	AbsD = NewMnemonic(
		Mnemonic_ABS_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	AbsS = NewMnemonic(
		Mnemonic_ABS_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	AddD = NewMnemonic(
		Mnemonic_ADD_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	AddS = NewMnemonic(
		Mnemonic_ADD_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	CCondD = NewMnemonic(
		Mnemonic_C_COND_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	CCondS = NewMnemonic(
		Mnemonic_C_COND_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	CvtDL = NewMnemonic(
		Mnemonic_CVT_D_L,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		nil)

	CvtDS = NewMnemonic(
		Mnemonic_CVT_D_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	CvtDW = NewMnemonic(
		Mnemonic_CVT_D_W,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		nil)

	CvtLD = NewMnemonic(
		Mnemonic_CVT_L_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	CvtLS = NewMnemonic(
		Mnemonic_CVT_L_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	CvtSD = NewMnemonic(
		Mnemonic_CVT_S_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	CvtSL = NewMnemonic(
		Mnemonic_CVT_S_L,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		nil)

	CvtSW = NewMnemonic(
		Mnemonic_CVT_S_W,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		nil)

	CvtWD = NewMnemonic(
		Mnemonic_CVT_W_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	CvtWS = NewMnemonic(
		Mnemonic_CVT_W_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	DivD = NewMnemonic(
		Mnemonic_DIV_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	DivS = NewMnemonic(
		Mnemonic_DIV_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	MovD = NewMnemonic(
		Mnemonic_MOV_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	MovS = NewMnemonic(
		Mnemonic_MOV_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	Movf = NewMnemonic(
		Mnemonic_MOVF,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x00000001, 0xfc0307ff),
		nil,
		nil)

	_Movf = NewMnemonic(
		Mnemonic__MOVF,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x44000011, 0xfc03003f),
		nil,
		nil)

	Movn = NewMnemonic(
		Mnemonic_MOVN,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x0000000b, 0xfc0007ff),
		nil,
		nil)

	_Movn = NewMnemonic(
		Mnemonic__MOVN,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x44000013, 0xfc00003f),
		nil,
		nil)

	_Movt = NewMnemonic(
		Mnemonic__MOVT,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x44010011, 0xfc03003f),
		nil,
		nil)

	Movz = NewMnemonic(
		Mnemonic_MOVZ,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x0000000a, 0xfc0007ff),
		nil,
		nil)

	_Movz = NewMnemonic(
		Mnemonic__MOVZ,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x44000012, 0xfc00003f),
		nil,
		nil)

	Mul = NewMnemonic(
		Mnemonic_MUL,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x70000002, 0xfc0007ff),
		nil,
		nil)

	TruncW = NewMnemonic(
		Mnemonic_TRUNC_W,
		[]StaticInstFlag{StaticInstFlag_UNIMPLEMENTED},
		NewDecodeMethod(0x4400000d, 0xfc1f003f),
		nil,
		nil)

	MulD = NewMnemonic(
		Mnemonic_MUL_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	MulS = NewMnemonic(
		Mnemonic_MUL_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	NegD = NewMnemonic(
		Mnemonic_NEG_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	NegS = NewMnemonic(
		Mnemonic_NEG_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	SqrtD = NewMnemonic(
		Mnemonic_SQRT_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	SqrtS = NewMnemonic(
		Mnemonic_SQRT_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	SubD = NewMnemonic(
		Mnemonic_SUB_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)

	SubS = NewMnemonic(
		Mnemonic_SUB_S,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		nil)

	J = NewMnemonic(
		Mnemonic_J,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x08000000, 0xfc000000),
		nil,
		nil)

	Jal = NewMnemonic(
		Mnemonic_JAL,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0c000000, 0xfc000000),
		nil,
		nil)

	Jalr = NewMnemonic(
		Mnemonic_JALR,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_INDIRECT_JUMP},
		NewDecodeMethod(0x00000009, 0xfc00003f),
		nil,
		nil)

	Jr = NewMnemonic(
		Mnemonic_JR,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_FUNCTION_RETURN, StaticInstFlag_INDIRECT_JUMP},
		NewDecodeMethod(0x00000008, 0xfc00003f),
		nil,
		nil)

	B = NewMnemonic(
		Mnemonic_B,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x10000000, 0xffff0000),
		nil,
		nil)

	Bal = NewMnemonic(
		Mnemonic_BAL,
		[]StaticInstFlag{StaticInstFlag_UNCONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x04110000, 0xffff0000),
		nil,
		nil)

	Bc1f = NewMnemonic(
		Mnemonic_BC1F,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL},
		NewDecodeMethod(0x45000000, 0xffe30000),
		nil,
		nil)

	Bc1fl = NewMnemonic(
		Mnemonic_BC1FL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bc1t = NewMnemonic(
		Mnemonic_BC1T,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL},
		NewDecodeMethod(0x45010000, 0xffe30000),
		nil,
		nil)

	Bc1tl = NewMnemonic(
		Mnemonic_BC1TL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Beq = NewMnemonic(
		Mnemonic_BEQ,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x10000000, 0xfc000000),
		nil,
		nil)

	Beql = NewMnemonic(
		Mnemonic_BEQL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bgez = NewMnemonic(
		Mnemonic_BGEZ,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x04010000, 0xfc1f0000),
		nil,
		nil)

	Bgezal = NewMnemonic(
		Mnemonic_BGEZAL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x04110000, 0xfc1f0000),
		nil,
		nil)

	Bgezall = NewMnemonic(
		Mnemonic_BGEZALL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bgezl = NewMnemonic(
		Mnemonic_BGEZL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bgtz = NewMnemonic(
		Mnemonic_BGTZ,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x1c000000, 0xfc1f0000),
		nil,
		nil)

	Bgtzl = NewMnemonic(
		Mnemonic_BGTZL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Blez = NewMnemonic(
		Mnemonic_BLEZ,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x18000000, 0xfc1f0000),
		nil,
		nil)

	Blezl = NewMnemonic(
		Mnemonic_BLEZL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bltz = NewMnemonic(
		Mnemonic_BLTZ,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x04000000, 0xfc1f0000),
		nil,
		nil)

	Bltzal = NewMnemonic(
		Mnemonic_BLTZAL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bltzall = NewMnemonic(
		Mnemonic_BLTZALL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_FUNCTION_CALL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bltzl = NewMnemonic(
		Mnemonic_BLTZL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Bne = NewMnemonic(
		Mnemonic_BNE,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x14000000, 0xfc000000),
		nil,
		nil)

	Bnel = NewMnemonic(
		Mnemonic_BNEL,
		[]StaticInstFlag{StaticInstFlag_CONDITIONAL, StaticInstFlag_DIRECT_JUMP},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Lb = NewMnemonic(
		Mnemonic_LB,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x80000000, 0xfc000000),
		nil,
		nil)

	Lbu = NewMnemonic(
		Mnemonic_LBU,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x90000000, 0xfc000000),
		nil,
		nil)

	Ldc1 = NewMnemonic(
		Mnemonic_LDC1,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xd4000000, 0xfc000000),
		nil,
		nil)

	Lh = NewMnemonic(
		Mnemonic_LH,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x84000000, 0xfc000000),
		nil,
		nil)

	Lhu = NewMnemonic(
		Mnemonic_LHU,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x94000000, 0xfc000000),
		nil,
		nil)

	Ll = NewMnemonic(
		Mnemonic_LL,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xc0000000, 0xfc000000),
		nil,
		nil)

	Lw = NewMnemonic(
		Mnemonic_LW,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x8c000000, 0xfc000000),
		nil,
		nil)

	Lwc1 = NewMnemonic(
		Mnemonic_LWC1,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xc4000000, 0xfc000000),
		nil,
		nil)

	Lwl = NewMnemonic(
		Mnemonic_LWL,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x88000000, 0xfc000000),
		nil,
		nil)

	Lwr = NewMnemonic(
		Mnemonic_LWR,
		[]StaticInstFlag{StaticInstFlag_LOAD, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0x98000000, 0xfc000000),
		nil,
		nil)

	Sb = NewMnemonic(
		Mnemonic_SB,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xa0000000, 0xfc000000),
		nil,
		nil)

	Sc = NewMnemonic(
		Mnemonic_SC,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xe0000000, 0xfc000000),
		nil,
		nil)

	Sdc1 = NewMnemonic(
		Mnemonic_SDC1,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xf4000000, 0xfc000000),
		nil,
		nil)

	Sh = NewMnemonic(
		Mnemonic_SH,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xa4000000, 0xfc000000),
		nil,
		nil)

	Sw = NewMnemonic(
		Mnemonic_SW,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xac000000, 0xfc000000),
		nil,
		nil)

	Swc1 = NewMnemonic(
		Mnemonic_SWC1,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xe4000000, 0xfc000000),
		nil,
		nil)

	Swl = NewMnemonic(
		Mnemonic_SWL,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xa8000000, 0xfc000000),
		nil,
		nil)

	Swr = NewMnemonic(
		Mnemonic_SWR,
		[]StaticInstFlag{StaticInstFlag_STORE, StaticInstFlag_DISPLACED_ADDRESSING},
		NewDecodeMethod(0xb8000000, 0xfc000000),
		nil,
		nil)

	Cfc1 = NewMnemonic(
		Mnemonic_CFC1,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x44400000, 0xffe007ff),
		nil,
		nil)

	Ctc1 = NewMnemonic(
		Mnemonic_CTC1,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x44c00000, 0xffe007ff),
		nil,
		nil)

	Mfc1 = NewMnemonic(
		Mnemonic_MFC1,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x44000000, 0xffe007ff),
		nil,
		nil)

	Mtc1 = NewMnemonic(
		Mnemonic_MTC1,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x44800000, 0xffe007ff),
		nil,
		nil)

	Break = NewMnemonic(
		Mnemonic_BREAK,
		[]StaticInstFlag{StaticInstFlag_TRAP},
		NewDecodeMethod(0x0000000d, 0xfc00003f),
		nil,
		nil)

	SystemCall = NewMnemonic(
		Mnemonic_SYSTEM_CALL,
		[]StaticInstFlag{StaticInstFlag_TRAP},
		NewDecodeMethod(0x0000000c, 0xfc00003f),
		nil,
		nil)

	Nop = NewMnemonic(
		Mnemonic_NOP,
		[]StaticInstFlag{StaticInstFlag_NOP},
		NewDecodeMethod(0x00000000, 0xffffffff),
		nil,
		nil)

	Unknown = NewMnemonic(
		Mnemonic_UNKNOWN,
		[]StaticInstFlag{StaticInstFlag_UNKNOWN},
		NewDecodeMethod(0x0, 0x0), //TODO: special support for unknown instruction
		nil,
		nil)
)
