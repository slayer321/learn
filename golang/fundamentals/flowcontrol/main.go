package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("This will be called on main initialization")
}

func main() {
	//workingSelect()
	//exercise74()
	//exercise75()
	//exercise77()
	//exercise79()
	//exercise84()
	exercise87()
}

func workingSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 41
	}()

	select {
	case val1 := <-ch1:
		println("ch1 val1:", val1)
	case val2 := <-ch2:
		println("ch2 val2:", val2)
	}
}

func exercise74() {
	x := rand.Intn(250)
	fmt.Printf("x is %v\n", x)

	if x > 0 && x < 100 {
		fmt.Printf("x is between 0 and 100 %d\n", x)
	} else if x > 101 && x < 200 {
		fmt.Printf("x is between 101 and 200 %d\n", x)
	} else if x > 201 && x < 250 {
		fmt.Printf("x is between 201 and 250 %d\n", x)
	}
}

func exercise75() {
	x := rand.Intn(250)
	fmt.Printf("x is %v\n", x)

	switch {
	case x <= 100:
		fmt.Printf("x is between 0 and 100 %d\n", x)
	case x > 101 && x <= 200:
		fmt.Printf("x is between 101 and 200 %d\n", x)
	case x > 201 && x <= 250:
		fmt.Printf("x is between 201 and 250 %d\n", x)
	default:
		fmt.Printf("x is not between 0 and 250 %d\n", x)
	}
}

func exercise77() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The value of x and y is %d and %d\n", x, y)
	if x < 4 && y < 4 {
		fmt.Printf("x and y are less than 4 %d , %d\n", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("x and y are greater than 6 %d , %d\n", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is greater than 4 and y is greater than 6 %d , %d\n", x, y)
	} else if y != 5 {
		fmt.Printf("y is not equal to 5 %d , %d\n", x, y)
	} else {
		fmt.Printf("None of the above cases %d , %d\n", x, y)
	}

}

func exercise79() {
	// for i := 0; i <= 100; i++ {
	// 	fmt.Printf("The number is %d\n", i)
	// }

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("iteration %v \t x and y are %v and %v\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("both are less than four")
		case x > 6 && y > 6:
			fmt.Println("both are greater than six")
		case x >= 4 && x <= 6:
			fmt.Println("x is from 4 to 6 inclusive of both numbers")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous were met")
		}
	}
}

func exercise84() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("The outer loop is %d\n", i)
		for j := 0; j <= 5; j++ {
			fmt.Printf("\t\tThe inner loop is %d\n", j)
		}
	}
}

func exercise87() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}
	for k, v := range m {
		fmt.Printf("The key is %v and the value is %v\n", k, v)
	}

	val, ok := m["Q"]
	if ok {
		fmt.Printf("The value is %v\n", val)
	} else {
		fmt.Println("The value is not present")
	}
}
