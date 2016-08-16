package acogo

func add(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RD)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] + context.Regs.Gprs[machInst.ValueOf(RT)]
}

func addi(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RT)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] + SignExtend(machInst.ValueOf(INTIMM))
}

func addiu(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RT)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] + SignExtend(machInst.ValueOf(INTIMM))
}

func addu(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RD)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] + context.Regs.Gprs[machInst.ValueOf(RT)]
}

func and(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RD)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] & context.Regs.Gprs[machInst.ValueOf(RT)]
}

func andi(context *Context, machInst MachInst) {
	context.Regs.Gprs[machInst.ValueOf(RT)] =
		context.Regs.Gprs[machInst.ValueOf(RS)] & ZeroExtend(machInst.ValueOf(INTIMM))
}

func div(context *Context, machInst MachInst) {
	var rs = context.Regs.Gprs[machInst.ValueOf(RS)]
	var rt = context.Regs.Gprs[machInst.ValueOf(RT)]

	if rt == 0 {
		context.Regs.Hi = 0
		context.Regs.Lo = 0
	} else {
		context.Regs.Hi = rs % rt
		context.Regs.Lo = rs / rt
	}
}

func divu(context *Context, machInst MachInst) {
	var rs = context.Regs.Gprs[machInst.ValueOf(RS)]
	var rt = context.Regs.Gprs[machInst.ValueOf(RT)]

	if rt == 0 {
		context.Regs.Hi = 0
		context.Regs.Lo = 0
	} else {
		context.Regs.Hi = rs % rt
		context.Regs.Lo = rs / rt
	}
}
