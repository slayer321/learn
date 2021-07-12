'''
Intersection of Two Linked Lists
'''

# time complexity O(m+n) and space O(n)


def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
    tmp1 = headA
    dic = {}
    while tmp1:
        dic[tmp1] = 0
        # print(dic)
        tmp1 = tmp1.next
    arr = []
    cc = headA
    while cc:
        # print(cc)
        arr.append(cc.val)
        cc = cc.next

    print(dic)
    tmp2 = headB

    while tmp2:
        if tmp2 in dic:
            # print(tmp2)
            return tmp2
        tmp2 = tmp2.next
    # print(1)
