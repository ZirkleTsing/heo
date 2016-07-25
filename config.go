package acogo

type NoCConfig struct {
	OutputDirectory         string

	NumNodes                int

	MaxCycles               int

	MaxPackets              int

	DrainPackets            bool

	RandSeed                int64

	Routing                 string

	Selection               string

	MaxInjectionBufferSize  int

	MaxInputBufferSize      int

	NumVirtualChannels      int

	LinkWidth               int
	LinkDelay               int

	AntPacketTraffic        string
	AntPacketSize           int
	AntPacketInjectionRate  float64

	AcoSelectionAlpha       float64
	ReinforcementFactor     float64

	DataPacketTraffic       string
	DataPacketInjectionRate float64
	DataPacketSize          int
}

func NewNoCConfig(outputDirectory string, numNodes int, maxCycles int, maxPackets int, drainPackets bool) *NoCConfig {
	var config = &NoCConfig{
		OutputDirectory:outputDirectory,

		NumNodes:numNodes,

		MaxCycles:maxCycles,

		MaxPackets:maxPackets,

		DrainPackets:drainPackets,

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
