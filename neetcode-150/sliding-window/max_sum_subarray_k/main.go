package main

import "fmt"

func maximumSumSubarray(arr []int, k int) int {

	i := 0
	j := 0
	sum := 0
	maxNum := 0
	for j < len(arr) {
		sum = sum + arr[j]
		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == k {
			maxNum = max(maxNum, sum)
			sum = sum - arr[i]
			i++
			j++
		}

	}
	return maxNum
}

func main() {

	arr := []int{100, 200, 300, 400}
	k := 2
	out := maximumSumSubarray(arr, k)
	fmt.Printf("out: %v\n", out)
}
