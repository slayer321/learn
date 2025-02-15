package main

import "fmt"

func reverseAString(s *string) string {

	b := []byte(*s)
	i := 0
	j := len(*s) - 1

	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return string(b)
}

func main() {

	s := "sachin"
	gg := reverseAString(&s)
	fmt.Printf("gg: %v\n", gg)
}
