package cpu

func (processor *Processor) addMnemonic(name MnemonicName, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, fuOperationType FUOperationType, staticInstType StaticInstType, staticInstFlags []StaticInstFlag, inputDependencies []StaticInstDependency, outputDependencies []StaticInstDependency, execute func(context *Context, machInst MachInst)) {
	var mnemonic = NewMnemonic(name, decodeMethod, decodeCondition, fuOperationType, staticInstType, staticInstFlags, inputDependencies, outputDependencies, execute)

	processor.Mnemonics = append(processor.Mnemonics, mnemonic)
}

func (processor *Processor) addMnemonics() {
	processor.addMnemonic(
		Mnemonic_NOP,
		NewDecodeMethod(0x00000000, 0xffffffff),
		nil,
		FUOperationType_NONE,
		StaticInstType_NOP,
		[]StaticInstFlag{
			StaticInstFlag_NOP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{},
		nop,
	)

	processor.addMnemonic(
		Mnemonic_BC1F,
		NewDecodeMethod(0x45000000, 0xffe30000),
		nil,
		FUOperationType_NONE,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		[]StaticInstDependency{},
		bc1f,
	)

	processor.addMnemonic(
		Mnemonic_BC1T,
		NewDecodeMethod(0x45010000, 0xffe30000),
		nil,
		FUOperationType_NONE,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		[]StaticInstDependency{},
		bc1t,
	)

	processor.addMnemonic(
		Mnemonic_MFC1,
		NewDecodeMethod(0x44000000, 0xffe007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		mfc1,
	)

	processor.addMnemonic(
		Mnemonic_MTC1,
		NewDecodeMethod(0x44800000, 0xffe007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		mtc1,
	)

	processor.addMnemonic(
		Mnemonic_CFC1,
		NewDecodeMethod(0x44400000, 0xffe007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		cfc1,
	)

	processor.addMnemonic(
		Mnemonic_CTC1,
		NewDecodeMethod(0x44c00000, 0xffe007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		ctc1,
	)

	processor.addMnemonic(
		Mnemonic_ABS_S,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		abs_s,
	)

	processor.addMnemonic(
		Mnemonic_ABS_D,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		abs_d,
	)

	processor.addMnemonic(
		Mnemonic_ADD,
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		add,
	)

	processor.addMnemonic(
		Mnemonic_ADD_S,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_ADD,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		add_s,
	)

	processor.addMnemonic(
		Mnemonic_ADD_D,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_ADD,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		add_d,
	)

	processor.addMnemonic(
		Mnemonic_ADDI,
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		addi,
	)

	processor.addMnemonic(
		Mnemonic_ADDIU,
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		addiu,
	)

	processor.addMnemonic(
		Mnemonic_ADDU,
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		addu,
	)

	processor.addMnemonic(
		Mnemonic_AND,
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		and,
	)

	processor.addMnemonic(
		Mnemonic_ANDI,
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		andi,
	)

	processor.addMnemonic(
		Mnemonic_B,
		NewDecodeMethod(0x10000000, 0xffff0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{},
		b,
	)

	processor.addMnemonic(
		Mnemonic_BAL,
		NewDecodeMethod(0x04110000, 0xffff0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_RA,
		},
		bal,
	)

	processor.addMnemonic(
		Mnemonic_BEQ,
		NewDecodeMethod(0x10000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		beq,
	)

	processor.addMnemonic(
		Mnemonic_BGEZ,
		NewDecodeMethod(0x04010000, 0xfc1f0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{},
		bgez,
	)

	processor.addMnemonic(
		Mnemonic_BGEZAL,
		NewDecodeMethod(0x04110000, 0xfc1f0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_RA,
		},
		bgezal,
	)

	processor.addMnemonic(
		Mnemonic_BGTZ,
		NewDecodeMethod(0x1c000000, 0xfc1f0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{},
		bgtz,
	)

	processor.addMnemonic(
		Mnemonic_BLEZ,
		NewDecodeMethod(0x18000000, 0xfc1f0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{},
		blez,
	)

	processor.addMnemonic(
		Mnemonic_BLTZ,
		NewDecodeMethod(0x04000000, 0xfc1f0000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{},
		bltz,
	)

	processor.addMnemonic(
		Mnemonic_BNE,
		NewDecodeMethod(0x14000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		bne,
	)

	processor.addMnemonic(
		Mnemonic_BREAK,
		NewDecodeMethod(0x0000000d, 0xfc00003f),
		nil,
		FUOperationType_NONE,
		StaticInstType_TRAP,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{},
		_break,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_D,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
			StaticInstDependency_REGISTER_FCSR,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		c_cond_d,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_S,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
			StaticInstDependency_REGISTER_FCSR,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_FCSR,
		},
		c_cond_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_S,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_d_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_W,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_d_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_L,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_d_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_D,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_s_d,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_W,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_s_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_L,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_s_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_S,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_w_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_D,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_CVT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		cvt_w_d,
	)

	processor.addMnemonic(
		Mnemonic_DIV,
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		FUOperationType_INT_DIV,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		div,
	)

	processor.addMnemonic(
		Mnemonic_DIV_S,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_DIV,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		div_s,
	)

	processor.addMnemonic(
		Mnemonic_DIV_D,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_DIV,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		div_d,
	)

	processor.addMnemonic(
		Mnemonic_DIVU,
		NewDecodeMethod(0x0000001b, 0xfc00003f),
		nil,
		FUOperationType_INT_DIV,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		divu,
	)

	processor.addMnemonic(
		Mnemonic_J,
		NewDecodeMethod(0x08000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{},
		j,
	)

	processor.addMnemonic(
		Mnemonic_JAL,
		NewDecodeMethod(0x0c000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_RA,
		},
		jal,
	)

	processor.addMnemonic(
		Mnemonic_JALR,
		NewDecodeMethod(0x00000009, 0xfc00003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_INDIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		jalr,
	)

	processor.addMnemonic(
		Mnemonic_JR,
		NewDecodeMethod(0x00000008, 0xfc00003f),
		nil,
		FUOperationType_NONE,
		StaticInstType_FUNC_RET,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_INDIRECT_JMP,
			StaticInstFlag_FUNC_RET,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{},
		jr,
	)

	processor.addMnemonic(
		Mnemonic_LB,
		NewDecodeMethod(0x80000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lb,
	)

	processor.addMnemonic(
		Mnemonic_LBU,
		NewDecodeMethod(0x90000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lbu,
	)

	processor.addMnemonic(
		Mnemonic_LDC1,
		NewDecodeMethod(0xd4000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FT,
		},
		ldc1,
	)

	processor.addMnemonic(
		Mnemonic_LH,
		NewDecodeMethod(0x84000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lh,
	)

	processor.addMnemonic(
		Mnemonic_LHU,
		NewDecodeMethod(0x94000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lhu,
	)

	processor.addMnemonic(
		Mnemonic_LL,
		NewDecodeMethod(0xc0000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		ll,
	)

	processor.addMnemonic(
		Mnemonic_LUI,
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lui,
	)

	processor.addMnemonic(
		Mnemonic_LW,
		NewDecodeMethod(0x8c000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lw,
	)

	processor.addMnemonic(
		Mnemonic_LWC1,
		NewDecodeMethod(0xc4000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FT,
		},
		lwc1,
	)

	processor.addMnemonic(
		Mnemonic_LWL,
		NewDecodeMethod(0x88000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lwl,
	)

	processor.addMnemonic(
		Mnemonic_LWR,
		NewDecodeMethod(0x98000000, 0xfc000000),
		nil,
		FUOperationType_READ_PORT,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		lwr,
	)

	processor.addMnemonic(
		Mnemonic_MADD,
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		FUOperationType_INT_MULT,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		madd,
	)

	processor.addMnemonic(
		Mnemonic_MFHI,
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		mfhi,
	)

	processor.addMnemonic(
		Mnemonic_MFLO,
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_LO,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		mflo,
	)

	processor.addMnemonic(
		Mnemonic_MOV_S,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_NONE,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		mov_s,
	)

	processor.addMnemonic(
		Mnemonic_MOV_D,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_NONE,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		mov_d,
	)

	processor.addMnemonic(
		Mnemonic_MSUB,
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		FUOperationType_INT_MULT,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		msub,
	)

	processor.addMnemonic(
		Mnemonic_MTLO,
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_LO,
		},
		mtlo,
	)

	processor.addMnemonic(
		Mnemonic_MUL_S,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_MULT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		mul_s,
	)

	processor.addMnemonic(
		Mnemonic_MUL_D,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_MULT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		mul_d,
	)

	processor.addMnemonic(
		Mnemonic_MULT,
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		mult,
	)

	processor.addMnemonic(
		Mnemonic_MULTU,
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_HI,
			StaticInstDependency_REGISTER_LO,
		},
		multu,
	)

	processor.addMnemonic(
		Mnemonic_NEG_S,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		neg_s,
	)

	processor.addMnemonic(
		Mnemonic_NEG_D,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_CMP,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		neg_d,
	)

	processor.addMnemonic(
		Mnemonic_NOR,
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		nor,
	)

	processor.addMnemonic(
		Mnemonic_OR,
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		or,
	)

	processor.addMnemonic(
		Mnemonic_ORI,
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		ori,
	)

	processor.addMnemonic(
		Mnemonic_SB,
		NewDecodeMethod(0xa0000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		sb,
	)

	processor.addMnemonic(
		Mnemonic_SC,
		NewDecodeMethod(0xe0000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		sc,
	)

	processor.addMnemonic(
		Mnemonic_SDC1,
		NewDecodeMethod(0xf4000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{},
		sdc1,
	)

	processor.addMnemonic(
		Mnemonic_SH,
		NewDecodeMethod(0xa4000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		sh,
	)

	processor.addMnemonic(
		Mnemonic_SLL,
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		sll,
	)

	processor.addMnemonic(
		Mnemonic_SLLV,
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		sllv,
	)

	processor.addMnemonic(
		Mnemonic_SLT,
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		slt,
	)

	processor.addMnemonic(
		Mnemonic_SLTI,
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		slti,
	)

	processor.addMnemonic(
		Mnemonic_SLTIU,
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		sltiu,
	)

	processor.addMnemonic(
		Mnemonic_SLTU,
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		sltu,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_S,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_SQRT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		sqrt_s,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_D,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_SQRT,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		sqrt_d,
	)

	processor.addMnemonic(
		Mnemonic_SRA,
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		sra,
	)

	processor.addMnemonic(
		Mnemonic_SRAV,
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		srav,
	)

	processor.addMnemonic(
		Mnemonic_SRL,
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		srl,
	)

	processor.addMnemonic(
		Mnemonic_SRLV,
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		srlv,
	)

	processor.addMnemonic(
		Mnemonic_SUB_S,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		FUOperationType_FP_ADD,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		sub_s,
	)

	processor.addMnemonic(
		Mnemonic_SUB_D,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		FUOperationType_FP_ADD,
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_FS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{
			StaticInstDependency_FD,
		},
		sub_d,
	)

	processor.addMnemonic(
		Mnemonic_SUBU,
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		subu,
	)

	processor.addMnemonic(
		Mnemonic_SW,
		NewDecodeMethod(0xac000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		sw,
	)

	processor.addMnemonic(
		Mnemonic_SWC1,
		NewDecodeMethod(0xe4000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_FT,
		},
		[]StaticInstDependency{},
		swc1,
	)

	processor.addMnemonic(
		Mnemonic_SWL,
		NewDecodeMethod(0xa8000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		swl,
	)

	processor.addMnemonic(
		Mnemonic_SWR,
		NewDecodeMethod(0xb8000000, 0xfc000000),
		nil,
		FUOperationType_WRITE_PORT,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{},
		swr,
	)

	processor.addMnemonic(
		Mnemonic_SYSCALL,
		NewDecodeMethod(0x0000000c, 0xfc00003f),
		nil,
		FUOperationType_NONE,
		StaticInstType_TRAP,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		[]StaticInstDependency{
			StaticInstDependency_REGISTER_V0,
		},
		[]StaticInstDependency{},
		_syscall,
	)

	processor.addMnemonic(
		Mnemonic_XOR,
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
			StaticInstDependency_RT,
		},
		[]StaticInstDependency{
			StaticInstDependency_RD,
		},
		xor,
	)

	processor.addMnemonic(
		Mnemonic_XORI,
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		FUOperationType_INT_ALU,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		[]StaticInstDependency{
			StaticInstDependency_RS,
		},
		[]StaticInstDependency{
			StaticInstDependency_RT,
		},
		xori,
	)
}
