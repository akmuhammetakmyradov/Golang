package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxRequests     = 5
	rateLimitPeriod = time.Second
)

type RateLimiter struct {
	mu         sync.Mutex
	requests   map[string]int
	timestamps map[string]time.Time
	limit      int
	interval   time.Duration
}

func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		requests:   make(map[string]int),
		timestamps: make(map[string]time.Time),
		limit:      maxRequests,
		interval:   rateLimitPeriod,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(rl.interval)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for key, ts := range rl.timestamps {
			if now.Sub(ts) > rl.interval {
				delete(rl.requests, key)
				delete(rl.timestamps, key)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Allow(deviceID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if ts, found := rl.timestamps[deviceID]; found {
		if now.Sub(ts) > rl.interval {
			rl.requests[deviceID] = 0
			rl.timestamps[deviceID] = now
		}
	} else {
		rl.timestamps[deviceID] = now
	}

	if rl.requests[deviceID] < rl.limit {
		rl.requests[deviceID]++
		return true
	}

	return false
}

func main() {
	// Example usage
	rl := NewRateLimiter()
	deviceID := "device1"
	deviceID2 := "device2"
	deviceID3 := "device3"

	for i := 0; i < 30; i++ {
		if rl.Allow(deviceID) && rl.Allow(deviceID2) && rl.Allow(deviceID3) {
			println("Request allowed")
		} else {
			println("Request denied")
			fmt.Println(i + 1)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
