package acogo

const CONFIG_JSON_FILE_NAME = "config.json"

type TrafficType string

const (
	TRAFFIC_UNIFORM = TrafficType("Uniform")
	TRAFFIC_TRANSPOSE = TrafficType("Transpose")
)

var TRAFFICS = []TrafficType{
	TRAFFIC_UNIFORM,
	TRAFFIC_TRANSPOSE,
}

type RoutingType string

const (
	ROUTING_XY = RoutingType("XY")
	ROUTING_ODD_EVEN = RoutingType("OddEven")
)

var ROUTINGS = []RoutingType{
	ROUTING_XY,
	ROUTING_ODD_EVEN,
}

type SelectionType string

const (
	SELECTION_BUFFER_LEVEL = SelectionType("BufferLevel")
	SELECTION_ACO = SelectionType("ACO")
)

var SELECTIONS = []SelectionType{
	SELECTION_BUFFER_LEVEL,
	SELECTION_ACO,
}

type Config struct {
	OutputDirectory         string

	NumNodes                int

	MaxCycles               int64

	MaxPackets              int64

	DrainPackets            bool

	Routing                 RoutingType

	Selection               SelectionType

	MaxInjectionBufferSize  int

	MaxInputBufferSize      int

	NumVirtualChannels      int

	LinkWidth               int
	LinkDelay               int

	DataPacketTraffic       TrafficType
	DataPacketInjectionRate float64
	DataPacketSize          int

	AntPacketTraffic        TrafficType
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

		Routing:ROUTING_ODD_EVEN,

		Selection:SELECTION_ACO,

		MaxInjectionBufferSize:32,

		MaxInputBufferSize:4,

		NumVirtualChannels:4,

		LinkWidth:4,
		LinkDelay:1,

		DataPacketTraffic:TRAFFIC_TRANSPOSE,
		DataPacketInjectionRate:0.01,
		DataPacketSize:16,

		AntPacketTraffic:TRAFFIC_UNIFORM,
		AntPacketInjectionRate:0.01,
		AntPacketSize:4,

		AcoSelectionAlpha:0.5,
		ReinforcementFactor:0.05,
	}

	return config
}

func (experiment *Experiment) DumpConfig() {
	WriteJsonFile(experiment.Config, experiment.Config.OutputDirectory, CONFIG_JSON_FILE_NAME)
}
