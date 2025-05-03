package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	myMap := make(map[int]int)

	for _, val := range nums {
		if _, ok := myMap[val]; ok {
			myMap[val] += 1
		} else {
			myMap[val] = 1
		}
	}
	fmt.Printf("myMap: %v\n", myMap)

	longestNum := 0
	newLongest := 0
	for _, val := range nums {
		longestNum = max(newLongest, longestNum)
		previousVal := val - 1
		_, ok := myMap[previousVal]
		if !ok {
			newLongest = 0
			newLongest = newLongest + myMap[val]
			newVal := val
			for {
				nextVal := newVal + 1
				if mapVal, ok := myMap[nextVal]; ok {
					newLongest = newLongest + mapVal
					newVal++
					continue
				}
				break
			}
		}
	}

	return longestNum

}

func longestConsecutiveWorking(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	longest := 0
	newSet := make(map[int]bool)

	for _, val := range nums {
		newSet[val] = true
	}

	for _, val := range nums {

		if !newSet[val-1] {
			newVal := val
			currentStreak := 1

			for newSet[newVal+1] {
				currentStreak++
				newVal++
			}

			// longest = max(longest, currentStreak)
			if currentStreak > longest {
				longest = currentStreak
			}
		}

	}

	return longest

}

func main() {
	arr := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}

	out := longestConsecutiveWorking(arr)
	fmt.Printf("out: %v\n", out)
}
