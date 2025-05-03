package main

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/description/
func productExceptSelf(nums []int) []int {

	result := []int{}
	var product int
	for _, i := range nums {
		product = nums[0]

		for _, val := range nums[1:] {
			if i == 0 && val == 0 {
				continue
			}
			product = product * val

		}
		output := product
		if i != 0 {
			output = product / i
		}
		result = append(result, output)
	}
	return result
}

func productExceptSelfWorking(nums []int) []int {

	prefix := 1
	output := make([]int, len(nums))

	for i, val := range nums {
		output[i] = prefix
		prefix = prefix * val
	}
	postfix := 1
	for j := len(nums) - 1; j >= 0; j-- {
		outVal := output[j]
		finalval := postfix * outVal
		output[j] = finalval
		postfix = nums[j] * postfix
	}

	return output
}
func main() {

	arr := []int{-1, 1, 0, -3, 3}

	out := productExceptSelfWorking(arr)
	fmt.Printf("out: %v\n", out)

}
