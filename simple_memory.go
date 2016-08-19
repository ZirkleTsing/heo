package acogo

import "encoding/binary"

type SimpleMemory struct {
	LittleEndian  bool
	ByteOrder     binary.ByteOrder
	Data          []byte
	ReadPosition  uint64
	WritePosition uint64
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

func (memory *SimpleMemory) ReadByteAt(virtualAddress uint64) byte {
	var buffer = make([]byte, 1)
	memory.access(virtualAddress, 1, &buffer, false)
	return buffer[0]
}

func (memory *SimpleMemory) ReadHalfWordAt(virtualAddress uint64) uint16 {
	var buffer = make([]byte, 2)
	memory.access(virtualAddress, 2, &buffer, false)
	return memory.ByteOrder.Uint16(buffer)
}

func (memory *SimpleMemory) ReadWordAt(virtualAddress uint64) uint32 {
	var buffer = make([]byte, 4)
	memory.access(virtualAddress, 4, &buffer, false)
	return memory.ByteOrder.Uint32(buffer)
}

func (memory *SimpleMemory) ReadDoubleWordAt(virtualAddress uint64) uint64 {
	var buffer = make([]byte, 8)
	memory.access(virtualAddress, 8, &buffer, false)
	return memory.ByteOrder.Uint64(buffer)
}

func (memory *SimpleMemory) ReadBlockAt(virtualAddress uint64, size uint64) []byte {
	var buffer = make([]byte, size)
	memory.access(virtualAddress, size, &buffer, false)
	return buffer
}

func (memory *SimpleMemory) ReadStringAt(virtualAddress uint64, size uint64) string {
	var data = memory.ReadBlockAt(virtualAddress, size)
	return string(data)
}

func (memory *SimpleMemory) WriteByteAt(virtualAddress uint64, data byte) {
	var buffer = make([]byte, 1)
	buffer[0] = data
	memory.access(virtualAddress, 1, &buffer, true)
}

func (memory *SimpleMemory) WriteHalfWordAt(virtualAddress uint64, data uint16) {
	var buffer = make([]byte, 2)
	memory.ByteOrder.PutUint16(buffer, data)
	memory.access(virtualAddress, 2, &buffer, true)
}

func (memory *SimpleMemory) WriteWordAt(virtualAddress uint64, data uint32) {
	var buffer = make([]byte, 4)
	memory.ByteOrder.PutUint32(buffer, data)
	memory.access(virtualAddress, 4, &buffer, true)
}

func (memory *SimpleMemory) WriteDoubleWordAt(virtualAddress uint64, data uint64) {
	var buffer = make([]byte, 8)
	memory.ByteOrder.PutUint64(buffer, data)
	memory.access(virtualAddress, 8, &buffer, true)
}

func (memory *SimpleMemory) WriteStringAt(virtualAddress uint64, data string) {
	var buffer = []byte(data)
	memory.access(virtualAddress, uint64(len(buffer)), &buffer, true)
}

func (memory *SimpleMemory) WriteBlockAt(virtualAddress uint64, size uint64, data []byte) {
	memory.access(virtualAddress, size, &data, true)
}

func (memory *SimpleMemory) ReadByte() byte {
	var data = memory.ReadByteAt(memory.ReadPosition)
	memory.ReadPosition++
	return data
}

func (memory *SimpleMemory) ReadHalfWord() uint16 {
	var data = memory.ReadHalfWordAt(memory.ReadPosition)
	memory.ReadPosition += 2
	return data
}

func (memory *SimpleMemory) ReadWord() uint32 {
	var data = memory.ReadWordAt(memory.ReadPosition)
	memory.ReadPosition += 4
	return data
}

func (memory *SimpleMemory) ReadDoubleWord() uint64 {
	var data = memory.ReadDoubleWordAt(memory.ReadPosition)
	memory.ReadPosition += 8
	return data
}

func (memory *SimpleMemory) ReadString(size uint64) string {
	var data = memory.ReadStringAt(memory.ReadPosition, size)
	memory.ReadPosition += size
	return data
}

func (memory *SimpleMemory) ReadBlock(size uint64) []byte {
	var data = memory.ReadBlockAt(memory.ReadPosition, size)
	memory.ReadPosition += size
	return data
}

func (memory *SimpleMemory) WriteByte(data byte) {
	memory.WriteByteAt(memory.WritePosition, data)
	memory.WritePosition++
}

func (memory *SimpleMemory) WriteHalfWord(data uint16) {
	memory.WriteHalfWordAt(memory.WritePosition, data)
	memory.WritePosition += 2
}

func (memory *SimpleMemory) WriteWord(data uint32) {
	memory.WriteWordAt(memory.WritePosition, data)
	memory.WritePosition += 4
}

func (memory *SimpleMemory) WriteDoubleWord(data uint64) {
	memory.WriteDoubleWordAt(memory.WritePosition, data)
	memory.WritePosition += 8
}

func (memory *SimpleMemory) WriteString(data string) {
	memory.WriteStringAt(memory.WritePosition, data)
	memory.WritePosition += uint64(len([]byte(data)))
}

func (memory *SimpleMemory) WriteBlock(size uint64, data []byte) {
	memory.WriteBlockAt(memory.WritePosition, size, data)
	memory.WritePosition += size
}

func (memory *SimpleMemory) access(virtualAddress uint64, size uint64, buffer *[]byte, write bool) {
	if write {
		copy(memory.Data[virtualAddress:virtualAddress + size], (*buffer)[0:size])
	} else {
		copy((*buffer)[0:size], memory.Data[virtualAddress:virtualAddress + size])
	}
}
