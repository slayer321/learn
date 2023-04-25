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

func SumIntsOrFloats[K comparable, V float64 | int64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum

}

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		fmt.Println("Stack is empty")
	}

	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func main() {
	// out := "sachin"

	// for _, val := range out {
	// 	//fmt.Println(val)
	// 	fmt.Printf("%x\n", val)
	// }
	// Initialize a map for the integer values
	// ints := map[string]int64{
	// 	"first":  34,
	// 	"second": 12,
	// }

	// // Initialize a map for the float values
	// floats := map[string]float64{
	// 	"first":  35.98,
	// 	"second": 26.99,
	// }

	// fmt.Printf("Generics %v, %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))

	s := make(Stack[int], 0)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	k := make(Stack[string], 0)
	k.Push("sachin")
	k.Push("kumar")
	fmt.Println(k.Pop())
	fmt.Println(k.Pop())
}
