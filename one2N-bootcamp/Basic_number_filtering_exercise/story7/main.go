package main

import "fmt"

var Result []int

func oddNumber(numberList []int) []int {

	outList := make([]int, 0)
	for _, val := range numberList {
		if val%2 != 0 {
			outList = append(outList, val)
		}
	}
	return outList
}

func EvenNumber(numberList []int) []int {

	outVal := make([]int, 0)
	for _, val := range numberList {
		if val%2 == 0 {
			outVal = append(outVal, val)
		}
	}
	return outVal
}

func GreatorThan5(numberList []int) []int {

	outVal := make([]int, 0)
	for _, val := range numberList {
		if val > 5 {
			outVal = append(outVal, val)
		}
	}
	return outVal
}

func MultipleOf3(numberList []int) []int {

	outVal := make([]int, 0)
	for _, val := range numberList {
		if val%3 == 0 {
			outVal = append(outVal, val)
		}
	}
	return outVal
}

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

func dynamicFunction(inputString []string) {

	availableString := []string{"odd", "even", "gt5", "mul3", "prime"}
	for _, val := range inputString {
		if Contains(val, availableString) {
			// call the function
			switch val {
			case "odd":
				Result = oddNumber(Result)
			case "even":
				Result = EvenNumber(Result)
			case "gt5":
				Result = GreatorThan5(Result)
			case "mul3":
				Result = MultipleOf3(Result)
			case "prime":
				Result = primeNumber(Result)
			default:
				fmt.Println("No function found")
			}
		}

	}
	return
}

func Contains(s string, arr []string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func removeIndex(index int, slice []int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// odd, greater than 5, multiple of 3
func main() {
	// val := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// gg := removeIndex(3, val)
	// fmt.Println(gg)
	Result = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	inputString := []string{"odd", "gt5", "mul3", "prime"}
	dynamicFunction(inputString)
	fmt.Println(Result)
}
