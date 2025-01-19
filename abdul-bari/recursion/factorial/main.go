package main

import "fmt"

func factorial(n int) int {

	if n <= 1 {
		return n
	}

	return factorial(n-1) * n

}

func Ifactorial(n int) int {
	out := 1
	for n >= 1 {

		out *= n
		n--
	}
	return out

}

func main() {

	fmt.Printf("%d\n", factorial(5))
	fmt.Printf("%d\n", Ifactorial(5))
}
