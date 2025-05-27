package main

import (
	"fmt"
	"slices"
)

type StackT struct {
	val   int
	index int
}

type Stack struct {
	arr []StackT
}

func (stack *Stack) Push(val, index int) {
	stack.arr = append(stack.arr, StackT{
		val:   val,
		index: index,
	})
}

func (stack *Stack) Pop() {
	stack.arr = stack.arr[:len(stack.arr)-1]
}

func (stack *Stack) Top() StackT {
	return stack.arr[len(stack.arr)-1]
}

func nseL(arr []int) []int {
	result := []int{}
	stack := Stack{
		arr: make([]StackT, 0),
	}

	for i, val := range arr {
		if len(stack.arr) == 0 {
			result = append(result, -1)
		} else if len(stack.arr) > 0 && stack.Top().val < val {
			result = append(result, stack.Top().index)
		} else if len(stack.arr) > 0 && stack.Top().val >= val {
			for len(stack.arr) > 0 && stack.Top().val >= val {
				stack.Pop()
			}

			if len(stack.arr) == 0 {
				result = append(result, -1)
			} else {
				result = append(result, stack.Top().index)
			}
		}

		stack.Push(val, i)
	}
	return result
}

func nseR(arr []int) []int {
	result := []int{}
	stack := Stack{
		arr: make([]StackT, 0),
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if len(stack.arr) == 0 {
			result = append(result, len(arr))
		} else if len(stack.arr) > 0 && stack.Top().val < arr[i] {
			result = append(result, stack.Top().index)
		} else if len(stack.arr) > 0 && stack.Top().val >= arr[i] {
			for len(stack.arr) > 0 && stack.Top().val >= arr[i] {
				stack.Pop()
			}

			if len(stack.arr) == 0 {
				result = append(result, len(arr))
			} else {
				result = append(result, stack.Top().index)
			}
		}

		stack.Push(arr[i], i)
	}
	slices.Reverse(result)
	return result
}

func largestRectangleArea(heights []int) int {

	right := nseR(heights)
	left := nseL(heights)

	widthSlice := make([]int, len(right))

	for i := 0; i < len(right); i++ {
		val := right[i] - left[i] - 1
		widthSlice[i] = val
	}
	maxLenght := 0
	for i, val := range heights {
		area := widthSlice[i] * val
		maxLenght = max(maxLenght, area)
	}

	return maxLenght

}
func BytesToInts(b []byte) []int {
	ints := make([]int, len(b))
	for i, v := range b {
		ints[i] = int(v)
	}
	return ints
}

func calNewHistorgramSlice(prevRow []byte, currentRow []byte) []int {
	sum := 0
	newSlice := make([]int, len(currentRow))
	for i, val := range currentRow {

		if val == 0 {
			newSlice[i] = 0
		} else {
			sum = int(currentRow[i]) + int(prevRow[i])
			newSlice[i] = sum
		}
	}

	return newSlice
}

func maximalRectangle(matrix [][]byte) int {
	// prevRow := []byte{}
	// newSlice := []int{}
	// finalLargesHistorgram := 0
	// for i, val := range matrix {
	// 	if i == 0 {
	// 		newSlice = BytesToInts(val)
	// 	} else {
	// 		newSlice = calNewHistorgramSlice(prevRow, val)
	// 	}

	// 	largesHistorgram := largestRectangleArea(newSlice)
	// 	// fmt.Printf("largesHistorgram: %v\n", largesHistorgram)
	// 	finalLargesHistorgram = max(finalLargesHistorgram, largesHistorgram)
	// 	prevRow = val
	// }
	// return finalLargesHistorgram
	ans := make([]int, len(matrix[0]))
	maxV := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				ans[j] = 0
			} else {
				ans[j]++
			}
		}
		maxV = max(maxV, largestRectangleArea(ans))
	}

	return maxV
}

func main() {

	heights := [][]byte{{1, 0, 1, 0, 0}, {1, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}
	// heights := [][]byte{{0}}
	out := maximalRectangle(heights)
	fmt.Printf("out: %v\n", out)
	// out := largestRectangleArea(heights)
}
