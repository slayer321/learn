package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {

	l := 1
	maxVal := -1
	for _, val := range piles {
		maxVal = max(val, maxVal)
	}

	r := maxVal
	res := maxVal

	for l <= r {
		mid := l + (r-l)/2

		hours := 0.0
		for i := 0; i < len(piles); i++ {
			hours += math.Ceil(float64(piles[i]) / float64(mid))
		}

		if int(hours) <= h {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}

	}

	return res

}

func main() {

	arr := []int{3, 6, 7, 11}
	h := 8

	out := minEatingSpeed(arr, h)
	fmt.Printf("out: %v\n", out)

}
