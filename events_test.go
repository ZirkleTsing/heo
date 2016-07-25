package acogo

import (
	"testing"
	"fmt"
)

func TestEvents(t *testing.T) {
	var cycleAccurateEventQueue = NewCycleAccurateEventQueue()

	for i := 99; i >= 0; i-- {
		var j = i
		cycleAccurateEventQueue.Schedule(func() {
			fmt.Printf("[%d] Hello world %d.\n", cycleAccurateEventQueue.CurrentCycle, j)
		}, i)
	}

	for i := 0; i < 100; i++ {
		cycleAccurateEventQueue.AdvanceOneCycle()
	}
}