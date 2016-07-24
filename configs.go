package acogo

type NoCConfig struct {
	OutputDirectory         string

	NumNodes                int

	MaxCycles               int

	MaxPackets              int

	NoDrain                 bool

	RandSeed                int

	Routing                 string

	Selection               string

	MaxInjectionBufferSize  int

	MaxInputBufferSize      int

	NumVirtualChannels      int

	LinkWidth               int
	LinkDelay               int

	AntPacketTraffic        string
	AntPacketSize           int
	AntPacketInjectionRate  float32

	AcoSelectionAlpha       float32
	ReinforcementFactor     float32

	DataPacketTraffic       string
	DataPacketInjectionRate float32
	DataPacketSize          int
}

func NewNoCConfig(outputDirectory string, numNodes int, maxCycles int, maxPackets int, noDrain bool) *NoCConfig {
	var config = &NoCConfig{
		OutputDirectory:outputDirectory,

		NumNodes:numNodes,

		MaxCycles:maxCycles,

		MaxPackets:maxPackets,

		NoDrain:noDrain,

		RandSeed:13,

		Routing:"oddEven",

		Selection:"aco",

		MaxInjectionBufferSize:32,

		MaxInputBufferSize:4,

		NumVirtualChannels:4,

		LinkWidth:4,
		LinkDelay:1,

		AntPacketTraffic:"uniform",
		AntPacketSize:4,
		AntPacketInjectionRate:0.01,

		AcoSelectionAlpha:0.5,
		ReinforcementFactor:0.05,

		DataPacketTraffic:"uniform",
		DataPacketInjectionRate:0.01,
		DataPacketSize:16,
	}

	return config
}
