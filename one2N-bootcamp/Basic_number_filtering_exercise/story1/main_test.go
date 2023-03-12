package main

import "testing"

func TestEvenNumber(t *testing.T) {

	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenList := EvenNumber(numberList)
	if len(evenList) != 5 {
		t.Errorf("Expected length of evenList to be 5, but got %d", len(evenList))
	}
}
