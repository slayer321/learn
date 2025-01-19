package main

//https://www.geeksforgeeks.org/problems/print-n-bit-binary-numbers-having-more-1s-than-0s0252/1
import "fmt"

// My solution
func NBitBinary(n int) {

	one := n
	zero := n
	op := ""
	counter := 0 // height of the tree

	solve(one, zero, op, counter, n)

}

func solve(one, zero int, op string, counter, n int) {

	if counter == n {
		fmt.Println(op)
		return
	}

	if one == zero {
		op1 := op + "1"
		solve(one-1, zero, op1, counter+1, n)
	} else if one < zero {
		op1 := op + "1"
		op2 := op + "0"

		solve(one-1, zero, op1, counter+1, n)
		solve(one, zero-1, op2, counter+1, n)
	}

}

// Adtiya Verma video Solution

func NBitBinaryAdtiya(n int) {
	one := 0
	zero := 0
	op := ""

	solve1(one, zero, n, op)

}

func solve1(one, zero, n int, op string) {
	if n == 0 {
		fmt.Println(op)
		return
	}

	if one == zero {
		op1 := op + "1"
		solve1(one+1, zero, n-1, op1)
	} else if one > zero {
		op1 := op + "1"
		op2 := op + "0"

		solve1(one+1, zero, n-1, op1)
		solve1(one, zero+1, n-1, op2)
	}
}

func main() {
	// NBitBinary(4)
	NBitBinaryAdtiya(3)
}
