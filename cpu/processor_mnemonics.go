package cpu

func (processor *Processor) addMnemonic(name MnemonicName, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, staticInstType StaticInstType, staticInstFlags []StaticInstFlag, execute func(context *Context, machInst MachInst)) {
	var mnemonic = NewMnemonic(name, decodeMethod, decodeCondition, staticInstType, staticInstFlags, execute)

	processor.Mnemonics = append(processor.Mnemonics, mnemonic)
}

func (processor *Processor) addMnemonics() {
	processor.addMnemonic(
		Mnemonic_NOP,
		NewDecodeMethod(0x00000000, 0xffffffff),
		nil,
		StaticInstType_NOP,
		[]StaticInstFlag{
			StaticInstFlag_NOP,
		},
		nop,
	)

	processor.addMnemonic(
		Mnemonic_BC1F,
		NewDecodeMethod(0x45000000, 0xffe30000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
		},
		bc1f,
	)

	processor.addMnemonic(
		Mnemonic_BC1T,
		NewDecodeMethod(0x45010000, 0xffe30000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
		},
		bc1t,
	)

	processor.addMnemonic(
		Mnemonic_MFC1,
		NewDecodeMethod(0x44000000, 0xffe007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mfc1,
	)

	processor.addMnemonic(
		Mnemonic_MTC1,
		NewDecodeMethod(0x44800000, 0xffe007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mtc1,
	)

	processor.addMnemonic(
		Mnemonic_CFC1,
		NewDecodeMethod(0x44400000, 0xffe007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		cfc1,
	)

	processor.addMnemonic(
		Mnemonic_CTC1,
		NewDecodeMethod(0x44c00000, 0xffe007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		ctc1,
	)

	processor.addMnemonic(
		Mnemonic_ABS_S,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		abs_s,
	)

	processor.addMnemonic(
		Mnemonic_ABS_D,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		abs_d,
	)

	processor.addMnemonic(
		Mnemonic_ADD,
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		add,
	)

	processor.addMnemonic(
		Mnemonic_ADD_S,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		add_s,
	)

	processor.addMnemonic(
		Mnemonic_ADD_D,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		add_d,
	)

	processor.addMnemonic(
		Mnemonic_ADDI,
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		addi,
	)

	processor.addMnemonic(
		Mnemonic_ADDIU,
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		addiu,
	)

	processor.addMnemonic(
		Mnemonic_ADDU,
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		addu,
	)

	processor.addMnemonic(
		Mnemonic_AND,
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		and,
	)

	processor.addMnemonic(
		Mnemonic_ANDI,
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		andi,
	)

	processor.addMnemonic(
		Mnemonic_B,
		NewDecodeMethod(0x10000000, 0xffff0000),
		nil,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		b,
	)

	processor.addMnemonic(
		Mnemonic_BAL,
		NewDecodeMethod(0x04110000, 0xffff0000),
		nil,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		bal,
	)

	processor.addMnemonic(
		Mnemonic_BEQ,
		NewDecodeMethod(0x10000000, 0xfc000000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		beq,
	)

	processor.addMnemonic(
		Mnemonic_BGEZ,
		NewDecodeMethod(0x04010000, 0xfc1f0000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		bgez,
	)

	processor.addMnemonic(
		Mnemonic_BGEZAL,
		NewDecodeMethod(0x04110000, 0xfc1f0000),
		nil,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		bgezal,
	)

	processor.addMnemonic(
		Mnemonic_BGTZ,
		NewDecodeMethod(0x1c000000, 0xfc1f0000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		bgtz,
	)

	processor.addMnemonic(
		Mnemonic_BLEZ,
		NewDecodeMethod(0x18000000, 0xfc1f0000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		blez,
	)

	processor.addMnemonic(
		Mnemonic_BLTZ,
		NewDecodeMethod(0x04000000, 0xfc1f0000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		bltz,
	)

	processor.addMnemonic(
		Mnemonic_BNE,
		NewDecodeMethod(0x14000000, 0xfc000000),
		nil,
		StaticInstType_COND,
		[]StaticInstFlag{
			StaticInstFlag_COND,
			StaticInstFlag_DIRECT_JMP,
		},
		bne,
	)

	processor.addMnemonic(
		Mnemonic_BREAK,
		NewDecodeMethod(0x0000000d, 0xfc00003f),
		nil,
		StaticInstType_TRAP,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		_break,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_D,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		c_cond_d,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_S,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		c_cond_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_S,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_d_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_W,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_d_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_L,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_d_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_D,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_s_d,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_W,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_s_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_L,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_s_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_S,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_w_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_D,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		cvt_w_d,
	)

	processor.addMnemonic(
		Mnemonic_DIV,
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		div,
	)

	processor.addMnemonic(
		Mnemonic_DIV_S,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		div_s,
	)

	processor.addMnemonic(
		Mnemonic_DIV_D,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		div_d,
	)

	processor.addMnemonic(
		Mnemonic_DIVU,
		NewDecodeMethod(0x0000001b, 0xfc00003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		divu,
	)

	processor.addMnemonic(
		Mnemonic_J,
		NewDecodeMethod(0x08000000, 0xfc000000),
		nil,
		StaticInstType_UNCOND,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
		},
		j,
	)

	processor.addMnemonic(
		Mnemonic_JAL,
		NewDecodeMethod(0x0c000000, 0xfc000000),
		nil,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_DIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		jal,
	)

	processor.addMnemonic(
		Mnemonic_JALR,
		NewDecodeMethod(0x00000009, 0xfc00003f),
		nil,
		StaticInstType_FUNC_CALL,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_INDIRECT_JMP,
			StaticInstFlag_FUNC_CALL,
		},
		jalr,
	)

	processor.addMnemonic(
		Mnemonic_JR,
		NewDecodeMethod(0x00000008, 0xfc00003f),
		nil,
		StaticInstType_FUNC_RET,
		[]StaticInstFlag{
			StaticInstFlag_UNCOND,
			StaticInstFlag_INDIRECT_JMP,
			StaticInstFlag_FUNC_RET,
		},
		jr,
	)

	processor.addMnemonic(
		Mnemonic_LB,
		NewDecodeMethod(0x80000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lb,
	)

	processor.addMnemonic(
		Mnemonic_LBU,
		NewDecodeMethod(0x90000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lbu,
	)

	processor.addMnemonic(
		Mnemonic_LDC1,
		NewDecodeMethod(0xd4000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		ldc1,
	)

	processor.addMnemonic(
		Mnemonic_LH,
		NewDecodeMethod(0x84000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lh,
	)

	processor.addMnemonic(
		Mnemonic_LHU,
		NewDecodeMethod(0x94000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lhu,
	)

	processor.addMnemonic(
		Mnemonic_LL,
		NewDecodeMethod(0xc0000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		ll,
	)

	processor.addMnemonic(
		Mnemonic_LUI,
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		lui,
	)

	processor.addMnemonic(
		Mnemonic_LW,
		NewDecodeMethod(0x8c000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lw,
	)

	processor.addMnemonic(
		Mnemonic_LWC1,
		NewDecodeMethod(0xc4000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwc1,
	)

	processor.addMnemonic(
		Mnemonic_LWL,
		NewDecodeMethod(0x88000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwl,
	)

	processor.addMnemonic(
		Mnemonic_LWR,
		NewDecodeMethod(0x98000000, 0xfc000000),
		nil,
		StaticInstType_LD,
		[]StaticInstFlag{
			StaticInstFlag_LD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwr,
	)

	processor.addMnemonic(
		Mnemonic_MADD,
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		madd,
	)

	processor.addMnemonic(
		Mnemonic_MFHI,
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mfhi,
	)

	processor.addMnemonic(
		Mnemonic_MFLO,
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mflo,
	)

	processor.addMnemonic(
		Mnemonic_MOV_S,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		mov_s,
	)

	processor.addMnemonic(
		Mnemonic_MOV_D,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		mov_d,
	)

	processor.addMnemonic(
		Mnemonic_MSUB,
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		msub,
	)

	processor.addMnemonic(
		Mnemonic_MTLO,
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mtlo,
	)

	processor.addMnemonic(
		Mnemonic_MUL_S,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		mul_s,
	)

	processor.addMnemonic(
		Mnemonic_MUL_D,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		mul_d,
	)

	processor.addMnemonic(
		Mnemonic_MULT,
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		mult,
	)

	processor.addMnemonic(
		Mnemonic_MULTU,
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		multu,
	)

	processor.addMnemonic(
		Mnemonic_NEG_S,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		neg_s,
	)

	processor.addMnemonic(
		Mnemonic_NEG_D,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		neg_d,
	)

	processor.addMnemonic(
		Mnemonic_NOR,
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		nor,
	)

	processor.addMnemonic(
		Mnemonic_OR,
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		or,
	)

	processor.addMnemonic(
		Mnemonic_ORI,
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		ori,
	)

	processor.addMnemonic(
		Mnemonic_SB,
		NewDecodeMethod(0xa0000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sb,
	)

	processor.addMnemonic(
		Mnemonic_SC,
		NewDecodeMethod(0xe0000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sc,
	)

	processor.addMnemonic(
		Mnemonic_SDC1,
		NewDecodeMethod(0xf4000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sdc1,
	)

	processor.addMnemonic(
		Mnemonic_SH,
		NewDecodeMethod(0xa4000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sh,
	)

	processor.addMnemonic(
		Mnemonic_SLL,
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		sll,
	)

	processor.addMnemonic(
		Mnemonic_SLLV,
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		sllv,
	)

	processor.addMnemonic(
		Mnemonic_SLT,
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		slt,
	)

	processor.addMnemonic(
		Mnemonic_SLTI,
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		slti,
	)

	processor.addMnemonic(
		Mnemonic_SLTIU,
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		sltiu,
	)

	processor.addMnemonic(
		Mnemonic_SLTU,
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		sltu,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_S,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		sqrt_s,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_D,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		sqrt_d,
	)

	processor.addMnemonic(
		Mnemonic_SRA,
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		sra,
	)

	processor.addMnemonic(
		Mnemonic_SRAV,
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		srav,
	)

	processor.addMnemonic(
		Mnemonic_SRL,
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		srl,
	)

	processor.addMnemonic(
		Mnemonic_SRLV,
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		srlv,
	)

	processor.addMnemonic(
		Mnemonic_SUB_S,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		sub_s,
	)

	processor.addMnemonic(
		Mnemonic_SUB_D,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		StaticInstType_FP_COMP,
		[]StaticInstFlag{
			StaticInstFlag_FP_COMP,
		},
		sub_d,
	)

	processor.addMnemonic(
		Mnemonic_SUBU,
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		subu,
	)

	processor.addMnemonic(
		Mnemonic_SW,
		NewDecodeMethod(0xac000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sw,
	)

	processor.addMnemonic(
		Mnemonic_SWC1,
		NewDecodeMethod(0xe4000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swc1,
	)

	processor.addMnemonic(
		Mnemonic_SWL,
		NewDecodeMethod(0xa8000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swl,
	)

	processor.addMnemonic(
		Mnemonic_SWR,
		NewDecodeMethod(0xb8000000, 0xfc000000),
		nil,
		StaticInstType_ST,
		[]StaticInstFlag{
			StaticInstFlag_ST,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swr,
	)

	processor.addMnemonic(
		Mnemonic_SYSCALL,
		NewDecodeMethod(0x0000000c, 0xfc00003f),
		nil,
		StaticInstType_TRAP,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		_syscall,
	)

	processor.addMnemonic(
		Mnemonic_XOR,
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
		},
		xor,
	)

	processor.addMnemonic(
		Mnemonic_XORI,
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		StaticInstType_INT_COMP,
		[]StaticInstFlag{
			StaticInstFlag_INT_COMP,
			StaticInstFlag_IMM,
		},
		xori,
	)
}
