package main

import (
	"fmt"
	"slices"
)

// Time complexitity O(2^n)
func fact(n int) int {
	if n <= 1 {
		return n
	}

	return fact(n-2) + fact(n-1)
}

// Time complexitity O(n)
func Ifact(n int) int {
	i, j := 0, 1
	for n >= 1 {

		i, j = j, i+j
		n--
	}

	return i

}

var arr = slices.Repeat([]int{-1}, 6)

// Time complexitity O(n)
func MemFact(n int) int {
	if n <= 1 {
		arr[n] = n
		return n
	} else {
		if arr[n-2] == -1 {
			arr[n-2] = MemFact(n - 2)
		}
		if arr[n-1] == -1 {
			arr[n-1] = MemFact(n - 1)
		}
		arr[n] = arr[n-2] + arr[n-1]
		fmt.Printf("arr: %v\n", arr)
		return arr[n-2] + arr[n-1]

	}

}

func main() {

	fmt.Printf("%d\n", fact(6))
	fmt.Printf("%d\n", Ifact(6))
	fmt.Printf("%d\n", MemFact(5))

}
