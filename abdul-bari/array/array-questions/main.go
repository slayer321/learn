package main

import (
	"fmt"
)

func insertInSortedArray(arr []int, target int) []int {

	for i, val := range arr {

		if target < val {

			arr = append(arr[:i+1], arr[i:]...)
			arr[i] = target
			break
		} else if i == len(arr)-1 && target > val {
			arr = append(arr[:i+1], arr[i:]...)
			arr[i+1] = target
			break
		}
	}

	return arr
}

func insertInSortedArrayOptimal(arr []int, target int) []int {

	// i := 0
	end := len(arr) - 1

	for arr[end] > target {
		arr = append(arr[:end+1], arr[end:]...)
		end--
	}
	arr[end] = target

	return arr

}

// Doesn't work for all the conditions
func insertInSortedArrayB(arr []int, target int) []int {

	l := 0
	h := len(arr) - 1

	return solve(l, h, arr, target)

}
func solve(l, h int, arr []int, target int) []int {
	if l <= h {

		mid := (l + h) / 2

		if target < arr[mid] && l == h {
			arr = append(arr[:mid+1], arr[mid:]...)
			arr[mid] = target
			return arr
		} else if target > arr[mid] && l == h {
			arr = append(arr[:mid+1], arr[mid:]...)
			arr[mid+1] = target
			return arr
		}

		if target < arr[mid] {
			return solve(l, mid-1, arr, target)

		} else if target > arr[mid] {
			return solve(mid+1, h, arr, target)
		}
	}
	return arr
}

func insert(arr []int, value int) []int {

	i := 0
	for i < len(arr) && arr[i] < value {
		i++
	}

	arr = append(arr, 0)
	fmt.Printf("arr: %v\n", arr)
	copy(arr[i+1:], arr[i:])
	fmt.Printf("arr: %v\n", arr)
	arr[i] = value
	fmt.Printf("arr: %v\n", arr)
	return arr

}

func arrayIsSorted(arr []int) bool {
	for i, val := range arr[:len(arr)-1] {
		if arr[i+1] < val {
			return false
		}
	}
	return true
}

func shiftNegativeToLeft(arr []int) []int {
	i := 0
	e := len(arr) - 1

	for i <= e {

		if arr[i] < 0 {
			i++
		} else if arr[i] > 0 && arr[e] < 0 {
			arr[i], arr[e] = arr[e], arr[i]
			i++
			e--

		} else if arr[e] > 0 {
			e--
		}
	}

	return arr

}

func mergeSortedArray(nums1, nums2 []int, m, n int) []int {

	i := 0
	j := 0
	out := make([]int, 0, m+n)
	// k := 0

	for i < m && j < n {

		if nums1[i] <= nums2[j] {
			out = append(out, nums1[i])
			i++
			// m--
		} else if nums2[j] <= nums1[i] {
			out = append(out, nums2[j])
			j++
			// n--
		}

	}

	for ; i < m; i++ {
		out = append(out, nums1[i])
	}
	for ; j < n; j++ {
		out = append(out, nums2[j])
	}

	// if n != 0 {
	// 	out = append(out, nums2[j])
	// } else if m != 0 {
	// 	out = append(out, nums1[i])

	// }

	return out
}

func mergeSortedArrayNew(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {

		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k--
			i--
		} else {
			nums1[k] = nums2[j]
			k--
			j--
		}
	}

	return nums1
	// for i <= m && j <= n {

	// 	if nums1[i] <= nums2[j] {
	// 		if nums1[i] == nums2[j] {
	// 			// nums1 = append(nums1[:i], nums2[j])

	// 			copy(nums1[i+1:], nums1[i:])
	// 			nums1[i] = nums2[j]
	// 			j++
	// 			i++
	// 		}
	// 		i++
	// 	} else if nums2[j] <= nums1[i] {
	// 		if nums2[j] == nums1[i] {
	// 			copy(nums2[i+1:], nums2[i:])
	// 			nums1[i] = nums2[j]
	// 			j++
	// 			i++
	// 		}
	// 		j++
	// 	}
	// }

	// // for ; i < m; i++ {
	// // 	nums1[i] =
	// // 	out = append(out, nums1[i])
	// // }
	// for ; j < n; j++ {
	// 	nums1[i] = nums2[j]
	// 	i++

	// 	// out = append(out, nums2[j])
	// }
	// return nums1
}

// [3,4,5,1,2]
// [1,2,3,4,5]

func main() {

	nums1 := []int{0}
	nums2 := []int{1}
	result := mergeSortedArrayNew(nums1, 0, nums2, len(nums2))
	fmt.Printf("result: %v\n", result)

	// nums1 := []int{4, 6, 9, 20, 23}
	// nums2 := []int{3, 8, 10, 11}
	// result := mergeSortedArray(nums1, nums2, len(nums1), len(nums2))
	// fmt.Printf("result: %v\n", result)

	// arr := []int{-6, 3, -8, 10, 5, -7, -9, 12, -4, 2}
	// gg := shiftNegativeToLeft(arr)
	// fmt.Printf("gg: %v\n", gg)
	// arr := []int{11, 22, 22, 34, 44, 54}
	// // index := 2
	// sortedValue := arrayIsSorted(arr)
	// fmt.Printf("sortedValue: %v\n", sortedValue)
	// value := 9

	// out := insert(arr, value)

	// fmt.Printf("out: %v\n", out)

	// arr := []int{2, 9, 12, 16, 25, 30}
	// fmt.Printf("arr: %v\n", arr)

	// out := insertInSortedArrayOptimal(arr, 15)
	// fmt.Printf("out: %v\n", out)
	// out := insertInSortedArray(arr, 29)
	// fmt.Printf("out: %v\n", out)
	// out := insertInSortedArrayB(arr, 1)
	// fmt.Printf("out: %v\n", out)
}
