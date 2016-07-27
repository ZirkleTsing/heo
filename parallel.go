package acogo

import (
	"runtime"
	"fmt"
	"time"
)

func RunInParallel(tasks []func()) {
	var c = make(chan bool, runtime.NumCPU())

	for i, task := range tasks {
		c <- true
		go func(i int, len int, t func(), c chan bool) {
			defer func() {
				<-c
			}()

			fmt.Printf("[%s] Task %d/%d started.\n", time.Now().Format("2006-01-02 15:04:05"), i + 1, len)

			t()

			fmt.Printf("[%s] Task %d/%d ended.\n", time.Now().Format("2006-01-02 15:04:05"), i + 1, len)
		}(i, len(tasks), task, c)
	}

	for i := 0; i < cap(c); i++ {
		c <- true
	}
}
