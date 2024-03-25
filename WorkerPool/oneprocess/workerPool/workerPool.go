package workerPool

import (
	"fmt"
	"sync"
	"time"
)

// Task defination
type Task struct {
	ID int
}

// Way to process the tasks
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	TasksChan   chan Task
	Wg          sync.WaitGroup
}

// Functions to execute the worker pool

func (wp *WorkerPool) worker() {
	for task := range wp.TasksChan {
		task.Process()
		wp.Wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.TasksChan = make(chan Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.Concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks channel
	wp.Wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.TasksChan <- task
	}
	close(wp.TasksChan)

	// Wait for all tasks to finish
	wp.Wg.Wait()
}
