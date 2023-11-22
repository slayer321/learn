package main

import (
	"fmt"
	"sync"
)

func main() {
	//testChannel()
	//mainWork()

	//c := gen(2, 3)
	// out := sq(c)

	// fmt.Println(<-out)
	// fmt.Println(<-out)

	// in := gen(2, 3)
	// c1 := sq(in)
	// c2 := sq(in)

	// // fmt.Println(<-c1)
	// // fmt.Println(<-c2)

	// for n := range merge(c1, c2) {
	// 	fmt.Println(n)
	// }
	// ---------------------
	// c := make(chan int, 1) // Allocate a channel.
	// // Start the sort in a goroutine; when it completes, signal on the channel.
	// go func() {
	// 	//list.Sort()
	// 	fmt.Println("Hello")
	// 	c <- 1 // Send a signal; value does not matter.
	// }()
	// doSomethingForAWhile()
	// fmt.Println(<-c) // Wait for sort to finish; discard sent value.
	// ---------------------
	// q := make(chan int)
	// c := gen1(q)
	// receive1(c, q)

	// fmt.Println("about to exit")
	// ---------------------
	exercise170()
}

func exercise170() {
	c := make(chan int, 100)

	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
	for val := range c {
		fmt.Println(val)
	}
}

func gen1(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	// for i := 0; i < 100; i++ {
	// 	c <- i
	// }
	// close(c)
	return c
}

func receive1(g <-chan int, q <-chan int) {
	// for val := range g {
	// 	//fmt.Println("Inside the receive1 function")
	// 	fmt.Println(val)
	// }
	for {
		select {
		case v := <-g:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func doSomethingForAWhile() {

}

func gen(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func mainWork() {
	even := make(chan int)
	odd := make(chan int)
	//quit := make(chan bool)
	quitInt := make(chan int)

	go send(even, odd, quitInt)
	receive(even, odd, quitInt)
}

func send(even, odd chan<- int, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i

		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd <-chan int, quit <-chan int) {
	for {
		select {
		case V := <-even:
			fmt.Println("From the even channel:", V)
		case V := <-odd:
			fmt.Println("From the odd channel:", V)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma ok:", i)
				return
			} else {
				fmt.Println("From comma ok:", i)
			}
		}
	}
}

func testChannel() {
	c := make(chan int, 1)
	c <- 42
	// go func() {
	// 	c <- 42
	// 	c <- 43
	// 	c <- 44
	// }()
	fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}
