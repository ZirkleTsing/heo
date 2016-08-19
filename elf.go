package acogo

import (
	"io/ioutil"
	"fmt"
	"encoding/binary"
)

type ElfFile struct {
	Data           *SimpleMemory
	Identification *ElfIdentification
	Header         *ElfHeader
	SectionHeaders []*ElfSectionHeader
	ProgramHeaders []*ElfProgramHeader
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

	elfFile.Identification = NewElfIdentification(elfFile)

	if elfFile.Identification.Clz != ElfClass32 {
		panic("ElfClass32 is not supported")
	}

	if elfFile.Identification.Data == ElfData2Lsb {
		elfFile.Data.LittleEndian = true
		elfFile.Data.ByteOrder = binary.LittleEndian
	} else {
		elfFile.Data.LittleEndian = false
		elfFile.Data.ByteOrder = binary.BigEndian
	}

	elfFile.Header = NewElfHeader(elfFile)

	if elfFile.Header.Machine != EM_MIPS {
		panic("Non-MIPS ELF file is not supported")
	}

	for i := uint16(0); i < elfFile.Header.SectionHeaderTableEntryCount; i++ {
		elfFile.Data.ReadPosition = uint64(
			elfFile.Header.SectionHeaderTableOffset +
				uint32(i * elfFile.Header.SectionHeaderTableEntrySize))
		elfFile.SectionHeaders = append(elfFile.SectionHeaders, NewElfSectionHeader(elfFile))
	}

	elfFile.StringTable = NewElfStringTable(elfFile, elfFile.SectionHeaders[elfFile.Header.SectionHeaderStringTableIndex])

	elfFile.Data.ReadPosition = uint64(elfFile.Header.ProgramHeaderTableOffset)

	for i := uint16(0); i < elfFile.Header.ProgramHeaderTableEntryCount; i++ {
		elfFile.ProgramHeaders = append(elfFile.ProgramHeaders, NewElfProgramHeader(elfFile))
	}

	return elfFile
}

type ElfClass string

const (
	ElfClassNone ElfClass = "ElfClassNone"
	ElfClass32 ElfClass = "ElfClass32"
	ElfClass64 ElfClass = "ElfClass64"
)

type ElfData string

const (
	ElfDataNone ElfData = "ElfDataNone"
	ElfData2Lsb ElfData = "ElfData2Lsb"
	ElfData2Msb ElfData = "ElfData2Msb"
)

type ElfIdentification struct {
	Clz  ElfClass
	Data ElfData
}

func NewElfIdentification(elfFile *ElfFile) *ElfIdentification {
	var elfIdentification = &ElfIdentification{
	}

	var eIdent = elfFile.Data.ReadBlock(16)

	if !(eIdent[0] == 0x7f && eIdent[1] == byte('E') && eIdent[2] == byte('L') && eIdent[3] == byte('F')) {
		panic("Not ELF file")
	}

	switch eIdent[4] {
	case 1:
		elfIdentification.Clz = ElfClass32
	case 2:
		elfIdentification.Clz = ElfClass64
	default:
		elfIdentification.Clz = ElfClassNone
	}

	switch eIdent[5] {
	case 1:
		elfIdentification.Data = ElfData2Lsb
	case 2:
		elfIdentification.Data = ElfData2Msb
	default:
		elfIdentification.Data = ElfDataNone
	}

	return elfIdentification
}

const (
	EM_MIPS uint16 = 8
)

type ElfHeader struct {
	HeaderType                    uint16
	Machine                       uint16
	Version                       uint32
	Entry                         uint32
	ProgramHeaderTableOffset      uint32
	SectionHeaderTableOffset      uint32
	Flags                         uint32
	ElfHeaderSize                 uint16
	ProgramHeaderTableEntrySize   uint16
	ProgramHeaderTableEntryCount  uint16
	SectionHeaderTableEntrySize   uint16
	SectionHeaderTableEntryCount  uint16
	SectionHeaderStringTableIndex uint16
}

func NewElfHeader(elfFile *ElfFile) *ElfHeader {
	var header = &ElfHeader{
	}

	header.HeaderType = elfFile.Data.ReadHalfWord()

	header.Machine = elfFile.Data.ReadHalfWord()
	header.Version = elfFile.Data.ReadWord()
	header.Entry = elfFile.Data.ReadWord()
	header.ProgramHeaderTableOffset = elfFile.Data.ReadWord()
	header.SectionHeaderTableOffset = elfFile.Data.ReadWord()
	header.Flags = elfFile.Data.ReadWord()

	header.ElfHeaderSize = elfFile.Data.ReadHalfWord()
	header.ProgramHeaderTableEntrySize = elfFile.Data.ReadHalfWord()
	header.ProgramHeaderTableEntryCount = elfFile.Data.ReadHalfWord()
	header.SectionHeaderTableEntrySize = elfFile.Data.ReadHalfWord()
	header.SectionHeaderTableEntryCount = elfFile.Data.ReadHalfWord()
	header.SectionHeaderStringTableIndex = elfFile.Data.ReadHalfWord()

	return header
}

type ElfSectionHeaderType uint32

const (
	SHT_NULL ElfSectionHeaderType = 0
	SHT_PROGBITS ElfSectionHeaderType = 1
	SHT_SYMTAB ElfSectionHeaderType = 2
	SHT_STRTAB ElfSectionHeaderType = 3
	SHT_RELA ElfSectionHeaderType = 4
	SHT_HASH ElfSectionHeaderType = 5
	SHT_DYNAMIC ElfSectionHeaderType = 6
	SHT_NOTE ElfSectionHeaderType = 7
	SHT_NOBITS ElfSectionHeaderType = 8
	SHT_REL ElfSectionHeaderType = 9
	SHT_SHLIB ElfSectionHeaderType = 10
	SHT_DYNSYM ElfSectionHeaderType = 11
)

type ElfSectionHeaderFlag uint32

const (
	SHF_WRITE ElfSectionHeaderFlag = 0x1
	SHF_ALLOC ElfSectionHeaderFlag = 0x2
	SHF_EXECINSTR ElfSectionHeaderFlag = 0x4
)

type ElfSectionHeader struct {
	NameIndex        uint32
	HeaderType       ElfSectionHeaderType
	Flags            uint32
	Address          uint32
	Offset           uint32
	Size             uint32
	Link             uint32
	Info             uint32
	AddressAlignment uint32
	EntrySize        uint32
	ElfFile          *ElfFile
	Name             string
}

func NewElfSectionHeader(elfFile *ElfFile) *ElfSectionHeader {
	var elfSectionHeader = &ElfSectionHeader{
	}

	elfSectionHeader.NameIndex = elfFile.Data.ReadWord()
	elfSectionHeader.HeaderType = ElfSectionHeaderType(elfFile.Data.ReadWord())
	elfSectionHeader.Flags = elfFile.Data.ReadWord()
	elfSectionHeader.Address = elfFile.Data.ReadWord()
	elfSectionHeader.Offset = elfFile.Data.ReadWord()
	elfSectionHeader.Size = elfFile.Data.ReadWord()
	elfSectionHeader.Link = elfFile.Data.ReadWord()
	elfSectionHeader.Info = elfFile.Data.ReadWord()
	elfSectionHeader.AddressAlignment = elfFile.Data.ReadWord()
	elfSectionHeader.EntrySize = elfFile.Data.ReadWord()

	return elfSectionHeader
}

func (elfSectionHeader *ElfSectionHeader) ReadContent(elfFile *ElfFile) []byte {
	return elfFile.Data.ReadBlockAt(uint64(elfSectionHeader.Offset), uint64(elfSectionHeader.Size))
}

type ElfProgramHeader struct {
	HeaderType      uint32
	Offset          uint32
	VirtualAddress  uint32
	PhysicalAddress uint32
	SizeInFile      uint32
	SizeInMemory    uint32
	Flags           uint32
	Alignment       uint32
}

func NewElfProgramHeader(elfFile *ElfFile) *ElfProgramHeader {
	var elfProgramHeader = &ElfProgramHeader{
	}

	elfProgramHeader.HeaderType = elfFile.Data.ReadWord()
	elfProgramHeader.Offset = elfFile.Data.ReadWord()
	elfProgramHeader.VirtualAddress = elfFile.Data.ReadWord()
	elfProgramHeader.PhysicalAddress = elfFile.Data.ReadWord()
	elfProgramHeader.SizeInFile = elfFile.Data.ReadWord()
	elfProgramHeader.SizeInMemory = elfFile.Data.ReadWord()
	elfProgramHeader.Flags = elfFile.Data.ReadWord()
	elfProgramHeader.Alignment = elfFile.Data.ReadWord()

	return elfProgramHeader
}

func (elfProgramHeader *ElfProgramHeader) ReadContent(elfFile *ElfFile) []byte {
	return elfFile.Data.ReadBlockAt(uint64(elfProgramHeader.Offset), uint64(elfProgramHeader.SizeInFile))
}

type ElfStringTable struct {
	Data []byte
}

func NewElfStringTable(elfFile *ElfFile, sectionHeader *ElfSectionHeader) *ElfStringTable {
	if sectionHeader.HeaderType != SHT_STRTAB {
		panic("Section is not a string table")
	}

	var elfStringTable = &ElfStringTable{
	}

	elfStringTable.Data = sectionHeader.ReadContent(elfFile)

	return elfStringTable
}