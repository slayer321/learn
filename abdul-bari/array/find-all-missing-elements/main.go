package main

import "fmt"

func findAllMissingElement(nums []int) []int {

	i := 0
	result := make([]int, 0)

	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			realPos := nums[i] - 1
			nums[i], nums[realPos] = nums[realPos], nums[i]
		} else {
			i++
		}
	}

	for i, val := range nums {
		if val != i+1 {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {

	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	gg := findAllMissingElement(nums)
	fmt.Printf("gg: %v\n", gg)

}
