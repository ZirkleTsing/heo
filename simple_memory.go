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

	if littleEndian {
		memory.ByteOrder = binary.LittleEndian
	} else {
		memory.ByteOrder = binary.BigEndian
	}

	return memory
}

func (memory *SimpleMemory) ReadByte(virtualAddress uint64) byte {
	var buffer = make([]byte, 1)
	memory.access(virtualAddress, 1, &buffer, false)
	return buffer[0]
}

func (memory *SimpleMemory) ReadHalfWord(virtualAddress uint64) uint16 {
	var buffer = make([]byte, 2)
	memory.access(virtualAddress, 2, &buffer, false)
	return memory.ByteOrder.Uint16(buffer)
}

func (memory *SimpleMemory) ReadWord(virtualAddress uint64) uint32 {
	var buffer = make([]byte, 4)
	memory.access(virtualAddress, 4, &buffer, false)
	return memory.ByteOrder.Uint32(buffer)
}

func (memory *SimpleMemory) ReadDoubleWord(virtualAddress uint64) uint64 {
	var buffer = make([]byte, 8)
	memory.access(virtualAddress, 8, &buffer, false)
	return memory.ByteOrder.Uint64(buffer)
}

func (memory *SimpleMemory) ReadBlock(virtualAddress uint64, size uint64) []byte {
	var buffer = make([]byte, size)
	memory.access(virtualAddress, size, &buffer, false)
	return buffer
}

func (memory *SimpleMemory) ReadString(virtualAddress uint64, size uint64) string {
	var data = memory.ReadBlock(virtualAddress, size)
	return string(data)
}

func (memory *SimpleMemory) WriteByte(virtualAddress uint64, data byte) {
	var buffer = make([]byte, 1)
	buffer[0] = data
	memory.access(virtualAddress, 1, &buffer, true)
}

func (memory *SimpleMemory) WriteHalfWord(virtualAddress uint64, data uint16) {
	var buffer = make([]byte, 2)
	memory.ByteOrder.PutUint16(buffer, data)
	memory.access(virtualAddress, 2, &buffer, true)
}

func (memory *SimpleMemory) WriteWord(virtualAddress uint64, data uint32) {
	var buffer = make([]byte, 4)
	memory.ByteOrder.PutUint32(buffer, data)
	memory.access(virtualAddress, 4, &buffer, true)
}

func (memory *SimpleMemory) WriteDoubleWord(virtualAddress uint64, data uint64) {
	var buffer = make([]byte, 8)
	memory.ByteOrder.PutUint64(buffer, data)
	memory.access(virtualAddress, 8, &buffer, true)
}

func (memory *SimpleMemory) WriteString(virtualAddress uint64, data string) {
	var buffer = []byte(data)
	memory.access(virtualAddress, uint64(len(buffer)), &buffer, true)
}

func (memory *SimpleMemory) WriteBlock(virtualAddress uint64, size uint64, data []byte) {
	memory.access(virtualAddress, size, &data, true)
}

func (memory *SimpleMemory) access(virtualAddress uint64, size uint64, buffer *[]byte, write bool) {
	if write {
		copy(memory.Data[virtualAddress:virtualAddress + size], (*buffer)[0:size])
	} else {
		copy((*buffer)[0:size], memory.Data[virtualAddress:virtualAddress + size])
	}
}
