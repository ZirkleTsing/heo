package acogo

type Context struct {
	Id     int
	Memory *Memory
	Regs   *ArchitecturalRegisterFile
}

func NewContext(id int, littleEndian bool) *Context {
	var context = &Context{
		Id: id,
		Memory:NewMemory(littleEndian),
		Regs:NewArchitecturalRegisterFile(littleEndian),
	}

	return context
}
