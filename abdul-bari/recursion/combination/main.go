package main

import "fmt"

// Combination formula
// nCr = n!/r!(n-r)!

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n-1) * n
}

func nCr(n, r int) int {
	num := 0
	dum := 0

	num = fact(n)
	dum = fact(r) * fact(n-r)

	return num / dum

}

func comb(n, r int) int {

	if n == r || r == 0 {
		return 1
	}

	return comb(n-1, r-1) + comb(n-1, r)

}

func main() {

	fmt.Printf("%d\n", comb(4, 2))
	fmt.Printf("%d\n", nCr(4, 2))

}
