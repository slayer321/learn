from typing import Counter


class Node:
    def __init__(self, data=None, next=None):
        self.data = data
        self.next = next


class Linkedlist:
    def __init__(self):
        self.head = None

    def insert_at_begining(self, data):
        node = Node(data, self.head)
        self.head = node

    def print_value(self):
        if self.head is None:
            print("Linkedlist is empty")
            return

        itr = self.head
        # print(itr.data)
        llstr = ''
        while (itr):
            # print(itr)
            llstr += str(itr.data) + '-->'
            itr = itr.next
        print(llstr)

    def insert_end(self, data):
        if self.head is None:
            self.head = Node(data, None)
            return

        tmp = self.head
        #print("start:", tmp.data, tmp.next)
        while (tmp.next):
            tmp = tmp.next
        #print("end", tmp.data, tmp.next)
        tmp.next = Node(data, None)

    def insert_value(self, data_list):
        self.head = None
        for data in data_list:
            self.insert_end(data)

    def get_length(self):
        count = 0
        tmp = self.head
        while tmp:
            count += 1
            tmp = tmp.next
        return count

    def remove_index(self, index):
        if index < 0 or index >= self.get_length():
            print("Out of bound")
            return
        count = 0
        tmp = self.head
        while tmp:

            if count == index - 1:
                tmp.next = tmp.next.next
                break
            tmp = tmp.next
            count += 1

        if index == 0:
            self.head = self.head.next

    def insert_at(self, index, data):
        if index < 0 or index >= self.get_length():
            print("Out of bound")
            return

        if index == 0:
            node = Node(data, self.head)
            self.head = node

        count = 0
        tmp = self.head
        while (tmp):

            if count == index - 1:
                tmp.next = Node(data, tmp.next)
                break
            tmp = tmp.next
            count += 1


if __name__ == "__main__":
    ll = Linkedlist()
    ll.insert_at_begining(23)
    ll.insert_at_begining(43)
    ll.print_value()
    ll.insert_end(65)
    ll.insert_end(5)
    #ll.insert_value([2, 2])
    # ll.remove_index(3)
    ll.insert_at(0, 444)
    ll.print_value()
