class Circularqueue:
    def __init__(self, k):
        self.k = k
        self.queue = [None]*k
        self.head = self.rear = -1

    def enqueue(self, data):
        if self.rear == len(self.queue) - 1:
            print("Queue is full\n")
        elif (self.head == -1):
            self.head = 0
            self.rear = 0
            self.queue[self.rear] = data
        else:
            self.rear = (self.rear + 1) % self.k
            self.queue[self.rear] = data

    def dequeue(self):
        if self.head == -1:
            print("Queue is empty")
        elif (self.head == self.rear):
            temp = self.queue[self.head]
            self.head = -1
            self.rear = -1
            return temp
        else:
            temp = self.queue[self.head]
            self.head = (self.head + 1) % self.k
            return temp

    def printCQueue(self):
        if(self.head == -1):
            print("No element in the circular queue")

        elif (self.rear >= self.head):
            for i in range(self.head, self.rear + 1):
                print(self.queue[i], end=" ")
            print()
        else:
            for i in range(self.head, self.k):
                print(self.queue[i], end=" ")
            for i in range(0, self.tail + 1):
                print(self.queue[i], end=" ")
            print()


obj = Circularqueue(5)
obj.enqueue(1)
obj.enqueue(2)
obj.enqueue(3)
obj.enqueue(4)
obj.enqueue(5)
print("Initial queue")
obj.printCQueue()

obj.dequeue()
print("After removing an element from the queue")
obj.printCQueue()
