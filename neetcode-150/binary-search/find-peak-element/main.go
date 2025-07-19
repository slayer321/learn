package main

import "fmt"

func findPeakElement(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	l := 0
	r := len(arr) - 1

	for l <= r {

		mid := l + (r-l)/2

		if (mid > 0 && mid < len(arr)-1) && (arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1]) {
			return mid
		}

		if mid == 0 && arr[mid] > arr[mid+1] {
			return mid
		}
		if mid == len(arr)-1 && arr[mid] > arr[mid-1] {
			return mid
		}

		if arr[mid+1] > arr[mid] {
			l = mid + 1
		} else if arr[mid-1] > arr[mid] {
			r = mid - 1
		}

	}

	return -1

}
func main() {
	arr := []int{1, 2, 1, 3, 5, 6, 4}

	out := findPeakElement(arr)
	fmt.Printf("out: %v\n", out)
}
