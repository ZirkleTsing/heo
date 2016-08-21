package regs

import (
	"encoding/binary"
	"math"
)

var GPR_NAMES = []string{
	"zero", "at", "v0", "v1", "a0", "a1", "a2", "a3",
	"t0", "t1", "t2", "t3", "t4", "t5", "t6", "t6",
	"s0", "s1", "s2", "s3", "s4", "s5", "s6", "s7",
	"t8", "t9", "k0", "k1", "gp", "sp", "fp", "ra",
}

const (
	NUM_INT_REGISTERS = 32

	NUM_FLOAT_REGISTERS = 32

	NUM_MISC_REGISTERS = 3

	REGISTER_ZERO = 0

	REGISTER_AT = 1

	REGISTER_V0 = 2

	REGISTER_V1 = 3

	REGISTER_A0 = 4

	REGISTER_A1 = 5

	REGISTER_A2 = 6

	REGISTER_A3 = 7

	REGISTER_T0 = 8

	REGISTER_T1 = 9

	REGISTER_T2 = 10

	REGISTER_T3 = 11

	REGISTER_T4 = 12

	REGISTER_T5 = 13

	REGISTER_T6 = 14

	REGISTER_T7 = 15

	REGISTER_S0 = 16

	REGISTER_S1 = 17

	REGISTER_S2 = 18

	REGISTER_S3 = 19

	REGISTER_S4 = 20

	REGISTER_S5 = 21

	REGISTER_S6 = 22

	REGISTER_S7 = 23

	REGISTER_T8 = 24

	REGISTER_T9 = 25

	REGISTER_K0 = 26

	REGISTER_K1 = 27

	REGISTER_GP = 28

	REGISTER_SP = 29

	REGISTER_FP = 30

	REGISTER_RA = 31

	REGISTER_MISC_LO = 0

	REGISTER_MISC_HI = 1

	REGISTER_MISC_FCSR = 2
)

type ArchitecturalRegisterFile struct {
	LittleEndian bool
	Pc           uint32
	Npc          uint32
	Nnpc         uint32
	Gpr          []uint32
	Fpr          *FloatingPointRegisters
	Hi           uint32
	Lo           uint32
	Fcsr         uint32
}

func NewArchitecturalRegisterFile(littleEndian bool) *ArchitecturalRegisterFile {
	var architecturalRegisterFile = &ArchitecturalRegisterFile{
		LittleEndian:littleEndian,
		Gpr:make([]uint32, 32),
		Fpr:NewFloatingPointRegisters(littleEndian),
	}

	return architecturalRegisterFile
}

func (architecturalRegisterFile *ArchitecturalRegisterFile) Clone() *ArchitecturalRegisterFile {
	var newArchitecturalRegisterFile = NewArchitecturalRegisterFile(architecturalRegisterFile.LittleEndian)

	newArchitecturalRegisterFile.Pc = architecturalRegisterFile.Pc
	newArchitecturalRegisterFile.Npc = architecturalRegisterFile.Npc
	newArchitecturalRegisterFile.Nnpc = architecturalRegisterFile.Nnpc

	copy(newArchitecturalRegisterFile.Gpr, architecturalRegisterFile.Gpr)

	newArchitecturalRegisterFile.Fpr = NewFloatingPointRegisters(architecturalRegisterFile.LittleEndian)
	copy(newArchitecturalRegisterFile.Fpr.data, architecturalRegisterFile.Fpr.data)

	newArchitecturalRegisterFile.Hi = architecturalRegisterFile.Hi
	newArchitecturalRegisterFile.Lo = architecturalRegisterFile.Lo
	newArchitecturalRegisterFile.Fcsr = architecturalRegisterFile.Fcsr

	return newArchitecturalRegisterFile
}

func (architecturalRegisterFile *ArchitecturalRegisterFile) Sgpr(i uint32) int32 {
	return int32(architecturalRegisterFile.Gpr[i])
}

type FloatingPointRegisters struct {
	LittleEndian bool
	ByteOrder    binary.ByteOrder
	data         []byte
}

func NewFloatingPointRegisters(littleEndian bool) *FloatingPointRegisters {
	var floatingPointRegisters = &FloatingPointRegisters{
		LittleEndian:littleEndian,
		data:make([]byte, 4 * 32),
	}

	if littleEndian {
		floatingPointRegisters.ByteOrder = binary.LittleEndian
	} else {
		floatingPointRegisters.ByteOrder = binary.BigEndian
	}

	return floatingPointRegisters
}

func (floatingPointRegisters *FloatingPointRegisters) Uint32(index uint32) uint32 {
	var size = uint32(4)

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return floatingPointRegisters.ByteOrder.Uint32(buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) SetUint32(index uint32, value uint32) {
	var size = uint32(4)

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint32(buffer, value)

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) Float32(index uint32) float32 {
	var size = uint32(4)

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return math.Float32frombits(floatingPointRegisters.ByteOrder.Uint32(buffer))
}

func (floatingPointRegisters *FloatingPointRegisters) SetFloat32(index uint32, value float32) {
	var size = uint32(4)

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint32(buffer, math.Float32bits(value))

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) Uint64(index uint32) uint64 {
	var size = uint32(8)

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return floatingPointRegisters.ByteOrder.Uint64(buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) SetUint64(index uint32, value uint64) {
	var size = uint32(8)

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint64(buffer, value)

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) Float64(index uint32) float64 {
	var size = uint32(8)

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return math.Float64frombits(floatingPointRegisters.ByteOrder.Uint64(buffer))
}

func (floatingPointRegisters *FloatingPointRegisters) SetFloat64(index uint32, value float64) {
	var size = uint32(8)

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint64(buffer, math.Float64bits(value))

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}
