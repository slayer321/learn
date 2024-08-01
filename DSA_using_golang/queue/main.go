package main

import "fmt"

type Queue struct {
	store []int
}

func (q *Queue) Add(val int) []int {
	q.store = append(q.store, val)
	//fmt.Println(q.store)
	return q.store

}

func (q *Queue) Remove() []int {

	q.store = q.store[1:]
	//fmt.Println(q.store)
	return q.store

}

func main() {

	queue := Queue{
		store: make([]int, 0, 10),
	}

	f := queue.Add(10)
	fmt.Println(f)

	f = queue.Add(32)
	fmt.Println(f)

	f = queue.Add(99)
	fmt.Println(f)

	f = queue.Add(44)
	fmt.Println(f)

	f = queue.Remove()
	fmt.Println(f)

	f = queue.Add(78)
	fmt.Println(f)

	f = queue.Remove()
	fmt.Println(f)

}
