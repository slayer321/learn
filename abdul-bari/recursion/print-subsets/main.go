package main

import (
	"fmt"
)

func printSubset(ip, op string, out *[]string) {
	if len(ip) == 0 {
		// fmt.Println(op)
		*out = append(*out, op)
		return
	}

	op1 := op
	op2 := op + string(ip[0])

	printSubset(ip[1:], op1, out)
	printSubset(ip[1:], op2, out)
}

func printArraySubset(nums []int) [][]int {

	// ip := nums
	// gg := int(math.Pow(2, float64(len(nums))))
	// op := make([][]int, gg)

	// for i := range op {
	// 	op[i] = make([]int, 0)
	// }

	op := [][]int{}
	currentSubset := []int{}
	fmt.Printf("op: %v\n", op)
	return solve1(nums, currentSubset, op)

}

func solve1(ip []int, currentSubset []int, op [][]int) [][]int {

	if len(ip) == 0 {
		op = append(op, append([]int{}, currentSubset...))
		return op
	}

	// op1 := currentSubset
	// op2 := append(currentSubset, ip[0])

	op = solve1(ip[1:], currentSubset, op)
	op = solve1(ip[1:], append(currentSubset, ip[0]), op)
	return op

}

func main() {
	//k()
	out := printArraySubset([]int{9, 0, 3, 5, 7})
	fmt.Printf("out: %v\n", out)

	// out := []string{}

	// printSubset("aab", "", &out)

	// fmt.Printf("out: %v\n", out)

}

// func solve(ip []int, op [][]int) [][]int {
// 	if len(ip) == 0 {
// 		op = append(op, op...)
// 		return op
// 	}

// 	gg := len(ip) - (len(ip) - 1)
// 	// kk := ip[0]
// 	op1 := op
// 	op2 := append(op, append([]int(nil), ip[0:gg]...))

// 	op = solve(ip[1:], op1)
// 	op = solve(ip[1:], op2)
// 	return op
// }

// func k() {
// 	kk := [][]int{{}}
// 	// gg := []int{12, 3}
// 	// kk = append(kk, append([]int(nil), gg...))

// 	fmt.Printf("kk: %v\n", kk)
// }
