package main

// Time complexitity O(2^n)
// It take 2^n -1 moves .. here n is the number of blocks
func toh(n, A, B, C int) int {

	if n == 1 {
		//fmt.Printf("from %d to %d\n", A, C)

		return 1
	}

	return toh(n-1, A, C, B)
	// fmt.Printf("from %d to %d\n", A, C)
	return toh(n-1, B, A, C)
}

func main() {

	toh(3, 1, 2, 3)

}
