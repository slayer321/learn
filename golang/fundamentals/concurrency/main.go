package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	osArch()
	//mutext()
	// p := person{}
	// saySomething(&p)
	// p1 := person{}
	// saySomething(p1)
}

func osArch() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

type human interface {
	speak()
}

type person struct {
}

func (p *person) speak() {

}

func saySomething(h human) {
	h.speak()
}

func mutext() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0
	//var counter uint64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println("Goroutines Onside :", runtime.NumGoroutine())
			mu.Lock()
			//atomic.AddUint64(&counter, 1)
			runtime.Gosched()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
