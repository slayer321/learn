package main

import "fmt"

func maxArea(height []int) int {

	i := 0
	j := len(height) - 1
	maxArea := 0

	for i < j {

		//if height[i] >= height[j] {
		newArea := (j - i) * min(height[j], height[i])
		maxArea = max(maxArea, newArea)
		// }

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea

}

func main() {
	arr := []int{8, 7, 2, 1}
	out := maxArea(arr)
	fmt.Printf("out: %v\n", out)
}
