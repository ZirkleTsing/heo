package acogo

import (
	"encoding/binary"
	"math"
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
