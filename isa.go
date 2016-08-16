package acogo

func add(context *Context, machInst MachInst) {
	var temp int32 = context.Regs.Sgpr(machInst.Rs()) + context.Regs.Sgpr(machInst.Rt())
	context.Regs.Gpr[machInst.Rd()] = uint32(temp)
}

func addi(context *Context, machInst MachInst) {
	var temp int32 = context.Regs.Sgpr(machInst.Rs()) + machInst.Imm()
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
}

func mflo(context *Context, machInst MachInst) {
}

func msub(context *Context, machInst MachInst) {
}

func mthi(context *Context, machInst MachInst) {
}

func mtlo(context *Context, machInst MachInst) {
}

func mult(context *Context, machInst MachInst) {
}

func multu(context *Context, machInst MachInst) {
}

func nor(context *Context, machInst MachInst) {
}

func or(context *Context, machInst MachInst) {
}

func ori(context *Context, machInst MachInst) {
}

func sll(context *Context, machInst MachInst) {
	context.Regs.Gpr[machInst.ValueOf(RD)] =
		context.Regs.Gpr[machInst.ValueOf(RT)] << machInst.ValueOf(SHIFT)
}

func sllv(context *Context, machInst MachInst) {
}

func slt(context *Context, machInst MachInst) {
}

func slti(context *Context, machInst MachInst) {
}

func sltiu(context *Context, machInst MachInst) {
}

func sltu(context *Context, machInst MachInst) {
}

func sra(context *Context, machInst MachInst) {
}

func srav(context *Context, machInst MachInst) {
}

func srl(context *Context, machInst MachInst) {
}

func srlv(context *Context, machInst MachInst) {
}

func sub(context *Context, machInst MachInst) {
}

func subu(context *Context, machInst MachInst) {
}

func xor(context *Context, machInst MachInst) {
}

func xori(context *Context, machInst MachInst) {
}

func absD(context *Context, machInst MachInst) {
}

func absS(context *Context, machInst MachInst) {
}

func addD(context *Context, machInst MachInst) {
}

func addS(context *Context, machInst MachInst) {
}

func cCondD(context *Context, machInst MachInst) {
}

func cCondS(context *Context, machInst MachInst) {
}

func cvtDL(context *Context, machInst MachInst) {
}

func cvtDS(context *Context, machInst MachInst) {
}

func cvtDW(context *Context, machInst MachInst) {
}

func cvtLD(context *Context, machInst MachInst) {
}

func cvtLS(context *Context, machInst MachInst) {
}

func cvtSD(context *Context, machInst MachInst) {
}

func cvtSL(context *Context, machInst MachInst) {
}

func cvtSW(context *Context, machInst MachInst) {
}

func cvtWD(context *Context, machInst MachInst) {
}

func cvtWS(context *Context, machInst MachInst) {
}

func divD(context *Context, machInst MachInst) {
}

func divS(context *Context, machInst MachInst) {
}

func movD(context *Context, machInst MachInst) {
}

func movS(context *Context, machInst MachInst) {
}

func movf(context *Context, machInst MachInst) {
}

func _movf(context *Context, machInst MachInst) {
}

func movn(context *Context, machInst MachInst) {
}

func _movn(context *Context, machInst MachInst) {
}

func _movt(context *Context, machInst MachInst) {
}

func movz(context *Context, machInst MachInst) {
}

func _movz(context *Context, machInst MachInst) {
}

func mul(context *Context, machInst MachInst) {
}

func truncW(context *Context, machInst MachInst) {
}

func mulD(context *Context, machInst MachInst) {
}

func mulS(context *Context, machInst MachInst) {
}

func negD(context *Context, machInst MachInst) {
}

func negS(context *Context, machInst MachInst) {
}

func sqrtD(context *Context, machInst MachInst) {
}

func sqrtS(context *Context, machInst MachInst) {
}

func subD(context *Context, machInst MachInst) {
}

func subS(context *Context, machInst MachInst) {
}

func j(context *Context, machInst MachInst) {
}

func jal(context *Context, machInst MachInst) {
}

func jalr(context *Context, machInst MachInst) {
}

func jr(context *Context, machInst MachInst) {
}

func b(context *Context, machInst MachInst) {
}

func bal(context *Context, machInst MachInst) {
}

func bc1f(context *Context, machInst MachInst) {
}

func bc1fl(context *Context, machInst MachInst) {
}

func bc1t(context *Context, machInst MachInst) {
}

func bc1tl(context *Context, machInst MachInst) {
}

func beq(context *Context, machInst MachInst) {
}

func beql(context *Context, machInst MachInst) {
}

func bgez(context *Context, machInst MachInst) {
}

func bgezal(context *Context, machInst MachInst) {
}

func bgezall(context *Context, machInst MachInst) {
}

func bgezl(context *Context, machInst MachInst) {
}

func bgtz(context *Context, machInst MachInst) {
}

func bgtzl(context *Context, machInst MachInst) {
}

func blez(context *Context, machInst MachInst) {
}

func blezl(context *Context, machInst MachInst) {
}

func bltz(context *Context, machInst MachInst) {
}

func bltzal(context *Context, machInst MachInst) {
}

func bltzall(context *Context, machInst MachInst) {
}

func bltzl(context *Context, machInst MachInst) {
}

func bne(context *Context, machInst MachInst) {
}

func bnel(context *Context, machInst MachInst) {
}

func lb(context *Context, machInst MachInst) {
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
