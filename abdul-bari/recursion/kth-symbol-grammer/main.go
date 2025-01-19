package main

import (
	"fmt"
	"math"
)

func kthGrammer(n, k int) int {

	if n == 1 && k == 1 {
		return 0
	}

	mid := int(math.Pow(2, float64(n-1))) / 2

	if k <= mid {
		return kthGrammer(n-1, k)
	} else {
		return 1 - kthGrammer(n-1, k-mid)
	}

}

func main() {

	fmt.Printf("%v\n", kthGrammer(1, 1))

}
