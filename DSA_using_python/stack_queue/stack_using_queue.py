# https: // leetcode.com/problems/implement-stack-using-queues/


class Mystack:
    def __init__(self, arr):
        self.arr = arr

    def push(self, data):
        self.arr.append(data)
        for i in range(len(self.arr)-1):
            self.arr.append(self.top())
            self.arr.pop(0)

    def pop(self):
        ans = self.arr.pop(0)
        return ans

    def top(self):
        val = self.arr[0]
        return val

    def isempty(self):
        if len(self.arr) == 0:
            return True
        return False


obj = Mystack([])
obj.push(3)
obj.push(2)
obj.push(4)

obj.push(1)
print(obj.top())
print(obj.pop())
print(obj.top())
