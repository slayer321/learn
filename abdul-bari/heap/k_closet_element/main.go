package main

import (
	"container/heap"
	"fmt"
	"slices"
)

type pair struct {
	key      int
	distance int
}

type intHeap []pair

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {

	if h[i].distance == h[j].distance {
		return h[i].key > h[j].key
	}
	return h[i].distance > h[j].distance
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *intHeap) Pop() any {
	old := *h
	length := len(old)
	lastVal := old[length-1]
	*h = old[:length-1]
	return lastVal
}

func AbsInt(z int) int {
	if z < 0 {
		return -z
	}
	return z
}

// func findClosestElements(arr []int, k int, x int) []int {

// 	newSlice := make([]int, len(arr))
// 	var Val int
// 	for i, val := range arr {

// 		if x < 0 {
// 			Val = AbsInt(x + val)
// 		} else {
// 			Val = AbsInt(x - val)

// 		}
// 		newSlice[i] = Val
// 	}

// 	fmt.Printf("newSlice: %v\n", newSlice)

// 	maxHeap := &intHeap{}

// 	for _, val := range newSlice {
// 		heap.Push(maxHeap, val)

// 		if len(*maxHeap) > k {
// 			heap.Pop(maxHeap)
// 		}
// 	}
// 	fmt.Printf("maxHeap: %v\n", maxHeap)
// 	result := make([]int, k)

// 	afterIndex := false
// 	for i, val := range *maxHeap {

// 		if afterIndex {
// 			out := x + val
// 			result[i] = out

// 		} else {
// 			out := x - val
// 			result[i] = out
// 			if val == 0 {
// 				afterIndex = true
// 			}
// 		}
// 	}

// 	return result
// }

func findClosestElementsWorking(arr []int, k int, x int) []int {

	maxHeap := &intHeap{}

	for _, val := range arr {
		heap.Push(maxHeap, pair{distance: AbsInt(x - val), key: val})

		if len(*maxHeap) > k {
			heap.Pop(maxHeap)
		}
	}
	fmt.Printf("maxHeap: %+v\n", maxHeap)
	result := []int{}

	for len(*maxHeap) > 0 {
		val := heap.Pop(maxHeap)
		result = append(result, val.(pair).key)
	}

	slices.Sort(result)

	// result []int{}

	return result

}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3

	out := findClosestElementsWorking(arr, k, x)
	fmt.Printf("out: %v\n", out)

}
