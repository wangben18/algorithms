package main

import (
	"fmt"
	"time"
)

func main() {
	requests := 5
	burstRequests := make(chan int, requests)
	for i := 1; i <= requests; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	burst := 3
	burstLimiter := make(chan time.Time, burst)
	for i := 1; i <= burst; i++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	for request := range burstRequests {
		<-burstLimiter
		fmt.Println("request", request, time.Now())
	}
}
