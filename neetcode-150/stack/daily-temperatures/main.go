package main

import (
	"fmt"
	"slices"
)

type StackFields struct {
	val   int
	index int
}

type Stack struct {
	arr []StackFields
}

func (stack *Stack) Push(val StackFields) {
	stack.arr = append(stack.arr, val)
}

func (stack *Stack) Pop() {
	stack.arr = stack.arr[:len(stack.arr)-1]
}

func (stack *Stack) Top() StackFields {
	return stack.arr[len(stack.arr)-1]
}

func dailyTemperatures(temperatures []int) []int {

	stack := Stack{
		arr: make([]StackFields, 0),
	}
	answer := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack.arr) == 0 {
			answer = append(answer, 0)
		} else if len(stack.arr) > 0 && stack.Top().val <= temperatures[i] {
			for len(stack.arr) > 0 && stack.Top().val <= temperatures[i] {
				stack.Pop()
			}

			if len(stack.arr) == 0 {
				answer = append(answer, 0)
			} else {
				answer = append(answer, stack.Top().index-i)
			}

		} else if len(stack.arr) > 0 && stack.Top().val > temperatures[i] {
			answer = append(answer, stack.Top().index-i)
		}

		stack.Push(StackFields{
			val:   temperatures[i],
			index: i,
		})
	}

	slices.Reverse(answer)
	return answer

}
func main() {
	temp := []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}
	out := dailyTemperatures(temp)
	fmt.Printf("out: %v\n", out)

}
