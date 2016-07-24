package acogo

type InjectionBuffer struct {
	Router  *Router
	packets *Queue
}

func NewInjectionBuffer(router *Router) *InjectionBuffer {
	var injectionBuffer = &InjectionBuffer{
		packets:NewQueue(router.Node.Network.Experiment.Config.MaxInjectionBufferSize),
	}

	return injectionBuffer
}

func (injectionBuffer *InjectionBuffer) Push(packet Packet) {
	injectionBuffer.packets.Push(packet)
}

func (injectionBuffer *InjectionBuffer) Peek() Packet {
	if injectionBuffer.packets.Count > 0 {
		return injectionBuffer.packets.Peek().(Packet)
	} else {
		return nil
	}
}

func (injectionBuffer *InjectionBuffer) Pop() {
	injectionBuffer.packets.Pop()
}

func (injectionBuffer *InjectionBuffer) Full() bool {
	return injectionBuffer.packets.Size <= injectionBuffer.packets.Count
}

func (injectionBuffer *InjectionBuffer) Size() int {
	return injectionBuffer.packets.Size
}

func (injectionBuffer *InjectionBuffer) Count() int {
	return injectionBuffer.packets.Count
}

func (injectionBuffer *InjectionBuffer) FreeSlots() int {
	return injectionBuffer.packets.Size - injectionBuffer.packets.Count
}

type InputBuffer struct {
	InputVirtualChannel *InputVirtualChannel
	flits               *Queue
}

func NewInputBuffer(inputVirtualChannel *InputVirtualChannel) *InputBuffer {
	var inputBuffer = &InputBuffer{
		InputVirtualChannel:inputVirtualChannel,
		flits:NewQueue(inputVirtualChannel.InputPort.Router.Node.Network.Experiment.Config.MaxInputBufferSize),
	}

	return inputBuffer
}

func (inputBuffer *InputBuffer) Push(flit *Flit) {
	inputBuffer.flits.Push(flit)
}

func (inputBuffer *InputBuffer) Peek() *Flit {
	if inputBuffer.flits.Count > 0 {
		return inputBuffer.flits.Peek().(*Flit)
	} else {
		return nil
	}
}

func (inputBuffer *InputBuffer) Pop() {
	inputBuffer.flits.Pop()
}

func (inputBuffer *InputBuffer) Full() bool {
	return inputBuffer.flits.Size <= inputBuffer.flits.Count
}

func (inputBuffer *InputBuffer) Size() int {
	return inputBuffer.flits.Size
}

func (inputBuffer *InputBuffer) Count() int {
	return inputBuffer.flits.Count
}

func (inputBuffer *InputBuffer) FreeSlots() int {
	return inputBuffer.flits.Size - inputBuffer.flits.Count
}
