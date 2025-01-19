package main

import "fmt"

// For this function time complexility O(2^n)
// space complexitiy O(n)
func treeRecursion(n int) {

	if n > 0 {
		fmt.Printf("%d\n", n)
		treeRecursion(n - 1)
		treeRecursion(n - 1)
	}
}

func main() {
	treeRecursion(3)
}
