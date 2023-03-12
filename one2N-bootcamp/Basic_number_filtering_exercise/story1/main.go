package main

import (
	"fmt"
)

func EvenNumber(numberList []int) []int {

	//outVal := []int{}
	outVal := make([]int, 0)

	// gg_val, _ := json.Marshal(gg)
	// var myslice []int

	//myslice, _ := json.Marshal(myslice)
	for _, val := range numberList {
		if val%2 == 0 {
			outVal = append(outVal, val)
		}
	}
	// fmt.Printf("%T", gg_val)
	// fmt.Println(myslice)
	return outVal
}

func main() {

	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenList := EvenNumber(numberList)
	fmt.Println(evenList)
}
