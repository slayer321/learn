package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	checkMap := make(map[rune]int)

	for _, val := range s {
		if _, ok := checkMap[val]; !ok {
			checkMap[val] = 1
		} else {
			checkMap[val] += 1
		}
	}

	for _, val := range t {
		if _, ok := checkMap[val]; ok {
			checkMap[val] -= 1
		}
	}

	for _, val := range checkMap {
		if val > 0 {
			return false
		}
	}

	return true
}

func isAnagramE(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countSlice := make([]int, 26)

	for i, val := range s {
		sValue := val - rune('a')
		tValue := rune(t[i]) - rune('a')
		countSlice[sValue] += 1
		countSlice[tValue] -= 1
	}

	for _, val := range countSlice {
		if val != 0 {
			return false
		}
	}

	return true
}

func main() {

	s := "anagram"
	t := "nagaram"

	gg := isAnagramE(s, t)
	fmt.Printf("gg: %v\n", gg)
}
