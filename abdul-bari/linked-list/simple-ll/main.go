package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedList struct {
	head *Node
}

func (ll *linkedList) Insert(val int) {

	if ll.head == nil {
		ll.head = &Node{
			data: val,
		}
		return
	}

	tmp := ll.head
	for tmp.next != nil {
		tmp = tmp.next

	}

	tmp.next = &Node{
		data: val,
	}

}

func (ll *linkedList) InsertAtFront(val int) {

	node := &Node{
		data: val,
	}

	if ll.head == nil {
		fmt.Println("linkedin list is empty")
	}

	currentHead := ll.head

	node.next = currentHead

	ll.head = node

}

func (ll *linkedList) DeleteNodeWithVal(val int) {

	if ll.head == nil {
		fmt.Println("linkedin list is empty")
	}

	tmp := ll.head
	var prev *Node

	for tmp != nil {
		if tmp.data == val {
			prev.next = tmp.next
			tmp.next = nil
			break
		}

		prev = tmp
		tmp = tmp.next
	}
}

func (ll *linkedList) Display() {
	tmp := ll.head
	for tmp.next != nil {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
	fmt.Println(tmp.data)

	fmt.Println("------------")
}

func main() {

	ll := linkedList{}
	ll.Insert(10)
	ll.Insert(1)
	ll.Insert(43)
	ll.Insert(23)
	ll.Display()
	ll.InsertAtFront(111)
	ll.Display()
	ll.DeleteNodeWithVal(43)
	ll.Display()

}
