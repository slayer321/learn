package main

import (
	"fmt"
	"log"
)

type SliceOperation struct {
	arr []int
}

func (s *SliceOperation) Add(val ...int) {
	s.arr = append(s.arr, val...)
}

func (s *SliceOperation) Insert(val, index int) {

	if index > len(s.arr) {
		log.Fatal("index is greater then the length of the slice")
	}
	s.arr = append(s.arr[:index+1], s.arr[index:]...)
	s.arr[index] = val

}

func (s *SliceOperation) Delete(index int) {
	if index > len(s.arr) {
		log.Fatal("index is greater then the length of the slice")
	}

	s.arr = append(s.arr[:index], s.arr[index+1:]...)
}

func (s *SliceOperation) LinerSearch(val int) {

	for i, value := range s.arr {
		if value == val {
			fmt.Printf("Element %d found at index %d\n", val, i)
			return
		}
	}
	fmt.Printf("Element %d not found in the slice \n", val)
}

func (s *SliceOperation) Get(index int) int {
	if index > len(s.arr) {
		log.Fatal("index is greater then the length of the slice")
	}

	return s.arr[index]
}

func (s *SliceOperation) Set(index, x int) {
	if index > len(s.arr) {
		log.Fatal("index is greater then the length of the slice")
	}

	// s.arr = append(s.arr[:index+1], s.arr[index:]...)
	s.arr[index] = x
}

// 8,3,14,54,12,33,6
func max(values []int) int {
	maxVal := values[0]

	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal

}

func min(values []int) int {
	maxVal := values[0]

	for _, val := range values {
		if val < maxVal {
			maxVal = val
		}
	}

	return maxVal

}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func avg(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum / 2
}

func ReverseArray(arr []int) {
	s := 0
	end := len(arr) - 1

	for s <= end {
		arr[s], arr[end] = arr[end], arr[s]
		s++
		end--
	}

}

func InsertInSortedArray(arr []int, k int) {

	for i := range arr[:len(arr)-1] {

		if arr[i] < k && arr[i+1] > k {
			arr = append(arr[:i+1], arr[i:]...)
			arr[i+1] = k
		} else if arr[i] < k {
			arr = append(arr[:i+1], arr[i:]...)
			arr[i] = k
		}
	}

}

func main() {

	gg := []int{8, 3, 14, 54, 12, 33, 6}

	fmt.Printf("gg: %v\n", gg)
	InsertInSortedArray(gg, 10)
	fmt.Printf("gg: %v\n", gg)
	// ReverseArray(gg)

	// fmt.Printf("gg: %v\n", gg)

	// out := max(gg)
	// fmt.Printf("out: %v\n", out)
	// minV := min(gg)
	// fmt.Printf("minV: %v\n", minV)

	// sumV := sum(gg)
	// fmt.Printf("sumV: %v\n", sumV)

	// avgV := avg(gg)
	// fmt.Printf("avgV: %v\n", avgV)

	// sliceOps := SliceOperation{
	// 	arr: make([]int, 0, 100),
	// }

	// sliceOps.Add(4, 23, 1, 5, 3)

	// // fmt.Printf("sliceOps.arr: %v\n", sliceOps.arr)

	// sliceOps.Set(4, 99)
	// fmt.Printf("sliceOps.arr: %v\n", sliceOps.arr)
	// sliceOps.Insert(100, 5)

	// fmt.Printf("sliceOps.arr: %v\n", sliceOps.arr)

	// sliceOps.Delete(2)

	// fmt.Printf("sliceOps.arr: %v\n", sliceOps.arr)

	// sliceOps.LinerSearch(100)
}
