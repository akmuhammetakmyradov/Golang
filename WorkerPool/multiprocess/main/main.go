package main

import (
	"fmt"

	wppk "github.com/multiprocess/workerPool"
)

func main() {
	// create new tasks
	tasks := []wppk.Task{
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/sample1.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/sample2.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/sample3.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/sample4.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/sample5.jpg"},
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
