package main

import (
	"fmt"

	wppk "github.com/multiprocess/workerPool"
)

func main() {
	// create new tasks
	tasks := []wppk.Task{
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/surat1.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/surat2.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/surat3.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/surat4.jpg"},
		&wppk.EmailTask{Email: "akmyradowakmuhammet21@gmail.com", Subject: "test", MessageBody: "test"},
		&wppk.ImageProcessingTask{ImageURL: "/images/surat5.jpg"},
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
