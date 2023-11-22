package main

import "fmt"

func main() {
	//testMap()
	//exercise118()
	exercise121()
}

func testMap() {

	m := make(map[string]string)

	fmt.Println(m)
	fmt.Println(len(m))

	m["name"] = "sachin"
	m["age"] = "30"

	fmt.Printf("Type is %T, value %s\n", m, m)
	delete(m, "age")
	//fmt.Println(cap(m))
	fmt.Println(m)
	delete(m, "age")
	val, ok := m["age"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("key not found")
	}
}

func exercise118() {
	m := map[string][]string{
		"bond_james":      []string{"shakens", "not stirred", "martinis"},
		"moneypenny_miss": []string{"james bond", "literature", "computer science"},
		"no_dr":           []string{"being evil", "ice cream", "sunsets"},
	}
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	delete(m, "no_dr")
	for k, v := range m {
		for i, val := range v {
			fmt.Println(k, i, val)
		}
	}
}

func exercise121() {
	words := []string{"alpha", "alpha", "beta", "gamma", "alpha", "beta", "gamma", "zeta", "eta", "theta"}
	m := make(map[string]int)
	for _, word := range words {
		if m[word] == 0 {
			m[word]++
		} else {
			m[word]++
		}
	}
	fmt.Println(m)
}
