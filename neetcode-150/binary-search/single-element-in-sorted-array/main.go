package main

import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {

		mid := l + (r-l)/2

		if (mid-1 < 0 || nums[mid] != nums[mid-1]) && (mid+1 == len(nums) || nums[mid] != nums[mid+1]) {
			return mid
		}

		leftSize := 0

		if nums[mid-1] == nums[mid] {
			leftSize = mid - 1
		} else {
			leftSize = mid
		}

		if leftSize%2 == 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}

		// if mid-1 <= l && mid+1 <= r && nums[mid] != nums[mid-1] {
		// 	if nums[mid] != nums[mid+1] {
		// 		return nums[mid]
		// 	}
		// }

		// even := false
		// odd := false
		// if mid%2 == 0 {
		// 	even = true
		// } else {
		// 	odd = true
		// }

		// if even && nums[mid-1] == nums[mid] {
		// 	r = mid - 1
		// } else if odd && nums[mid+1] == nums[mid] {
		// 	l = mid + 1
		// }
	}

	return -1

}

func main() {

	nums := []int{3, 3, 7, 7, 10, 11, 11}
	out := singleNonDuplicate(nums)
	fmt.Printf("out: %v\n", out)

}
