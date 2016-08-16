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
		ADD,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		add)

	Addi = NewMnemonic(
		ADDI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		addi)

	Addiu = NewMnemonic(
		ADDIU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		addiu)

	Addu = NewMnemonic(
		ADDU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		addu)

	And = NewMnemonic(
		AND,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		and)

	Andi = NewMnemonic(
		ANDI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		andi)

	Div = NewMnemonic(
		DIV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		div)

	Divu = NewMnemonic(
		DIVU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		divu)

	Lui = NewMnemonic(
		LUI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		nil)

	Madd = NewMnemonic(
		MADD,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		nil)

	Mfhi = NewMnemonic(
		MFHI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		nil)

	Mflo = NewMnemonic(
		MFLO,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		nil)

	Msub = NewMnemonic(
		MSUB,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		nil)

	Mthi = NewMnemonic(
		MTHI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		nil)

	Mtlo = NewMnemonic(
		MTLO,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		nil)

	Mult = NewMnemonic(
		MULT,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		nil)

	Multu = NewMnemonic(
		MULTU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		nil)

	Nor = NewMnemonic(
		NOR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		nil)

	Or = NewMnemonic(
		OR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		nil)

	Ori = NewMnemonic(
		ORI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		nil)

	Sll = NewMnemonic(
		SLL,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		nil)

	Sllv = NewMnemonic(
		SLLV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		nil)

	Slt = NewMnemonic(
		SLT,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		nil)

	Slti = NewMnemonic(
		SLTI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		nil)

	Sltiu = NewMnemonic(
		SLTIU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		nil)

	Sltu = NewMnemonic(
		SLTU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		nil)

	Sra = NewMnemonic(
		SRA,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		nil)

	Srav = NewMnemonic(
		SRAV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		nil)

	Srl = NewMnemonic(
		SRL,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		nil)

	Srlv = NewMnemonic(
		SRLV,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		nil)

	Sub = NewMnemonic(
		SUB,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x0, 0x0), // TODO: missing decoding information
		nil,
		nil)

	Subu = NewMnemonic(
		SUBU,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		nil)

	Xor = NewMnemonic(
		XOR,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION},
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		nil)

	Xori = NewMnemonic(
		XORI,
		[]StaticInstFlag{StaticInstFlag_INTEGER_COMPUTATION, StaticInstFlag_IMMEDIATE},
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		nil)

	AbsD = NewMnemonic(
		ABS_D,
		[]StaticInstFlag{StaticInstFlag_FLOAT_COMPUTATION},
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		nil)
)
