package main

import (
	"fmt"
)

func main() {
	// err := New("EOF")
	// if err != nil {
	// 	fmt.Println(err.Error())

	// }
	//---------------------------------
	f()
	fmt.Println("Returned normally from f.")

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// func New(text string) error {
// 	return &errorString{text}
// }
