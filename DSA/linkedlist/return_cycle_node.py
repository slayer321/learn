'''
Linked List Cycle II
'''

# Time O(n) , space O(n)


def detectCycle(self, head: ListNode) -> ListNode:
     dic = {}
      tmp = head
       while tmp:
            if tmp in dic:
                return tmp
            dic[tmp] = 0
            tmp = tmp.next
        return None



