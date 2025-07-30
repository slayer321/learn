package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr

		curr = nxt
	}

	return prev
}

func main() {

}
