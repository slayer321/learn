package main

import "fmt"

func longestSubarrayWithSumK(arr []int, k int) int {

	i := 0
	j := 0
	longestSubarray := 0
	sum := 0

	for j < len(arr) {
		// newVal := arr[i] + arr[j]
		sum = sum + arr[j]
		for sum > k {
			sum = sum - arr[i]
			i++
		}

		if sum == k {
			size := j - i + 1
			longestSubarray = max(longestSubarray, size)
			// i++
			j++
			// sum = 0
		} else {
			j++
		}
	}

	return longestSubarray

}

func main() {
	arr := []int{10, -10, 20, 30}
	k := 5
	out := longestSubarrayWithSumK(arr, k)
	fmt.Printf("out: %v\n", out)
}
