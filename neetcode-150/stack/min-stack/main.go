package main

import "fmt"

var (
	FinalMinVal int
)

type MinStack struct {
	arr []int
	// helperStack []int
}

func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0),
		// helperStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {

	// if len(this.arr) == 0 {
	// 	this.helperStack = append(this.helperStack, val)
	// } else if val <= this.helperStack[len(this.helperStack)-1] {
	// 	this.helperStack = append(this.helperStack, val)
	// }
	// this.arr = append(this.arr, val)

	if len(this.arr) == 0 {
		this.arr = append(this.arr, val)
		FinalMinVal = val
	} else if len(this.arr) > 0 && val < FinalMinVal {
		newVal := 2*val - FinalMinVal
		this.arr = append(this.arr, newVal)
		FinalMinVal = val
	} else if val >= FinalMinVal {
		this.arr = append(this.arr, val)

	}

	// fmt.Printf("FinalMinVal: %v\n", FinalMinVal)
	// if len(this.arr) > 0 && val <= FinalMinVal {
	// 	newVal := 2*val - FinalMinVal
	// 	this.arr = append(this.arr, newVal)
	// 	FinalMinVal = val
	// } else {
	// 	this.arr = append(this.arr, val)
	// }

}

func (this *MinStack) Pop() {

	// if this.Top() == this.helperStack[len(this.helperStack)-1] {
	// 	this.helperStack = this.helperStack[:len(this.helperStack)-1]
	// }
	// this.arr = this.arr[:len(this.arr)-1]
	// fmt.Printf("FinalMinVal: %v\n", FinalMinVal)

	if this.arr[len(this.arr)-1] >= FinalMinVal {
		this.arr = this.arr[:len(this.arr)-1]

	} else if this.arr[len(this.arr)-1] < FinalMinVal {

		newMin := (2 * FinalMinVal) - this.arr[len(this.arr)-1]
		FinalMinVal = newMin
		this.arr = this.arr[:len(this.arr)-1]
	}

}

func (this *MinStack) Top() int {

	if this.arr[len(this.arr)-1] >= FinalMinVal {
		return this.arr[len(this.arr)-1]
	} else if this.arr[len(this.arr)-1] < FinalMinVal {
		return FinalMinVal
	}

	return -1
}

func (this *MinStack) GetMin() int {

	if len(this.arr) == 0 {
		return -1
	}
	return FinalMinVal
	// minVal := this.Top()
	// for i := 0; i < len(this.arr); i++ {
	// 	minVal = min(minVal, this.arr[i])
	// }

	// return this.helperStack[len(this.helperStack)-1]

}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())
	stack.Pop()
	fmt.Printf("stack.Top(): %v\n", stack.Top())
	fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())

	// stack.Push(2)
	// stack.Push(0)
	// stack.Push(3)
	// stack.Push(0)
	// fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())
	// stack.Pop()
	// fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())
	// stack.Pop()
	// fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())
	// stack.Pop()
	// fmt.Printf("stack.GetMin(): %v\n", stack.GetMin())
}
