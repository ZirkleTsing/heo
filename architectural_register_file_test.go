package acogo

import (
	"testing"
	"fmt"
)

func TestFloatingPointRegisters(t *testing.T) {
	var regs = NewArchitecturalRegisterFile(true)

	regs.Fprs.PutUint32(0, 100)
	regs.Fprs.PutUint64(1, 20000043)

	fmt.Printf("%d\n", regs.Fprs.GetUint32(0))
	fmt.Printf("%d\n", regs.Fprs.GetUint64(1))

	regs.Fprs.PutFloat32(4, 100.1)
	regs.Fprs.PutFloat64(5, 20000.043)

	fmt.Printf("%f\n", regs.Fprs.GetFloat32(4))
	fmt.Printf("%f\n", regs.Fprs.GetFloat64(5))
}
