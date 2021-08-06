'''
Add two numbers as LinkedList
'''

# time complexity O(l1 l2)  space O(n)


def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
    dummy = ListNode(0)
    tmp = dummy
    carry = 0
    while (l1 != None or l2 != None or carry == 1):
        sum_ = 0
        if l1:
            sum_ += l1.val
            l1 = l1.next

        if l2:
            sum_ += l2.val
            l2 = l2.next

        sum_ += carry
        carry = sum_//10
        node = ListNode(sum_ % 10)
        tmp.next = node
        tmp = tmp.next

    return dummy.next
