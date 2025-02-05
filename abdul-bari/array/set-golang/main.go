package main

import "fmt"

type Set struct {
	set map[int]struct{}
}

func NewSet() Set {
	return Set{
		set: make(map[int]struct{}),
	}
}

func (s Set) Add(val int) {
	s.set[val] = struct{}{}
}

func (s Set) Remove(val int) {
	delete(s.set, val)
}

func (s Set) Contains(val int) bool {

	_, found := s.set[val]
	return found

	// if s.set[val] != struct{}{} {
	// 	return true
	// }
	// return false
}

func (s Set) Size() int {
	return len(s.set)
}

func (s Set) List() []int {
	result := make([]int, 0)
	for k := range s.set {
		result = append(result, k)
	}
	return result
}

func (s Set) Union(other Set) Set {
	set := NewSet()
	for k := range other.set {
		if _, found := s.set[k]; !found {
			set.Add(k)
		}
	}

	for k := range s.set {
		set.Add(k)
	}
	return set
}

func (s Set) Intersection(other Set) Set {
	set := NewSet()

	for k := range other.set {
		if s.Contains(k) {
			set.Add(k)
		}
	}

	return set

}

func (s Set) Difference(other Set) Set {
	set := NewSet()

	for k := range s.set {
		if !other.Contains(k) {
			set.Add(k)
		}
	}

	return set

}

func main() {

	s := NewSet()

	s.Add(1)
	s.Add(2)
	s.Add(3)

	fmt.Printf("s.Size(): %v\n", s.Size())

	// s.Remove(1)
	// s.Remove(2)

	// fmt.Printf("s.Contains(\"sachin\"): %v\n", s.Contains(1))
	// fmt.Printf("s.Size(): %v\n", s.Size())

	unionSet := NewSet()
	for i := 0; i < 10; i++ {
		unionSet.Add(i)
	}

	newUnionSet := s.Union(unionSet)
	fmt.Printf("newUnionSet.Size(): %v\n", newUnionSet.Size())
	fmt.Printf("newUnionSet.List(): %v\n", newUnionSet.List())

	// fmt.Printf("s.Size(): %v\n", s.Size())

	// fmt.Printf("s.List(): %v\n", s.List())

	set2 := NewSet()
	set2.Add(2)
	set2.Add(1)

	fmt.Printf("set2.Intersection(set2): %v\n", s.Intersection(set2))

	fmt.Printf("s.Difference(set2): %v\n", s.Difference(set2))

}
