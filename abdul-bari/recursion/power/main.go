package main

import "fmt"

func power(m, n int) int {
	if n == 0 {
		return 1
	}
	return power(m, n-1) * m
}

func Ipower(m, n int) int {
	out := 1
	for n > 0 {

		out *= m
		n--
	}

	return out

}

func main() {

	fmt.Printf("%v\n", power(2, 4))
	fmt.Printf("%v\n", Ipower(2, 4))

}
