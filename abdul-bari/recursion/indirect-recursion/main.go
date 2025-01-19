package main

import (
	"fmt"
)

func funcA(n int) {
	if n > 0 {
		fmt.Printf("%d\n", n)
		funcB(n - 1)
	}
}

func funcB(n int) {
	if n > 1 {
		fmt.Printf("%d\n", n)
		funcA(n / 2)
	}
}
func main() {
	funcA(10)
}
