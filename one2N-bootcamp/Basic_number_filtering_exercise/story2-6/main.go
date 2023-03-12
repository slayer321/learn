package main

func oddNumber(numberList []int) []int {

	outList := make([]int, 0)
	for val := range numberList {
		if val%2 != 0 {
			outList = append(outList, val)
		}
	}
	return outList
}

// 1,2,3,4,5,6,7,8,9,10
func primeNumber(numberList []int) []int {
	outList := make([]int, 0)
	var flag int

	// Here if the flag is zero then the number is prime
	for _, val := range numberList {
		flag = 0
		for i := 2; i <= val/2; i++ {
			if val%i == 0 {
				flag = 1
				break
			}

		}
		if flag == 0 && val > 1 {
			outList = append(outList, val)
		}

	}
	return outList
}

func oddPrimeNumber(numberList []int) []int {
	outList := make([]int, 0)
	var flag int

	// Here if the flag is zero then the number is prime
	for _, val := range numberList {
		flag = 0
		for i := 2; i <= val/2; i++ {
			if val%i == 0 {
				flag = 1
				break
			}

		}
		if flag == 0 && val > 1 && val%2 != 0 {
			outList = append(outList, val)
		}

	}
	return outList
}

func evenMul5(arr []int) []int {

	outVal := make([]int, 0)

	for _, val := range arr {
		if val%2 == 0 && val%5 == 0 {
			outVal = append(outVal, val)
		}
	}

	return outVal
}

func oddMul3(arr []int) []int {
	outVal := make([]int, 0)

	for _, val := range arr {
		if val%2 != 0 && val%3 == 0 && val > 10 {
			outVal = append(outVal, val)
		}
	}

	return outVal
}

func main() {

}
