package acogo

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
	Pc           int
	Npc          int
	Nnpc         int
	Gprs         [32]int
	Fprs         *FloatingPointRegisters
	Hi           int
	Lo           int
	Fcsr         int
}

func NewArchitecturalRegisterFile(littleEndian bool) *ArchitecturalRegisterFile {
	var architecturalRegisterFile = &ArchitecturalRegisterFile{
		LittleEndian:littleEndian,
		Fprs:NewFloatingPointRegisters(littleEndian),
	}

	return architecturalRegisterFile
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

func (floatingPointRegisters *FloatingPointRegisters) GetUint32(index int) uint32 {
	var size = 4

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return floatingPointRegisters.ByteOrder.Uint32(buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) PutUint32(index int, value uint32) {
	var size = 4

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint32(buffer, value)

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) GetFloat32(index int) float32 {
	var size = 4

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return math.Float32frombits(floatingPointRegisters.ByteOrder.Uint32(buffer))
}

func (floatingPointRegisters *FloatingPointRegisters) PutFloat32(index int, value float32) {
	var size = 4

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint32(buffer, math.Float32bits(value))

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) GetUint64(index int) uint64 {
	var size = 8

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return floatingPointRegisters.ByteOrder.Uint64(buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) PutUint64(index int, value uint64) {
	var size = 8

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint64(buffer, value)

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}

func (floatingPointRegisters *FloatingPointRegisters) GetFloat64(index int) float64 {
	var size = 8

	var buffer = make([]byte, size)

	copy(buffer, floatingPointRegisters.data[index * size:index * size + size])

	return math.Float64frombits(floatingPointRegisters.ByteOrder.Uint64(buffer))
}

func (floatingPointRegisters *FloatingPointRegisters) PutFloat64(index int, value float64) {
	var size = 8

	var buffer = make([]byte, size)

	floatingPointRegisters.ByteOrder.PutUint64(buffer, math.Float64bits(value))

	copy(floatingPointRegisters.data[index * size:index * size + size], buffer)
}
