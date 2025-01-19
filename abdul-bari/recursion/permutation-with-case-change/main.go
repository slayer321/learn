package main

import (
	"fmt"
	"strings"
)

func permutationiWithCaseChange(s string) {

	op := ""
	ip := s

	solve(ip, op)

}

func solve(ip, op string) {

	if len(ip) == 0 {
		fmt.Println(op)
		return
	}

	op1 := op + strings.ToLower(string(ip[0]))
	op2 := op + strings.ToUpper(string(ip[0]))

	solve(ip[1:], op1)
	solve(ip[1:], op2)

}

func main() {

	permutationiWithCaseChange("AB")

}
