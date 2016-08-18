package acogo

type Kernel struct {
	Pipes         []*Pipe
	SystemEvents  []SystemEvent
	SignalActions []*SignalAction
	Contexts      []*Context
	Processes     []*Process

	CurrentCycle  uint64
}

func NewKernel() *Kernel {
	var kernel = &Kernel{
	}

	return kernel
}

func (kernel *Kernel) MustProcessSignal(context *Context, signal uint32) bool {
	return false //TODO
}