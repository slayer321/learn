package main

import "fmt"

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return "usr/name"
	}

	return c.path
}

type SomeStruct struct{}

func negate(s [5]int) {
	for i := range s {
		s[i] = -s[i]
	}
}

func main() {

	var a = [5]int{1, 2, 3, 4, 5}
	negate(a)
	fmt.Println(a) // prints [-1 -2 -3 -4 -5]
	// var c1 *Config
	// var c2 = &Config{
	// 	path: "my/home",
	// }

	// fmt.Println(c1.Path(), c2.Path())

	// var s string
	// var c complex128
	// var st struct{}
	// var i int16

	// type S struct {
	// 	a uint16
	// 	b uint32
	// }

	// var gg S

	// fmt.Println(unsafe.Sizeof(s))
	// fmt.Println(unsafe.Sizeof(c))
	// fmt.Println(unsafe.Sizeof(st))
	// fmt.Println(unsafe.Sizeof(i))
	// fmt.Println(unsafe.Sizeof(gg))

	// s := &SomeStruct{}
	// v := SomeStruct{}
	// g := &v
	// f := new(SomeStruct)

	// gg := make(map[int]string)
	// var jj map[string]string
	// fmt.Println(s)
	// fmt.Println(v)
	// fmt.Println(g)
	// fmt.Println(f)

	// fmt.Println(gg)
	// fmt.Println(jj)
	// fmt.Println(reflect.DeepEqual(gg, jj))

	// gg[1] = "sachin"
	// fmt.Println(gg)
	// jj["sachin"] = "raju"
	// fmt.Println(jj)

}
