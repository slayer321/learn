package main

import "fmt"

func nextGreatestletters(letters []byte, target byte) byte {

	l := 0
	r := len(letters) - 1
	res := byte('{')

	for l <= r {
		mid := l + (r-l)/2

		if letters[mid] == target {
			// res = min(res, letters[mid])
			l = mid + 1
		}

		if target > letters[mid] {
			l = mid + 1
		} else if target < letters[mid] {
			res = letters[mid]
			r = mid - 1
		}

	}

	if res > byte('z') {
		return letters[0]
	}

	return res

}

func main() {

	letters := []byte{'c', 'f', 'j'}
	target := byte('c')

	out := nextGreatestletters(letters, target)
	fmt.Printf("out: %v\n", string(out))

}
