package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data int
}

type DLL struct {
	head *Node
}

func (dll *DLL) Insert(val *Node) {

	if dll.head == nil {
		dll.head = val
		return
	}

	tmp := dll.head

	for tmp.next != nil {

		tmp = tmp.next
	}
	tmp.next = val
	val.prev = tmp

}

func (dll *DLL) Delete() {

	if dll.head == nil {
		return
	}
	if dll.head.next == nil {
		dll.head = nil
		return
	}

	tmp := dll.head

	for tmp.next.next != nil {
		tmp = tmp.next
	}

	tmp.next = nil

}

func (dll *DLL) InsetAtIdx(val *Node, idx int) {

	if idx == 0 {
		newHead := val
		newHead.next = dll.head
		dll.head.prev = newHead

		dll.head = newHead

		return
	}

	count := 0

	tmp := dll.head

	for tmp != nil {

		if count == idx {
			currPrev := tmp.prev
			tmp.prev = val
			currPrev.next = val
			val.next = tmp

		}
		count++
		tmp = tmp.next
	}
}

func (dll *DLL) Display() {
	tmp := dll.head

	for tmp != nil {
		fmt.Printf("%v->", tmp.data)
		tmp = tmp.next
	}

	fmt.Println()
}

func main() {
	head := DLL{}
	head.Insert(&Node{data: 10})
	head.Insert(&Node{data: 23})
	head.Insert(&Node{data: 45})

	head.Display()

	head.InsetAtIdx(&Node{data: 100}, 2)

	head.Display()

}
