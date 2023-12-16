package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	processForMutex()
}

type Bucket struct {
	Mutex  sync.Mutex
	Tokens []int
	Done   chan int
}

func processForMutex() {
	bucket := &Bucket{
		Mutex:  sync.Mutex{},
		Tokens: nil,
		Done:   make(chan int),
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for i := 0; i < 2; i++ {
			go bucket.producer(i)
			time.Sleep(time.Second * 1)
		}
		fmt.Println("close producer")
		wg.Done()
	}()

	go bucket.consumer(&wg)

	// go func() {
	// 	//time.Sleep(time.Second * 15)
	// 	fmt.Println("close producer")
	// 	wg.Done()
	// }()

	wg.Wait()
	fmt.Println("done")
}

func (b *Bucket) producer(idx int) {
	b.Mutex.Lock()
	defer func() {
		b.Mutex.Unlock()
		b.Done <- idx
	}()

	tokens := []int{rand.Intn(100), rand.Intn(100)}
	fmt.Printf("Put: Index: %d tokens: %v\r\n", idx, tokens)
	b.Tokens = append(b.Tokens, tokens...)
}

func (b *Bucket) consumer(wg *sync.WaitGroup) {
	for {
		fmt.Println("Waiting for producer...")
		select {
		case idx := <-b.Done:
			b.Mutex.Lock()
			fmt.Printf("Get: Index: %d tokens: %v\r\n", idx, b.Tokens)
			b.Tokens = nil
			b.Mutex.Unlock()

		}
	}
}
