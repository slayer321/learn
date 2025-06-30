package main

import (
	"fmt"
	"time"
)

func play(player string, table chan int) {

	var lastValue int = 1

	for {
		fmt.Println("I'm stuck at the top")
		ball := <-table
		fmt.Println(player, ball)
		ball += lastValue
		if ball < lastValue {
			break
		}

		lastValue = ball
		table <- ball
		time.Sleep(2 * time.Second)
	}
}

func main() {

	table := make(chan int)

	go func() {
		table <- 1
	}()

	go play("A:", table)
	play("B:", table)
}
