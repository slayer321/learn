package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ByteSize int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	//oddEven()
	//printBit()
	//size()
	// exercise10()
	//exercise11()
	//exercise12()
	exercise13()
}

func size() {
	fmt.Printf("%d \t %b\n", KB, KB)
	fmt.Printf("%d \t %b\n", MB, MB)
	fmt.Printf("%d \t %b\n", GB, GB)
	fmt.Printf("%d \t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}

func oddEven() {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Println(num, "is odd")
		} else {
			fmt.Println(num, "is even")
		}
	}
}

func printBit() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
}

func exercise10() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)

}

func exercise11() {
	str := "Hello World"
	intVal := 10
	floatVal := 3.14

	fmt.Printf("%T %v\n", str, str)
	fmt.Printf("%T %v\n", intVal, intVal)
	fmt.Printf("%T %v\n", floatVal, floatVal)
}

func exercise12() {
	d := 747
	b := 911
	hex := 90210

	fmt.Printf("%d\n", d)
	fmt.Printf("%#b\n", b)
	fmt.Printf("%#x\n", hex)
}

func exercise13() {
	var maxInt8 int8 = math.MaxInt8
	var minInt8 int8 = math.MinInt8

	fmt.Printf("Maximum value for int8: %d\n", maxInt8)
	fmt.Printf("Minimum value for int8: %d\n", minInt8)

	var first int8 = 127
	var second uint8 = 255

	fmt.Printf("first %d\n", first)
	fmt.Printf("second %d\n", second)
}
