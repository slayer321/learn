package main

import (
	"fmt"
	"strconv"
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

func (stack *Stack) CheckOperatorAndCal(val string) {

	switch val {
	case "+":
		second := stack.Top()
		stack.Pop()
		first := stack.Top()
		stack.Pop()
		finalVal := first + second
		stack.Push(finalVal)
	case "-":
		second := stack.Top()
		stack.Pop()
		first := stack.Top()
		stack.Pop()
		finalVal := first - second
		stack.Push(finalVal)
	case "*":
		second := stack.Top()
		stack.Pop()
		first := stack.Top()
		stack.Pop()
		finalVal := first * second
		stack.Push(finalVal)
	case "/":
		second := stack.Top()
		stack.Pop()
		first := stack.Top()
		stack.Pop()
		finalVal := first / second
		stack.Push(finalVal)
	}
}

func isOperator(operator string) bool {
	if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
		return true
	}
	return false
}

func evalPRN(tokens []string) int {
	stack := Stack{
		arr: make([]int, 0),
	}

	for _, val := range tokens {
		if isOperator(val) {
			stack.CheckOperatorAndCal(val)
		} else {
			parseFloat, _ := strconv.ParseFloat(val, 64)
			stack.Push(int(parseFloat))
		}
	}

	return stack.Top()
}

func main() {
	arr := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	out := evalPRN(arr)
	fmt.Printf("out: %v\n", out)

	// s := "+"

	// gg, err := strconv.Atoi(s)

	// // gg, err := strconv.ParseFloat(s, 64)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("gg: %v\n", gg)
}
