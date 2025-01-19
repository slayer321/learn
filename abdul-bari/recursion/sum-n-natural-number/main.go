package main

import "fmt"

func work(n int) int {
	if n == 0 {
		return 0
	}

	return work(n-1) + n
}

func usingLoops(n int) int {
	sum := 0
	for n >= 0 {
		sum += n
		n--
	}

	return sum

}

func main() {

	fmt.Printf("%v\n", work(5))
	fmt.Printf("%v\n", usingLoops(5))

}
