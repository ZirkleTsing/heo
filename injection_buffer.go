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
	if injectionBuffer.Full() {
		panic("Injection buffer is full")
	}

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
