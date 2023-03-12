package main

import "testing"

func TestAllFunction(t *testing.T) {
	t.Run("Test oddNumber function", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{1, 3, 5, 7, 9}
		actual := oddNumber(input)
		if len(actual) != len(expected) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
	t.Run("Test primeNumber", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{2, 3, 5, 7}
		actual := primeNumber(input)
		if len(actual) != len(expected) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
	t.Run("Test oddPrimeNumber", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{3, 5, 7}
		actual := oddPrimeNumber(input)
		if len(actual) != len(expected) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
	t.Run("Test evenMul5", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		expected := []int{10, 20}
		actual := evenMul5(input)
		if len(actual) != len(expected) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	})
	t.Run("Test oddMul3", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		expected := []int{15}
		actual := oddMul3(input)
		if len(actual) != len(expected) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}

	})

}
