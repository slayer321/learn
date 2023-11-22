package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, playground")
	//deferval()
	//byteBuffer()
	//fmt.Printf("foo %d\n", foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	// p1 := Person{
	// 	first: "sachin",
	// 	age:   30,
	// }
	// p1.speak()

	// s := Square{
	// 	length: 10,
	// 	width:  10,
	// }
	// c := Circle{
	// 	radius: 2,
	// }
	// ss := info(s)
	// fmt.Printf("Area of square is %f\n", ss)
	// cc := info(c)
	// fmt.Printf("Area of circle is %f\n", cc)

	// val := returnFunc()
	// fmt.Println(val())
}

// func deferval() {
// 	fmt.Println("start")
// 	defer fmt.Println("middle")
// 	fmt.Println("end")
// 	defer fmt.Println("end2")

// }

// func byteBuffer() {

// 	bf := bytes.NewBuffer([]byte("hello world"))
// 	fmt.Println(bf.String())

// 	bf.Reset()
// 	fmt.Println(bf.String())

// }

// func foo(vals ...int) int {
// 	sum := 0
// 	for _, val := range vals {
// 		sum += val
// 	}
// 	return sum
// }

// type Person struct {
// 	first string
// 	age   int
// }

// func (p Person) speak() {
// 	fmt.Println("I am ", p.first, " and I am ", p.age, " years old")
// }

// Interfaces

type Shape interface {
	Area() float64
}

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s Shape) float64 {
	return s.Area()
}

func returnFunc() func() int {
	return func() int {
		return 42
	}
}
