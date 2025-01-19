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

	s.arr = s.arr[:len(s.arr)-1]
}

func (s *Stack) Top() int {

	return s.arr[len(s.arr)-1]
}

func (s *Stack) Display() {
	for i := len(s.arr) - 1; i >= 0; i-- {
		fmt.Println(s.arr[i])

	}
	fmt.Println("-------------------------")
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func DeleteMiddle(stack *Stack) {
	if stack.IsEmpty() {
		return
	}

	middle := (len(stack.arr) / 2) + 1

	solve(stack, middle)
}

func solve(stack *Stack, middleValue int) {

	if middleValue == 1 {
		stack.Pop()
		return

	}

	val := stack.arr[len(stack.arr)-1]
	stack.Pop()

	solve(stack, middleValue-1)
	stack.Push(val)
}

func main() {
	stack := Stack{
		arr: make([]int, 0, 10),
	}

	stack.Push(121)
	stack.Push(23)
	stack.Push(13)
	stack.Push(324)
	// stack.Push(51)

	stack.Display()

	DeleteMiddle(&stack)

	stack.Display()

}
