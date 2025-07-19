package main

import "fmt"

func findPages(arr []int, k int) int {

	if k > len(arr) {
		return -1
	}

	sum := 0
	maxValInArray := -1
	res := -1
	for _, val := range arr {
		sum += val
		maxValInArray = max(maxValInArray, val)
	}

	if k == 1 {
		return sum
	}

	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("maxValInArray: %v\n", maxValInArray)

	start := maxValInArray
	end := sum

	for start <= end {
		mid := start + (end-start)/2

		if isValid(arr, len(arr), k, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return res

}

func isValid(arr []int, n, k, maxVal int) bool {
	student := 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum > maxVal {
			student++
			sum = arr[i]
		}

		if student > k {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{12, 34, 67, 90}
	k := 2

	out := findPages(arr, k)
	fmt.Printf("out: %v\n", out)
}
