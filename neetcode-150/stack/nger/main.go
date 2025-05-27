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

func ngeRight(arr []int) []int {
	stack := Stack{
		arr: make([]int, 0),
	}

	result := []int{}

	for i := len(arr) - 1; i >= 0; i-- {

		if len(stack.arr) == 0 {
			result = append(result, -1)
		} else if len(stack.arr) > 0 && stack.Top() > arr[i] {
			result = append(result, stack.Top())
		} else if len(stack.arr) > 0 && stack.Top() < arr[i] {
			for stack.Top() < arr[i] {
				stack.Pop()
			}

			if len(stack.arr) == 0 {
				result = append(result, -1)
			} else {
				result = append(result, stack.Top())
			}
		}

		stack.Push(arr[i])

	}

	slices.Reverse(result)

	return result

}

func main() {
	arr := []int{1, 3, 2, 4}
	out := ngeRight(arr)
	fmt.Printf("out: %v\n", out)
}
