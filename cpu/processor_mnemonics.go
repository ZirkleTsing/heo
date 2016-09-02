package cpu

func (processor *Processor) addMnemonic(name MnemonicName, decodeMethod *DecodeMethod, decodeCondition *DecodeCondition, flags []StaticInstFlag, execute func(context *Context, machInst MachInst)) {
	var mnemonic = &Mnemonic{
		Name:name,
		DecodeMethod:decodeMethod,
		DecodeCondition:decodeCondition,
		Bits:decodeMethod.Bits,
		Mask:decodeMethod.Mask,
		Flags:flags,
		Execute:execute,
	}

	if decodeCondition != nil {
		mnemonic.ExtraBitField = decodeCondition.BitField
		mnemonic.ExtraBitFieldValue = decodeCondition.Value
	}

	processor.Mnemonics = append(processor.Mnemonics, mnemonic)
}

func (processor *Processor) addMnemonics() {
	processor.addMnemonic(
		Mnemonic_NOP,
		NewDecodeMethod(0x00000000, 0xffffffff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_NOP,
		},
		nop,
	)

	processor.addMnemonic(
		Mnemonic_BC1F,
		NewDecodeMethod(0x45000000, 0xffe30000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
		},
		bc1f,
	)

	processor.addMnemonic(
		Mnemonic_BC1T,
		NewDecodeMethod(0x45010000, 0xffe30000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
		},
		bc1t,
	)

	processor.addMnemonic(
		Mnemonic_MFC1,
		NewDecodeMethod(0x44000000, 0xffe007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mfc1,
	)

	processor.addMnemonic(
		Mnemonic_MTC1,
		NewDecodeMethod(0x44800000, 0xffe007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mtc1,
	)

	processor.addMnemonic(
		Mnemonic_CFC1,
		NewDecodeMethod(0x44400000, 0xffe007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		cfc1,
	)

	processor.addMnemonic(
		Mnemonic_CTC1,
		NewDecodeMethod(0x44c00000, 0xffe007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		ctc1,
	)

	processor.addMnemonic(
		Mnemonic_ABS_S,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		abs_s,
	)

	processor.addMnemonic(
		Mnemonic_ABS_D,
		NewDecodeMethod(0x44000005, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		abs_d,
	)

	processor.addMnemonic(
		Mnemonic_ADD,
		NewDecodeMethod(0x00000020, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		add,
	)

	processor.addMnemonic(
		Mnemonic_ADD_S,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		add_s,
	)

	processor.addMnemonic(
		Mnemonic_ADD_D,
		NewDecodeMethod(0x44000000, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		add_d,
	)

	processor.addMnemonic(
		Mnemonic_ADDI,
		NewDecodeMethod(0x20000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		addi,
	)

	processor.addMnemonic(
		Mnemonic_ADDIU,
		NewDecodeMethod(0x24000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		addiu,
	)

	processor.addMnemonic(
		Mnemonic_ADDU,
		NewDecodeMethod(0x00000021, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		addu,
	)

	processor.addMnemonic(
		Mnemonic_AND,
		NewDecodeMethod(0x00000024, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		and,
	)

	processor.addMnemonic(
		Mnemonic_ANDI,
		NewDecodeMethod(0x30000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		andi,
	)

	processor.addMnemonic(
		Mnemonic_B,
		NewDecodeMethod(0x10000000, 0xffff0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		b,
	)

	processor.addMnemonic(
		Mnemonic_BAL,
		NewDecodeMethod(0x04110000, 0xffff0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bal,
	)

	processor.addMnemonic(
		Mnemonic_BEQ,
		NewDecodeMethod(0x10000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		beq,
	)

	processor.addMnemonic(
		Mnemonic_BGEZ,
		NewDecodeMethod(0x04010000, 0xfc1f0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bgez,
	)

	processor.addMnemonic(
		Mnemonic_BGEZAL,
		NewDecodeMethod(0x04110000, 0xfc1f0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_FUNCTION_CALL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bgezal,
	)

	processor.addMnemonic(
		Mnemonic_BGTZ,
		NewDecodeMethod(0x1c000000, 0xfc1f0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bgtz,
	)

	processor.addMnemonic(
		Mnemonic_BLEZ,
		NewDecodeMethod(0x18000000, 0xfc1f0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		blez,
	)

	processor.addMnemonic(
		Mnemonic_BLTZ,
		NewDecodeMethod(0x04000000, 0xfc1f0000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bltz,
	)

	processor.addMnemonic(
		Mnemonic_BNE,
		NewDecodeMethod(0x14000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_CONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		bne,
	)

	processor.addMnemonic(
		Mnemonic_BREAK,
		NewDecodeMethod(0x0000000d, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		_break,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_D,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		c_cond_d,
	)

	processor.addMnemonic(
		Mnemonic_C_COND_S,
		NewDecodeMethod(0x44000030, 0xfc0000f0),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		c_cond_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_S,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_d_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_W,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_d_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_D_L,
		NewDecodeMethod(0x44000021, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_d_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_D,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_s_d,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_W,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_WORD),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_s_w,
	)

	processor.addMnemonic(
		Mnemonic_CVT_S_L,
		NewDecodeMethod(0x44000020, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_LONG),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_s_l,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_S,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_w_s,
	)

	processor.addMnemonic(
		Mnemonic_CVT_W_D,
		NewDecodeMethod(0x44000024, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		cvt_w_d,
	)

	processor.addMnemonic(
		Mnemonic_DIV,
		NewDecodeMethod(0x0000001a, 0xfc00ffff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		div,
	)

	processor.addMnemonic(
		Mnemonic_DIV_S,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		div_s,
	)

	processor.addMnemonic(
		Mnemonic_DIV_D,
		NewDecodeMethod(0x44000003, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		div_d,
	)

	processor.addMnemonic(
		Mnemonic_DIVU,
		NewDecodeMethod(0x0000001b, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		divu,
	)

	processor.addMnemonic(
		Mnemonic_J,
		NewDecodeMethod(0x08000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_DIRECT_JUMP,
		},
		j,
	)

	processor.addMnemonic(
		Mnemonic_JAL,
		NewDecodeMethod(0x0c000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_FUNCTION_CALL,
			StaticInstFlag_DIRECT_JUMP,
		},
		jal,
	)

	processor.addMnemonic(
		Mnemonic_JALR,
		NewDecodeMethod(0x00000009, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_FUNCTION_CALL,
			StaticInstFlag_INDIRECT_JUMP,
		},
		jalr,
	)

	processor.addMnemonic(
		Mnemonic_JR,
		NewDecodeMethod(0x00000008, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNCONDITIONAL,
			StaticInstFlag_FUNCTION_RETURN,
			StaticInstFlag_INDIRECT_JUMP,
		},
		jr,
	)

	processor.addMnemonic(
		Mnemonic_LB,
		NewDecodeMethod(0x80000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lb,
	)

	processor.addMnemonic(
		Mnemonic_LBU,
		NewDecodeMethod(0x90000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lbu,
	)

	processor.addMnemonic(
		Mnemonic_LDC1,
		NewDecodeMethod(0xd4000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		ldc1,
	)

	processor.addMnemonic(
		Mnemonic_LH,
		NewDecodeMethod(0x84000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lh,
	)

	processor.addMnemonic(
		Mnemonic_LHU,
		NewDecodeMethod(0x94000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lhu,
	)

	processor.addMnemonic(
		Mnemonic_LL,
		NewDecodeMethod(0xc0000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		ll,
	)

	processor.addMnemonic(
		Mnemonic_LUI,
		NewDecodeMethod(0x3c000000, 0xffe00000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		lui,
	)

	processor.addMnemonic(
		Mnemonic_LW,
		NewDecodeMethod(0x8c000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lw,
	)

	processor.addMnemonic(
		Mnemonic_LWC1,
		NewDecodeMethod(0xc4000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwc1,
	)

	processor.addMnemonic(
		Mnemonic_LWL,
		NewDecodeMethod(0x88000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwl,
	)

	processor.addMnemonic(
		Mnemonic_LWR,
		NewDecodeMethod(0x98000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_LOAD,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		lwr,
	)

	processor.addMnemonic(
		Mnemonic_MADD,
		NewDecodeMethod(0x70000000, 0xfc00ffff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		madd,
	)

	processor.addMnemonic(
		Mnemonic_MFHI,
		NewDecodeMethod(0x00000010, 0xffff07ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mfhi,
	)

	processor.addMnemonic(
		Mnemonic_MFLO,
		NewDecodeMethod(0x00000012, 0xffff07ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mflo,
	)

	processor.addMnemonic(
		Mnemonic_MOV_S,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		mov_s,
	)

	processor.addMnemonic(
		Mnemonic_MOV_D,
		NewDecodeMethod(0x44000006, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		mov_d,
	)

	processor.addMnemonic(
		Mnemonic_MOVF,
		NewDecodeMethod(0x00000001, 0xfc0307ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		movf,
	)

	processor.addMnemonic(
		Mnemonic__MOVF,
		NewDecodeMethod(0x44000011, 0xfc03003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		_movf,
	)

	processor.addMnemonic(
		Mnemonic_MOVN,
		NewDecodeMethod(0x0000000b, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		movn,
	)

	processor.addMnemonic(
		Mnemonic__MOVN,
		NewDecodeMethod(0x44000013, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		_movn,
	)

	processor.addMnemonic(
		Mnemonic__MOVT,
		NewDecodeMethod(0x44010011, 0xfc03003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		_movt,
	)

	processor.addMnemonic(
		Mnemonic_MOVZ,
		NewDecodeMethod(0x0000000a, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		movz,
	)

	processor.addMnemonic(
		Mnemonic__MOVZ,
		NewDecodeMethod(0x44000012, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		_movz,
	)

	processor.addMnemonic(
		Mnemonic_MSUB,
		NewDecodeMethod(0x70000004, 0xfc00ffff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		msub,
	)

	processor.addMnemonic(
		Mnemonic_MTLO,
		NewDecodeMethod(0x00000013, 0xfc1fffff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mtlo,
	)

	processor.addMnemonic(
		Mnemonic_MUL,
		NewDecodeMethod(0x70000002, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		mul,
	)

	processor.addMnemonic(
		Mnemonic_MUL_S,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		mul_s,
	)

	processor.addMnemonic(
		Mnemonic_MUL_D,
		NewDecodeMethod(0x44000002, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		mul_d,
	)

	processor.addMnemonic(
		Mnemonic_MULT,
		NewDecodeMethod(0x00000018, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		mult,
	)

	processor.addMnemonic(
		Mnemonic_MULTU,
		NewDecodeMethod(0x00000019, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		multu,
	)

	processor.addMnemonic(
		Mnemonic_NEG_S,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		neg_s,
	)

	processor.addMnemonic(
		Mnemonic_NEG_D,
		NewDecodeMethod(0x44000007, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		neg_d,
	)

	processor.addMnemonic(
		Mnemonic_NOR,
		NewDecodeMethod(0x00000027, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		nor,
	)

	processor.addMnemonic(
		Mnemonic_OR,
		NewDecodeMethod(0x00000025, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		or,
	)

	processor.addMnemonic(
		Mnemonic_ORI,
		NewDecodeMethod(0x34000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		ori,
	)

	processor.addMnemonic(
		Mnemonic_SB,
		NewDecodeMethod(0xa0000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sb,
	)

	processor.addMnemonic(
		Mnemonic_SC,
		NewDecodeMethod(0xe0000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sc,
	)

	processor.addMnemonic(
		Mnemonic_SDC1,
		NewDecodeMethod(0xf4000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sdc1,
	)

	processor.addMnemonic(
		Mnemonic_SH,
		NewDecodeMethod(0xa4000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sh,
	)

	processor.addMnemonic(
		Mnemonic_SLL,
		NewDecodeMethod(0x00000000, 0xffe0003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		sll,
	)

	processor.addMnemonic(
		Mnemonic_SLLV,
		NewDecodeMethod(0x00000004, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		sllv,
	)

	processor.addMnemonic(
		Mnemonic_SLT,
		NewDecodeMethod(0x0000002a, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		slt,
	)

	processor.addMnemonic(
		Mnemonic_SLTI,
		NewDecodeMethod(0x28000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		slti,
	)

	processor.addMnemonic(
		Mnemonic_SLTIU,
		NewDecodeMethod(0x2c000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		sltiu,
	)

	processor.addMnemonic(
		Mnemonic_SLTU,
		NewDecodeMethod(0x0000002b, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		sltu,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_S,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		sqrt_s,
	)

	processor.addMnemonic(
		Mnemonic_SQRT_D,
		NewDecodeMethod(0x44000004, 0xfc1f003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		sqrt_d,
	)

	processor.addMnemonic(
		Mnemonic_SRA,
		NewDecodeMethod(0x00000003, 0xffe0003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		sra,
	)

	processor.addMnemonic(
		Mnemonic_SRAV,
		NewDecodeMethod(0x00000007, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		srav,
	)

	processor.addMnemonic(
		Mnemonic_SRL,
		NewDecodeMethod(0x00000002, 0xffe0003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		srl,
	)

	processor.addMnemonic(
		Mnemonic_SRLV,
		NewDecodeMethod(0x00000006, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		srlv,
	)

	processor.addMnemonic(
		Mnemonic_SUB_S,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_SINGLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		sub_s,
	)

	processor.addMnemonic(
		Mnemonic_SUB_D,
		NewDecodeMethod(0x44000001, 0xfc00003f),
		NewDecodeCondition(FMT, FMT_DOUBLE),
		[]StaticInstFlag{
			StaticInstFlag_FLOAT_COMPUTATION,
		},
		sub_d,
	)

	processor.addMnemonic(
		Mnemonic_SUBU,
		NewDecodeMethod(0x00000023, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		subu,
	)

	processor.addMnemonic(
		Mnemonic_SW,
		NewDecodeMethod(0xac000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		sw,
	)

	processor.addMnemonic(
		Mnemonic_SWC1,
		NewDecodeMethod(0xe4000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swc1,
	)

	processor.addMnemonic(
		Mnemonic_SWL,
		NewDecodeMethod(0xa8000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swl,
	)

	processor.addMnemonic(
		Mnemonic_SWR,
		NewDecodeMethod(0xb8000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_STORE,
			StaticInstFlag_DISPLACED_ADDRESSING,
		},
		swr,
	)

	processor.addMnemonic(
		Mnemonic_SYSCALL,
		NewDecodeMethod(0x0000000c, 0xfc00003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_TRAP,
		},
		_syscall,
	)

	processor.addMnemonic(
		Mnemonic_TRUNC_W,
		NewDecodeMethod(0x4400000d, 0xfc1f003f),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_UNIMPLEMENTED,
		},
		trunc_w,
	)

	processor.addMnemonic(
		Mnemonic_XOR,
		NewDecodeMethod(0x00000026, 0xfc0007ff),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
		},
		xor,
	)

	processor.addMnemonic(
		Mnemonic_XORI,
		NewDecodeMethod(0x38000000, 0xfc000000),
		nil,
		[]StaticInstFlag{
			StaticInstFlag_INTEGER_COMPUTATION,
			StaticInstFlag_IMMEDIATE,
		},
		xori,
	)
}
