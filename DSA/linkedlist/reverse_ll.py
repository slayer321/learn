# class Node:
#     def __init__(self, data=None, next=None):
#         self.data = data
#         self.next = next


# class Linkedlist:
#     def __init__(self):
#         self.head = None

#     def insert_end(self, data):
#         if self.head is None:
#             self.head = Node(data, None)
#             return

#         tmp = self.head
#         #print("start:", tmp.data, tmp.next)
#         while (tmp.next):
#             tmp = tmp.next
#         #print("end", tmp.data, tmp.next)
#         tmp.next = Node(data, None)

#     def print_value(self):
#         if self.head is None:
#             print("Linkedlist is empty")
#             return

#         itr = self.head
#         # print(itr.data)
#         llstr = ''
#         while (itr):
#             # print(itr)
#             llstr += str(itr.data) + '-->'
#             itr = itr.next
#         print(llstr)

#     def reverse_ll(self,):
#         tmp = self.head
#         count = 0
#         while tmp:
#             tmp = tmp.next
#             if tmp.next == None:
