package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {

	i := 0
	j := 0

	// maxNumber := 0
	result := []int{}
	queue := []int{}

	for j < len(nums) {
		for len(queue) > 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])

		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == k {
			result = append(result, queue[0])
			if nums[i] == queue[0] {
				queue = queue[1:]
			}
			i++
			j++
		}
	}

	return result

}
func main() {

	arr := []int{1, -1}
	k := 1

	out := maxSlidingWindow(arr, k)
	fmt.Printf("out: %v\n", out)
}
