package main

import (
	"fmt"
)

func main() {
	//exercise122()
	//exercise123()
	//exercise124()
	//exercise125()
	//exercise126()
	//exercise127()
	//exercise128()
	exercise129()
}

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func exercise126() {
	sachin := person{
		firstName:   "sachin",
		lastName:    "kumar",
		favIceCream: []string{"vanilla", "chocolate", "strawberry"},
	}
	vishal := person{
		firstName:   "sachin",
		lastName:    "kumar",
		favIceCream: []string{"vanilla", "chocolate", "strawberry"},
	}

	for _, ice := range sachin.favIceCream {
		fmt.Println(ice)
	}

	for _, ice := range vishal.favIceCream {
		fmt.Println(ice)
	}
}

func exercise127() {
	sachin := person{
		firstName:   "sachin",
		lastName:    "kumar",
		favIceCream: []string{"vanilla", "chocolate", "strawberry"},
	}
	vishal := person{
		firstName:   "vishal",
		lastName:    "maurya",
		favIceCream: []string{"vanilla", "chocolate", "strawberry"},
	}
	m := make(map[string]person)

	m["kumar"] = sachin
	m["maurya"] = vishal

	for _, v := range m {
		fmt.Println(v)
	}
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func exercise128() {
	v := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "Model S",
		doors: 4,
		color: "black",
	}
	fmt.Println(v)

}

func exercise129() {
	val := []struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		{first: "sachin",
			friends: map[string]int{
				"vishal": 123,
				"rahul":  456,
			},
			favDrinks: []string{"coke", "pepsi"},
		},
	}

	fmt.Println(val)
}
