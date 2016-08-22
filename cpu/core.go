package cpu

type Core struct {
	Processor *Processor
	Threads   []*Thread
	Num       uint32
}

func NewCore(processor *Processor, num uint32) *Core {
	var core = &Core{
		Processor:processor,
		Num:num,
	}

	return core
}

func (core *Core) AdvanceOneCycle() {
	for _, thread := range core.Threads {
		thread.AdvanceOneCycle()
	}
}
