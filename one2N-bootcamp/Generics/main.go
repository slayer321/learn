package main

import "fmt"

func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}

func main() {
	out := "sachin"

	for _, val := range out {
		//fmt.Println(val)
		fmt.Printf("%x\n", val)
	}
}
