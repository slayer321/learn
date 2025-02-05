package main

import "fmt"

func findDuplicatesUsingMap(nums []int) []int {
	outMap := make(map[int]struct{})
	out := make([]int, 0)

	for _, val := range nums {

		if _, found := outMap[val]; !found {
			outMap[val] = struct{}{}
		} else {
			out = append(out, val)
		}
	}

	return out
}

func findDuplicatesUsing(nums []int) []int {
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
			result = append(result, val)
		}
	}

	return result

}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	gg := findDuplicatesUsing(nums)
	// gg := findDuplicatesUsingMap(nums)
	fmt.Printf("gg: %v\n", gg)

}
