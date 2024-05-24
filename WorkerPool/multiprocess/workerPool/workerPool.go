package workerPool

import (
	"fmt"
	"sync"
	"time"
)

// Task defination
type Task interface {
	Process()
}

// Email task defination
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Way to process the Email task
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Image task defination
type ImageProcessingTask struct {
	ImageURL string
}

func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing image %s\n", t.ImageURL)
	// Simulate a time consuming process
	time.Sleep(5 * time.Second)
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
