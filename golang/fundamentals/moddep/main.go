package main

import "fmt"

var name string = "Sachin"

const age = 23

func main() {
	exercise14()
}

func exercise14() {
	var college string = "IIT Bombay"
	marks := 9.5

	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v, Type %T\n", age, age)
	fmt.Printf("College: %v\n", college)
	fmt.Printf("Marks: %v\n", marks)
}
