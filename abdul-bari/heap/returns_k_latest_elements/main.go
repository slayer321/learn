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
	return h[i] < h[j]
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

func k_largest_element(arr []int, k int) []int {

	min_heap := &intHeap{}
	result := make([]int, k)
	for _, val := range arr {
		heap.Push(min_heap, val)

		if len(*min_heap) > k {
			heap.Pop(min_heap)
		}
	}

	for i, val := range *min_heap {
		result[i] = val
	}

	return result

}

func main() {

	arr := []int{7, 10, 4, 3, 20, 15}
	k := 3

	out := k_largest_element(arr, k)
	fmt.Printf("out: %v\n", out)

}
