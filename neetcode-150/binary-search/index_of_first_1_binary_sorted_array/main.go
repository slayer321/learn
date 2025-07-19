package main

import "fmt"

func firstOneIndex(arr []int) int {

	l := 0
	r := 1

	for arr[r] == 0 {
		l = r
		r = r * 2
	}

	return solve(arr, l, r)
}

func solve(arr []int, l, r int) int {
	res := -1
	for l <= r {
		mid := l + (r-l)/2

		if (l >= 0 || arr[mid-1] != 1) && arr[mid] == 1 {
			res = mid
			r = mid - 1
		}

		if arr[mid+1] == 1 {
			l = mid + 1
		}

	}

	return res

}

func main() {

	arr := []int{1, 1, 1, 1, 1, 1}
	out := firstOneIndex(arr)
	fmt.Printf("out: %v\n", out)

}
