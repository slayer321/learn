package main

import "fmt"

func main() {
	arrayPrint()
}

func arrayPrint() {
	x := [...]int{1, 2, 3, 5}
	x[0] = 21
	fmt.Println(x, len(x))
	fmt.Printf("%T\n", x)
}
