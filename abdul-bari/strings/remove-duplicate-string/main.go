package main

import "fmt"

func removeDuplicate(s string) string {

	mapString := make(map[rune]struct{})

	for _, val := range s {
		if _, ok := mapString[val]; !ok {
			mapString[val] = struct{}{}
		}
	}

	result := ""
	for k := range mapString {
		result += string(k)
	}

	return result
}

func main() {

	s := "bcabc"
	gg := removeDuplicate(s)
	fmt.Printf("gg: %v\n", gg)

}
