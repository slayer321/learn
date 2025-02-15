package main

import (
	"fmt"
	"slices"
	"strings"
)

func CountVowels(s string) (int, int) {
	vowelsSlice := []rune{'a', 'e', 'i', 'o', 'u'}
	countVol := 0
	for _, val := range s {
		if slices.Contains(vowelsSlice, val) {
			countVol++
		}
	}

	return countVol, len(strings.Split(s, " "))
}

func CountVowels2(s string) {

	g := []byte(s)
	count := 1
	fmt.Printf("g: %v\n", g)

	for i := 0; i < len(g); i++ {
		if g[i] == 32 {
			count++
		}
	}

	fmt.Println(count)
}

func main() {

	ss := "Take u forward is Awesome"

	gg, jj := CountVowels(strings.ToLower(ss))
	fmt.Printf("gg: %v, jj: %v\n", gg, jj)

	CountVowels2(ss)
}
