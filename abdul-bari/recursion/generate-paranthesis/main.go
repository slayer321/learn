package main

//https://leetcode.com/problems/generate-parentheses/description/
import "fmt"

func generateParanthesis(n int) []string {

	o := n
	c := n
	op := ""
	result := make([]string, 0)

	return solve(o, c, op, result)
}

func solve(o, c int, op string, result []string) []string {

	if o == 0 && c == 0 {
		// fmt.Println(op)
		result = append(result, op)
		return result
	}

	if o == c {
		op1 := op + "("
		result = solve(o-1, c, op1, result)
	} else if o != 0 && o < c {
		op1 := op + "("
		op2 := op + ")"
		result = solve(o-1, c, op1, result)
		result = solve(o, c-1, op2, result)
	} else if o == 0 {
		op1 := op + ")"
		result = solve(o, c-1, op1, result)
	}
	return result

}
func main() {
	result := generateParanthesis(3)
	fmt.Printf("result: %v\n", result)
}
