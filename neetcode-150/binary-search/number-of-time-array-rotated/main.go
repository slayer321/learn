package main

import "fmt"

func numberOfTimeRotated(arr []int) int {
	l := 0
	r := len(arr) - 1
	return solve(arr, l, r)
}

func solve(arr []int, l, r int) int {
	n := len(arr) - 1
	for l <= r {
		mid := l + (r-l)/2

		next := (mid + 1) % n
		prev := (mid + n - 1) % n
		// if mid == 0 || mid == len(arr)-1 {
		// 	return mid
		if arr[mid] < arr[prev] && arr[mid] < arr[next] {
			return mid
		}

		if arr[mid] > arr[r] {
			l = mid + 1
		} else if arr[mid] < arr[r] {
			r = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{15, 7, 9, 11, 12}

	out := numberOfTimeRotated(arr)
	fmt.Printf("out: %v\n", out)
}
