package isa

import (
	"math"
	"github.com/mcai/acogo/cpu/regs"
	"github.com/mcai/acogo/cpu/cpuutil"
	"github.com/mcai/acogo/cpu"
)

func add(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + context.Regs.Sgpr(machInst.Rt())
	context.Regs.Gpr[machInst.Rd()] = uint32(temp)
}

func addi(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Sgpr(machInst.Rs()) + machInst.Imm()
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func addiu(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = uint32(context.Regs.Sgpr(machInst.Rs()) + machInst.Imm())
}

func addu(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] + context.Regs.Gpr[machInst.Rt()]
}

func and(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] & context.Regs.Gpr[machInst.Rt()]
}

func andi(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] & machInst.Uimm()
}

func div(context *cpu.Context, machInst MachInst) {
	if machInst.Rt() == 0 {
		context.Regs.Lo = uint32(context.Regs.Sgpr(machInst.Rs()) / context.Regs.Sgpr(machInst.Rt()))
		context.Regs.Hi = uint32(context.Regs.Sgpr(machInst.Rs()) % context.Regs.Sgpr(machInst.Rt()))
	}
}

func divu(context *cpu.Context, machInst MachInst) {
	if machInst.Rt() == 0 {
		context.Regs.Lo = uint32(context.Regs.Gpr[machInst.Rs()] / context.Regs.Gpr[machInst.Rt()])
		context.Regs.Hi = uint32(context.Regs.Gpr[machInst.Rs()] % context.Regs.Gpr[machInst.Rt()])
	}
}

func lui(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rt()] = machInst.Uimm() << 16
}

func madd(context *cpu.Context, machInst MachInst) {
	var temp, temp1, temp2, temp3 int64
	temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	temp3 = (int64(context.Regs.Hi << 32) | int64(context.Regs.Lo))
	temp = temp1 * temp2 + temp3
	context.Regs.Hi = uint32(cpuutil.Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(cpuutil.Bits64(uint64(temp), 31, 0))
}

func mfhi(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Hi
}

func mflo(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Lo
}

func msub(context *cpu.Context, machInst MachInst) {
	var temp, temp1, temp2, temp3 int64
	temp1 = int64(context.Regs.Sgpr(machInst.Rs()))
	temp2 = int64(context.Regs.Sgpr(machInst.Rt()))
	temp3 = int64(context.Regs.Hi << 32) | int64(context.Regs.Lo)
	temp = temp3 - temp1 * temp2 + temp3
	context.Regs.Hi = uint32(cpuutil.Bits64(uint64(temp), 63, 32))
	context.Regs.Lo = uint32(cpuutil.Bits64(uint64(temp), 31, 0))
}

func mthi(context *cpu.Context, machInst MachInst) {
	context.Regs.Hi = context.Regs.Gpr[machInst.Rd()]
}

func mtlo(context *cpu.Context, machInst MachInst) {
	context.Regs.Lo = context.Regs.Gpr[machInst.Rd()]
}

func mult(context *cpu.Context, machInst MachInst) {
	var temp = uint64(int64(context.Regs.Sgpr(machInst.Rs())) * int64(context.Regs.Sgpr(machInst.Rt())))
	context.Regs.Lo = uint32(cpuutil.Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(cpuutil.Bits64(temp, 63, 32))
}

func multu(context *cpu.Context, machInst MachInst) {
	var temp = uint64(context.Regs.Sgpr(machInst.Rs())) * uint64(context.Regs.Sgpr(machInst.Rt()))
	context.Regs.Lo = uint32(cpuutil.Bits64(temp, 31, 0))
	context.Regs.Hi = uint32(cpuutil.Bits64(temp, 63, 32))
}

func nor(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
	context.Regs.Gpr[machInst.Rd()] = ^temp
}

func or(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] | context.Regs.Gpr[machInst.Rt()]
}

func ori(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] | machInst.Uimm()
}

func sll(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] << machInst.Shift()
}

func sllv(context *cpu.Context, machInst MachInst) {
	var s uint32 = cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] << s
}

func slt(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < context.Regs.Sgpr(machInst.Rt()) {
		context.Regs.Gpr[machInst.Rd()] = 1
	} else {
		context.Regs.Gpr[machInst.Rd()] = 0
	}
}

func slti(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < machInst.Imm() {
		context.Regs.Gpr[machInst.Rt()] = 1
	} else {
		context.Regs.Gpr[machInst.Rt()] = 0
	}
}

func sltiu(context *cpu.Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] < uint32(machInst.Imm()) {
		context.Regs.Gpr[machInst.Rt()] = 1
	} else {
		context.Regs.Gpr[machInst.Rt()] = 0
	}
}

func sltu(context *cpu.Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] < context.Regs.Gpr[machInst.Rt()] {
		context.Regs.Gpr[machInst.Rd()] = 1
	} else {
		context.Regs.Gpr[machInst.Rd()] = 0
	}
}

func sra(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = uint32(context.Regs.Sgpr(machInst.Rt()) >> machInst.Shift())
}

func srav(context *cpu.Context, machInst MachInst) {
	var s = int32(cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0))
	context.Regs.Gpr[machInst.Rd()] = uint32(context.Regs.Sgpr(machInst.Rt() >> uint32(s)))
}

func srl(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] >> machInst.Shift()
}

func srlv(context *cpu.Context, machInst MachInst) {
	var s = cpuutil.Bits32(context.Regs.Gpr[machInst.Rs()], 4, 0)
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rt()] >> s
}

func sub(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] - context.Regs.Gpr[machInst.Rt()]
}

func subu(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] - context.Regs.Gpr[machInst.Rt()]
}

func xor(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] ^ context.Regs.Gpr[machInst.Rt()]
}

func xori(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Gpr[machInst.Rs()] ^ machInst.Uimm()
}

func absD(context *cpu.Context, machInst MachInst) {
	var temp float64

	var fs = context.Regs.Fpr.Float64(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func absS(context *cpu.Context, machInst MachInst) {
	var temp float32

	var fs = context.Regs.Fpr.Float32(machInst.Fs())

	if fs < 0.0 {
		temp = -fs
	} else {
		temp = fs
	}

	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func addD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(
		machInst.Fd(), context.Regs.Fpr.Float64(machInst.Fs()) + context.Regs.Fpr.Float64(machInst.Ft()))
}

func addS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(
		machInst.Fd(), context.Regs.Fpr.Float32(machInst.Fs()) + context.Regs.Fpr.Float32(machInst.Ft()))
}

func cCondD(context *cpu.Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float64(machInst.Fs())
	var op2 = context.Regs.Fpr.Float64(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func cCondS(context *cpu.Context, machInst MachInst) {
	var op1 = context.Regs.Fpr.Float32(machInst.Fs())
	var op2 = context.Regs.Fpr.Float32(machInst.Ft())

	var less = op1 < op2
	var equal = op1 == op2

	var unordered = false

	cCond(context, machInst, less, equal, unordered)
}

func cCond(context *cpu.Context, machInst MachInst, less bool, equal bool, unordered bool) {
	var cc = machInst.Cc()

	var condition = (cpuutil.GetBit32(machInst.Cond(), 2) != 0 && less) ||
		(cpuutil.GetBit32(machInst.Cond(), 1) != 0 && equal) ||
		(cpuutil.GetBit32(machInst.Cond(), 0) != 0 && unordered)

	if cc != 0 {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 24 + cc, condition)
	} else {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 23, condition)
	}
}

func cvtDL(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Uint64(machInst.Fs())))
}

func cvtDS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Float32(machInst.Fs())))
}

func cvtDW(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		float64(context.Regs.Fpr.Uint32(machInst.Fs())))
}

func cvtLD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetUint64(machInst.Fd(),
		uint64(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtLS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetUint64(machInst.Fd(),
		uint64(context.Regs.Fpr.Float32(machInst.Fs())))
}

func cvtSD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtSL(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Uint64(machInst.Fs())))
}

func cvtSW(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		float32(context.Regs.Fpr.Uint32(machInst.Fs())))
}

func cvtWD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetUint32(machInst.Fd(),
		uint32(context.Regs.Fpr.Float64(machInst.Fs())))
}

func cvtWS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetUint32(machInst.Fd(),
		uint32(context.Regs.Fpr.Float32(machInst.Fs())))
}

func divD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) / context.Regs.Fpr.Float64(machInst.Ft()))
}

func divS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) / context.Regs.Fpr.Float32(machInst.Ft()))
}

func movD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(), context.Regs.Fpr.Float64(machInst.Fs()))
}

func movS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(), context.Regs.Fpr.Float32(machInst.Fs()))
}

func movf(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movf(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func movn(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movn(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movt(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func movz(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func _movz(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func mul(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func truncW(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func mulD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) * context.Regs.Fpr.Float64(machInst.Ft()))
}

func mulS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) * context.Regs.Fpr.Float32(machInst.Ft()))
}

func negD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(), -context.Regs.Fpr.Float64(machInst.Fs()))
}

func negS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(), -context.Regs.Fpr.Float32(machInst.Fs()))
}

func sqrtD(context *cpu.Context, machInst MachInst) {
	var temp = math.Sqrt(context.Regs.Fpr.Float64(machInst.Fs()))
	context.Regs.Fpr.SetFloat64(machInst.Fd(), temp)
}

func sqrtS(context *cpu.Context, machInst MachInst) {
	var temp = float32(math.Sqrt(float64(context.Regs.Fpr.Float32(machInst.Fs()))))
	context.Regs.Fpr.SetFloat32(machInst.Fd(), temp)
}

func subD(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat64(machInst.Fd(),
		context.Regs.Fpr.Float64(machInst.Fs()) - context.Regs.Fpr.Float64(machInst.Ft()))
}

func subS(context *cpu.Context, machInst MachInst) {
	context.Regs.Fpr.SetFloat32(machInst.Fd(),
		context.Regs.Fpr.Float32(machInst.Fs()) - context.Regs.Fpr.Float32(machInst.Ft()))
}

func branch(context *cpu.Context, v uint32) {
	context.Regs.Nnpc = v
}

func relBranch(context *cpu.Context, v int32) {
	context.Regs.Nnpc = uint32(int32(context.Regs.Pc) + v + 4)
}

func j(context *cpu.Context, machInst MachInst) {
	var dest = (cpuutil.Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	branch(context, dest)
}

func jal(context *cpu.Context, machInst MachInst) {
	var dest = (cpuutil.Bits32(context.Regs.Pc + 4, 32, 28) << 28) | (machInst.Target() << 2)
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
	branch(context, dest)
}

func jalr(context *cpu.Context, machInst MachInst) {
	branch(context, context.Regs.Gpr[machInst.Rs()])
	context.Regs.Gpr[machInst.Rd()] = context.Regs.Pc + 8
}

func jr(context *cpu.Context, machInst MachInst) {
	branch(context, context.Regs.Gpr[machInst.Rs()])
}

func b(context *cpu.Context, machInst MachInst) {
	relBranch(context, machInst.Imm() << 2)
}

func bal(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
	relBranch(context, machInst.Imm() << 2)
}

func fPCC(context *cpu.Context, c uint32) uint32 {
	if c != 0 {
		return cpuutil.GetBit32(context.Regs.Fcsr, 24 + c)
	} else {
		return cpuutil.GetBit32(context.Regs.Fcsr, 23)
	}
}

func SetFPCC(context *cpu.Context, c uint32, v bool) {
	if c != 0 {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 24 + c, v)
	} else {
		context.Regs.Fcsr = cpuutil.SetBitValue32(context.Regs.Fcsr, 23, v)
	}
}

func bc1f(context *cpu.Context, machInst MachInst) {
	if fPCC(context, machInst.BranchCc()) == 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bc1fl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bc1t(context *cpu.Context, machInst MachInst) {
	if fPCC(context, machInst.BranchCc()) != 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bc1tl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func beq(context *cpu.Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] == context.Regs.Gpr[machInst.Rt()] {
		relBranch(context, machInst.Imm() << 2)
	}
}

func beql(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgez(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) >= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgezal(context *cpu.Context, machInst MachInst) {
	context.Regs.Gpr[regs.REGISTER_RA] = context.Regs.Pc + 8
	if context.Regs.Sgpr(machInst.Rs()) >= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgezall(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgezl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bgtz(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) > 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bgtzl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func blez(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) <= 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func blezl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltz(context *cpu.Context, machInst MachInst) {
	if context.Regs.Sgpr(machInst.Rs()) < 0 {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bltzal(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltzall(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bltzl(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func bne(context *cpu.Context, machInst MachInst) {
	if context.Regs.Gpr[machInst.Rs()] != context.Regs.Gpr[machInst.Rt()] {
		relBranch(context, machInst.Imm() << 2)
	}
}

func bnel(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}

func lb(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp byte = context.Process.Memory.ReadByteAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(cpuutil.Sext32(uint32(temp), 8))
}

func lbu(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp byte = context.Process.Memory.ReadByteAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func ldc1(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint64 = context.Process.Memory.ReadDoubleWordAt(addr)
	context.Regs.Fpr.SetFloat64(machInst.Ft(), float64(temp))
}

func lh(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint16 = context.Process.Memory.ReadHalfWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(cpuutil.Sext32(uint32(temp), 16))
}

func lhu(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint16 = context.Process.Memory.ReadHalfWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = uint32(temp)
}

func ll(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint32 = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = temp
}

func lw(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint32 = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Gpr[machInst.Rt()] = temp
}

func lwc1(context *cpu.Context, machInst MachInst) {
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	var temp uint32 = context.Process.Memory.ReadWordAt(addr)
	context.Regs.Fpr.SetFloat32(machInst.Ft(), float32(temp))
}

func lwl(context *cpu.Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = uint64(4 - (addr & 3))

	var src = context.Process.Memory.ReadBlockAt(addr, size)

	for i := uint64(0); i < size; i++ {
		dst[3 - i] = src[i]
	}

	context.Process.Memory.ByteOrder.PutUint32(dst, context.Regs.Gpr[machInst.Rt()])
}

func lwr(context *cpu.Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = uint64(1 + (addr & 3))

	var src = context.Process.Memory.ReadBlockAt(addr - size + 1, size)

	for i := uint64(0); i < size; i++ {
		dst[size - i - 1] = src[i]
	}

	context.Process.Memory.ByteOrder.PutUint32(dst, context.Regs.Gpr[machInst.Rt()])
}

func sb(context *cpu.Context, machInst MachInst) {
	var temp byte = byte(context.Regs.Gpr[machInst.Rt()])
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteByteAt(addr, temp)
}

func sc(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
	context.Regs.Gpr[machInst.Rt()] = 1
}

func sdc1(context *cpu.Context, machInst MachInst) {
	var dbl = context.Regs.Fpr.Float64(machInst.Ft())
	var temp = uint64(dbl)
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteDoubleWordAt(addr, temp)
}

func sh(context *cpu.Context, machInst MachInst) {
	var temp = uint16(context.Regs.Gpr[machInst.Rt()])
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteHalfWordAt(addr, temp)
}

func sw(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
}

func swc1(context *cpu.Context, machInst MachInst) {
	var f = context.Regs.Fpr.Float32(machInst.Ft())
	var temp = uint32(f)
	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())
	context.Process.Memory.WriteWordAt(addr, temp)
}

func swl(context *cpu.Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = uint64(4 - (addr & 3))

	var src = make([]byte, 4)
	context.Process.Memory.ByteOrder.PutUint32(src, context.Regs.Gpr[machInst.Rt()])

	for i := uint64(0); i < size; i++ {
		dst[i] = src[3 - i]
	}

	context.Process.Memory.WriteBlockAt(addr, size, dst)
}

func swr(context *cpu.Context, machInst MachInst) {
	var dst = make([]byte, 4)

	var addr = uint64(int32(context.Regs.Gpr[machInst.Rs()]) + machInst.Imm())

	var size = uint64(1 + (addr & 3))

	var src = make([]byte, 4)
	context.Process.Memory.ByteOrder.PutUint32(src, context.Regs.Gpr[machInst.Rt()])

	for i := uint64(0); i < size; i++ {
		dst[i] = src[size - i - 1]
	}

	context.Process.Memory.WriteBlockAt(addr - size + 1, size, dst)
}

func cfc1(context *cpu.Context, machInst MachInst) {
	if machInst.Fs() == 31 {
		var temp = context.Regs.Fcsr
		context.Regs.Gpr[machInst.Rt()] = temp
	}
}

func ctc1(context *cpu.Context, machInst MachInst) {
	if machInst.Fs() != 0 {
		var temp = context.Regs.Gpr[machInst.Rt()]
		context.Regs.Fcsr = temp
	}
}

func mfc1(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Fpr.Uint32(machInst.Fs())
	context.Regs.Gpr[machInst.Rt()] = temp
}

func mtc1(context *cpu.Context, machInst MachInst) {
	var temp = context.Regs.Gpr[machInst.Rt()]
	context.Regs.Fpr.SetUint32(machInst.Fs(), temp)
}

func _break(context *cpu.Context, machInst MachInst) {
	//TODO
}

func systemCall(context *cpu.Context, machInst MachInst) {
	//TODO
}

func nop(context *cpu.Context, machInst MachInst) {
}

func unknown(context *cpu.Context, machInst MachInst) {
	panic("Unimplemented")
}
