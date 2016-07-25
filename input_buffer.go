package acogo

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
	if inputBuffer.Full() {
		panic("Input buffer is full")
	}

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
