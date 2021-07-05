'''
Implementing Queue form scratch using List
'''


class Queue:
    def __init__(self, capacity):
        self.front = self.size = 0
        self.capacity = capacity
        self.rear = capacity - 1
        self.Q = [None] * capacity

    def isFull(self):
        return self.size == self.capacity

    def isEmpty(self):
        return self.size == 0

    def enqueue(self, item):
        if self.isFull():
            print("Full")
            return
        self.rear = (self.rear + 1) % (self.capacity)
        self.Q[self.rear] = item
        self.size += 1
        print(f"{item} enqueued to queue")

    def dequeue(self):
        if self.isEmpty():
            print("Empty")
        print(f"{self.Q[self.front]} dequeued from queue")
        self.front = (self.front + 1) % self.capacity
        self.size -= 1

    def que_front(self):
        if self.isEmpty():
            print("Queue is empty")

        print(f"Front item is {self.Q[self.front]}")

    def que_rear(self):
        if self.isEmpty():
            print("Queue is Empty")

        print(f"Rear item is {self.Q[self.rear]}")

    def print_que(self):
        print(f"Queue is {self.Q}")


if __name__ == '__main__':

    queue = Queue(30)
    queue.enqueue(10)
    queue.enqueue(20)
    queue.enqueue(30)
    queue.enqueue(40)
    queue.dequeue()
    queue.que_front()
    queue.que_rear()
    queue.print_que()
