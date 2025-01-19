package main

import "fmt"

func gameOfDeath(n, k int) {

	arr := make([]int, 0)

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	index := 0
	k--
	solve(index, k, arr)

}

func solve(index, k int, arr []int) {

	if len(arr) == 1 {
		fmt.Println(arr[0])
		return
	}

	currentIndex := (index + k) % len(arr)

	arr = append(arr[:currentIndex], arr[currentIndex+1:]...)
	solve(currentIndex, k, arr)

}

func main() {

	gameOfDeath(5, 2)
}
