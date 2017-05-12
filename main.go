package main

import (
	"fmt"
	"github.com/mcai/acogo/noc"
	"github.com/mcai/acogo/simutil"
)

var (
	numNodes int = 8 * 8
	maxCycles int64 = int64(100000000)
	maxPackets int64 = int64(-1)
	drainPackets = false
)

func NewTraceDrivenExperiment(outputDirectoryPrefix string, traffic noc.TrafficType, dataPacketInjectionRate float64, routing noc.RoutingType, selection noc.SelectionType, antPacketInjectionRate float64, acoSelectionAlpha float64, reinforcementFactor float64, traceFileNames string) simutil.Experiment {
	var outputDirectory string

	switch {
	case selection == noc.SELECTION_ACO:
		outputDirectory = fmt.Sprintf("trace_driven_results/%s/t_%s/j_%f/r_%s/s_%s/aj_%f/a_%f/rf_%f/",
			outputDirectoryPrefix, traffic, dataPacketInjectionRate, routing, selection, antPacketInjectionRate, acoSelectionAlpha, reinforcementFactor)
	default:
		outputDirectory = fmt.Sprintf("trace_driven_results/%s/t_%s/j_%f/r_%s/s_%s/",
			outputDirectoryPrefix, traffic, dataPacketInjectionRate, routing, selection)
	}

	var config = noc.NewNoCConfig(
		outputDirectory,
		numNodes,
		maxCycles,
		maxPackets,
		drainPackets)

	config.DataPacketTraffic = traffic
	config.DataPacketInjectionRate = dataPacketInjectionRate
	config.Routing = routing
	config.Selection = selection

	if selection == noc.SELECTION_ACO {
		config.AntPacketInjectionRate = antPacketInjectionRate
		config.AcoSelectionAlpha = acoSelectionAlpha
		config.ReinforcementFactor = reinforcementFactor
	}

	config.TraceFileName = traceFileNames

	return noc.NewNoCExperiment(config)
}

func main() {
	var dataPacketInjectionRate = 0.015
	var antPacketInjectionRate = 0.0002

	var acoSelectionAlpha = 0.45
	var reinforcementFactor = 0.001

	var outputDirectoryPrefix = "trafficsAndDataPacketInjectionRates"

	var traceFileNames []string

	traceFileNames = append(traceFileNames, "/home/poo/Downloads/pin-3.2-81205-gcc-linux/source/tools/MemTrace/traces/blackscholes.trace.13996.txt")

	var experiment = NewTraceDrivenExperiment(
		outputDirectoryPrefix,
		noc.TRAFFIC_TRACE,
		dataPacketInjectionRate,
		noc.ROUTING_ODD_EVEN, noc.SELECTION_BUFFER_LEVEL,
		antPacketInjectionRate,
		acoSelectionAlpha,
		reinforcementFactor,
		traceFileNames,
	)

	var experiments []simutil.Experiment

	experiments = append(experiments, experiment)

	simutil.RunExperiments(experiments, false)
}
