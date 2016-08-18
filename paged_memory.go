package acogo

import (
	"math"
	"encoding/binary"
)

type MemoryPage struct {
	Memory          *PagedMemory
	Id              uint32
	PhysicalAddress uint64
	Buffer          []byte
}

func NewMemoryPage(memory *PagedMemory, id uint32) *MemoryPage {
	var page = &MemoryPage{
		Memory:memory,
		Id:id,
		PhysicalAddress:uint64(id) << uint64(memory.GetPageSizeInLog2()),
		Buffer:make([]byte, memory.GetPageSize()),
	}

	return page
}

func (page *MemoryPage) Access(virtualAddress uint64, buffer *[]byte, offset uint64, size uint64, write bool) {
	var displacement = page.Memory.GetDisplacement(virtualAddress)

	if write {
		copy(page.Buffer[displacement:displacement + size], (*buffer)[offset:offset + size])
	} else {
		copy((*buffer)[offset:offset + size], page.Buffer[displacement:displacement + size])
	}
}

type PagedMemory struct {
	LittleEndian bool
	ByteOrder    binary.ByteOrder
	Pages        map[uint64]*MemoryPage
	Geometry     *CacheGeometry
	NumPages     uint32
}

func NewPagedMemory(littleEndian bool) *PagedMemory {
	var memory = &PagedMemory{
		LittleEndian:littleEndian,
		Pages:make(map[uint64]*MemoryPage),
		Geometry:NewCacheGeometry(0xffffffff, 1, 1 << 12),
	}

	if littleEndian {
		memory.ByteOrder = binary.LittleEndian
	} else {
		memory.ByteOrder = binary.BigEndian
	}

	return memory
}

func (memory *PagedMemory) ReadByte(virtualAddress uint64) byte {
	var buffer = make([]byte, 1)
	memory.access(virtualAddress, 1, &buffer, false, true)
	return buffer[0]
}

func (memory *PagedMemory) ReadHalfWord(virtualAddress uint64) uint16 {
	var buffer = make([]byte, 2)
	memory.access(virtualAddress, 2, &buffer, false, true)
	return memory.ByteOrder.Uint16(buffer)
}

func (memory *PagedMemory) ReadWord(virtualAddress uint64) uint32 {
	var buffer = make([]byte, 4)
	memory.access(virtualAddress, 4, &buffer, false, true)
	return memory.ByteOrder.Uint32(buffer)
}

func (memory *PagedMemory) ReadDoubleWord(virtualAddress uint64) uint64 {
	var buffer = make([]byte, 8)
	memory.access(virtualAddress, 8, &buffer, false, true)
	return memory.ByteOrder.Uint64(buffer)
}

func (memory *PagedMemory) ReadBlock(virtualAddress uint64, size uint64) []byte {
	var buffer = make([]byte, size)
	memory.access(virtualAddress, size, &buffer, false, true)
	return buffer
}

func (memory *PagedMemory) ReadString(virtualAddress uint64, size uint64) string {
	var data = memory.ReadBlock(virtualAddress, size)
	return string(data)
}

func (memory *PagedMemory) WriteByte(virtualAddress uint64, data byte) {
	var buffer = make([]byte, 1)
	buffer[0] = data
	memory.access(virtualAddress, 1, &buffer, true, true)
}

func (memory *PagedMemory) WriteHalfWord(virtualAddress uint64, data uint16) {
	var buffer = make([]byte, 2)
	memory.ByteOrder.PutUint16(buffer, data)
	memory.access(virtualAddress, 2, &buffer, true, true)
}

func (memory *PagedMemory) WriteWord(virtualAddress uint64, data uint32) {
	var buffer = make([]byte, 4)
	memory.ByteOrder.PutUint32(buffer, data)
	memory.access(virtualAddress, 4, &buffer, true, true)
}

func (memory *PagedMemory) WriteDoubleWord(virtualAddress uint64, data uint64) {
	var buffer = make([]byte, 8)
	memory.ByteOrder.PutUint64(buffer, data)
	memory.access(virtualAddress, 8, &buffer, true, true)
}

func (memory *PagedMemory) WriteString(virtualAddress uint64, data string) {
	var buffer = []byte(data)
	memory.access(virtualAddress, uint64(len(buffer)), &buffer, true, true)
}

func (memory *PagedMemory) WriteBlock(virtualAddress uint64, size uint64, data []byte) {
	memory.access(virtualAddress, size, &data, true, true)
}

func (memory *PagedMemory) access(virtualAddress uint64, size uint64, buffer *[]byte, write bool, createNewPageIfNecessary bool) {
	var offset uint64 = 0

	var pageSize = memory.GetPageSize()

	for size > 0 {
		var chunkSize = uint64(math.Min(float64(size), float64(pageSize - memory.GetDisplacement(virtualAddress))))
		memory.accessPageBoundary(virtualAddress, chunkSize, buffer, offset, write, createNewPageIfNecessary)

		size -= chunkSize
		offset += chunkSize
		virtualAddress += uint64(chunkSize)
	}
}

func (memory *PagedMemory) accessPageBoundary(virtualAddress uint64, size uint64, buffer *[]byte, offset uint64, write bool, createNewPageIfNecessary bool) {
	var page = memory.GetPage(virtualAddress)

	if page == nil && createNewPageIfNecessary {
		page = memory.addPage(memory.GetTag(virtualAddress))
	}

	if page != nil {
		page.Access(virtualAddress, buffer, offset, size, write)
	}
}

func (memory *PagedMemory) GetPage(virtualAddress uint64) *MemoryPage {
	var index = memory.GetIndex(virtualAddress)

	if page, ok := memory.Pages[index]; ok {
		return page
	}

	return nil
}

func (memory *PagedMemory) addPage(virtualAddress uint64) *MemoryPage {
	var index = memory.GetIndex(virtualAddress)

	var page = NewMemoryPage(memory, memory.NumPages)

	memory.NumPages++

	memory.Pages[index] = page

	return page
}

func (memory *PagedMemory) removePage(virtualAddress uint64) {
	var index = memory.GetIndex(virtualAddress)

	delete(memory.Pages, index)
}

func (memory *PagedMemory) GetPhysicalAddress(virtualAddress uint64) uint64 {
	return memory.GetPage(virtualAddress).PhysicalAddress + memory.GetDisplacement(virtualAddress)
}

func (memory *PagedMemory) GetDisplacement(virtualAddress uint64) uint64 {
	return memory.Geometry.GetDisplacement(virtualAddress)
}

func (memory *PagedMemory) GetTag(virtualAddress uint64) uint64 {
	return memory.Geometry.GetTag(virtualAddress)
}

func (memory *PagedMemory) GetIndex(virtualAddress uint64) uint64 {
	return memory.Geometry.GetLineId(virtualAddress)
}

func (memory *PagedMemory) GetPageSizeInLog2() uint32 {
	return memory.Geometry.LineSizeInLog2
}

func (memory *PagedMemory) GetPageSize() uint64 {
	return memory.Geometry.LineSize
}