package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string) // unbuffered channel

	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			orders <- order // Blocks until barista is ready to accept
			fmt.Println("Placed: ", order)
		}
		close(orders)
	}()

	for order := range orders {
		fmt.Printf("Preparing order: %s\n", order)
		time.Sleep(2 * time.Second) // Time taken for preparing order
		fmt.Printf("Coffee served: %s\n", order)
	}
}
