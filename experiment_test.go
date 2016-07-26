package acogo

import "testing"

func TestExperiment(t *testing.T) {
	var numNodes = 64
	var maxCycles = int64(2000)
	var maxPackets = int64(-1)
	var drainPackets = true

	var config = NewConfig("test_results/synthetic/aco", numNodes, maxCycles, maxPackets, drainPackets)

	config.Routing = ROUTING_ODD_EVEN
	config.Selection = SELECTION_ACO

	config.DataPacketTraffic = TRAFFIC_TRANSPOSE
	config.DataPacketInjectionRate = 0.06

	config.AntPacketTraffic = TRAFFIC_UNIFORM
	config.AntPacketInjectionRate = 0.0002

	config.AcoSelectionAlpha = 0.45
	config.ReinforcementFactor = 0.001

	var experiment = NewExperiment(config)

	experiment.Run()
}
