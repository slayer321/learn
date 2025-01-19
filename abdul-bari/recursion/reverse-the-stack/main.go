package main

import (
	"fmt"
	"log"
)

type Stack struct {
	arr []int
}

func (s *Stack) Push(val int) {

	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() {
	if len(s.arr) == 0 {
		log.Fatal("stack is empty")
	}

	s.arr = s.arr[0 : len(s.arr)-1]
}

func (s *Stack) Last() int {
	return s.arr[0]
}

func (s *Stack) Top() int {
	return s.arr[len(s.arr)-1]
}

func (s *Stack) Display() {
	for i := len(s.arr) - 1; i >= 0; i-- {
		fmt.Println(s.arr[i])
	}

	fmt.Println("-----------------------------")
}

// here we are working with removing the first element and then pushing it back
func ReverseStack1(stack *Stack) {

	if len(stack.arr) == 0 {
		return
	}

	val := stack.arr[0]
	stack.arr = stack.arr[1:]

	ReverseStack1(stack)
	stack.Push(val)

}

// here we are using two recursive functions to reverse the stack
func ReverseStack2(stack *Stack) {
	if len(stack.arr) == 1 {
		return
	}

	val := stack.Top()
	stack.Pop()
	ReverseStack2(stack)
	insert(stack, val)
	//stack.Push(val)

}

func insert(stack *Stack, val int) {
	if len(stack.arr) == 0 {
		stack.Push(val)
		return
	}
	top := stack.Top()
	stack.Pop()
	insert(stack, val)
	stack.Push(top)
}

func main() {

	stack := Stack{
		arr: make([]int, 0, 10),
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Display()

	//ReverseStack1(&stack)

	ReverseStack2(&stack)
	stack.Display()

}
