# Rotate Linked List


# Time O(n) , O(1)

def rotateRight(self, head: ListNode, k: int) -> ListNode:
    if head == None or head.next == None:
        return head
    tmp1 = head
    count = 1
    while tmp1.next:
        count += 1
        tmp1 = tmp1.next
    # print(count)
    # print(tmp1)
    k = k % count
    tmp1.next = head

    run = count - k
    # print(run)
    # for _ in range(run):
    # print("out")
    #     tmp2 = tmp2.next
    # print(tmp2.next)
    # out = tmp2.next
    # tmp2.next = None
    # print(tmp2.next)
    # print(out)

    for _ in range(run):
        tmp1 = tmp1.next
    head = tmp1.next
    tmp1.next = None
    return head
