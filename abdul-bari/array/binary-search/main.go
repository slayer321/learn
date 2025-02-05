package main

import "fmt"

func binarySearch(arr []int, target int) int {

	l := 0
	h := len(arr) - 1
	return solve(l, h, arr, target)

	// mid := len(arr) / 2

	// if target == arr[mid] {
	// 	fmt.Printf("found target %d at index %d\n", target, mid)
	// }

	// if target < arr[mid] {
	// 	binarySearch(arr[:mid], target)
	// } else if target > arr[mid] {
	// 	binarySearch(arr[mid+1:], target)
	// }
}

func solve(l, h int, arr []int, target int) int {

	if l <= h {

		mid := (l + h) / 2

		if arr[mid] == target {
			return mid
		}

		if target < arr[mid] {
			return solve(l, mid-1, arr, target)
		} else if target > arr[mid] {
			return solve(mid+1, h, arr, target)
		}

	}

	return -1
}

func binarySearchLoop(arr []int, target int) int {

	l := 0
	h := len(arr) - 1

	for l <= h {

		mid := (l + h) / 2

		if target == arr[mid] {
			return mid
		}

		if target < arr[mid] {
			h = mid - 1
		} else if target > arr[mid] {
			l = mid + 1
		}
	}

	return -1
}

func main() {

	arr := []int{4, 6, 9, 10, 32, 77}

	// target := binarySearch(arr, 77)
	// fmt.Printf("target: %v\n", target)

	out := binarySearchLoop(arr, 7)
	fmt.Printf("out: %v\n", out)
}
