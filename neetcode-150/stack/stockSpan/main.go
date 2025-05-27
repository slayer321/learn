package main

import "fmt"

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

func stockSpan(arr []int) []int {
	stack := Stack{
		arr: make([]int, 0),
	}

	result := []int{}

	for i := 0; i < len(arr); i++ {
		if len(stack.arr) == 0 || stack.Top() > arr[i] {
			result = append(result, 1)
		} else if len(stack.arr) > 0 && stack.Top() <= arr[i] {
			count := 1
			newIndex := i - 1
			for len(stack.arr) > 0 && arr[newIndex] <= arr[i] {
				// stack.Pop()
				newIndex--
				count++
			}
			result = append(result, count)
		}
		stack.Push(arr[i])
	}
	return result
}

func main() {
	arr := []int{100, 80, 60, 70, 60, 75, 85}
	out := stockSpan(arr)
	fmt.Printf("out: %v\n", out)

}
