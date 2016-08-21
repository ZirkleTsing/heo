package acogo

import "syscall"

type Process struct {
	Id                   uint32
	ContextMapping       *ContextMapping
	Environments         []string
	StdInFileDescriptor  uint32
	StdOutFileDescriptor uint32
	StackBase            uint32
	StackSize            uint32
	TextSize             uint32
	EnvironmentBase      uint32
	HeapTop              uint32
	DataTop              uint32
	ProgramEntry         uint32
	LittleEndian         bool
	Memory               *PagedMemory
}

func NewProcess(kernel *Kernel, contextMapping *ContextMapping) *Process {
	var process = &Process{
		ContextMapping:contextMapping,

	}

	//TODO

	return process
}

func (process *Process) LoadProgram(kernel *Kernel, contextMapping *ContextMapping) {
	//TODO
}

func (process *Process) TranslateFileDescriptor(fileDescriptor uint32) uint32 {
	if fileDescriptor == 1 || fileDescriptor == 2 {
		return process.StdOutFileDescriptor
	} else if fileDescriptor == 0 {
		return process.StdInFileDescriptor
	} else {
		return fileDescriptor
	}
}

func (process *Process) CloseProgram() {
	if process.StdInFileDescriptor != 0 {
		syscall.Close(int(process.StdInFileDescriptor))
	}

	if process.StdOutFileDescriptor > 2 {
		syscall.Close(int(process.StdOutFileDescriptor))
	}
}

func (process *Process) Decode(machInst MachInst) *StaticInst {
	for _, mnemonic := range Mnemonics {
		if machInst != 0 && mnemonic.Mask == mnemonic.Bits && (mnemonic.ExtraBitField == nil || machInst.ValueOf(mnemonic.ExtraBitField) == mnemonic.ExtraBitFieldValue) {
			return NewStaticInst(mnemonic, machInst)
		}
	}

	panic("Impossible")
}