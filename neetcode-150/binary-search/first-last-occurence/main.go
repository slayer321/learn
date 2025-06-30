package main

import "fmt"

func firstOcc(nums []int, target int) int {
	res := -1
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			r = mid - 1
			res = mid
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
	res := -1
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			l = mid + 1
			res = mid
		}

		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		}
	}
	return res
}

func searchRange(nums []int, target int) []int {
	result := []int{}
	first := firstOcc(nums, target)
	second := secondOcc(nums, target)
	result = append(result, first, second)
	return result
}

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	target := 8

	out := searchRange(arr, target)
	fmt.Printf("out: %v\n", out)

}
