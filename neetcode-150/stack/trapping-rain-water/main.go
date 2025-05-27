package main

import "fmt"

func trap(height []int) int {
	maxL := make([]int, len(height))

	maxLVal := 0
	for i, val := range height {
		maxLVal = max(val, maxLVal)
		maxL[i] = maxLVal
	}

	maxR := make([]int, len(height))

	maxR[len(height)-1] = height[len(height)-1]

	for i := len(height) - 2; i >= 0; i-- {
		maxR[i] = max(height[i], maxR[i+1])
	}
	// maxRVal := 0
	// for i, val := range height {
	// 	maxRVal = max(val, maxRVal)
	// 	maxR[i] = maxRVal
	// }

	finalSum := 0
	for i, val := range height {
		currentMin := min(maxL[i], maxR[i])
		holdingVal := currentMin - val
		finalSum += holdingVal

	}

	return finalSum

}

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	out := trap(arr)
	fmt.Printf("out: %v\n", out)
}
