package acogo

type NoCConfig struct {
	OutputDirectory         string

	NumNodes                int

	MaxCycles               int64

	MaxPackets              int64

	DrainPackets            bool

	RandSeed                int64

	Routing                 string //TODO

	Selection               string //TODO

	MaxInjectionBufferSize  int

	MaxInputBufferSize      int

	NumVirtualChannels      int

	LinkWidth               int
	LinkDelay               int

	DataPacketTraffic       string //TODO
	DataPacketInjectionRate float64
	DataPacketSize          int

	AntPacketTraffic        string //TODO
	AntPacketInjectionRate  float64 //TODO
	AntPacketSize           int

	AcoSelectionAlpha       float64
	ReinforcementFactor     float64
}

func NewNoCConfig(outputDirectory string, numNodes int, maxCycles int64, maxPackets int64, drainPackets bool) *NoCConfig {
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
