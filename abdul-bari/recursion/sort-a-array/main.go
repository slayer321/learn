package main

import "fmt"

func sortAArray(arr []int) {

	if len(arr) == 1 {
		return
	}

	temp := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1]

	sortAArray(arr)
	insert(arr, temp)

}

func insert(arr []int, temp int) []int {

	if len(arr) == 0 || arr[len(arr)-1] <= temp {
		arr = append(arr, temp)
		return arr
	}

	val := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1]

	arr = insert(arr, temp)
	arr = append(arr, val)
	return arr

}

func main() {

	arr := []int{1, 0, 5, 2}
	fmt.Printf("arr: %v\n", arr)
	sortAArray(arr)
	fmt.Printf("newArrr %v\n", arr)
}
