package acogo

import (
	"reflect"
	"fmt"
)

type TrafficType string

const (
	TRAFFIC_TRANSPOSE = TrafficType("Transpose")
	TRAFFIC_UNIFORM = TrafficType("Uniform")
)

type RoutingType string

const (
	ROUTING_XY = RoutingType("XY")
	ROUTING_ODD_EVEN = RoutingType("Odd Even")
)

type SelectionType string

const (
	SELECTION_BUFFER_LEVEL = SelectionType("Buffer Level")
	SELECTION_ACO = SelectionType("ACO")
)

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
	s := reflect.ValueOf(experiment.Config).Elem()
	typeOfT := s.Type()

	fmt.Println("Config:")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("  %s: %v\n", typeOfT.Field(i).Name, f.Interface())
	}
}
