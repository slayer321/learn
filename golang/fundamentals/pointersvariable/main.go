package main

// https://stackoverflow.com/questions/50830676/set-an-int-pointer-to-an-int-valu
import (
	"fmt"
	"reflect"
)

func main() {
	//failed()
	//worked()
	newExample()
}

func worked() {
	var guess *int
	fmt.Println(reflect.TypeOf(guess))
	guess = new(int)
	fmt.Printf("%v === %v\n", guess, *guess)
	*guess = 12345
	fmt.Printf("%v === %v\n", guess, *guess)
}

func failed() {
	var guess *int
	fmt.Println(reflect.TypeOf(guess))
	*guess = 12345
	fmt.Println(guess)
}

type Vehicle struct {
	brand          string
	name           string
	numberOfWheels int
}

func newExample() {
	//new returns a pointer to an instance of the Vehicle struct
	//var vehicle = new(Vehicle)
	vehicle := Vehicle{}
	fmt.Println(vehicle)
	//setting values of the fields
	vehicle.brand = "Tata Motors"
	vehicle.name = "Nexxon"
	vehicle.numberOfWheels = 4

	fmt.Printf("Before anotherFunc: %v\n", vehicle)
	anotherFunc(vehicle)
	//printing Vehicle

	fmt.Printf("After anotherFunc: %v\n", vehicle)
}

func anotherFunc(vehicle Vehicle) {
	//setting values of the fields
	vehicle.numberOfWheels = 12
}
