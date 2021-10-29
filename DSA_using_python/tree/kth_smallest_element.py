
# A binary tree node
class Node:

    # Constructor to create a new node
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


# iterative solution

def solve(root, k):
    s = []
    while True:
        while root:
            s.append(root)
            root = root.right
        root = s.pop()
        k -= 1
        if not k:
            return root.data
        root = root.left


# recursive solution


# def solve(root, k):
#     arr = []
#     s1(root, arr)
#     for i in range(1, len(arr)+1):
#         if len(arr)+1 - k == arr[i]:
#             return arr[i]


def s1(root, arr):
    if root == None:
        return

    s1(root.left, arr)
    arr.append(root.data)
    s1(root.right, arr)

    return arr


# Driver program to test above function
root = Node(3)
root.left = Node(1)
root.right = Node(4)
#root.left.left = Node(4)
root.left.right = Node(2)
#root.left.left.left = Node(15)


print((solve(root, 4)))
