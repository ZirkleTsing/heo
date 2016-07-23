package acogo

type NoCConfig struct {
	outputDirectory string

	numNodes int

	maxCycles int

	maxPackets int

	noDrain bool

	randSeed int

	routing string

	selection string

	maxInjectionBufferSize int

	maxInputBufferSize int

	numVirtualChannels int

	linkWidth int
	linkDelay int

	antPacketTraffic string
	antPacketSize int
	antPacketInjectionRate float32

	acoSelectionAlpha float32
	reinforcementFactor float32

	dataPacketTraffic string
	dataPacketInjectionRate float32
	dataPacketSize int
}

func NewNoCConfig(outputDirectory string, numNodes int, maxCycles int, maxPackets int, noDrain bool) *NoCConfig {
	var config = &NoCConfig{
		outputDirectory:outputDirectory,

		numNodes:numNodes,

		maxCycles:maxCycles,

		maxPackets:maxPackets,

		noDrain:noDrain,

		randSeed:13,

		routing:"oddEven",

		selection:"aco",

		maxInjectionBufferSize:32,

		maxInputBufferSize:4,

		numVirtualChannels:4,

		linkWidth:4,
		linkDelay:1,

		antPacketTraffic:"uniform",
		antPacketSize:4,
		antPacketInjectionRate:0.01,

		acoSelectionAlpha:0.5,
		reinforcementFactor:0.05,

		dataPacketTraffic:"uniform",
		dataPacketInjectionRate:0.01,
		dataPacketSize:16,
	}

	return config
}
