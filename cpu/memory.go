package cpu

type Memory interface {
	ReadByteAt(virtualAddress uint64) byte
	ReadHalfWordAt(virtualAddress uint64) uint16
	ReadWordAt(virtualAddress uint64) uint32
	ReadDoubleWordAt(virtualAddress uint64) uint64
	ReadBlockAt(virtualAddress uint64, size uint64) []byte
	ReadStringAt(virtualAddress uint64, size uint64) string
	WriteByteAt(virtualAddress uint64, data byte)
	WriteHalfWordAt(virtualAddress uint64, data uint16)
	WriteWordAt(virtualAddress uint64, data uint32)
	WriteDoubleWordAt(virtualAddress uint64, data uint64)
	WriteStringAt(virtualAddress uint64, data string)
	WriteBlockAt(virtualAddress uint64, size uint64, data []byte)
}

type MemoryReader interface {
	ReadByte() byte
	ReadHalfWord() uint16
	ReadWord() uint32
	ReadDoubleWord() uint64
	ReadString(size uint64) string
	ReadBlock(size uint64) []byte
}

type MemoryWriter interface {
	WriteByte(data byte)
	WriteHalfWord(data uint16)
	WriteWord(data uint32)
	WriteDoubleWord(data uint64)
	WriteString(data string)
	WriteBlock(size uint64, data []byte)
}
