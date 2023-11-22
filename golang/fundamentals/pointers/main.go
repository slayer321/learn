package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	// d1 := dog{"Henry"}
	// d1.walk()
	// d1.run()
	// //youngRun(d1) // doesn't run

	// d2 := &dog{"Padget"}
	// d2.walk()
	// d2.run()
	// youngRun(d2)
	defrefPointers()

}

type first struct {
	name string
}

func firstFun() first {
	f := first{"sachin"}
	fmt.Println(f.name)
	return f

}

func secondFun() *first {
	f := &first{"raju"}
	fmt.Println(f.name)
	return f
}

func defrefPointers() {
	var x *int
	var y *string
	fmt.Printf("Type : %T, memory: %x\n", x, x)
	fmt.Printf("Type : %T, memory: %x\n", y, y)

	x = new(int)
	y = new(string)

	fmt.Printf("Type : %T, memory: %#x\n", x, x)
	fmt.Printf("Type : %T, memory: %#x\n", y, y)

	*x = 42
	*y = "James Bond"

	fmt.Printf("Type : %T, memory: %d\n", x, *x)
	fmt.Printf("Type : %T, memory: %s\n", y, *y)

}
