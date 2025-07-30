package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList(head *Node) *LinkedList {
	return &LinkedList{
		head: head,
	}
}

func (ll *LinkedList) Insert(node *Node) {

	tmp := ll.head
	for tmp.next != nil {
		tmp = tmp.next

	}

	tmp.next = node

}

func (ll *LinkedList) InsertFront(node *Node) {

	node.next = ll.head

	ll.head = node

}

func (ll *LinkedList) InsertBetween(node *Node, idx int) {

	if idx == 0 {
		node.next = ll.head
		ll.head = node
		return
	}

	i := 0

	tmp := ll.head
	var prev *Node

	for tmp != nil {

		if i == idx {
			prev.next = node
			node.next = tmp
		}

		i++
		prev = tmp
		tmp = tmp.next
	}
}

func (ll *LinkedList) Display() {

	tmp := ll.head

	for tmp != nil {
		data := tmp.data
		fmt.Printf("%d-> ", data)
		tmp = tmp.next
	}
	fmt.Println()
}

func main() {

	ll := LinkedList{
		head: &Node{
			data: 12,
		},
	}

	ll.Insert(&Node{data: 31})
	ll.Insert(&Node{data: 32})
	ll.Insert(&Node{data: 43})
	ll.Insert(&Node{data: 56})

	ll.Display()

	ll.InsertBetween(&Node{data: 100}, 1)

	// ll.InsertFront(&Node{data: 100})

	ll.Display()

}
