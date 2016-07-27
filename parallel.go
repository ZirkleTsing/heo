package acogo

import (
	"fmt"
	"time"
)

func RunInParallel(tasks []func()) {
	var done = make(chan bool)

	for i, task := range tasks {
		go func(i int, len int, t func(), c chan bool) {
			fmt.Printf("[%s] Task %d/%d started.\n",
				time.Now().Format("2006-01-02 15:04:05"), i + 1, len)

			t()

			done <- true

			fmt.Printf("[%s] Task %d/%d ended.\n",
				time.Now().Format("2006-01-02 15:04:05"), i + 1, len)
		}(i, len(tasks), task, done)
	}

	for i := 0; i < len(tasks); i++ {
		<-done

		fmt.Printf("[%s] There are %d tasks to be run.\n",
			time.Now().Format("2006-01-02 15:04:05"), len(tasks) - i - 1)
	}
}