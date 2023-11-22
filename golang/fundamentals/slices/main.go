package main

import "fmt"

func main() {
	//slicePrint()
	//deleteSlice()
	//exercise111()
	exercise112()
}

func slicePrint() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func deleteSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}

func exercise111() {
	x := make([]string, 0, 50)
	//y := []string{}

	fmt.Println(x, len(x), cap(x))
}

func exercise112() {
	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James"},
	}

	for _, v := range x {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
