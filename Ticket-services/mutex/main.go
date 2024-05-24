package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicket(wg *sync.WaitGroup, userID int, ticketCounts *int) {
	defer wg.Done()
	mutex.Lock()
	if *ticketCounts > 0 {
		*ticketCounts--
		fmt.Printf("User %d purchased a ticket. Tickets remaining %d \n", userID, *ticketCounts)
	} else {
		fmt.Printf("User %d found no tickets.\n", userID)
	}
	mutex.Unlock()
}

func main() {
	var tickets int = 500

	var wg sync.WaitGroup

	for userID := 0; userID < 1000; userID++ {
		wg.Add(1)
 
		go buyTicket(&wg, userID, &tickets)
	}

	wg.Wait()
}

// you can test with the following commands
// go run main.go | grep purchased | wc -l
// for run in {1..10}; do go run main.go | grep purchased | wc -l; done
