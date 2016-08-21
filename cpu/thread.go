package cpu

type Thread struct {
	Num             int
	Context         *Context
	Core            *Core
	NumInstructions int64
}

func NewThread(core *Core, num int) *Thread {
	var thread = &Thread{
		Core:core,
		Num:num,
	}

	return thread
}

func (thread *Thread) AdvanceOneCycle() {
	//TODO
}
