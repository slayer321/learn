package main

import "fmt"

func insertmax(arr []int, n int) {

	i := n
	temp := arr[i]

	for i > 1 && temp > arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}

	arr[i] = temp
}

func insertmin(arr []int, n int) {
	i := n
	temp := arr[i]

	for i > 1 && temp < arr[i/2] {
		arr[i] = arr[i/2]
		i = i / 2
	}
	arr[i] = temp
}

func main() {

	arr := []int{0, 10, 20, 30, 25, 5, 40, 35}

	for i := 2; i <= len(arr)-1; i++ {
		insertmin(arr, i)
	}

	fmt.Println(arr)

}
