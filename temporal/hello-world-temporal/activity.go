package app

import (
	"context"
)

func OddNumber(ctx context.Context, numberList []int) ([]int, error) {
	var result []int
	for _, number := range numberList {
		if number%2 != 0 {
			result = append(result, number)
		}
	}
	return result, nil
}

func EvenNumber(ctx context.Context, numberList []int) ([]int, error) {
	var result []int
	for _, number := range numberList {
		if number%2 == 0 {
			result = append(result, number)
		}
	}
	return result, nil
}
