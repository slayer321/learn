package main

import "fmt"

func maxProfit(prices []int) int {

	i := 0
	j := i + 1
	maxProfit := 0
	for j < len(prices) {
		if prices[i] < prices[j] {
			currentProfit := prices[j] - prices[i]
			maxProfit = max(currentProfit, maxProfit)
			j++
		} else {
			i = j
			j++
		}
	}

	return maxProfit
}
func main() {

	arr := []int{7, 1, 5, 3, 6, 4}
	out := maxProfit(arr)
	fmt.Printf("out: %v\n", out)

}
