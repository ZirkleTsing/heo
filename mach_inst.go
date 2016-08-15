package acogo

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	MachInstType_R = "r"
	MachInstType_I = "i"
	MachInstType_J = "j"
	MachInstType_F = "f"
)

type MachInstType string

type MachInst uint32

func (machInst MachInst) ValueOf(field *BitField) uint32 {
	return Bits(uint32(machInst), field.Hi, field.Lo)
}

func (machInst MachInst) GetType() MachInstType {
	var opcode = machInst.ValueOf(OPCODE)

	switch opcode {
	case 0:
		return MachInstType_R
	case 0x02, 0x03:
		return MachInstType_J
	case 0x11:
		return MachInstType_F
	default:
		return MachInstType_I
	}
}

func (machInst MachInst) IsRMt() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ == 0x10 || func_ == 0x11
}

func (machInst MachInst) IsRMf() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ == 0x12 || func_ == 0x13
}

func (machInst MachInst) IsROneOp() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ == 0x08 || func_ == 0x09
}

func (machInst MachInst) IsRTwoOp() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ >= 0x18 && func_ <= 0x1b
}

func (machInst MachInst) IsLoadStore() bool {
	var opcode = machInst.ValueOf(OPCODE)
	return opcode >= 0x20 && opcode <= 0x2e || opcode == 0x30 || opcode == 0x38
}

func (machInst MachInst) IsFPLoadStore() bool {
	var opcode = machInst.ValueOf(OPCODE)
	return opcode == 0x31 || opcode == 0x39
}

func (machInst MachInst) IsOneOpBranch() bool {
	var opcode = machInst.ValueOf(OPCODE)
	return opcode == 0x00 || opcode == 0x01 || opcode == 0x06 || opcode == 0x07
}

func (machInst MachInst) IsShift() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ == 0x00 || func_ == 0x01 || func_ == 0x03
}

func (machInst MachInst) IsCVT() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ == 32 || func_ == 33 || func_ == 36
}

func (machInst MachInst) IsCompare() bool {
	var func_ = machInst.ValueOf(FUNC)
	return func_ >= 48
}

func (machInst MachInst) IsGprFpMove() bool {
	var rs = machInst.ValueOf(RS)
	return rs == 0 || rs == 4
}

func (machInst MachInst) IsGprFcrMove() bool {
	var rs = machInst.ValueOf(RS)
	return rs == 2 || rs == 6
}

func (machInst MachInst) IsFpBranch() bool {
	var rs = machInst.ValueOf(RS)
	return rs == 8
}

func (machInst MachInst) IsSystemCall() bool {
	var opcodeLo = machInst.ValueOf(OPCODE_LO)
	var funcHi = machInst.ValueOf(FUNC_HI)
	var funcLo = machInst.ValueOf(FUNC_LO)
	return opcodeLo == 0x0 && funcHi == 0x1 && funcLo == 0x4
}

func Disassemble(pc uint, staticInst *StaticInst) string {
	var buf bytes.Buffer

	var machInst = staticInst.MachInst

	buf.WriteString(fmt.Sprintf("0x%08x: 0x%08x %s ",
		pc, machInst, strings.ToLower(string(staticInst.Mnemonic.Name))))

	if machInst == 0x00000000 {
		return buf.String()
	}

	var machInstType = machInst.GetType()

	var imm = SignExtend(machInst.ValueOf(INTIMM))

	var rs = machInst.ValueOf(RS)
	var rt = machInst.ValueOf(RT)
	var rd = machInst.ValueOf(RD)

	var fs = machInst.ValueOf(FS)
	var ft = machInst.ValueOf(FT)
	var fd = machInst.ValueOf(FD)

	var shift = machInst.ValueOf(SHIFT)

	var target = machInst.ValueOf(TARGET)

	switch machInstType {
	case MachInstType_J:
		buf.WriteString(fmt.Sprintf("%x", target))
	case MachInstType_I:
		if (machInst.IsOneOpBranch()) {
			buf.WriteString(fmt.Sprintf("$%s, %d", GPR_NAMES[rs], imm));
		} else if (machInst.IsLoadStore()) {
			buf.WriteString(fmt.Sprintf("$%s, %d($%s)", GPR_NAMES[rt], imm, GPR_NAMES[rs]));
		} else if (machInst.IsFPLoadStore()) {
			buf.WriteString(fmt.Sprintf("$f%d, %d($%s)", ft, imm, GPR_NAMES[rs]));
		} else {
			buf.WriteString(fmt.Sprintf("$%s, $%s, %d", GPR_NAMES[rt], GPR_NAMES[rs], imm));
		}
	case MachInstType_F:
		if (machInst.IsCVT()) {
			buf.WriteString(fmt.Sprintf("$f%d, $f%d", fd, fs));
		} else if (machInst.IsCompare()) {
			buf.WriteString(fmt.Sprintf("%d, $f%d, $f%d", fd >> 2, fs, ft));
		} else if (machInst.IsFpBranch()) {
			buf.WriteString(fmt.Sprintf("%d, %d", fd >> 2, imm));
		} else if (machInst.IsGprFpMove()) {
			buf.WriteString(fmt.Sprintf("$%s, $f%d", GPR_NAMES[rt], fs));
		} else if (machInst.IsGprFcrMove()) {
			buf.WriteString(fmt.Sprintf("$%s, $%d", GPR_NAMES[rt], fs));
		} else {
			buf.WriteString(fmt.Sprintf("$f%d, $f%d, $f%d", fd, fs, ft));
		}
	case MachInstType_R:
		if (!machInst.IsSystemCall()) {
			if (machInst.IsShift()) {
				buf.WriteString(fmt.Sprintf("$%s, $%s, %d", GPR_NAMES[rd], GPR_NAMES[rt], shift));
			} else if (machInst.IsROneOp()) {
				buf.WriteString(fmt.Sprintf("$%s", GPR_NAMES[rs]));
			} else if (machInst.IsRTwoOp()) {
				buf.WriteString(fmt.Sprintf("$%s, $%s", GPR_NAMES[rs], GPR_NAMES[rt]));
			} else if (machInst.IsRMt()) {
				buf.WriteString(fmt.Sprintf("$%s", GPR_NAMES[rs]));
			} else if (machInst.IsRMf()) {
				buf.WriteString(fmt.Sprintf("$%s", GPR_NAMES[rd]));
			} else {
				buf.WriteString(fmt.Sprintf("$%s, $%s, $%s", GPR_NAMES[rd], GPR_NAMES[rs], GPR_NAMES[rt]));
			}
		}
	default:
		panic("Impossible!")
	}

	return buf.String()
}
