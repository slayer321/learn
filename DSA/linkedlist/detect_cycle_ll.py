
'''
Linked List Cycle
'''


# Time complexity O(n) , space O(n)

def hasCycle(self, head: ListNode) -> bool:
    dic = {}
      tmp = head
       while tmp:
            if tmp in dic:
                return True
            dic[tmp] = 0
            tmp = tmp.next

        return False



# Time complexity O(n) , space O(1)


def hasCycle(self, head: ListNode) -> bool:
        fast = head
        slow = head
        while fast and slow and fast.next:
            fast = fast.next.next
            slow = slow.next
            if fast == slow  :
                return True
            
        return False
