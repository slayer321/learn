package main

import (
	"fmt"
	"slices"
)

func missingNumber(arr []int) int {

	slices.Sort(arr)

	i := 0
	for ; i < len(arr); i++ {
		if arr[i] != i {
			return i
		}
	}
	return i

}

func missingNumberEfficient(arr []int) int {
	n := len(arr)
	wholeNumber := n * (n + 1) / 2
	sum := 0

	for _, val := range arr {
		sum += val
	}

	result := wholeNumber - sum
	return result
}

func duplicateInSortedArray(nums []int) int {
	for i, val := range nums[:len(nums)-1] {
		if val == nums[i+1] {
			return val
		}
	}
	return -1
}

func duplicateInSortedArrayHashing(nums []int) int {
	res := make(map[int]bool)
	for _, val := range nums {
		if res[val] {
			return val
		}
		res[val] = true

	}
	return -1
}

func removeDuplicates(nums []int) []int {
	i := 0
	j := i + 1

	for j < len(nums) {
		if nums[i] == nums[j] {
			nums = append(nums[:j], nums[j+1:]...)

		} else {
			i++
			j++
		}
	}

	return nums

}



func main() {

	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	gg := removeDuplicates(arr)
	fmt.Printf("gg: %v\n", gg)

	// arr := []int{1, 2, 2, 3, 3, 4, 5, 6}
	// gg := duplicateInSortedArray(arr)
	// fmt.Printf("gg: %v\n", gg)

	// arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	// // gg := missingNumber(arr)
	// gg := missingNumberEfficient(arr)
	// fmt.Printf("gg: %v\n", gg)
}
