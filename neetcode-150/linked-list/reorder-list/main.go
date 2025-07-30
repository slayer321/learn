package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderListNotWorking(head *ListNode) {

	arr := []int{}
	tmp := head
	count := 0
	for tmp != nil {
		fmt.Printf("tmp.Val: %v\n", tmp.Val)
		if count%2 != 0 {
			arr = append(arr, tmp.Val)
		}

		count++
		tmp = tmp.Next
	}

	fmt.Println(arr)

	dummy := &ListNode{}
	// dummy := helper
	newCounter := 0
	gg := head
	for gg != nil {
		lastVal := len(arr) - 1

		fmt.Println("counter", newCounter)
		if newCounter%2 != 0 {
			fmt.Println("arr", arr[lastVal])
			nextE := gg
			dummy.Next = &ListNode{
				Val:  arr[lastVal],
				Next: nextE,
			}
			dummy = dummy.Next
			arr = arr[:lastVal]
			fmt.Println(arr)
		} else {
			dummy.Val = gg.Val
		}
		gg = gg.Next

		newCounter++
		// helper = helper
	}
	head = dummy.Next

	for dummy != nil {
		fmt.Print("-> ", dummy.Val)
		dummy = dummy.Next
	}
}

func reorderList(head *ListNode) {

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	prev := &ListNode{}
	slow.Next = nil
	prev = prev.Next

	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	first := head
	second = prev

	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}

	for head != nil {
		fmt.Print("-> ", head.Val)
		head = head.Next
	}
}
func main() {

	head := &ListNode{
		Val: 1,
	}

	head.Next = &ListNode{
		Val: 2,
	}

	head.Next.Next = &ListNode{
		Val: 3,
	}

	head.Next.Next.Next = &ListNode{
		Val: 4,
	}

	reorderList(head)

}
