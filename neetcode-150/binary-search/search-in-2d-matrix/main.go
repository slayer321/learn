package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {

	// rows := len(matrix)
	coloum := len(matrix[0])

	// fmt.Printf("rows: %v\n", rows)
	// fmt.Printf("coloum: %v\n", coloum)

	for i, arr := range matrix {
		lastColoums := matrix[i][coloum-1]
		if lastColoums >= target {
			return solve(arr, target)
		}
	}

	return false

}

func solve(arr []int, target int) bool {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return true
		}

		if target > arr[mid] {
			l = mid + 1
		} else if target < arr[mid] {
			r = mid - 1
		}
	}

	return false
}

func optimizedSolution(matrix [][]int, target int) bool {
	row := 0
	rowN := len(matrix) - 1
	coloum := len(matrix[0]) - 1
	coloumN := len(matrix[0]) - 1

	for (row >= 0 && row <= rowN) && (coloum >= 0 && coloum <= coloumN) {
		if target == matrix[row][coloum] {
			return true
		} else if target < matrix[row][coloum] {
			coloum--
		} else if target > matrix[row][coloum] {
			row++
		}
	}

	return false

}

func main() {

	arr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}

	// out := searchMatrix(arr, 301)
	out := optimizedSolution(arr, 100)
	fmt.Printf("out: %v\n", out)
}
