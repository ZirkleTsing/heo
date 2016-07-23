package acogo

import (
	"fmt"
	"testing"
)

func TestCycleAccurateEventQueue(t *testing.T) {
	var cycleAccurateEventQueue = NewCycleAccurateEventQueue()

	var _ = NewNetwork(16, cycleAccurateEventQueue)

	cycleAccurateEventQueue.Schedule(func() {
		fmt.Printf("[%d] Welcome to ACOGo, haha!\n", cycleAccurateEventQueue.currentCycle)
	}, 2)

	for i := 0; i < 10; i++ {
		cycleAccurateEventQueue.AdvanceOneCycle()
	}
}