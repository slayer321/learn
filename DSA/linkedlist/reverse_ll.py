'''
Reverse a LL
'''


def reverseList(self, head: ListNode) -> ListNode:
    if head is None or head.next is None:
        return head
    new_head = self.reverseList(head.next)
    head.next.next = head
    head.next = None
    return new_head


class Solution:
    # Function to reverse a linked list.
    def reverseList(self, head):
        # Code here
        prev = None

        curr = head
        while (curr != None):
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        return prev
