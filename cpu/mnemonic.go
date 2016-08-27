package cpu

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
	DecodeMethod       *DecodeMethod
	DecodeCondition    *DecodeCondition
	Bits               uint32
	Mask               uint32
	ExtraBitField      *BitField
	ExtraBitFieldValue uint32
	Execute            func(context *Context, machInst MachInst)
}

func NewMnemonic(name MnemonicName, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, execute func(context *Context, machInst MachInst)) *Mnemonic {
	var mnemonic = &Mnemonic{
		Name:name,
		DecodeMethod:decodeMethod,
		DecodeCondition:decodeCondition,
		Bits:decodeMethod.Bits,
		Mask:decodeMethod.Mask,
		Execute:execute,
	}

	if decodeCondition != nil {
		mnemonic.ExtraBitField = decodeCondition.BitField
		mnemonic.ExtraBitFieldValue = decodeCondition.Value
	}

	Mnemonics = append(Mnemonics, mnemonic)

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
	Mnemonics = []*Mnemonic{}
)

var (
	Add = NewMnemonic(
		Mnemonic_ADD,
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		add)

	Addi = NewMnemonic(
		Mnemonic_ADDI,
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		addi)

	Addiu = NewMnemonic(
		Mnemonic_ADDIU,
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		addiu)

	Addu = NewMnemonic(
		Mnemonic_ADDU,
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		addu)

	And = NewMnemonic(
		Mnemonic_AND,
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		and)

	Andi = NewMnemonic(
		Mnemonic_ANDI,
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		andi)

	Div = NewMnemonic(
		Mnemonic_DIV,
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		div)

	Divu = NewMnemonic(
		Mnemonic_DIVU,
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		divu)

	Lui = NewMnemonic(
		Mnemonic_LUI,
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		lui)

	Madd = NewMnemonic(
		Mnemonic_MADD,
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		madd)

	Mfhi = NewMnemonic(
		Mnemonic_MFHI,
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		mfhi)

	Mflo = NewMnemonic(
		Mnemonic_MFLO,
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		mflo)

	Msub = NewMnemonic(
		Mnemonic_MSUB,
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		msub)

	Mthi = NewMnemonic(
		Mnemonic_MTHI,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		mthi)

	Mtlo = NewMnemonic(
		Mnemonic_MTLO,
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		mtlo)

	Mult = NewMnemonic(
		Mnemonic_MULT,
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		mult)

	Multu = NewMnemonic(
		Mnemonic_MULTU,
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		multu)

	Nor = NewMnemonic(
		Mnemonic_NOR,
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		nor)

	Or = NewMnemonic(
		Mnemonic_OR,
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		or)

	Ori = NewMnemonic(
		Mnemonic_ORI,
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		ori)

	Sll = NewMnemonic(
		Mnemonic_SLL,
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		sll)

	Sllv = NewMnemonic(
		Mnemonic_SLLV,
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		sllv)

	Slt = NewMnemonic(
		Mnemonic_SLT,
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		slt)

	Slti = NewMnemonic(
		Mnemonic_SLTI,
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		slti)

	Sltiu = NewMnemonic(
		Mnemonic_SLTIU,
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		sltiu)

	Sltu = NewMnemonic(
		Mnemonic_SLTU,
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		sltu)

	Sra = NewMnemonic(
		Mnemonic_SRA,
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		sra)

	Srav = NewMnemonic(
		Mnemonic_SRAV,
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		srav)

	Srl = NewMnemonic(
		Mnemonic_SRL,
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		srl)

	Srlv = NewMnemonic(
		Mnemonic_SRLV,
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		srlv)

	Sub = NewMnemonic(
		Mnemonic_SUB,
		NewDecodeMethod(0x0, 0x0), // TODO: missing decoding information
		nil,
		sub)

	Subu = NewMnemonic(
		Mnemonic_SUBU,
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		subu)

	Xor = NewMnemonic(
		Mnemonic_XOR,
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		xor)

	Xori = NewMnemonic(
		Mnemonic_XORI,
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		xori)

	AbsD = NewMnemonic(
		Mnemonic_ABS_D,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		absD)

	AbsS = NewMnemonic(
		Mnemonic_ABS_S,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		absS)

	AddD = NewMnemonic(
		Mnemonic_ADD_D,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		addD)

	AddS = NewMnemonic(
		Mnemonic_ADD_S,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		addS)

	CCondD = NewMnemonic(
		Mnemonic_C_COND_D,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		cCondD)

	CCondS = NewMnemonic(
		Mnemonic_C_COND_S,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_SINGLE),
		cCondS)

	CvtDL = NewMnemonic(
		Mnemonic_CVT_D_L,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		cvtDL)

	CvtDS = NewMnemonic(
		Mnemonic_CVT_D_S,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		cvtDS)

	CvtDW = NewMnemonic(
		Mnemonic_CVT_D_W,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		cvtDW)

	CvtLD = NewMnemonic(
		Mnemonic_CVT_L_D,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		cvtLD)

	CvtLS = NewMnemonic(
		Mnemonic_CVT_L_S,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		cvtLS)

	CvtSD = NewMnemonic(
		Mnemonic_CVT_S_D,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		cvtSD)

	CvtSL = NewMnemonic(
		Mnemonic_CVT_S_L,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		cvtSL)

	CvtSW = NewMnemonic(
		Mnemonic_CVT_S_W,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		cvtSW)

	CvtWD = NewMnemonic(
		Mnemonic_CVT_W_D,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		cvtWD)

	CvtWS = NewMnemonic(
		Mnemonic_CVT_W_S,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		cvtWS)

	DivD = NewMnemonic(
		Mnemonic_DIV_D,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		divD)

	DivS = NewMnemonic(
		Mnemonic_DIV_S,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		divS)

	MovD = NewMnemonic(
		Mnemonic_MOV_D,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		movD)

	MovS = NewMnemonic(
		Mnemonic_MOV_S,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		movS)

	Movf = NewMnemonic(
		Mnemonic_MOVF,
		NewDecodeMethod(0x00000001, 0xfc0307ff),
		nil,
		movf)

	_Movf = NewMnemonic(
		Mnemonic__MOVF,
		NewDecodeMethod(0x44000011, 0xfc03003f),
		nil,
		_movf)

	Movn = NewMnemonic(
		Mnemonic_MOVN,
		NewDecodeMethod(0x0000000b, 0xfc0007ff),
		nil,
		movn)

	_Movn = NewMnemonic(
		Mnemonic__MOVN,
		NewDecodeMethod(0x44000013, 0xfc00003f),
		nil,
		_movn)

	_Movt = NewMnemonic(
		Mnemonic__MOVT,
		NewDecodeMethod(0x44010011, 0xfc03003f),
		nil,
		_movt)

	Movz = NewMnemonic(
		Mnemonic_MOVZ,
		NewDecodeMethod(0x0000000a, 0xfc0007ff),
		nil,
		movz)

	_Movz = NewMnemonic(
		Mnemonic__MOVZ,
		NewDecodeMethod(0x44000012, 0xfc00003f),
		nil,
		_movz)

	Mul = NewMnemonic(
		Mnemonic_MUL,
		NewDecodeMethod(0x70000002, 0xfc0007ff),
		nil,
		mul)

	TruncW = NewMnemonic(
		Mnemonic_TRUNC_W,
		NewDecodeMethod(0x4400000d, 0xfc1f003f),
		nil,
		truncW)

	MulD = NewMnemonic(
		Mnemonic_MUL_D,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		mulD)

	MulS = NewMnemonic(
		Mnemonic_MUL_S,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		mulS)

	NegD = NewMnemonic(
		Mnemonic_NEG_D,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		negD)

	NegS = NewMnemonic(
		Mnemonic_NEG_S,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		negS)

	SqrtD = NewMnemonic(
		Mnemonic_SQRT_D,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		sqrtD)

	SqrtS = NewMnemonic(
		Mnemonic_SQRT_S,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		sqrtS)

	SubD = NewMnemonic(
		Mnemonic_SUB_D,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		subD)

	SubS = NewMnemonic(
		Mnemonic_SUB_S,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		subS)

	J = NewMnemonic(
		Mnemonic_J,
		NewDecodeMethod(0x08000000, 0xfc000000),
		nil,
		j)

	Jal = NewMnemonic(
		Mnemonic_JAL,
		NewDecodeMethod(0x0c000000, 0xfc000000),
		nil,
		jal)

	Jalr = NewMnemonic(
		Mnemonic_JALR,
		NewDecodeMethod(0x00000009, 0xfc00003f),
		nil,
		jalr)

	Jr = NewMnemonic(
		Mnemonic_JR,
		NewDecodeMethod(0x00000008, 0xfc00003f),
		nil,
		jr)

	B = NewMnemonic(
		Mnemonic_B,
		NewDecodeMethod(0x10000000, 0xffff0000),
		nil,
		b)

	Bal = NewMnemonic(
		Mnemonic_BAL,
		NewDecodeMethod(0x04110000, 0xffff0000),
		nil,
		bal)

	Bc1f = NewMnemonic(
		Mnemonic_BC1F,
		NewDecodeMethod(0x45000000, 0xffe30000),
		nil,
		bc1f)

	Bc1fl = NewMnemonic(
		Mnemonic_BC1FL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bc1fl)

	Bc1t = NewMnemonic(
		Mnemonic_BC1T,
		NewDecodeMethod(0x45010000, 0xffe30000),
		nil,
		bc1t)

	Bc1tl = NewMnemonic(
		Mnemonic_BC1TL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bc1tl)

	Beq = NewMnemonic(
		Mnemonic_BEQ,
		NewDecodeMethod(0x10000000, 0xfc000000),
		nil,
		beq)

	Beql = NewMnemonic(
		Mnemonic_BEQL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		beql)

	Bgez = NewMnemonic(
		Mnemonic_BGEZ,
		NewDecodeMethod(0x04010000, 0xfc1f0000),
		nil,
		bgez)

	Bgezal = NewMnemonic(
		Mnemonic_BGEZAL,
		NewDecodeMethod(0x04110000, 0xfc1f0000),
		nil,
		bgezal)

	Bgezall = NewMnemonic(
		Mnemonic_BGEZALL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bgezall)

	Bgezl = NewMnemonic(
		Mnemonic_BGEZL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bgezl)

	Bgtz = NewMnemonic(
		Mnemonic_BGTZ,
		NewDecodeMethod(0x1c000000, 0xfc1f0000),
		nil,
		bgtz)

	Bgtzl = NewMnemonic(
		Mnemonic_BGTZL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bgtzl)

	Blez = NewMnemonic(
		Mnemonic_BLEZ,
		NewDecodeMethod(0x18000000, 0xfc1f0000),
		nil,
		blez)

	Blezl = NewMnemonic(
		Mnemonic_BLEZL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		blezl)

	Bltz = NewMnemonic(
		Mnemonic_BLTZ,
		NewDecodeMethod(0x04000000, 0xfc1f0000),
		nil,
		bltz)

	Bltzal = NewMnemonic(
		Mnemonic_BLTZAL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bltzal)

	Bltzall = NewMnemonic(
		Mnemonic_BLTZALL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bltzall)

	Bltzl = NewMnemonic(
		Mnemonic_BLTZL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bltzl)

	Bne = NewMnemonic(
		Mnemonic_BNE,
		NewDecodeMethod(0x14000000, 0xfc000000),
		nil,
		bne)

	Bnel = NewMnemonic(
		Mnemonic_BNEL,
		NewDecodeMethod(0x0, 0x0), //TODO: missing decoding information
		nil,
		bnel)

	Lb = NewMnemonic(
		Mnemonic_LB,
		NewDecodeMethod(0x80000000, 0xfc000000),
		nil,
		lb)

	Lbu = NewMnemonic(
		Mnemonic_LBU,
		NewDecodeMethod(0x90000000, 0xfc000000),
		nil,
		lbu)

	Ldc1 = NewMnemonic(
		Mnemonic_LDC1,
		NewDecodeMethod(0xd4000000, 0xfc000000),
		nil,
		ldc1)

	Lh = NewMnemonic(
		Mnemonic_LH,
		NewDecodeMethod(0x84000000, 0xfc000000),
		nil,
		lh)

	Lhu = NewMnemonic(
		Mnemonic_LHU,
		NewDecodeMethod(0x94000000, 0xfc000000),
		nil,
		lhu)

	Ll = NewMnemonic(
		Mnemonic_LL,
		NewDecodeMethod(0xc0000000, 0xfc000000),
		nil,
		ll)

	Lw = NewMnemonic(
		Mnemonic_LW,
		NewDecodeMethod(0x8c000000, 0xfc000000),
		nil,
		lw)

	Lwc1 = NewMnemonic(
		Mnemonic_LWC1,
		NewDecodeMethod(0xc4000000, 0xfc000000),
		nil,
		lwc1)

	Lwl = NewMnemonic(
		Mnemonic_LWL,
		NewDecodeMethod(0x88000000, 0xfc000000),
		nil,
		lwl)

	Lwr = NewMnemonic(
		Mnemonic_LWR,
		NewDecodeMethod(0x98000000, 0xfc000000),
		nil,
		lwr)

	Sb = NewMnemonic(
		Mnemonic_SB,
		NewDecodeMethod(0xa0000000, 0xfc000000),
		nil,
		sb)

	Sc = NewMnemonic(
		Mnemonic_SC,
		NewDecodeMethod(0xe0000000, 0xfc000000),
		nil,
		sc)

	Sdc1 = NewMnemonic(
		Mnemonic_SDC1,
		NewDecodeMethod(0xf4000000, 0xfc000000),
		nil,
		sdc1)

	Sh = NewMnemonic(
		Mnemonic_SH,
		NewDecodeMethod(0xa4000000, 0xfc000000),
		nil,
		sh)

	Sw = NewMnemonic(
		Mnemonic_SW,
		NewDecodeMethod(0xac000000, 0xfc000000),
		nil,
		sw)

	Swc1 = NewMnemonic(
		Mnemonic_SWC1,
		NewDecodeMethod(0xe4000000, 0xfc000000),
		nil,
		swc1)

	Swl = NewMnemonic(
		Mnemonic_SWL,
		NewDecodeMethod(0xa8000000, 0xfc000000),
		nil,
		swl)

	Swr = NewMnemonic(
		Mnemonic_SWR,
		NewDecodeMethod(0xb8000000, 0xfc000000),
		nil,
		swr)

	Cfc1 = NewMnemonic(
		Mnemonic_CFC1,
		NewDecodeMethod(0x44400000, 0xffe007ff),
		nil,
		cfc1)

	Ctc1 = NewMnemonic(
		Mnemonic_CTC1,
		NewDecodeMethod(0x44c00000, 0xffe007ff),
		nil,
		ctc1)

	Mfc1 = NewMnemonic(
		Mnemonic_MFC1,
		NewDecodeMethod(0x44000000, 0xffe007ff),
		nil,
		mfc1)

	Mtc1 = NewMnemonic(
		Mnemonic_MTC1,
		NewDecodeMethod(0x44800000, 0xffe007ff),
		nil,
		mtc1)

	Break = NewMnemonic(
		Mnemonic_BREAK,
		NewDecodeMethod(0x0000000d, 0xfc00003f),
		nil,
		_break)

	SystemCall = NewMnemonic(
		Mnemonic_SYSTEM_CALL,
		NewDecodeMethod(0x0000000c, 0xfc00003f),
		nil,
		systemCall)

	Nop = NewMnemonic(
		Mnemonic_NOP,
		NewDecodeMethod(0x00000000, 0xffffffff),
		nil,
		nop)

	Unknown = NewMnemonic(
		Mnemonic_UNKNOWN,
		NewDecodeMethod(0x0, 0x0), //TODO: special support for unknown instruction
		nil,
		unknown)
)
