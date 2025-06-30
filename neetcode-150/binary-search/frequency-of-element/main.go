package main

import "fmt"

func firstOcc(nums []int, target int) int {
	res := 0
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			r = mid - 1
			res++
		}

		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		}
	}
	return res
}

func secondOcc(nums []int, target int) int {
	res := 0
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			l = mid + 1
			res++
		}

		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		}
	}
	return res
}

func numberofOccurance(nums []int, target int) int {

	first := firstOcc(nums, target)
	second := secondOcc(nums, target)
	fmt.Printf("first: %v\n", first)
	fmt.Printf("second: %v\n", second)
	result := first + second
	return result
}

func main() {
	arr := []int{1, 1, 2, 2, 2, 2, 3}
	target := 4

	out := numberofOccurance(arr, target)
	fmt.Printf("out: %v\n", out)

}
