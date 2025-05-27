package main

import "fmt"

type Stack struct {
	arr []rune
}

func (stack *Stack) Push(s rune) {
	stack.arr = append(stack.arr, s)
}

func (stack *Stack) Pop() {
	stack.arr = stack.arr[:len(stack.arr)-1]
}

func (stack *Stack) Top() rune {

	return stack.arr[len(stack.arr)-1]
}

func isValid(s string) bool {

	if len(s) == 1 {
		return false
	}
	stack := Stack{
		arr: make([]rune, 0),
	}

	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack.Push(val)
		}
		if val == ')' || val == '}' || val == ']' {
			if len(stack.arr) == 0 {
				return false
			}
			topVal := stack.Top()
			switch val {
			case ')':
				if topVal != '(' {
					return false
				}
				stack.Pop()
			case ']':
				if topVal != '[' {
					return false
				}
				stack.Pop()
			case '}':
				if topVal != '{' {
					return false
				}
				stack.Pop()
			}
			// if topVal != val {
			// 	return false
			// }
			// stack.Pop()
		}
	}
	if len(stack.arr) != 0 {
		return false
	}
	return true

}

func main() {
	s := "(]"
	out := isValid(s)
	fmt.Printf("out: %v\n", out)
}
