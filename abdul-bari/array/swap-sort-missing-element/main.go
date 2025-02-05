package main

import "fmt"

func repeatedNumber(nums []int) []int {

	i := 0
	result := make([]int, 0)

	for i < len(nums) {

		if nums[i] == nums[i]-1 {
			i++

		} else if nums[i] > i+1 {
			swapPosition := nums[i] - 1
			nums[i], nums[swapPosition] = nums[swapPosition], nums[i]
			if nums[i] == nums[i+1] {
				i++
			}
			// realValue := nums[i] - 1
			// if nums[realValue] == nums[i] {
			// 	i++
			// }
		} else {
			i++
		}
	}

	fmt.Println(nums)
	// for i, val := range nums {
	// 	if val != i+1 {
	// 		result = append(result, val)
	// 		result = append(result, i+1)
	// 	}
	// }
	return result
}

func repeatedNumberOneMore(nums []int) []int {
	i := 0
	result := make([]int, 0)

	for i < len(nums) {

		if nums[i] != nums[nums[i]-1] {
			swapPosition := nums[i] - 1
			nums[i], nums[swapPosition] = nums[swapPosition], nums[i]
		} else {
			i++
		}
	}

	for i, val := range nums {
		if val != i+1 {
			result = append(result, val)
			result = append(result, i+1)
		}
	}

	return result

}

func main() {
	nums := []int{3, 1, 2, 5, 3}
	out := repeatedNumberOneMore(nums)
	fmt.Printf("out: %v\n", out)
}
