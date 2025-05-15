package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	i := 0
	j := len(numbers) - 1
	out := make([]int, 2)
	for i <= j {
		sum := numbers[i] + numbers[j]

		if sum == target {
			out[0] = i + 1
			out[1] = j + 1
			return out
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}

	return out
}

func main() {
	arr := []int{2, 3, 4}
	target := 6

	out := twoSum(arr, target)
	fmt.Printf("out: %v\n", out)
}
