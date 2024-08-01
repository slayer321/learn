package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	//node *Node
}

func (l *LinkedList) InsertEnd(val int) {

	if l.head == nil {
		l.head = &Node{
			data: val,
			next: nil,
		}
		return
	}

	tmp := l.head
	for {

		if tmp.next == nil {
			break
		}
		tmp = tmp.next

	}

	tmp.next = &Node{
		data: val,
		next: nil,
	}

}

func (ll *LinkedList) InsertAt(index int, val int) {

	tmp := ll.head
	count := 0
	for {
		count++
		if tmp.next == nil {
			break
		}
		// tmp = tmp.next

		if count == index {
			nextemp := tmp.next
			tmp.next = &Node{
				data: val,
				next: nil,
			}
			tmp.next.next = nextemp

		}

		// count++
		// if tmp.next == nil {
		// 	break
		// }
		tmp = tmp.next

	}

}

func (ll *LinkedList) InsertBeginneing(val int) {
	tmp := ll.head

	gg := &Node{
		data: val,
		next: tmp,
	}
	ll.head = gg
}

func (ll *LinkedList) PrintVal() {
	//fmt.Println("Inside the print statement")
	tmp := ll.head

	for {
		fmt.Printf("%d-->", tmp.data)
		if tmp.next == nil {
			break
		}

		tmp = tmp.next
	}
	fmt.Println("")
}

func main() {

	ll := &LinkedList{
		head: &Node{
			data: 22,
			next: nil,
		},
	}
	ll.InsertEnd(10)
	ll.InsertEnd(132)
	ll.InsertEnd(4)
	ll.InsertEnd(54)
	ll.InsertEnd(99)
	ll.PrintVal()
	ll.InsertBeginneing(101)
	ll.InsertAt(1, 333)
	ll.PrintVal()

}
