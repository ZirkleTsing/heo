package acogo

type InputVirtualChannel struct {
	InputPort            *InputPort
	Num                  int
	InputBuffer          *InputBuffer
	Route                Direction
	OutputVirtualChannel *OutputVirtualChannel
}

func NewInputVirtualChannel(inputPort *InputPort, num int) *InputVirtualChannel {
	var inputVirtualChannel = &InputVirtualChannel{
		InputPort:inputPort,
		Num:num,
		Route:Direction(-1),
	}

	inputVirtualChannel.InputBuffer = NewInputBuffer(inputVirtualChannel)

	return inputVirtualChannel
}
