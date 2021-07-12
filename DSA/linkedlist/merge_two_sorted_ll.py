'''
	Function to merge two sorted lists in one
	using constant space.
	
	Function Arguments: head_a and head_b (head reference of both the sorted lists)
	Return Type: head of the obtained list after merger.

	{
		# Node Class
		class Node:
		    def __init__(self, data):   # data -> value stored in node
		        self.data = data
		        self.next = None
	}

'''


def sortedMerge(head1, head2):
    # code here

    dummy = Node(0)
    tail = dummy

    while True:

        if head1 is None:
            tail.next = head2
            break
        if head2 is None:
            tail.next = head1
            break

        if head1.data <= head2.data:
            tail.next = head1
            head1 = head1.next
        else:
            tail.next = head2
            head2 = head2.next

        tail = tail.next

    return dummy.next
