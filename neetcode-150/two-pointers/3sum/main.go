package main

import (
	"fmt"
	"slices"
)

func twoSum(first, i, j int, nums []int) (bool, int, int) {

	for i < j {

		sum := nums[i] + nums[j]
		realVal := first + sum

		if realVal == 0 {
			return true, i, j
		}

		if sum > first {
			j--
		} else {
			i++
		}
	}
	return false, 0, 0
}

func threeSum(nums []int) [][]int {

	result := [][]int{}

	slices.Sort(nums)

	for i, val := range nums[:len(nums)-2] {
		j := len(nums) - 1
		if i > 0 && val == nums[i-1] {
			continue
		}
		isTwoSum, secondI, thirdI := twoSum(val, i+1, j, nums)
		if isTwoSum {
			newSlice := []int{val, nums[secondI], nums[thirdI]}
			result = append(result, newSlice)
		}
	}

	return result

}

func threeSumWorking(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	slices.Sort(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}

	out := threeSumWorking(arr)
	fmt.Printf("out: %v\n", out)

}
