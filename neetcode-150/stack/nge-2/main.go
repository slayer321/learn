package main

import (
	"fmt"
	"slices"
)

type Stack struct {
	arr []int
}

func (stack *Stack) Push(val int) {
	stack.arr = append(stack.arr, val)
}

func (stack *Stack) Pop() {
	stack.arr = stack.arr[:len(stack.arr)-1]
}

func (stack *Stack) Top() int {
	return stack.arr[len(stack.arr)-1]
}

func nextGreaterElements(nums []int) []int {

	stack := Stack{
		arr: make([]int, 0),
	}
	out := []int{}
	numsLength := len(nums)
	isHelper := func(i int, out []int) ([]int, bool) {
		start := i / numsLength

		isIn := false
		for start < i {
			if nums[start] > nums[i] {
				out = append(out, nums[start])
				stack.Push(nums[i])
				isIn = true
				break
			}
			start++
		}
		return out, isIn
	}

	for i := numsLength - 1; i >= 0; i-- {
		if len(stack.arr) == 0 {
			// start := 0
			// start := i / numsLength

			// isIn := false
			// for start < i {
			// 	if nums[start] > nums[i] {
			// 		out = append(out, nums[start])
			// 		stack.Push(nums[i])
			// 		isIn = true
			// 		break
			// 	}
			// 	start++
			// }
			out, isIn := isHelper(i, out)
			if !isIn {
				stack.Push(nums[i])
				out = append(out, -1)
			}
		} else if len(stack.arr) > 0 && stack.Top() > nums[i] {
			out, isIn := isHelper(i, out)
			if !isIn {
				out = append(out, stack.Top())
				stack.Push(nums[i])
			}
		} else if len(stack.arr) > 0 && stack.Top() < nums[i] {
			for len(stack.arr) > 0 && stack.Top() <= nums[i] {
				stack.Pop()
			}

			if len(stack.arr) == 0 {
				stack.Push(nums[i])

				out = append(out, -1)
			} else {
				stack.Push(nums[i])
				out = append(out, stack.Top())
			}
		}
	}
	slices.Reverse(out)
	return out
}

func circularArra(arr []int) {

	start := 0

	for i := 0; i < len(arr); i++ {
		index := (start + i) % len(arr)
		fmt.Println(arr[index])
	}

}

func monotonicIncreasing(arr []int) {
	stack := Stack{
		arr: make([]int, 0),
	}

	for _, val := range arr {

		for len(stack.arr) > 0 && stack.Top() > val {
			stack.Pop()
		}
		stack.Push(val)
	}

	fmt.Printf("stack.arr: %v Inc\n", stack.arr)

}

func monotonicDecreasing(arr []int) {
	stack := Stack{
		arr: make([]int, 0),
	}

	for _, val := range arr {

		for len(stack.arr) > 0 && stack.Top() < val {
			stack.Pop()
		}
		stack.Push(val)
	}

	fmt.Printf("stack.arr: %v Des\n", stack.arr)

}

func main() {
	// nums := []int{5, 4, 3, 2, 1}
	// out := nextGreaterElements(nums)
	// fmt.Printf("out: %v\n", out)
	// circularArra(nums)

	arr := []int{5, 9, 2, 7, 12}
	monotonicIncreasing(arr)
	monotonicDecreasing(arr)

}
