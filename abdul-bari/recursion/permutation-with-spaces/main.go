package main

import (
	"fmt"
	"slices"
)

func permutationSpace(s string) []string {

	op := string(s[0])
	ip := s[1:]

	out := []string{}

	solve(ip, op, &out)
	slices.Sort(out)

	return out

}

func solve(ip, op string, out *[]string) {

	if len(ip) == 0 {
		// fmt.Println(op)
		*out = append(*out, op)
		return
	}

	op1 := op + "_" + string(ip[0])
	op2 := op + string(ip[0])

	solve(ip[1:], op1, out)
	solve(ip[1:], op2, out)

}

func main() {

	gg := permutationSpace("ABC")
	fmt.Printf("gg: %v\n", gg)
}
