package cpu

type Core struct {
	Processor *Processor
	Threads   []*Thread
	Num       int32
}

func NewCore(processor *Processor, num int32) *Core {
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
