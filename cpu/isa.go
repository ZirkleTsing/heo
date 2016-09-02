package cpu

import (
	"math"
	"github.com/mcai/acogo/cpu/regs"
	"github.com/mcai/acogo/cpu/cpuutil"
)

func add(context *Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + context.Regs.Sgpr(machInst.Rt())
	context.Regs.Gpr[machInst.Rd()] = uint32(temp)
}

func addi(context *Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + machInst.Imm()
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func addiu(context *Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + machInst.Imm()
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func addu(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] + context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = temp
}

func and(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] & context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = temp
}

func andi(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = context.Regs.Gpr[machInst.Rs()] & machInst.Uimm()
}

func div(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rt()] != 0 {
		context.Regs.Lo = uint32(context.Regs.Sgpr(machInst.Rs()) / context.Regs.Sgpr(machInst.Rt()))
		context.Regs.Hi = uint32(context.Regs.Sgpr(machInst.Rs()) % context.Regs.Sgpr(machInst.Rt()))
	}
}

func divu(context *Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rt()] != 0 {
		context.Regs.Lo = context.Regs.Gpr[machInst.Rs()] / context.Regs.Gpr[machInst.Rt()]
		context.Regs.Hi = context.Regs.Gpr[machInst.Rs()] % context.Regs.Gpr[machInst.Rt()]
	}
}

func lui(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = machInst.Uimm() << 16
}

func madd(context *Context, machInst MachInst) {
	var temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	var temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	var temp3 = (int64(context.Regs.Hi << 32) | int64(context.Regs.Lo))
	var temp = temp1 * temp2 + temp3
	context.Regs.Hi = uint32(cpuutil.Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(cpuutil.Bits64(uint64(temp), 31, 0))
}

func mfhi(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Hi
}

func mflo(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Lo
}

func msub(context *Context, machInst MachInst) {
	var temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	var temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	var temp3 = int64(context.Regs.Hi << 32) | int64(context.Regs.Lo)
	var temp = temp3 - temp1 * temp2 + temp3
	context.Regs.Hi = uint32(cpuutil.Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(cpuutil.Bits64(uint64(temp), 31, 0))
}

func mthi(context *Context, machInst MachInst) {
	context.Regs.Hi = context.Regs.Gpr[machInst.Rs()]
}

func mtlo(context *Context, machInst MachInst) {
	context.Regs.Lo = context.Regs.Gpr[machInst.Rs()]
}

func mult(context *Context, machInst MachInst) {
	var temp = uint64(int64(context.Regs.Sgpr(machInst.Rs())) * int64(context.Regs.Sgpr(machInst.Rt())))
	context.Regs.Lo = uint32(cpuutil.Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(cpuutil.Bits64(temp, 63, 32))
}

func multu(context *Context, machInst MachInst) {
	var temp = uint64(context.Regs.Gpr[machInst.Rs()]) * uint64(context.Regs.Gpr[machInst.Rt()])
	context.Regs.Lo = uint32(cpuutil.Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(cpuutil.Bits64(temp, 63, 32))
}

func nor(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = ^temp
}

func or(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = temp
}

func ori(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] | machInst.Uimm()
	context.Regs.Gpr[machInst.Rt()] = temp
}

func sll(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()] << machInst.Shift()
	context.Regs.Gpr[machInst.Rd()] = temp
}

func sllv(context *Context, machInst MachInst) {
	var s = cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
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
	var s = int32(cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0))
	context.Regs.Gpr[machInst.Rd()] = uint32(context.Regs.Sgpr(machInst.Rt()) >> uint32(s))
}

func srl(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] >> machInst.Shift()
}

func srlv(context *Context, machInst MachInst) {
	var s = cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
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
	context.Regs.Gpr[machInst.Rt()] = context.Regs.Gpr[machInst.Rs()] ^ machInst.Uimm()
}

func abs_d(context *Context, machInst MachInst) {
	var temp float64

	var fs = context.Regs.Fpr.Float64(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func abs_s(context *Context, machInst MachInst) {
	var temp float32

	var fs = context.Regs.Fpr.Float32(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func add_d(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float64(machInst.Fs()) + context.Regs.Fpr.Float64(machInst.Ft())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func add_s(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float32(machInst.Fs()) + context.Regs.Fpr.Float32(machInst.Ft())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func c_cond_d(context *Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float64(machInst.Fs())
	var op2 = context.Regs.Fpr.Float64(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func c_cond_s(context *Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float32(machInst.Fs())
	var op2 = context.Regs.Fpr.Float32(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func cCond(context *Context, machInst MachInst, less bool, equal bool, unordered bool) {
	var cc = machInst.Cc()

	var condition = (cpuutil.GetBit32(machInst.Cond(), 2) != 0 && less) ||
		(cpuutil.GetBit32(machInst.Cond(), 1) != 0 && equal) ||
		(cpuutil.GetBit32(machInst.Cond(), 0) != 0 && unordered)

	SetFPCC(context, cc, condition)
}

func cvt_d_l(context *Context, machInst MachInst) {
	var temp = math.Float64frombits(context.Regs.Fpr.Uint64(machInst.Fs()))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func cvt_d_s(context *Context, machInst MachInst) {
	var temp = float64(context.Regs.Fpr.Float32(machInst.Fs()))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func cvt_d_w(context *Context, machInst MachInst) {
	var temp = math.Float64frombits(uint64(context.Regs.Fpr.Uint32(machInst.Fs())))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func cvtLD(context *Context, machInst MachInst) {
	var temp = math.Float64bits(context.Regs.Fpr.Float64(machInst.Fs()))
	context.Regs.Fpr.SetUint64(machInst.Fd(), temp)
}

func cvtLS(context *Context, machInst MachInst) {
	var temp = math.Float64bits(float64(context.Regs.Fpr.Float32(machInst.Fs())))
	context.Regs.Fpr.SetUint64(machInst.Fd(), temp)
}

func cvt_s_d(context *Context, machInst MachInst) {
	var temp = float32(context.Regs.Fpr.Float64(machInst.Fs()))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func cvt_s_l(context *Context, machInst MachInst) {
	var temp = math.Float32frombits(uint32(context.Regs.Fpr.Uint64(machInst.Fs())))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func cvt_s_w(context *Context, machInst MachInst) {
	var temp = math.Float32frombits(context.Regs.Fpr.Uint32(machInst.Fs()))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func cvt_w_d(context *Context, machInst MachInst) {
	var temp = math.Float32bits(float32(context.Regs.Fpr.Float64(machInst.Fs())))
	context.Regs.Fpr.SetUint32(machInst.Fd(), temp)
}

func cvt_w_s(context *Context, machInst MachInst) {
	var temp = math.Float32bits(context.Regs.Fpr.Float32(machInst.Fs()))
	context.Regs.Fpr.SetUint32(machInst.Fd(), temp)
}

func div_d(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float64(machInst.Fs()) / context.Regs.Fpr.Float64(machInst.Ft())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func div_s(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float32(machInst.Fs()) / context.Regs.Fpr.Float32(machInst.Ft())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func mov_d(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float64(machInst.Fs())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func mov_s(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float32(machInst.Fs())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
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

func trunc_w(context *Context, machInst MachInst) {
	panic("Unimplemented")
}

func mul_d(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float64(machInst.Fs()) * context.Regs.Fpr.Float64(machInst.Ft())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func mul_s(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float32(machInst.Fs()) * context.Regs.Fpr.Float32(machInst.Ft())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func neg_d(context *Context, machInst MachInst) {
	var temp = -context.Regs.Fpr.Float64(machInst.Fs())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func neg_s(context *Context, machInst MachInst) {
	var temp = -context.Regs.Fpr.Float32(machInst.Fs())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func sqrt_d(context *Context, machInst MachInst) {
	var temp = math.Sqrt(context.Regs.Fpr.Float64(machInst.Fs()))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func sqrt_s(context *Context, machInst MachInst) {
	var temp = float32(math.Sqrt(float64(context.Regs.Fpr.Float32(machInst.Fs()))))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func sub_d(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float64(machInst.Fs()) - context.Regs.Fpr.Float64(machInst.Ft())
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func sub_s(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Float32(machInst.Fs()) - context.Regs.Fpr.Float32(machInst.Ft())
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func branch(context *Context, v uint32) {
	context.Regs.Nnpc = v
}

func relBranch(context *Context, v int32) {
	context.Regs.Nnpc = uint32(int32(context.Regs.Pc) + v + 4)
}

func j(context *Context, machInst MachInst) {
	var dest = (cpuutil.Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	branch(context, dest)
}

func jal(context *Context, machInst MachInst) {
	var dest = (cpuutil.Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
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
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
	relBranch(context, machInst.Imm() << 2)
}

func fPCC(context *Context, c uint32) uint32 {
	if c != 0 {
		return cpuutil.GetBit32(context.Regs.Fcsr, 24 + c)
	} else {
		return cpuutil.GetBit32(context.Regs.Fcsr, 23)
	}
}

func SetFPCC(context *Context, c uint32, v bool) {
	if c != 0 {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 24 + c, v)
	} else {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 23, v)
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
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
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
	var temp = context.Process.Memory.ReadByteAt(addr)
	context.Regs.Gpr[machInst.Rt()] = cpuutil.Sext32(uint32(temp), 8)
}

func lbu(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadByteAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func ldc1(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadDoubleWordAt(addr)
	context.Regs.Fpr.SetFloat64(machInst.Ft(), math.Float64frombits(temp))
}

func lh(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadHalfWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = cpuutil.Sext32(uint32(temp), 16)
}

func lhu(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadHalfWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func ll(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = temp
}

func lw(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = temp
}

func lwc1(context *Context, machInst MachInst) {
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Fpr.SetFloat32(machInst.Ft(), math.Float32frombits(temp))
}

func lwl(context *Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = 4 - (addr & 3)

	var src = context.Process.Memory.ReadBlockAt(addr, size)

	for i := uint32(0); i < size; i++ {
		dst[3 - i] = src[i]
	}

	context.Process.Memory.ByteOrder.PutUint32(dst, context.Regs.Gpr[machInst.Rt()])
}

func lwr(context *Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = 1 + (addr & 3)

	var src = context.Process.Memory.ReadBlockAt(addr - size + 1, size)

	for i := uint32(0); i < size; i++ {
		dst[size - i - 1] = src[i]
	}

	context.Process.Memory.ByteOrder.PutUint32(dst, context.Regs.Gpr[machInst.Rt()])
}

func sb(context *Context, machInst MachInst) {
	var temp byte = byte(context.Regs.Gpr[machInst.Rt()])
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteByteAt(addr, temp)
}

func sc(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
	context.Regs.Gpr[machInst.Rt()] = 1
}

func sdc1(context *Context, machInst MachInst) {
	var dbl = context.Regs.Fpr.Float64(machInst.Ft())
	var temp = math.Float64bits(dbl)
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteDoubleWordAt(addr, temp)
}

func sh(context *Context, machInst MachInst) {
	var temp = uint16(context.Regs.Gpr[machInst.Rt()])
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteHalfWordAt(addr, temp)
}

func sw(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
}

func swc1(context *Context, machInst MachInst) {
	var f = context.Regs.Fpr.Float32(machInst.Ft())
	var temp = math.Float32bits(f)
	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
}

func swl(context *Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = 4 - (addr & 3)

	var src = make([]byte, 4)
	context.Process.Memory.ByteOrder.PutUint32(src, context.Regs.Gpr[machInst.Rt()])

	for i := uint32(0); i < size; i++ {
		dst[i] = src[3 - i]
	}

	context.Process.Memory.WriteBlockAt(addr, size, dst)
}

func swr(context *Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint32(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = 1 + (addr & 3)

	var src = make([]byte, 4)
	context.Process.Memory.ByteOrder.PutUint32(src, context.Regs.Gpr[machInst.Rt()])

	for i := uint32(0); i < size; i++ {
		dst[i] = src[size - i - 1]
	}

	context.Process.Memory.WriteBlockAt(addr - size + 1, size, dst)
}

func cfc1(context *Context, machInst MachInst) {
	if machInst.Fs() == 31 {
		var temp = context.Regs.Fcsr
		context.Regs.Gpr[machInst.Rt()] = temp
	}
}

func ctc1(context *Context, machInst MachInst) {
	if machInst.Fs() != 0 {
		var temp = context.Regs.Gpr[machInst.Rt()]
		context.Regs.Fcsr = temp
	}
}

func mfc1(context *Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Uint32(machInst.Fs())
	context.Regs.Gpr[machInst.Rt()] = temp
}

func mtc1(context *Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	context.Regs.Fpr.SetUint32(machInst.Fs(), temp)
}

func _break(context *Context, machInst MachInst) {
	context.Finish()
}

func _syscall(context *Context, machInst MachInst) {
	context.Kernel.SyscallEmulation.DoSyscall(context.Regs.Gpr[regs.REGISTER_V0], context)
}

func nop(context *Context, machInst MachInst) {
}

func unknown(context *Context, machInst MachInst) {
	panic("Unimplemented")
}
