package main

import (
	"fmt"

	wppk "github.com/oneprocess/workerPool"
)

func main() {
	// create new tasks
	tasks := make([]wppk.Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = wppk.Task{
			ID: i + 1,
		}
	}

	// Create a worker pool
	wp := wppk.WorkerPool{
		Tasks:       tasks,
		Concurrency: 5, // Number of workers that can run at a time
	}

	// Run the pool
	wp.Run()
	fmt.Println("All tasks have been processed!")
}
