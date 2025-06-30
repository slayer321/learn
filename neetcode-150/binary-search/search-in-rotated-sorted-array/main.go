package main

import "fmt"

func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	minindex := getMin(nums, l, r)
	fmt.Printf("minindex: %v\n", minindex)
	// fmt.Printf("l: %v\n", l)
	// fmt.Printf("r: %v\n", r)
	ch := make(chan int)
	go solve(nums, l, minindex-1, target, ch)
	go solve(nums, minindex, r, target, ch)

	first := <-ch
	second := <-ch
	if first != -1 {
		return first
	} else if second != -1 {
		return second
	}
	return -1
}

func solve(arr []int, l, r int, target int, ch chan int) {
	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {

			ch <- mid
		}

		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		}
	}
	ch <- -1
}

func getMin(arr []int, l int, r int) int {
	if len(arr) == 1 {
		return 0
	}
	n := len(arr) - 1
	for l <= r {

		mid := l + (r-l)/2
		next := (mid + 1) % n
		prev := (mid + n - 1) % n

		if arr[mid] < arr[next] && arr[mid] < arr[prev] {
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
	arr := []int{1, 3}
	target := 0
	out := search(arr, target)
	fmt.Printf("out: %v\n", out)
}
