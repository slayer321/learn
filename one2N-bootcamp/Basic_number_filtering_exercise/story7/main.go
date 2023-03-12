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

func EvenNumber(numberList []int) []int {

	outVal := make([]int, 0)
	for _, val := range numberList {
		if val%2 == 0 {
			outVal = append(outVal, val)
		}
	}
	return outVal
}

func dynamicFunction(inputString []string) {

}

func main() {

}
