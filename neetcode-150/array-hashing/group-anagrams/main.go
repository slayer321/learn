package main

import (
	"fmt"
	"slices"
)

func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	countSlice := make([]int, 26)

	for i, val := range s {
		sValue := rune(val) - rune('a')
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

func groupAnagrams(str []string) [][]string {

	result := make([][]string, 0)
	skipIndex := make([]int, 0)

	for i := 0; i < len(str); i++ {

		if slices.Contains(skipIndex, i) {
			continue
		}
		subSlice := make([]string, 0)
		for j := i + 1; j < len(str); j++ {
			val := isAnagram(str[i], str[j])
			if val {
				subSlice = append(subSlice, str[j])
				skipIndex = append(skipIndex, j)
			}
		}
		subSlice = append(subSlice, str[i])
		result = append(result, subSlice)

	}

	return result
}

func groupAnagramsOptimized(strs []string) [][]string {

	res := make(map[[26]int][]string)

	for _, s := range strs {
		count := [26]int{}
		for _, c := range s {
			count[c-'a'] += 1
		}

		res[count] = append(res[count], s)
	}

	result := make([][]string, 0)

	for _, val := range res {
		result = append(result, val)
	}

	return result
}

func main() {

	// s := "anagram"
	// t := "nagaram"

	// out := isAnagram(s, t)
	// fmt.Printf("out: %v\n", out)
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagramsOptimized(strs)
	fmt.Printf("result: %v\n", result)

}
