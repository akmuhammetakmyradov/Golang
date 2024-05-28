package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order %d", i)
			orders <- order
			fmt.Printf("Placed %s\n", order)
		}
		close(orders)
	}()

	for order := range orders {
		fmt.Printf("Prepearing order %s\n", order)
		time.Sleep(2 * time.Second) // get time for preparing coffee
		fmt.Printf("Coffee served %s\n", order)
	}
}
