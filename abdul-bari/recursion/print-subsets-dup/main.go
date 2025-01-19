package main

import (
	"fmt"
	"slices"
)

func printSubsetDup(nums []int) [][]int {

	op := [][]int{}

	subsets := []int{}
	value := solve(nums, subsets, op)
	fmt.Printf("value: %v\n", value)

	slices.SortFunc(value, slices.Compare)
	value = slices.CompactFunc(value, slices.Equal)

	return value

}

func solve(ip []int, subset []int, op [][]int) [][]int {
	if len(ip) == 0 {
		op = append(op, subset)
		return op
	}

	op = solve(ip[1:], subset, op)
	op = solve(ip[1:], append(subset, ip[0]), op)
	return op

}

func sliceWork() {

	// ints := make([]int, 25, 500)
	// fmt.Printf("len: %d, capacity: %d\n", len(ints), cap(ints))
	// ints = slices.Clip(ints)
	// fmt.Printf("len: %d, capacity: %d\n", len(ints), cap(ints))

	// original := []int{0, 1, 2, 4, 5}
	// // cloned := original
	// // cloned[0] = 22
	// // fmt.Printf("original: %v\n", original)
	// // fmt.Printf("cloned: %v\n", cloned)
	// cloned := slices.Clone(original)
	// fmt.Printf("original: %p\n", &original)
	// fmt.Printf("cloned: %p\n", &cloned)

	// strings := []string{"Apple", "Orange", "Apple", "Banana"}
	// ints := []int{0, 1, 1, 2, 0, 5, 7, 5}

	// strings = slices.Compact(strings)
	// ints = slices.Compact(ints)

	// fmt.Printf("Not Sorted - Compact does not work\n")
	// fmt.Println(strings)
	// fmt.Println(ints)

	// fmt.Printf("\nSorted - Compact now works as expected\n")

	// slices.Sort(strings)
	// slices.Sort(ints)

	// strings = slices.Compact(strings)
	// ints = slices.Compact(ints)

	// fmt.Printf("strings: %v\n", strings)
	// fmt.Printf("ints: %v\n", ints)

	gg := [][]int{
		{1},
		{41, 2},
		{},
		{5, 3},
		{55, 9},
	}
	fmt.Printf("gg: %v\n", gg)
	kk := slices.Delete(gg, 0, 0+1)
	fmt.Printf("gg: %v\n", kk)
}

func main() {

	//sliceWork()

	nums := []int{4, 4, 4, 1, 4}

	gg := printSubsetDup(nums)
	fmt.Printf("gg: %v\n", gg)

}
