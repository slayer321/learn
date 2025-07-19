package main

import "fmt"

func searchInSortedArray(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {

		mid := l + (r-l)/2

		if arr[mid] == target || (mid-1 >= l && arr[mid-1] == target) || (mid+1 <= r && arr[mid+1] == target) {
			return mid
		}

		if arr[mid] < target {
			l = mid + 2
		} else if arr[mid] > target {
			r = mid - 2
		}

	}

	return -1

}

func main() {

	arr := []int{10, 3, 40, 20, 50, 80, 70}
	target := 90

	out := searchInSortedArray(arr, target)
	fmt.Printf("out: %v\n", out)

}
