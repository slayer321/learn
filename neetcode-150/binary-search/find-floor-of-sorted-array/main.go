package main

import "fmt"

func searchFloorSortedArray(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	lessMin := -1
	for l <= r {

		mid := l + (r-l)/2

		if arr[mid] <= target {
			lessMin = max(lessMin, mid)

		}

		if mid-1 >= l && arr[mid-1] <= target {
			lessMin = max(lessMin, mid-1)

		}

		if mid+1 <= r && arr[mid+1] <= target {
			lessMin = max(lessMin, mid+1)
		}

		if arr[mid] < target {
			l = mid + 1
		} else if arr[mid] > target {
			r = mid - 1
		}

	}

	return lessMin

}

func main() {

	arr := []int{11, 2, 8, 10, 10, 12, 19}
	target := 20

	out := searchFloorSortedArray(arr, target)
	fmt.Printf("out: %v\n", out)

}
