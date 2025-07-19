package main

//https://www.geeksforgeeks.org/dsa/find-position-element-sorted-array-infinite-numbers/
import "fmt"

func infiniteSortedArray(arr []int, target int) int {
	startL := 0
	startR := 1

	start, end := getRange(startL, startR, target, arr)
	fmt.Printf("start: %v\n", start)
	fmt.Printf("end: %v\n", end)
	return solve(start, end, target, arr)
}

func getRange(startL, startR int, target int, arr []int) (int, int) {

	start := startL
	end := startR
	for arr[end] < target {
		if arr[end] < target {
			newEnd := end * 2
			start = end
			end = newEnd
		}

	}

	return start, end

}

func solve(l, r int, target int, arr []int) int {

	if l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			return solve(mid+1, r, target, arr)
		} else if target < arr[mid] {
			return solve(l, mid-1, target, arr)
		}
	}

	return -1
}

func main() {
	arr := []int{3, 5, 7, 9, 10, 90, 100, 130, 140, 160, 170}
	target := 170
	out := infiniteSortedArray(arr, target)
	fmt.Printf("out: %v\n", out)
}
