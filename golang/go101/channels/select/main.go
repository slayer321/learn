package main

import (
	"fmt"
	"time"
)

func someWork() {

	type Book struct{ id int }
	bookCh := make(chan Book, 3)

	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < cap(bookCh)*2; i++ {
		select {
		case bookCh <- Book{i}:
			fmt.Printf("book send %d\n", i)
			time.Sleep(time.Second)
		case <-ticker.C:
			close(bookCh)
		case <-bookCh:
			fmt.Println("returning because the channel is closed")
			return
		default:
			fmt.Println("failed to send book")

		}
	}

	for i := 0; i < cap(bookCh)*2; i++ {
		select {
		case val := <-bookCh:
			fmt.Printf("book received %d\n", val.id)
		default:
			fmt.Println("failed to receive book")

		}
	}

}

func main() {

	someWork()

}
