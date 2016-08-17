package acogo

import "math"

func add(context *Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + context.Regs.Sgpr(machInst.Rt())
	context.Regs.Gpr[machInst.Rd()] = uint32(temp)
}

func addi(context *Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + machInst.Imm()
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func addiu(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = uint32(context.Regs.Sgpr(machInst.Rs()) + machInst.Imm())
}

func addu(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] + context.Regs.Gpr[machInst.Rt()]
}

func and(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] & context.Regs.Gpr[machInst.Rt()]
}

func andi(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] & machInst.Uimm()
}

func div(context *Context, machInst MachInst) {
	if machInst.Rt() == 0 {
		context.Regs.Lo = uint32(context.Regs.Sgpr(machInst.Rs()) / context.Regs.Sgpr(machInst.Rt()))
		context.Regs.Hi = uint32(context.Regs.Sgpr(machInst.Rs()) % context.Regs.Sgpr(machInst.Rt()))
	}
}

func divu(context *Context, machInst MachInst) {
	if machInst.Rt() == 0 {
		context.Regs.Lo = uint32(context.Regs.Gpr[machInst.Rs()] / context.Regs.Gpr[machInst.Rt()])
		context.Regs.Hi = uint32(context.Regs.Gpr[machInst.Rs()] % context.Regs.Gpr[machInst.Rt()])
	}
}

func lui(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = machInst.Uimm() << 16
}

func madd(context *Context, machInst MachInst) {
	var temp, temp1, temp2, temp3 int64
	temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	temp3 = (int64(context.Regs.Hi << 32) | int64(context.Regs.Lo))
	temp = temp1 * temp2 + temp3
	context.Regs.Hi = uint32(Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(Bits64(uint64(temp), 31, 0))
}

func mfhi(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Hi
}

func mflo(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Lo
}

func msub(context *Context, machInst MachInst) {
	var temp, temp1, temp2, temp3 int64
	temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	temp3 = int64(context.Regs.Hi << 32) | int64(context.Regs.Lo)
	temp = temp3 - temp1 * temp2 + temp3
	context.Regs.Hi = uint32(Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(Bits64(uint64(temp), 31, 0))
}

func mthi(context *Context, machInst MachInst) {
	context.Regs.Hi = context.Regs.Gpr[machInst.Rd()]
}

func mtlo(context *Context, machInst MachInst) {
	context.Regs.Lo = context.Regs.Gpr[machInst.Rd()]
}

func mult(context *Context, machInst MachInst) {
	var temp = uint64(int64(context.Regs.Sgpr(machInst.Rs())) * int64(context.Regs.Sgpr(machInst.Rt())))
	context.Regs.Lo = uint32(Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(Bits64(temp, 63, 32))
}

func multu(context *Context, machInst MachInst) {
	var temp = uint64(context.Regs.Sgpr(machInst.Rs())) * uint64(context.Regs.Sgpr(machInst.Rt()))
	context.Regs.Lo = uint32(Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(Bits64(temp, 63, 32))
}

func nor(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = ^temp
}

func or(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
}

func ori(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] | machInst.Uimm()
}

func sll(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] << machInst.Shift()
}

func sllv(context *Context, machInst MachInst) {
	var s uint32 = Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] << s
}

func slt(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < context.Regs.Sgpr(machInst.Rt()) {
		context.Regs.Gpr[machInst.Rd()] = 1
	} else {
		context.Regs.Gpr[machInst.Rd()] = 0
	}
}

func slti(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < machInst.Imm() {
		context.Regs.Gpr[machInst.Rt()] = 1
	} else {
		context.Regs.Gpr[machInst.Rt()] = 0
	}
}

func sltiu(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] < uint32(machInst.Imm()) {
		context.Regs.Gpr[machInst.Rt()] = 1
	} else {
		context.Regs.Gpr[machInst.Rt()] = 0
	}
}

func sltu(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] < context.Regs.Gpr[machInst.Rt()] {
		context.Regs.Gpr[machInst.Rd()] = 1
	} else {
		context.Regs.Gpr[machInst.Rd()] = 0
	}
}

func sra(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = uint32(context.Regs.Sgpr(machInst.Rt()) >> machInst.Shift())
}

func srav(context *Context, machInst MachInst) {
	var s = int32(Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0))
	context.Regs.Gpr[machInst.Rd()] = uint32(context.Regs.Sgpr(machInst.Rt() >> uint32(s)))
}

func srl(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] >> machInst.Shift()
}

func srlv(context *Context, machInst MachInst) {
	var s = Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] >> s
}

func sub(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] - context.Regs.Gpr[machInst.Rt()]
}

func subu(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] - context.Regs.Gpr[machInst.Rt()]
}

func xor(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] ^ context.Regs.Gpr[machInst.Rt()]
}

func xori(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] ^ machInst.Uimm()
}

func absD(context *Context, machInst MachInst) {
	var temp float64

	var fs = context.Regs.Fpr.Float64(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func absS(context *Context, machInst MachInst) {
	var temp float32

	var fs = context.Regs.Fpr.Float32(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func addD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(
		machInst.Fd(), context.Regs.Fpr.Float64(machInst.Fs()) + context.Regs.Fpr.Float64(machInst.Ft()))
}

func addS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(
		machInst.Fd(), context.Regs.Fpr.Float32(machInst.Fs()) + context.Regs.Fpr.Float32(machInst.Ft()))
}

func cCondD(context *Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float64(machInst.Fs())
	var op2 = context.Regs.Fpr.Float64(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func cCondS(context *Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float32(machInst.Fs())
	var op2 = context.Regs.Fpr.Float32(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func cCond(context *Context, machInst MachInst, less bool, equal bool, unordered bool) {
	var cc = machInst.Cc()

	var condition = (GetBit32(machInst.Cond(), 2) != 0 && less) ||
		(GetBit32(machInst.Cond(), 1) != 0 && equal) ||
		(GetBit32(machInst.Cond(), 0) != 0 && unordered)

	if cc != 0 {
		context.Regs.Fcsr = SetBitValue32(context.Regs.Fcsr, 24 + cc, condition)
	} else {
		context.Regs.Fcsr = SetBitValue32(context.Regs.Fcsr, 23, condition)
	}
}

func cvtDL(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Uint64(machInst.Fs())))
}

func cvtDS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Float32(machInst.Fs())))
}

func cvtDW(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Uint32(machInst.Fs())))
}

func cvtLD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetUint64(machInst.Fd(),
		uint64(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtLS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetUint64(machInst.Fd(),
		uint64(context.Regs.Fpr.Float32(machInst.Fs())))
}

func cvtSD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtSL(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Uint64(machInst.Fs())))
}

func cvtSW(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Uint32(machInst.Fs())))
}

func cvtWD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetUint32(machInst.Fd(),
		uint32(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtWS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetUint32(machInst.Fd(),
		uint32(context.Regs.Fpr.Float32(machInst.Fs())))
}

func divD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) / context.Regs.Fpr.Float64(machInst.Ft()))
}

func divS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) / context.Regs.Fpr.Float32(machInst.Ft()))
}

func movD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(), context.Regs.Fpr.Float64(machInst.Fs()))
}

func movS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(), context.Regs.Fpr.Float32(machInst.Fs()))
}

func movf(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movf(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func movn(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movn(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movt(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func movz(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movz(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func mul(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func truncW(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func mulD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) * context.Regs.Fpr.Float64(machInst.Ft()))
}

func mulS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) * context.Regs.Fpr.Float32(machInst.Ft()))
}

func negD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(), -context.Regs.Fpr.Float64(machInst.Fs()))
}

func negS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(), -context.Regs.Fpr.Float32(machInst.Fs()))
}

func sqrtD(context *Context, machInst MachInst) {
	var temp = math.Sqrt(context.Regs.Fpr.Float64(machInst.Fs()))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func sqrtS(context *Context, machInst MachInst) {
	var temp = float32(math.Sqrt(float64(context.Regs.Fpr.Float32(machInst.Fs()))))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func subD(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) - context.Regs.Fpr.Float64(machInst.Ft()))
}

func subS(context *Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) - context.Regs.Fpr.Float32(machInst.Ft()))
}

func branch(context *Context, v uint32) {
	context.Regs.Nnpc = v
}

func relBranch(context *Context, v int32) {
	context.Regs.Nnpc = uint32(int32(context.Regs.Pc) + v + 4)
}

func j(context *Context, machInst MachInst) {
	var dest = (Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	branch(context, dest)
}

func jal(context *Context, machInst MachInst) {
	var dest = (Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	context.Regs.Gpr[REGISTER_RA] = context.Regs.Pc + 8
	branch(context, dest)
}

func jalr(context *Context, machInst MachInst) {
	branch(context, context.Regs.Gpr[machInst.Rs()])
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Pc + 8
}

func jr(context *Context, machInst MachInst) {
	branch(context, context.Regs.Gpr[machInst.Rs()])
}

func b(context *Context, machInst MachInst) {
	relBranch(context, machInst.Imm() << 2)
}

func bal(context *Context, machInst MachInst) {
	context.Regs.Gpr[REGISTER_RA] = context.Regs.Pc + 8
	relBranch(context, machInst.Imm() << 2)
}

func fPCC(context *Context, c uint32) uint32 {
	if c != 0 {
		return GetBit32(context.Regs.Fcsr, 24 + c)
	} else {
		return GetBit32(context.Regs.Fcsr, 23)
	}
}

func SetFPCC(context *Context, c uint32, v bool) {
	if c != 0 {
		context.Regs.Fcsr = SetBitValue32(context.Regs.Fcsr, 24 + c, v)
	} else {
		context.Regs.Fcsr = SetBitValue32(context.Regs.Fcsr, 23, v)
	}
}

func bc1f(context *Context, machInst MachInst) {
	if fPCC(context, machInst.BranchCc()) == 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bc1fl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bc1t(context *Context, machInst MachInst) {
	if fPCC(context, machInst.BranchCc()) != 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bc1tl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func beq(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] == context.Regs.Gpr[machInst.Rt()] {
		relBranch(context, machInst.Imm() << 2)
	}
}

func beql(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgez(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) >= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgezal(context *Context, machInst MachInst) {
	context.Regs.Gpr[REGISTER_RA] = context.Regs.Pc + 8
	if context.Regs.Sgpr(machInst.Rs()) >= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgezall(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgezl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgtz(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) > 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgtzl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func blez(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) <= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func blezl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltz(context *Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bltzal(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltzall(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltzl(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func bne(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] != context.Regs.Gpr[machInst.Rt()] {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bnel(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func lb(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp byte = context.Memory.ReadByte(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(Sext32(uint32(temp), 8))
}

func lbu(context *Context, machInst MachInst) {
}

func ldc1(context *Context, machInst MachInst) {
}

func lh(context *Context, machInst MachInst) {
}

func lhu(context *Context, machInst MachInst) {
}

func ll(context *Context, machInst MachInst) {
}

func lw(context *Context, machInst MachInst) {
}

func lwc1(context *Context, machInst MachInst) {
}

func lwl(context *Context, machInst MachInst) {
}

func lwr(context *Context, machInst MachInst) {
}

func sb(context *Context, machInst MachInst) {
}

func sc(context *Context, machInst MachInst) {
}

func sdc1(context *Context, machInst MachInst) {
}

func sh(context *Context, machInst MachInst) {
}

func sw(context *Context, machInst MachInst) {
}

func swc1(context *Context, machInst MachInst) {
}

func swl(context *Context, machInst MachInst) {
}

func swr(context *Context, machInst MachInst) {
}

func cfc1(context *Context, machInst MachInst) {
}

func ctc1(context *Context, machInst MachInst) {
}

func mfc1(context *Context, machInst MachInst) {
}

func mtc1(context *Context, machInst MachInst) {
}

func _break(context *Context, machInst MachInst) {
}

func systemCall(context *Context, machInst MachInst) {
}

func nop(context *Context, machInst MachInst) {
}

func unknown(context *Context, machInst MachInst) {
}
