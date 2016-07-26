package acogo

import (
	"runtime"
	"fmt"
	"time"
)

func RunInParallel(tasks []func()) {
	var c = make(chan bool, runtime.NumCPU() - 1)

	for i, task := range tasks {
		var j = i

		go func(t func()) {
			defer func() {
				<-c
			}()

			fmt.Printf("[%s] Task %d/%d started.\n", time.Now().Format("2006-01-02 15:04:05"), j+1, len(tasks))

			t()

			fmt.Printf("[%s] Task %d/%d ended.\n", time.Now().Format("2006-01-02 15:04:05"), j+1, len(tasks))
		}(task)

		c <- true
	}
}
