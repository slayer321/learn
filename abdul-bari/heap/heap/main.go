package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	length := len(old)
	lastVal := old[length-1]
	*h = old[:length-1]
	return lastVal
}

func main() {

	nums := &intHeap{3, 1, 4, 5, 1, 1, 2, 6}

	heap.Init(nums)
	fmt.Printf("nums: %v\n", nums)

	heap.Push(nums, 20)
	fmt.Printf("nums: %v\n", nums)

	gg := heap.Pop(nums)
	fmt.Printf("gg: %v\n", gg)
	fmt.Printf("nums: %v\n", nums)

}
