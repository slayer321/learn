# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


'''
Remove Nth Node From End of Linked List | Amazon | Microsoft | Adobe

'''
# Time complexity O(2n) , space O(1)

# def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
#     dummy = ListNode(0)
#     dummy.next = head
#      first = head
#       count = 0
#        while first:
#             count += 1
#             first = first.next

#         count -= n
#         first = dummy
#         while count:
#             count -= 1
#             first = first.next

#         first.next = first.next.next

#         return dummy.next


# Time complexity O(n) , space O(1)

def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
    dummy = ListNode(0)
    dummy.next = head
    fast = dummy
    slow = dummy
    for i in range(n):
        fast = fast.next

    while fast and fast.next:
        slow = slow.next
        fast = fast.next

    slow.next = slow.next.next

    return dummy.next
