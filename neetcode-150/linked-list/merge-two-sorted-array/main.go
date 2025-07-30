package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	t1 := list1
	t2 := list2
	result := &ListNode{}
	dummy := result

	for t1 != nil && t2 != nil {
		if t1.Val <= t2.Val {
			dummy.Next = t1
			dummy = t1
			t1 = t1.Next
		} else if t1.Val >= t2.Val {
			dummy.Next = t2
			dummy = t2
			t2 = t2.Next
		}
	}

	if t1 != nil {
		dummy.Next = t1
	} else {
		dummy.Next = t2
	}

	return result.Next

}
