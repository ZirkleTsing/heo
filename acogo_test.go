package acogo

import (
	"fmt"
	"testing"
)

func TestCycleAccurateEventQueue(t *testing.T) {
	var numNodes = 64
	var maxCycles = 20000
	var maxPackets = -1
	var noDrain = false

	var config = NewNoCConfig("test_results/synthetic/aco", numNodes, maxCycles, maxPackets, noDrain)

	config.routing = "oddEven"
	config.selection = "aco"

	config.dataPacketTraffic = "transpose"
	config.dataPacketInjectionRate = 0.06

	config.antPacketTraffic = "uniform"
	config.antPacketInjectionRate = 0.0002
	config.acoSelectionAlpha = 0.45
	config.reinforcementFactor = 0.001

	var experiment = NewNoCExperiment(config)

	experiment.cycleAccurateEventQueue.Schedule(func() {
		fmt.Printf("[%d] Welcome to ACOGo!\n", experiment.cycleAccurateEventQueue.currentCycle)
	}, 4)

	experiment.run()

	fmt.Printf("[%d] Simulation ended!\n", experiment.cycleAccurateEventQueue.currentCycle)
}