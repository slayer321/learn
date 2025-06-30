package main

import (
	"fmt"
	"time"
)

var (
	rateLimitTime    = 30 * time.Second
	rateLimitRequest = 10
)

func handle(request int) {
	fmt.Printf("request: %v\n", request)

}
func handleRateLimit(request chan int) {

	startTime := time.Now()
	fmt.Printf("startTime: %v\n", startTime)

	ticker := time.NewTicker(rateLimitTime)
	quota := make(chan int, rateLimitRequest)

	for {
		select {
		case <-ticker.C:
			afterStartTime := time.Since(startTime)
			fmt.Printf("afterStartTime: %v\n", afterStartTime)
			for i := 0; i < rateLimitRequest; i++ {
				<-quota
			}

		default:
			if len(quota) != rateLimitRequest {
				val := <-request
				quota <- val
				go handle(val)

			}

		}
	}

}

func main() {

	requestCh := make(chan int)

	go handleRateLimit(requestCh)

	for i := 1; ; i++ {
		time.Sleep(100 * time.Millisecond)
		requestCh <- i
	}

}
