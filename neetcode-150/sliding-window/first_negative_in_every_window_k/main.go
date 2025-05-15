package main

import "fmt"

func firstNegInt(arr []int, k int) []int {
	i := 0
	j := 0

	result := []int{}
	firstNegative := 0

	for j < len(arr) {
		// if (arr[j] < 0 || arr[i] < 0) && firstNegative >= 0 {
		// 	if arr[i] < 0 {
		// 		firstNegative = arr[i]
		// 	} else {
		// 		firstNegative = arr[j]
		// 	}
		// }

		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == k {
			k := i
			for k <= j {
				if arr[k] < 0 && firstNegative >= 0 {
					firstNegative = arr[k]
				}
				k++
			}

			result = append(result, firstNegative)
			i++
			j++
			firstNegative = 0
		}
	}
	return result
}

func firstNegIntOptimize(arr []int, k int) []int {

	i := 0
	j := 0

	negativeSlice := []int{}
	result := []int{}

	for j < len(arr) {

		if arr[j] < 0 {
			negativeSlice = append(negativeSlice, arr[j])
		}

		if (j - i + 1) < k {
			j++
		} else if (j - i + 1) == k {
			if len(negativeSlice) == 0 {
				result = append(result, 0)
			} else {
				result = append(result, negativeSlice[0])
			}
			if arr[i] < 0 {
				negativeSlice = negativeSlice[1:]
			}
			i++
			j++
		}
	}

	return result

}

func main() {

	arr := []int{-8, 2, 3, -6, 10}
	k := 2

	out := firstNegIntOptimize(arr, k)
	fmt.Printf("out: %v\n", out)

}
