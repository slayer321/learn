package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//workChannel()
	//nilChannel()
	closeChannel()
}

func closeChannel() {

	var c = make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			wg.Done()
		}()

		fmt.Println("-----------")

	}

	wg.Wait()
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

func nilChannel() {
	//var ch chan bool
	gg := make(chan bool, 1)
	//ch <- true
	gg <- true
}

func workChannel() {

	finish := make(chan bool)
	var done sync.WaitGroup
	done.Add(1)

	go func() {
		select {
		case <-time.After(1 * time.Hour):
			fmt.Println("Exit before")
		case <-finish:
		}

		done.Done()
	}()
	time.Sleep(1 * time.Second)
	t0 := time.Now()

	finish <- true
	done.Wait()
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
}
