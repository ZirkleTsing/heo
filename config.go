package acogo

type Config struct {
	OutputDirectory         string

	NumNodes                int

	MaxCycles               int64

	MaxPackets              int64

	DrainPackets            bool

	RandSeed                int64

	Routing                 string

	Selection               string

	MaxInjectionBufferSize  int

	MaxInputBufferSize      int

	NumVirtualChannels      int

	LinkWidth               int
	LinkDelay               int

	DataPacketTraffic       string
	DataPacketInjectionRate float64
	DataPacketSize          int

	AntPacketTraffic        string
	AntPacketInjectionRate  float64
	AntPacketSize           int

	AcoSelectionAlpha       float64
	ReinforcementFactor     float64
}

func NewConfig(outputDirectory string, numNodes int, maxCycles int64, maxPackets int64, drainPackets bool) *Config {
	var config = &Config{
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

		DataPacketTraffic:"transpose",
		DataPacketInjectionRate:0.01,
		DataPacketSize:16,

		AntPacketTraffic:"uniform",
		AntPacketInjectionRate:0.01,
		AntPacketSize:4,

		AcoSelectionAlpha:0.5,
		ReinforcementFactor:0.05,
	}

	return config
}
