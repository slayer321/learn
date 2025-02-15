package main

import (
	"fmt"
	"strings"
)

func checkPalindrome(s string) bool {

	g := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		if g[i] != g[j] {
			return false
		}

	}
	return true
}

func checkForSentance(s string) bool {

	if s == "" {
		return false
	}
	// g := []rune(strings.ToLower(s))
	g := strings.ToLower(s)
	var result strings.Builder
	for i := 0; i < len(g); i++ {
		b := g[i]
		if (b >= 'a' && b <= 'z') ||
			(b >= 'A' && b <= 'Z') ||
			(b >= '0' && b <= '9') {
			result.WriteByte(b)
		}
	}

	returnString := result.String()

	fmt.Printf("returnString: %v\n", returnString)

	i := 0
	j := len(returnString) - 1

	// runeBytes := []rune(returnString)
	for i < j {

		if returnString[i] != returnString[j] {
			return false
		}
		i++
		j--

	}

	// i := 0
	// j := len(s) - 1

	// for i < j {

	// 	if (g[i] < 97 || g[i] > 122) || (g[i] < 48 || g[i] > 57) {
	// 		i++
	// 	} else if (g[j] < 97 || g[j] > 122) || (g[j] < 48 || g[j] > 57) {
	// 		j--

	// 	} else if (g[i] >= 97 && g[i] <= 122) || (g[j] >= 97 && g[j] <= 122) || (g[i] >= 48 && g[i] <= 57) || (g[j] >= 48 && g[j] <= 57) {

	// 		if g[i] != g[j] {
	// 			fmt.Printf("g[i] %c  g[j] %c \n", g[i], g[j])
	// 			return false
	// 		} else {
	// 			i++
	// 			j--
	// 		}

	// 	}

	// }
	return true

}

func main() {
	// s := "madam"

	// gg := checkPalindrome(s)
	// fmt.Printf("gg: %v\n", gg)

	ss := "A man, a plan, a canal: Panama"

	k := checkForSentance(ss)
	fmt.Printf("k: %v\n", k)
}
