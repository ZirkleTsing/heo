package acogo

import (
	"math"
	"encoding/binary"
)

type MemoryPage struct {
	Memory          *Memory
	Id              int
	PhysicalAddress int
	Buffer          []byte
}

func NewMemoryPage(memory *Memory, id int) *MemoryPage {
	var page = &MemoryPage{
		Memory:memory,
		Id:id,
		PhysicalAddress:id << uint(memory.GetPageSizeInLog2()),
		Buffer:make([]byte, memory.GetPageSize()),
	}

	return page
}

func (page *MemoryPage) Access(virtualAddress int, buffer *[]byte, offset int, size int, write bool) {
	var displacement = page.Memory.GetDisplacement(virtualAddress)

	if write {
		copy(page.Buffer[displacement:displacement + size], (*buffer)[offset:offset + size])
	} else {
		copy((*buffer)[offset:offset + size], page.Buffer[displacement:displacement + size])
	}
}

type Memory struct {
	LittleEndian bool
	ByteOrder    binary.ByteOrder
	Pages        map[int]*MemoryPage
	Geometry     *CacheGeometry
	NumPages     int
}

func NewMemory(littleEndian bool) *Memory {
	var memory = &Memory{
		LittleEndian:littleEndian,
		Pages:make(map[int]*MemoryPage),
		Geometry:NewCacheGeometry(-1, 1, 1 << 12),
	}

	if littleEndian {
		memory.ByteOrder = binary.LittleEndian
	} else {
		memory.ByteOrder = binary.BigEndian
	}

	return memory
}

func (memory *Memory) ReadByte(virtualAddress int) byte {
	var buffer = make([]byte, 1)
	memory.Access(virtualAddress, 1, &buffer, false, true)
	return buffer[0]
}

func (memory *Memory) ReadHalfWord(virtualAddress int) uint16 {
	var buffer = make([]byte, 2)
	memory.Access(virtualAddress, 2, &buffer, false, true)
	return memory.ByteOrder.Uint16(buffer)
}

func (memory *Memory) ReadWord(virtualAddress int) uint32 {
	var buffer = make([]byte, 4)
	memory.Access(virtualAddress, 4, &buffer, false, true)
	return memory.ByteOrder.Uint32(buffer)
}

func (memory *Memory) ReadDoubleWord(virtualAddress int) uint64 {
	var buffer = make([]byte, 8)
	memory.Access(virtualAddress, 8, &buffer, false, true)
	return memory.ByteOrder.Uint64(buffer)
}

func (memory *Memory) ReadBlock(virtualAddress int, size int) []byte {
	var buffer = make([]byte, size)
	memory.Access(virtualAddress, size, &buffer, false, true)
	return buffer
}

func (memory *Memory) ReadString(virtualAddress int, size int) string {
	var data = memory.ReadBlock(virtualAddress, size)
	return string(data)
}

func (memory *Memory) WriteByte(virtualAddress int, data byte) {
	var buffer = make([]byte, 1)
	buffer[0] = data
	memory.Access(virtualAddress, 1, &buffer, true, true)
}

func (memory *Memory) WriteHalfWord(virtualAddress int, data uint16) {
	var buffer = make([]byte, 2)
	memory.ByteOrder.PutUint16(buffer, data)
	memory.Access(virtualAddress, 2, &buffer, true, true)
}

func (memory *Memory) WriteWord(virtualAddress int, data uint32) {
	var buffer = make([]byte, 4)
	memory.ByteOrder.PutUint32(buffer, data)
	memory.Access(virtualAddress, 4, &buffer, true, true)
}

func (memory *Memory) WriteDoubleWord(virtualAddress int, data uint64) {
	var buffer = make([]byte, 8)
	memory.ByteOrder.PutUint64(buffer, data)
	memory.Access(virtualAddress, 8, &buffer, true, true)
}

func (memory *Memory) WriteString(virtualAddress int, data string) {
	var buffer = []byte(data)
	memory.Access(virtualAddress, len(buffer), &buffer, true, true)
}

func (memory *Memory) Access(virtualAddress int, size int, buffer *[]byte, write bool, createNewPageIfNecessary bool) {
	var offset = 0

	var pageSize = memory.GetPageSize()

	for size > 0 {
		var chunkSize = int(math.Min(float64(size), float64(pageSize - memory.GetDisplacement(virtualAddress))))
		memory.accessPageBoundary(virtualAddress, chunkSize, buffer, offset, write, createNewPageIfNecessary)

		size -= chunkSize
		offset += chunkSize
		virtualAddress += chunkSize
	}
}

func (memory *Memory) accessPageBoundary(virtualAddress int, size int, buffer *[]byte, offset int, write bool, createNewPageIfNecessary bool) {
	var page = memory.GetPage(virtualAddress)

	if page == nil && createNewPageIfNecessary {
		page = memory.addPage(memory.GetTag(virtualAddress))
	}

	if page != nil {
		page.Access(virtualAddress, buffer, offset, size, write)
	}
}

func (memory *Memory) GetPage(virtualAddress int) *MemoryPage {
	var index = memory.GetIndex(virtualAddress)

	if page, ok := memory.Pages[index]; ok {
		return page
	}

	return nil
}

func (memory *Memory) addPage(virtualAddress int) *MemoryPage {
	var index = memory.GetIndex(virtualAddress)

	var page = NewMemoryPage(memory, memory.NumPages)

	memory.NumPages++

	memory.Pages[index] = page

	return page
}

func (memory *Memory) removePage(virtualAddress int) {
	var index = memory.GetIndex(virtualAddress)

	delete(memory.Pages, index)
}

func (memory *Memory) GetPhysicalAddress(virtualAddress int) int {
	return memory.GetPage(virtualAddress).PhysicalAddress + memory.GetDisplacement(virtualAddress)
}

func (memory *Memory) GetDisplacement(virtualAddress int) int {
	return memory.Geometry.GetDisplacement(virtualAddress)
}

func (memory *Memory) GetTag(virtualAddress int) int {
	return memory.Geometry.GetTag(virtualAddress)
}

func (memory *Memory) GetIndex(virtualAddress int) int {
	return memory.Geometry.GetLineId(virtualAddress)
}

func (memory *Memory) GetPageSizeInLog2() int {
	return memory.Geometry.LineSizeInLog2
}

func (memory *Memory) GetPageSize() int {
	return memory.Geometry.LineSize
}