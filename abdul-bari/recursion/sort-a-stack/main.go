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
		return
	}

	s.arr = s.arr[0 : len(s.arr)-1]
}

func (s *Stack) Display() {
	for _, val := range s.arr {
		fmt.Println(val)
	}
	fmt.Println("------------------------------")
}

func (s *Stack) Top() int {
	return s.arr[len(s.arr)-1]
}

func sortAStack(stack *Stack) {
	if len(stack.arr) == 1 {
		return
	}

	temp := stack.Top()
	stack.Pop()

	sortAStack(stack)
	insert(stack, temp)

}

func insert(stack *Stack, temp int) {
	if len(stack.arr) == 0 || stack.arr[len(stack.arr)-1] >= temp {
		stack.arr = append(stack.arr, temp)
		return
	}

	val := stack.Top()
	stack.Pop()
	insert(stack, temp)
	stack.Push(val)
}

func main() {

	stack := Stack{
		arr: make([]int, 0, 10),
	}

	stack.Push(1)
	stack.Push(0)
	stack.Push(5)
	stack.Push(2)
	stack.Push(-2)
	stack.Push(20)
	stack.Display()

	sortAStack(&stack)

	stack.Display()

}
