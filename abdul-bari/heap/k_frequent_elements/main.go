package main

import (
	"container/heap"
	"fmt"
)

type pair struct {
	key       int
	frequency int
}

type intHeap []pair

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {

	// if h[i].distance == h[j].distance {
	// 	return h[i].key > h[j].key
	// }
	return h[i].frequency < h[j].frequency
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

func topKFrequent(nums []int, k int) []int {
	freqencyMap := make(map[int]int)

	for _, val := range nums {
		if _, ok := freqencyMap[val]; ok {
			freqencyMap[val] += 1
		} else {
			freqencyMap[val] = 1
		}
	}

	fmt.Printf("freqencyMap: %v\n", freqencyMap)

	minHeap := &intHeap{}

	for key, freq := range freqencyMap {

		heap.Push(minHeap, pair{key: key, frequency: freq})

		if len(*minHeap) > k {
			heap.Pop(minHeap)
		}

	}
	result := []int{}
	for len(*minHeap) > 0 {
		out := heap.Pop(minHeap)
		result = append(result, out.(pair).key)
	}

	return result

}

func main() {

	arr := []int{1, 1, 1, 2, 2, 3}
	k := 2

	out := topKFrequent(arr, k)
	fmt.Printf("out: %v\n", out)

}
