package cpu

type RegisterRenameTable struct {
	Name string
	Entries map[uint32]*PhysicalRegister
}

func NewRegisterRenameTable(name string) *RegisterRenameTable {
	var registerRenameTable = &RegisterRenameTable{
		Name:name,
	}

	return registerRenameTable
}
