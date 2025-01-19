package main

import "fmt"

func func1(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return func1(func1(n + 11))
	}
}

func main() {
	fmt.Printf("%d\n", func1(95))
}
