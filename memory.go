package acogo

type Memory interface {
	ReadByte(virtualAddress uint64) byte
	ReadHalfWord(virtualAddress uint64) uint16
	ReadWord(virtualAddress uint64) uint32
	ReadDoubleWord(virtualAddress uint64) uint64
	ReadBlock(virtualAddress uint64, size uint64) []byte
	ReadString(virtualAddress uint64, size uint64) string
	WriteByte(virtualAddress uint64, data byte)
	WriteHalfWord(virtualAddress uint64, data uint16)
	WriteWord(virtualAddress uint64, data uint32)
	WriteDoubleWord(virtualAddress uint64, data uint64)
	WriteString(virtualAddress uint64, data string)
	WriteBlock(virtualAddress uint64, size uint64, data []byte)
}
