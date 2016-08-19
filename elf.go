package acogo

import (
	"io/ioutil"
	"fmt"
)

type ElfFile struct {
	Data           *SimpleMemory
	Identification *ElfIdentification
	Header         *ElfHeader
	SectionHeaders []ElfSectionHeader
	ProgramHeaders []ElfProgramHeader
	StringTable    *ElfStringTable
}

func NewElfFile(fileName string) *ElfFile {
	var elfFile = &ElfFile{
	}

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(fmt.Sprintf("Cannot read ELF file (%s)", err))
	}

	elfFile.Data = NewSimpleMemory(false, data)

	return elfFile
}

func (elfFile *ElfFile) Dump() {
	//TODO
}

const (
	ElfClassNone = 0
	ElfClass32 = 1
	ElfClass64 = 2
)

const (
	ElfDataNone = 0
	ElfData2Lsb = 1
	ElfData2Msb = 2
)

type ElfIdentification struct {
	Clz     uint32
	Data    uint32
	Version uint32
}

func (elfIdentification *ElfIdentification) Read(elfFile *ElfFile) {
	//var e_ident = make([]byte, 16)
	//elfFile.Read()
}

type ElfHeader struct {
	HeaderType                    uint32
	Machine                       uint32
	Version                       uint64
	Entry                         uint64
	ProgramHeaderTableOffset      uint64
	SectionHeaderTableOffset      uint64
	Flags                         uint64
	ElfHeaderSize                 uint32
	ProgramHeaderTableEntrySize   uint32
	ProgramHeaderTableEntryCount  uint32
	SectionHeaderTableEntrySize   uint32
	SectionHeaderStringTableIndex uint32
}

type ElfSectionHeader struct {
	NameIndex        uint64
	HeaderType       uint64
	Flags            uint64
	Address          uint64
	Offset           uint64
	Size             uint64
	Link             uint64
	Info             uint64
	AddressAlignment uint64
	EntrySize        uint64
	ElfFile          *ElfFile
	Name             string
}

type ElfProgramHeader struct {
	HeaderType      uint64
	Offset          uint64
	VirtualAddress  uint64
	PhysicalAddress uint64
	SizeInFile      uint64
	SizeInMemory    uint64
	Flags           uint64
	Alignment       uint64
}

type ElfStringTable struct {
	Data []byte
}