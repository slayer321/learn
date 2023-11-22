package main

import "fmt"

// const (
// 	Red int = iota
// 	Orange
// 	Yellow
// 	Green
// 	Blue
// 	Indigo
// 	Violet
// )

// const (
// 	Blue int = iota
// 	Green
// 	Indigo
// 	Orange
// 	Red
// 	Violet
// 	Yellow
// )

const (
	_ = iota
	a
	b
	c
	d
	e
	f
)

const (
	c0 = iota
	c1 = iota
	c2 = iota
)

const (
	c3 = iota
	c4
	c5
	c6
)

func main() {
	// 	fmt.Printf("The value of Red    is %v\n", Red)
	// 	fmt.Printf("The value of Orange is %v\n", Orange)
	// 	fmt.Printf("The value of Yellow is %v\n", Yellow)
	// 	fmt.Printf("The value of Green  is %v\n", Green)
	// 	fmt.Printf("The value of Blue   is %v\n", Blue)
	// 	fmt.Printf("The value of Indigo is %v\n", Indigo)
	// 	fmt.Printf("The value of Violet is %v\n", Violet)
	// }
	//fmt.Printf("%T %v\n", 'b', 'b')
	//printAll()
	printC()
}

func printC() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
}

func printAll() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
