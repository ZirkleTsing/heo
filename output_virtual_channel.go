package acogo

type OutputVirtualChannel struct {
	OutputPort          *OutputPort
	Num                 int
	InputVirtualChannel *InputVirtualChannel
	Credits             int
}

func NewOutputVirtualChannel(outputPort *OutputPort, num int) *OutputVirtualChannel {
	var outputVirtualChannel = &OutputVirtualChannel{
		OutputPort:outputPort,
		Num: num,
		Credits:10,
	}

	return outputVirtualChannel
}
