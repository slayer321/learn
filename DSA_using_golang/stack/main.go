package main

import (
	"fmt"
	"log"
)

type Stack struct {
	store []int
}

func (s *Stack) Push(val int) []int {

	s.store = append(s.store, val)
	//fmt.Println(s.store)
	return s.store
}

func (s *Stack) Pop() []int {
	if len(s.store) == 0 {
		log.Println("Stack is empty no data to pop")
		return nil
	}

	s.store = s.store[:len(s.store)-1]
	//fmt.Println(s.store)
	return s.store

}

func main() {

	stack := Stack{
		store: make([]int, 0, 2),
	}
	f := stack.Push(4)
	fmt.Println(f)
	f = stack.Push(54)
	fmt.Println(f)
	f = stack.Push(12)
	fmt.Println(f)
	f = stack.Pop()
	fmt.Println(f)
	f = stack.Push(99)
	fmt.Println(f)
	f = stack.Push(55)
	fmt.Println(f)
	f = stack.Pop()
	fmt.Println(f)

}
