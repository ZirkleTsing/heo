package acogo

import "encoding/binary"

type SimpleMemory struct {
	LittleEndian bool
	ByteOrder    binary.ByteOrder
	Data        []byte
}

func NewSimpleMemory(littleEndian bool, data []byte) *SimpleMemory {
	var memory = &SimpleMemory{
		LittleEndian:littleEndian,
		Data:data,
	}

	return memory
}
