package main

import (
	"container/heap"
	"fmt"
)

// https://www.geeksforgeeks.org/nearly-sorted-algorithm/

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

func sort_k_sorted_array(arr []int, k int) []int {
	min_heap := &intHeap{}

	result := []int{}

	for _, val := range arr {
		heap.Push(min_heap, val)

		if len(*min_heap) > 3 {
			out := heap.Pop(min_heap)
			result = append(result, out.(int))
		}
	}

	for _, val := range *min_heap {
		result = append(result, val)
	}

	return result

}

func main() {

	arr := []int{7, 10, 4, 3, 20, 15}
	k := 3

	out := sort_k_sorted_array(arr, k)
	fmt.Printf("out: %v\n", out)
}
