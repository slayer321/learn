package main

//https://leetcode.com/problems/letter-case-permutation/description/
import (
	"fmt"
	"strconv"
	"strings"
)

func letterCasePermutation(s string) []string {

	op := ""
	ip := s
	//counter := 0

	out := make([]string, 0)
	return solve(ip, op, out)

}

func solve(ip, op string, out []string) []string {

	if len(ip) == 0 {
		// fmt.Println(op)//
		out = append(out, op)
		return out
	}

	if _, err := strconv.ParseInt(string(ip[0]), 10, 64); err != nil {
		op1 := op + strings.ToLower(string(ip[0]))
		op2 := op + strings.ToUpper(string(ip[0]))

		out = solve(ip[1:], op1, out)
		out = solve(ip[1:], op2, out)

	} else {
		op1 := op + string(ip[0])
		out = solve(ip[1:], op1, out)

	}

	return out

	// var op1, op2, opp1, opp2 string
	// //var op2 string

	// for ; counter < len(ip); counter++ {
	// 	fmt.Println(counter)
	// 	_, err := strconv.ParseInt(string(ip[counter]), 10, 32)

	// 	if err != nil {
	// 		break
	// 		// op1 = op + strings.ToLower(string(ip[i]))
	// 		// op2 = op + strings.ToUpper(string(ip[0]))
	// 	}
	// 	opp1 = op + strings.ToLower(string(ip[counter]))
	// 	opp2 = op + strings.ToUpper(string(ip[counter]))
	// }
	// op1 = opp1 + strings.ToLower(string(ip[counter]))
	// op2 = opp2 + strings.ToUpper(string(ip[counter]))

	// solve(ip[counter+1:], op1, counter+1)
	// solve(ip[counter+1:], op2, counter+1)

}

func main() {

	gg := "a1b2"

	result := letterCasePermutation(gg)
	fmt.Printf("result: %v\n", result)
}
