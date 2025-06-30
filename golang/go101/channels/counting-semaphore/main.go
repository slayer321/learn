package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {

	log.Println("cusomer#", c, "enter the bar")
	seat := <-bar
	log.Println("++ customer#", c, "drinks at seat", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Println("-- customer#", c, "frees seat#", seat)
	bar <- seat

}

func main() {

	bar24x7 := make(Bar, 10)

	for SeatID := 0; SeatID < cap(bar24x7); SeatID++ {
		bar24x7 <- Seat(SeatID)
	}

	for customerID := 0; customerID < 20; customerID++ {
		go bar24x7.ServeCustomer(customerID)
	}

	time.Sleep(2 * time.Second)

	// for {
	// 	time.Sleep(time.Second)
	// }

}
